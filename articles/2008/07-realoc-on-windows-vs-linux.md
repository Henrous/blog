Id: 3002
Title: realloc() on Windows vs. Linux
Tags: programming,optimization
Date: 2008-07-27T06:31:00-07:00
Format: Markdown
--------------
Short version: realloc() on Windows can have much worse performance than
on Linux/Mac.

As part of working on [Sumatra](https://www.sumatrapdfreader.org/), I sometimes work
on [mupdf](http://ccxvii.net/fitz/), underlying PDF parsing/rendering
engine. I’m not an expert on that code but I try to fix bugs and try to
improve performance.

In testing I noticed that on some PDFs it’s pathetically slow, so I
profiled the code and that told me that 99% of the time was spent in
realloc(). Also, the same PDF was rendered almost instantaneously on
Linux.

The code was reading a stream of unknown (a priori) length and the way
it worked was by reading 4k chunks of data into a buffer that was grown
with realloc() call.

If you’re lucky, realloc() will return the original block of memory and
just fiddle with internal allocator structures to make it bigger.

If you’re unlucky, realloc() will have to allocate a new block of memory
and copy the date from old memory.

If you’re reading 1MB file and growing file data buffer by 4k, you’ll
have to call realloc() 256 times and each time the buffer will be
bigger. If you’re unlucky and realloc() has to copy the data every time,
that’s a lot of cpu time spent data copying.

From observed behavior I have to conclude that realloc() implementation
on Linux and Mac is more likely to be able to reuse original memory than
implementation in Windows’ CRT (which most likely just calls
[HeapRealloc()](http://msdn.microsoft.com/en-us/library/aa366704.aspx)).

Once I knew what the problem was, the fix was simple.

[My first
implementation](http://code.google.com/p/sumatrapdf/source/detail?r=132&path=/fitz/stream/stm_misc.c)
would read into separate 4k buffers and at the end allocate the result
buffer and copy the data just once.

[My second
implementation](http://github.com/kjk/mupdf/commit/5d812a306fd0a192cfb5d28be58019c86e74b52b)
was even simpler: it would grow the buffer size by doubling it, instead
of just increasing by 4k. It was slower than first implementation (more
copying) but simpler and fast enough.

A side note: a simple and incorrect conclusion would be that memory
allocator implementation on Linux is much better than on Windows. Things
aren’t that simple, though. It’s easy to write an allocator that has
excellent performance for realloc() but is terrible for everything else.
Engineering is an art of compromises and optimizing realloc() apparently
wasn’t high on Microsoft’s list of things to shoot for when they were
designing allocator in Windows, which is, btw, comparable in performance
to allocator in Linux or Mac.
