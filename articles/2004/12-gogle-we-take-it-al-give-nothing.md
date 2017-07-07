Id: 1884
Title: Google - we take it all, give nothing back
Date: 2004-12-29T19:10:34-08:00
Format: Markdown
Status: deleted
--------------
I have a new hobby: [inventing advertising
slogans](../../../2004/12/25/google-ultimate-hypocrite.html) for Google.

Being the nice guy that I am, I grant Google non-exclusive, non-transferable
right to use those slogans in their marketing efforts.

Today, beside few new Google slogans, I give you the following shocking
revelations:

-   open source doesn't work as well as they tell you
-   Microsoft creates more open-source code than Google
-   how to bring down Microsoft

This post is for mature audiences only, ridicules Google and everyone
who gets in the way, doesn't use web standards and is not XHTML
compliant.

**Open-source - not working as advertised.**

The popular theory ("myth" would be a better name) is that open-source
works because of this positive feedback loop:

-   source code for product foo is released
-   it's free so it gets used
-   if it doesn't fully meet someone's needs, that someone can code the
    functionality (since the code is open) and submit the changes back
    to project (something not possible if you use closed products like
    Windows or Office or Google)
-   those contributions improve the product for everyone else, so more
    people use it so more people contribute the code and so on. Sky is
    hardly the limit.

The good thing in this theory is that it doesn't rely on kindness
of strangers but on enlightened self-interest of those who benefit
from free software. The bad thing about this theory is that in theory
it works much better than in practice.

It's all because of a [weblog post](http://www.adambosworth.net/archives/000038.html) by
Google's Adam Bosworth. Read it yourself, but the gist
of it is that, according to Adam, commercial database vendors don't
understand the needs of companies like Google or Amazon or Federal
Express.

Relational database rely on static schemas and there are no good ways
to dynamically reconfigure databases without the disruption in service.
Adam ends with a plea to open-source fairy:

> My message is to the Open Source community that has, so ably, built
> LAMP (Linux, Apache and Tomcat and MySQL and PHP and PERL and Python).
> Please finish the job. Do for databases what you did for web servers.
> Give us dynamism and robustness. Give us systems that scale linearly,
> are flexible and dynamically reconfigurable and load balanced and
> easy to use.

This is why the theory of open source doesn't work in real world. A
multi-billion company has a clear need for software that works well
for them but instead of investing in existing open-source projects
like PostgreSQL or MySQL to make them do what they need, all they
do is ask some magic, undefined entity they call Open Source community
to do the work for them. For free.

**Google - we take it all, give nothing back. Come work for us.**

Let's estimate how much money did Google save by using open source
software that they would otherwise have to purchase. The operating
system for tens of thousands of their computers. Web servers they use.
All the Unix utilities they use. Editors, compilers and debuggers they
use to write their code. E-mail smtp server. E-mail pop servers.
Languages like Perl and Python. Databases like MySQL and PostgreSQL.
It's safe to say that if Richard Stallman was never born, the licenses
for those kinds of software would cost them tens of millions of dollars.

And what does Google contribute back? Where are their patches to gcc,
gdb, python, postgresql, sendmail, emacs?

**Google - we leave open-source to Microsoft. Come work for us.**

