Id: 1355
Title: Principle of good design: discoverability
Tags: ui design
Date: 2002-07-26T01:31:44-07:00
Format: Markdown
--------------
One of the principles of good design is discoverability: it should be
easy for users to discover what the program can do just by fooling
around.

Menus and toolbars add discoverability: they don't interfere if users
don't care about them but they're always there, visible, for users
to try out new features of the program.

Emacs, in general, has poor discoverability. Most of the work is being
done with arbitrary keystrokes and obscure lisp functions and the only
way to find out what a given keystroke does is by reading copious
manuals or using one of the obscure lisp functions to list current key
bindings.

This is efficient for experienced users but presents a very high barrier
to entry for new users.

There are two problems:

* it's hard to find out what a given keystroke does
* it's hard to find out what is a keystroke that does what I want

It would help discoverability if users could easily find those 2 things
out. Here enters my idea: we need to add a new window that will list all
keystrokes that are reachable from a given point.

Normally it would be empty. When users presses "Alt" key it would list
all keystrokes that start with "Alt" key. Entries should be sorted in
the order of usefulness.

That way users could see what a given keystroke does and with time learn
other useful keystroke by looking at the list of available options.

This window could be small, permanent window (with the ability to
hide/show it) or something that only pops-up when needed (similar to how
the auto-completion usually works). This idea is not limited to Emacs -
many programs implement key bindings as a short-cut for frequently used
functions and could benefit from a system like this.
