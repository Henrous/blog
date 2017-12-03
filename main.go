package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	_ "net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/kjk/u"
	"github.com/rs/xid"
	"github.com/skratchdot/open-golang/open"

	"golang.org/x/crypto/acme/autocert"
)

var (
	analyticsCode = "UA-194516-1"

	logger       *ServerLogger
	dataDir      string
	store        *ArticlesStore
	sha1ver      string
	randomCookie string
)

func genRandomCookie() {
	randomCookie = xid.New().String()
}

func getDataDir() string {
	if dataDir != "" {
		return dataDir
	}

	dirsToCheck := []string{"/data", u.ExpandTildeInPath("~/data/blog")}
	for _, dir := range dirsToCheck {
		if u.PathExists(dir) {
			dataDir = dir
			return dataDir
		}
	}

	log.Fatalf("data directory (%v) doesn't exist", dirsToCheck)
	return ""
}

func isTopLevelURL(url string) bool {
	return 0 == len(url) || "/" == url
}

// this list was determined by watching /logs
var noLog404 = map[string]bool{
	"/crossdomain.xml":                                               true,
	"/article/Exercise-links-1.html":                                 true,
	"/article/Ecco-for-free.html":                                    true,
	"/article/Disappointed-by-The-Bat.html":                          true,
	"/article/Comments-need-not-apply.html":                          true,
	"/article/Browsing-Newton.html":                                  true,
	"/article/Perl-and-lisp-programmers.html":                        true,
	"/article/iPod-competition.html":                                 true,
	"/article/Programming-Jabber.html":                               true,
	"/article/Good-software-design-contradicts-eXtreme-Program.html": true,
	"/article/Bloglines-vs-Google-Reader-the-verdict.html":           true,
	"/2002/07/30/stuid-coding-mistake-of-the-day.html":               true,
	"/article/Corman-Lisp.html":                                      true,
	"/article/Offshore-outsourcing.html":                             true,
	"/article/Nabble-hosted-forums.html":                             true,
}

func shouldLog404(s string) bool {
	if strings.HasPrefix(s, "/apple-touch-icon") {
		return false
	}
	_, ok := noLog404[s]
	return !ok
}

func setContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}

func writeResponse(w http.ResponseWriter, responseBody string) {
	w.Header().Set("Content-Length", strconv.FormatInt(int64(len(responseBody)), 10))
	io.WriteString(w, responseBody)
}

func textResponse(w http.ResponseWriter, text string) {
	setContentType(w, "text/plain")
	writeResponse(w, text)
}

var (
	flgHTTPAddr        string
	flgProduction      bool
	flgNetlifyBuild    bool
	flgUpdateNotes     bool
	flgNewArticleTitle string
)

func parseCmdLineFlags() {
	flag.StringVar(&flgHTTPAddr, "addr", ":5020", "HTTP server address")
	flag.BoolVar(&flgProduction, "production", false, "are we running in production")
	flag.BoolVar(&flgUpdateNotes, "update-notes", false, "if true, we make sure that all notes have ids")
	flag.BoolVar(&flgNetlifyBuild, "netlify-build", false, "if true, builds and deploys to netlify")
	flag.StringVar(&flgNewArticleTitle, "newarticle", "", "create a new article")
	flag.Parse()
}

func isTmpFile(path string) bool {
	return strings.HasSuffix(path, ".tmp")
}

func sanitizeForFile(s string) string {
	var res []byte
	toRemove := "/\\#()[]{},?+.'\""
	var prev rune
	buf := make([]byte, 3)
	for _, c := range s {
		if strings.ContainsRune(toRemove, c) {
			continue
		}
		switch c {
		case ' ', '_':
			c = '-'
		}
		if c == prev {
			continue
		}
		prev = c
		n := utf8.EncodeRune(buf, c)
		for i := 0; i < n; i++ {
			res = append(res, buf[i])
		}
	}
	if len(res) > 32 {
		res = res[:32]
	}
	s = string(res)
	s = strings.Trim(s, "_- ")
	s = strings.ToLower(s)
	return s
}

func findUniqueArticleID(articles []*Article) string {
	existingIDs := make(map[string]bool)
	for _, a := range articles {
		existingIDs[a.ID] = true
	}

	for i := 1; i < 10000; i++ {
		s := u.EncodeBase64(i)
		if !existingIDs[s] {
			return strconv.Itoa(i)
		}
	}
	u.PanicIf(true, "couldn't find unique article id")
	return ""
}