It's very ironic that I can find more open-source code created
by Microsoft and its employees (
[RSS Bandit](http://rssbandit.org/),
[IronPython](http://ironpython.com/),
[Windows Installer XML (WiX)](http://sourceforge.net/projects/wix/),
[FlexWiki](http://sourceforge.net/projects/flexwiki/)) than
by Google employees. Not saying that there aren't any but they are
certainly not easy to find, even when I use mighty search engine
trying to find [google
open-source](http://www.google.com/search?q=google+open-source&sourceid=mozilla-search&start=0&start=0&ie=utf-8&oe=utf-8&client=firefox-a&rls=org.mozilla:en-US:official).

**Google - we like our hardware cheap and our software free. Come
work for us.**

If you're into this stuff you know that Google is known for it's highly
tuned process of selecting hardware components (i.e. all those
thousands of computers they need to index and store the web) to hit the best
price/performance ratio. In a way, they use the cheapest thing, when
you define the cost as the total cost of ownership (as opposed to simply
the cost of buying the hardware). Thanks to Adam's admission:
> Indeed, in these days of open source, I wonder if the software
> itself, should cost at all?\
we also know, that they like their software free.

As a side note, it's a surprising statement coming from Adam who knows
very well that writing software costs a lot. Open-source doesn't
eliminate this cost, it just shift the costs and allows unlimited
number of free-riders, like Google.

I'm picking on Google, but they are not alone. Amazon, Yahoo, eBay,
aol. Any large business that uses web as means of providing services and
making revenues is enjoying enormous savings by using open source
stack on their back end. And what do they contribute back? A good
approximation of zero compared to benefits they reap.

Another irony: compared to this bunch Microsoft looks good, they
are the only ones that did pay for the software they use on their
back-end

The thing is: I actually believe that positive feedback loop in open
source is possible. Except that it requires enlightenment on the part
of executives at all those companies. They should recognize that
they should invest in improving in open source software that they
use as their infrastructure. Not out of gratitude but because it makes
good sense in the long term. It's the cheapest way for them to get
software they use for building their services.

But Adam's example shows that there's a fat chance of this happening.
Adam is not a rank Google employee. He was not hired to give free
massage to stressed Google employees. Before Google Adam was a high-ranked
executive at Microsoft and BEA. He led teams that created successful products
(IE, Access among them). He's in position to influence what Google does.
He understands technology, he understand the cost and difficulty of making
software. He has a weblog and deep thoughts. If only he understood the
strategic value of open source.

If someone like Adam cannot see further than the tip of his own nose
and his ideas are as bold as asking others to write the software
he needs for free, then I don't have much hope for anyone at
aol to get it either.

**The Art of War - how to bring down Microsoft.**

As an additional bonus, my prescription on how amazon, yahoo, ebay, aol
and ibm could band together and bring down Microsoft. The reason Microsoft
is a competition to almost anybody who gets big, including yahoo, aol,
ibm and google, is that in order to fuel their growth, they have to
enter every profitable, software-related business. They are late to
almost every party, be it web browser, game consoles, instant
messaging, free web e-mail, webloging, decent search but they're still a
formidable enemy because they can fund those businesses longer than anybody else
thanks to their cash cows: Windows, Office, SQL Server and Exchange.

Kill those cash cows and Microsoft has nothing. They can no longer
support money-loosing businesses for years.

At this point no other company can hope to make good business in those
markets. There are open source alternatives but they are not good
enough to make switching to a non-Microsoft solution a possibility.

Combine the resources of google, yahoo, aol, ibm and invest serious
money into open source projects. Improve Linux, drivers, desktop
apps to be as good as Windows. Invest in Open Office to kill Office.
Invest in Firefox so that there's no reason at all to use IE. Invest
in Thunderbird so that it can replace Outlook, including Exchange
connectivity. Invest in PostgreSQL to kill SQLServer. Invest in
something to replace Exchange. And be prepared to hire 50 thousand
available Microsoft developers.

**Google - do no Evil. Do no Good. Just like everybody else. Come work for us.**

Let's consider the moral side of the story. Is Google (or any other
company) that benefits so much from open-source obliged to contribute
back? And I don't mean in a token way, like donating $100 to Mozilla
foundation. I'm talking about investing at least 10% of estimated
savings back into open source. Either by donating money to established venues
like EFF, FSF, Mozilla Foundation or by hiring developers to spend
all their time developing the code for open-source projects.

In those days of focus on corporate profits (where there any other
days?), Google's motto "Do no Evil" is refreshing.

Or is it? It's a nice soundbite, but when you think about it, it's
really a low requirement. There are very little things that deserve
to be called Evil. If a senior citizen is taking a nap outside his house
on a sunny day and you kick him in the groin - that's
Evil. Most other things are bad or neutral. Not doing Evil is easy.
Doing Good is the hard thing.
