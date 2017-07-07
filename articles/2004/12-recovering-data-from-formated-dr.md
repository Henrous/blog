Id: 1877
Title: Recovering data from formatted drives
Date: 2004-12-13T04:06:27-08:00
Format: Markdown
Status: deleted
--------------
I've made a small, but tragic in consequences, mistake. I have my
biggest hard-drive in a removable hard-drive bay. It got disconnected,
the drive vanished from Windows, so I pushed it to connect it. The drive
didn't immediately show up in so I thought Windows XP just needs a
little push in recognizing it. I used Computer Management console, saw
the drive there but not mounted. And that's when I made fatal mistake:
I've added it as a new disk. The fatal mistake was the fact, that this
causes initializing the disk, which is almost as bad as formatting it.
In retrospect, I think I should have just rebooted the computer and
Windows should detect the drive.

So the data was on the drive, I just destroyed the crucial data that
allowed Windows to actually see it. The only option at this point is
formatting. But before that I wanted to recover the files from that
hard-drive.

Googling revealed that there are multiple programs to do that. I've
tried [Restorer2000 Professional](http://www.bitmart.net/r2k.shtml)
($50) and I'm happy to say that it worked.

Restorer2000 scans the whole hard-drive looking for partitions and file
systems and is able to reconstruct files from file system info (in my
case the disk used NTFS).

The are a few gotchas.

First, the process takes a long time, directly proportional to the size
of the disk. First, Restorer2000 has to scan the whole hard-drive (read
every sector on it) to detect all the files. Then, during actual
recovery, it has to read the files again. We're looking at countless
hours for big hard-drives (mine was >200 GB).

The other thing is that it cannot recover files in-place i.e. you need
as much free space on other hard-drives as the size of files you want to
recover. I wanted to recover everything so that was a bit of a problem.

Actually, you need more free space than that since Restorer2000 often
detects more than one file with the same file name. I guess it makes
sense since its first purpose is to detect deleted files and moving
files around or defragmenting hard-drives might leave the file system
info that Restorer2000 picks up. During recovery you can either tell
Restorer2000 to skip duplicate files or rename them. The second option
requires more hard-drive space.

I can recommend Restorer2000 since it did what it was supposed to do. I
have no idea how it compares to other software of this type since that's
the only one I've tried and I hope I won't have to try anything else in
the future.

I wish it had two more options, though.

First, I would like it to be able to pause and restart the hard-drive
scan and recovery process. Those can be very long so it would be nice to
be able to e.g. exit the program in the middle of recover, restart it
and have it start where it left off.
