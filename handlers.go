package main

import (
	"net/http"
	"time"
)

func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}

	mux.HandleFunc("/", withAnalyticsLogging(handleMainPage))
	mux.HandleFunc("/favicon.ico", handleFavicon)
	mux.HandleFunc("/robots.txt", handleRobotsTxt)
	mux.HandleFunc("/contactme.html", withAnalyticsLogging(handleContactme))

	mux.HandleFunc("/book/go-cookbook.html", withAnalyticsLogging(handleGoCookbook))
	mux.HandleFunc("/articles/go-cookbook.html", withAnalyticsLogging(handleGoCookbook))

	mux.HandleFunc("/app/crashsubmit", withAnalyticsLogging(handleCrashSubmit))
	mux.HandleFunc("/app/debug", handleDebug)
	mux.HandleFunc("/app/sendmsg", handleSendMsg)
	mux.HandleFunc("/atom.xml", withAnalyticsLogging(handleAtom))
	mux.HandleFunc("/atom-all.xml", withAnalyticsLogging(handleAtomAll))
	mux.HandleFunc("/sitemap.xml", withAnalyticsLogging(handleSiteMap))
	mux.HandleFunc("/archives.html", withAnalyticsLogging(handleArchives))
	mux.HandleFunc("/software", withAnalyticsLogging(handleSoftware))
	mux.HandleFunc("/software/", withAnalyticsLogging(handleSoftware))
	mux.HandleFunc("/extremeoptimizations/", withAnalyticsLogging(handleExtremeOpt))
	mux.HandleFunc("/article/", withAnalyticsLogging(handleArticle))
	mux.HandleFunc("/kb/", withAnalyticsLogging(handleArticle))
	mux.HandleFunc("/blog/", withAnalyticsLogging(handleArticle))
	mux.HandleFunc("/forum_sumatra/", withAnalyticsLogging(forumRedirect))
	mux.HandleFunc("/articles/", withAnalyticsLogging(handleArticles))
	mux.HandleFunc("/book/", withAnalyticsLogging(handleArticles))
	mux.HandleFunc("/tag/", withAnalyticsLogging(handleTag))
	mux.HandleFunc("/static/", withAnalyticsLogging(handleStatic))
	mux.HandleFunc("/dailynotes-atom.xml", withAnalyticsLogging(handleNotesFeed))
	mux.HandleFunc("/dailynotes/week/", withAnalyticsLogging(handleNotesWeek))
	mux.HandleFunc("/dailynotes/tag/", withAnalyticsLogging(handleNotesTag))
	mux.HandleFunc("/dailynotes/note/", withAnalyticsLogging(handleNotesNote))
	mux.HandleFunc("/dailynotes", withAnalyticsLogging(handleNotesIndex))
	mux.HandleFunc("/tools/generate-unique-id", withAnalyticsLogging(handleGenerateUniqueID))
	mux.HandleFunc("/worklog", handleWorkLog)

	// not logged because not interesting for visitor analytics
	mux.HandleFunc("/ping", handlePing)
	mux.HandleFunc("/css/", handleCSS)
	mux.HandleFunc("/js/", handleJs)
	mux.HandleFunc("/gfx/", handleGfx)

	mux.HandleFunc("/djs/", withAnalyticsLogging(handleDjs))

	// websocket is only for dev mode, used for refreshing the pages if
	// they change on disk
	if !flgProduction {
		mux.HandleFunc("/ws", serveWs)
	}

	// https://blog.gopheracademy.com/advent-2016/exposing-go-on-the-internet/
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	return srv
}
