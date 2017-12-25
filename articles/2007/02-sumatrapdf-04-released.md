Id: 1974
Title: SumatraPDF 0.4 released
Tags: sumatra
Date: 2007-02-19T21:24:41-08:00
Format: Markdown
--------------
Weekend came, weekend went but what it left is another release of
SumatraPDF - my humble reader for PDF files.

This release adds some of the most requested (in the forums) features
and fixes.

This is the first release to have printing, a feature contributed in
great part by Omar. It had limited testing (there are only so many pages
I'm willing to print).

It has improved support for portable usage (i.e. running off of e.g. USB
drive). When it detects it's not running from Program Files directory,
it assumes portable usage and doesn't write any files or registry items
on the computer. It writes its preferences files in the same directory
as the executable. Because of that I changed the name of preferences
file from prefs.txt to sumatrapdfprefs.txt, to make it clear it is
related to Sumatra. Unfortunate side effect is that settings from 0.3
will be lost.

It no longer blindly registers itself as a default handler for PDF
files. It asks the user nicely first.

Under the covers a lot of effort went into supporting alternative PDF
rendering engine. This new engine is significantly faster than poppler
so I've enabled it by default. You can switch back to poppler by
unchecking menu item.

And a bunch of fixes, most of them reported at some point on those
forums:

-   scrolling with mouse wheel works
-   fix toolbar issues on win2k
-   improve the way fonts directory is found (hopefully)
-   uninstaller completely removes the program

A side effect of using new engine is that support for link inside PDF
files was disabled. Well, one more reason to wait for the next
version...
