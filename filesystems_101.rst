File systems
************

Background
==========

A filesystem is the street grid of your hard drive. It's a map of addresses to
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

Navigating the filesystem
=========================

When you log into a Unix system, you will be given a command line by the
:doc:`shell <shells_101>` which may look something like this:

.. code-block:: console

  bash-4.0$

By default you will be in the "current working directory" if the process that
spawned the shell. Normally this is the home directory of your user. 
It can be different in some edge cases, such as if you manually change the
current working directory, but these cases are rare until you start doing more
advanced things.

You can find the name of the current directory with the ``pwd`` command:

.. code-block:: console

  bash-4.0$ pwd
  /home/opsschool

You can see the list of files and directories in this directory with the ``ls``
command:

.. code-block:: console

  bash-4.0$ ls
  file1.txt file2.txt tmpdir

The ``ls`` command also accepts the ``-l`` argument to provide a long-listing,
which will show you permissions, dates, ownership and other information:

.. code-block:: console

  bash-4.0$ ls -l
  -rw-r--r--  1 opsschool opsgroup   2444 Mar 29  2012 file1.txt
  -rw-r--r--  1 opsschool opsgroup  32423 Jun 03  2011 file2.txt
  drwxr-xr-x 15 opsschool opsgroup   4096 Apr 22  2012 tmpdir

You can see the contents of other directories, by giving the name of the
directory:

.. code-block:: console

  bash-4.0$ ls -l /
  dr-xr-xr-x    2 root root  4096 Apr 26  2012 bin
  dr-xr-xr-x    6 root root  1024 Sep 18 14:09 boot
  drwxr-xr-x   19 root root  8660 Jan  8 16:57 dev
  drwxr-xr-x  112 root root 12288 Feb  8 06:56 etc
  drwxr-xr-x   67 root root  4096 Feb  7 19:43 home
  dr-xr-xr-x   13 root root  4096 Mar  6  2012 lib
  drwx------    2 root root 16384 Sep 18  2011 lost+found
  drwxr-xr-x    5 root root  4096 Nov 19 18:53 mnt
  drwxr-xr-x    4 root root  4096 Sep  4 15:15 opt
  dr-xr-xr-x 1011 root root     0 Sep 23  2011 proc
  dr-xr-x---   10 root root  4096 Jan 23 23:14 root
  dr-xr-xr-x    2 root root 12288 Oct 16 22:23 sbin
  drwxr-xr-x   13 root root     0 Sep 23  2011 sys
  drwxrwxrwt   65 root root 16384 Feb 11 04:37 tmp
  drwxr-xr-x   16 root root  4096 Feb  8  2012 usr
  drwxr-xr-x   27 root root  4096 Nov  4 03:47 var

You may have noticed that the names of directories follow a pattern. ``/`` is
also called the root directory. All directories and files are contained under
it. From the first example, the ``/`` directory contains the ``/home``
directory, which in turn contains the ``/home/opsschool`` directory.

To change directories, use the ``cd`` command:

.. code-block:: console

  bash-4.0$ cd /tmp
  bash-4.0$ pwd
  /tmp

There may be times you need to find a file on your filesystem, based on its
name, date, size, or other particulars. For this you can use the ``find``
command:

.. code-block:: console

  bash-4.0$ find /home/opsschool -type f -name file3.txt
  /home/opsschool/tmpdir/file3.txt


Working with disks in Linux
===========================
Disks in Linux are normally named ``/dev/sda``, ``/dev/sdb``, etc.
If you are in a VM, they may be named ``/dev/xvda``, ``/dev/xvdb``, etc.
The last letter ("a", "b", "c"..) relates to the physical hard drive in your
computer. "a" is the first drive, "b" is the second.

If you have an already configured system, you will likely see entries like
this:

.. code-block:: console

    -bash-4.1$ ls -la /dev/sd*
    brw-rw---- 1 root disk 8, 0 Jul  6 16:51 /dev/sda
    brw-rw---- 1 root disk 8, 1 Sep 18  2011 /dev/sda1
    brw-rw---- 1 root disk 8, 2 Sep 18  2011 /dev/sda2
    brw-rw---- 1 root disk 8, 3 Sep 18  2011 /dev/sda3

The number at the end of each drive maps to the partition on the drive.
A partition refers to a fixed amount of space on the physical drive. Drives must
have at least one partition. Depending on your specific needs, you might want
more than one partition, but to start with, we'll assume you just need one big
partition.

Configuring your drive with partitions
======================================
man parted

Formatting partitions with new file systems
===========================================
man mkfs

Mounting a filesystem
=====================
.. todo:: explain different kinds of mounts, autofs, /etc/fstab

Filesystem options
==================
noatime
nobarriers

How filesystems work
====================
Files, directories, inodes

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

Fragmentation in unix filesystems
=================================

Filesystem objects
==================
Filesystem contain more than just files and directories.
Talk about devices (mknod), pipes (mkfifo), sockets, etc.
