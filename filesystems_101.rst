File systems
************

Background
==========

A filesystem is the street grid of your hard drive. It’s a map of addresses to
where data is located on your drive. Your operating system uses the filesystem
to store data on the drive.

There are a number of different types of filesystems. Some are better at
handling many small files (ReiserFS), some are much better at large files and
deleting files quickly (XFS, EXT4).

The version of Unix you use will have picked a filesystem which is used by
default, on Linux this is often EXT3.

Understanding the way filesystems work is important when you have to fix issues
related to disk space, performance issues with reading and writing to disk, and
a host of other issues.

In this section we will discuss creating partitions, file systems on those
partitions, and then mounting those file systems so your operating system can
use them.

Working with disks in Linux
===========================
Disks in Linux are normally named ``/dev/sda``, ``/dev/sdb``, etc.
If you are in a VM, they may be named ``/dev/xvda``, ``/dev/xvdb``, etc.
The last letter ("a", "b", "c"..) relates to the physical hard drive in your
computer. "a" is the first drive, "b" is the second.

If you have an already configured system, you will likely see entries like
this::

    -bash-4.1$ ls -la /dev/sd*
    brw-rw---- 1 root disk 8, 0 Jul  6 16:51 /dev/sda
    brw-rw---- 1 root disk 8, 1 Sep 18  2011 /dev/sda1
    brw-rw---- 1 root disk 8, 2 Sep 18  2011 /dev/sda2
    brw-rw---- 1 root disk 8, 3 Sep 18  2011 /dev/sda3

The number at the end of each drive maps to the partition on the drive.
A partition refers to a fixed amount of space on the physical drive. Drives must
have at least one partition. Depending on your specific needs, you might want
more than one partition, but to start with, we’ll assume you just need one big
partition.

Configuring your drive with partitions
======================================
man parted

Formatting partitions with new file systems
===========================================
man mkfs

How filesystems work
====================
Files, directories, inodes

Navigating the filesystem
=========================
``cd``, ``ls``, ``rm``, ``find``

Inodes
======
What the contain, how they work

The POSIX standard dictates files must have the following attributes:

* File size in bytes.
* A device id.
* User ID of file's owner.
* Group ID of file.
* The file's mode (permissions).
* Additional system and user flags (e.g. append only or ACLs).
* Timestamps when the inode was last modified (ctime), file content last modified/accessed (mtime/atime).
* Link count of how many hard links point to the inode.
* Pointers to the file's contents.

http://en.wikipedia.org/wiki/Inode

File system layout
==================
File system hierarchy standard is a reference on managing a Unix filesystem or directory structure.

http://www.pathname.com/fhs/

Filesystem options
==================
noatime
nobarriers

Fragmentation in unix filesystems
=================================

Filesystem objects
==================
Filesystem contain more than just files and directories.
Talk about devices (mknod), pipes (mkfifo), sockets, etc.