func genNewArticle(title string) {
	fmt.Printf("genNewArticle: %q\n", title)
	store, err := NewArticlesStore()
	if err != nil {
		log.Fatalf("NewStore() failed with %s", err)
	}
	newID := findUniqueArticleID(store.articlesWithDrafts)
	t := time.Now()
	dir := "articles"
	yyyy := fmt.Sprintf("%04d", t.Year())
	month := t.Month()
	sanitizedTitle := sanitizeForFile(title)
	name := fmt.Sprintf("%02d-%s.md", month, sanitizedTitle)
	fmt.Printf("new id: %s, name: %s\n", newID, name)
	path := filepath.Join(dir, yyyy, name)
	s := fmt.Sprintf(`---
Id: %s
Title: %s
Date: %s
Format: Markdown
---`, newID, title, t.Format(time.RFC3339))
	for i := 1; i < 10; i++ {
		if !u.PathExists(path) {
			break
		}
		name = fmt.Sprintf("%02d-%s-%d.md", month, sanitizedTitle, i)
		path = filepath.Join(dir, yyyy, name)
	}
	u.PanicIf(u.PathExists(path))
	fmt.Printf("path: %s\n", path)
	u.CreateDirForFileMust(path)
	ioutil.WriteFile(path, []byte(s), 0644)
}

func loadArticles() {
	var err error
	if store, err = NewArticlesStore(); err != nil {
		log.Fatalf("NewStore() failed with %s", err)
	}
	articles := store.GetArticles(true)
	articlesJs, articlesJsSha1 = buildArticlesJSON(articles)
}

func main() {
	genRandomCookie()

	parseCmdLineFlags()

	if false {
		testAnalyticsStats("/Users/kjk/Downloads/2017-06-02.txt.gz")
		os.Exit(0)
	}

	if !flgProduction {
		updateMainCSSSha1Must()
	}

	if flgNewArticleTitle != "" {
		genNewArticle(flgNewArticleTitle)
		return
	}

	if flgProduction {
		reloadTemplates = false
		flgHTTPAddr = ":80"
	} else {
		analyticsCode = ""
	}

	logger = NewServerLogger(256, 256)

	rand.Seed(time.Now().UnixNano())

	if flgUpdateNotes {
		notesGenIDIfNecessary()
		return
	}

	loadArticles()

	readRedirects()

	if flgNetlifyBuild {
		netlifyBuild()
		return
	}

	analyticsPath := filepath.Join(getDataDir(), "analytics", "2006-01-02.txt")
	initAnalyticsMust(analyticsPath)

	var wg sync.WaitGroup
	var httpsSrv *http.Server

	if flgProduction {
		hostPolicy := func(ctx context.Context, host string) error {
			allowedDomain := "kowalczyk.info"
			if strings.HasSuffix(host, allowedDomain) {
				return nil
			}
			return fmt.Errorf("acme/autocert: only *.%s hosts are allowed", allowedDomain)
		}

		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: hostPolicy,
			Cache:      autocert.DirCache(getDataDir()),
		}

		httpsSrv = makeHTTPServer()
		httpsSrv.Addr = ":443"
		httpsSrv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
		logger.Noticef("Starting https server on %s\n", httpsSrv.Addr)
		go func() {
			wg.Add(1)
			err := httpsSrv.ListenAndServeTLS("", "")
			// mute error caused by Shutdown()
			if err == http.ErrServerClosed {
				err = nil
			}
			u.PanicIfErr(err)
			fmt.Printf("HTTPS server shutdown gracefully\n")
			wg.Done()
		}()
	}

	httpSrv := makeHTTPServer()
	httpSrv.Addr = flgHTTPAddr
	logger.Noticef("Starting http server on %s, in production: %v, ver: github.com/kjk/blog/commit/%s", httpSrv.Addr, flgProduction, sha1ver)
	go func() {
		wg.Add(1)
		err := httpSrv.ListenAndServe()
		// mute error caused by Shutdown()
		if err == http.ErrServerClosed {
			err = nil
		}
		u.PanicIfErr(err)
		fmt.Printf("HTTP server shutdown gracefully\n")
		wg.Done()
	}()

	if flgProduction {
		sendBootMail()
	}

	if !flgProduction {
		open.Run("http://" + flgHTTPAddr)
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt /* SIGINT */, syscall.SIGTERM)
	sig := <-c
	fmt.Printf("Got signal %s\n", sig)
	ctx := context.Background()
	if httpsSrv != nil {
		httpsSrv.Shutdown(ctx)
	}
	if httpSrv != nil {
		// Shutdown() needs a non-nil context
		httpSrv.Shutdown(ctx)
	}
	wg.Wait()
	analyticsClose()
	fmt.Printf("Exited\n")
}
