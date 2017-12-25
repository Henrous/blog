Id: 969
Title: Few things I've learned when writing Sumatra PDF
Tags: sumatra
Date: 2007-04-13T17:52:35-07:00
Format: Markdown
--------------

1\. [Good software takes ten years - get used to
it](http://www.joelonsoftware.com/articles/fog0000000017.html).

Version 0.1 had 573 trickling downloads.

Version 0.2 - about 7k, 0.3 - 7.5k, 0.4 - 15k.

It wasn't until version 0.5 when the downloads took off and as of today reached 80k.

It took several months and several program revisions before it was noticed,
linked to and got a decent ranking in google.

2\. People like .zip downloads.

This is Windows software so I focused on providing great installer.

I only added a zip archive because it was easy to do.

I expected that a majority of people will use the installer but it
turns out that about 30% of downloads are of zip archive.

Even though those numbers are a bit inflated (see below), it's still an unexpectedly big percentage.

3\. There are many websites that catalog software for Windows.

They write a little blurb and point to a program.

It's a win-win situation: software gets additional exposure, the website
provides a service to users and can make some money from ads. Except
that those websites sometimes get overboard in their focus on keeping
users on their website (so that they can make more money from ads).

One practice that I don't like is download link directly pointing to the
file instead of the home page for my program. It doesn't serve the
program author since it cuts people off from product and, for example,
they won't find out that they can leave a comment in the forum that will
help me improve the program in the future.

It's also not good for the users. I carefully designed my website to
provide useful information about the program. My website always has the
latest version (while others might link to an older version). One
website linked only to a zip version of my program so people didn't even
know that they could download an easier to use installer (and as stats
show, at least 70% of them would choose installer if given the choice).

I have little faith that an appeal to moral values can change behavior
of cowboys in our World Wild West, but there is a simple technical
solution to this problem. A little bit of mod\_rewrite magic based on
referrer field and now all request to older versions and all requests
not originating from my own websites are redirected to a generic
download page. Fortunately most of those operator are too cheap and lazy
to host the downloads themselves.
