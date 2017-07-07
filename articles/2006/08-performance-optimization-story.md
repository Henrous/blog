Id: 950
Title: Performance optimization story
Tags: programming,optimization
Date: 2006-08-14T12:27:26-07:00
Format: Markdown
--------------
The story you’re about to read makes those major points about
optimizing software for speed:

* it’s good to read other people’s sources. You will learn new tricks.
* performance work is driven by data. Don’t guess what is slow, measure it.
* a good profiler is extremely helpful in getting the data
* lots of allocation of small objects isn’t good in a C/C++ program

When working on my [Sumatra PDF viewer](https://www.sumatrapdfreader.org/) for
Windows, I decided to take a look at the performance. I profiled the
code to parse a rather large (\~8MB) PDF. I found a rather surprising
thing: a lot of time was spent inside malloc()/free() (they were in the
top 10 most expensive functions in the profile) and a large portion of
those allocations/frees was for strings. The code in question has it’s
own, simple GooString class.

To get more data I instrumented GooString destructor to find out what
are the typical sizes of the strings. An allocation histogram told me
that about 90% of them is 16 bytes or less.

Then I looked at the implementation. GooString is a very typical
implementation. It keeps track of the size of string and a pointer to\
allocated string i.e. (to paraphrase):

```c++
class GooString {
 int length;
 char * str;
};
```

It does have an interesting trick. Most typical implementations allocate
more memory than strictly needed for the string, which avoids frequent
re-allocation when you add data to string. So they also have to keep
track of how big is the actual allocated area e.g.:

```
class DumberString {
 int allocated; /* the real size of ‘str’ buffer */
 int length;
 char *str;
};
```

GooString gets rid of ‘allocated’ variable by using a rounding function
based on size e.g.:

```c++
static inline int rounded\_size(int len) {
 int delta;
 delta = len < 256 ? 7 : 255;
 return ((len + 1) + delta) & ~delta;
}
```

That way `GooString` saves 4 bytes per object. Not that it usually
matters, as we’ll find out very shortly, but it illustrates that reading
other people’s code is useful. I’ve seen a couple of string
implementation but this is the first time I noticed that particular
trick and I would probably never have come up with that trick by myself.

The problem with `GooString` is that creating an instance causes 2
allocations: one for the object and another for the str pointer.

You might think that the amount of memory taken from the system for a
1-byte string (an empty string that only contains terminating zero)
would be `sizeof(GooString) (8) + 1` i.e. 9 bytes.

This is not so. First, most systems rounds allocation. You can find out
a rounding of your system with:
```
  int n = (char*)malloc(1) - (char*)malloc(1);
  printf(“rounding: %d\n”,n);
```

On Ubuntu Linux this turns out to be 16. Allocating 1 byte or 16
bytes uses the same amount of memory. One instance of `GooString` uses 32 bytes.

But that’s not all. The OS has to somehow keep track of each allocation.
How it’s done and what’s the exact overhead are highly implementation
dependent, but we can safely assume at least 8 bytes (that’s just 2
32-bit pointers). So the real cost of allocating a 1-byte GooString is
at least 48 bytes. And we thought it was 9.

There is a better way. A trick used in dynamic string implementation in\
venerable Tcl language uses a static buffer that is a part of the
string:

```c++
#define STR_STATIC_SIZE 16
class BetterString {
 char sStatic[STR_STATIC_SIZE];
 int length;
 char * ‘s’;
};
```

If the size of the string is less than `STR_STATIC_SIZE`, `str` points to
`sStatic`. If it’s bigger, we allocate the string as before. That way
for strings smaller than `STR_STATIC_SIZE` we don’t have to allocate
memory (halving the cost of allocations). It doesn’t even cost us more
memory in most cases since for small strings we avoid the minimum 24
bytes cost of allocating at least 1 byte, and for larger strings the
overhead is small compared to the total size.

You can tweak `STR_STATIC_SIZE`. The bigger it is, the faster we’ll be
(less cases where we need to allocate additional storage) but more
memory we’ll use.

In my particular case, implementing this trick reduced allocations due
to string by 45% (since 90% of strings were less than `STR_STATIC_SIZE`)
which improved loading time by 10%. And that was [a very simple
change](https://bugs.freedesktop.org/attachment.cgi?id=6496).

So let’s recap the things we can learn from this story.

The only way to know what is slow is to get data i.e. profile the app.
Pdf parser and renderer I use is a complex piece of code. It would be
pointless for me to try to guess which part of it is slow.

A good profiler is essential to giving the right data. An hour spent
profiling and reading the results pointed me in the right direction.

It’s important for a programmer to read other people’s source code. I’ve
learn new tricks from reading the source of `GooString`. I’ve learn new
tricks from reading Tcl’s implementation. In the end it’s much cheaper
than trying to come up with those ideas on my own.

And finally, as you can see, allocating small objects in C\\C++ has a
huge overhead, so try not to do it. Unfortunately naive implementation
of common data structures (strings, nodes in trees or lists) has a
problem of requiring lots of small allocation. A good answer to this
problem is custom allocator that pre-allocates large numbers of a given
object and uses a bitmap to keep track of which ones are used (1-bit
overhead per object as opposed to 8 + whatever rounding to 16 takes).
And, if done right, they should be faster than a standard OS allocator.
But that’s a story for another day.
