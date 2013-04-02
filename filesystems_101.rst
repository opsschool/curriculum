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

The ``parted`` tool allows you to create partitions on your disk. It also
can create disk labels. Disk labels describe the partitioning scheme. Most Linux
systems will have the msdos partitioning table, although newer systems with EFI
will be using the gpt partitioning table. Msdos labeled drives will have a
Master Boot Record, or MBR, at the beginning of the drive. This is where the
bootloader is installed. GPT labeled drives will usually have a FAT formatted
partition at the beginning of the disk for EFI programs and the bootloader.

``parted`` has a subshell interface, call it with the name of a block device:

.. code-block:: console


  roott@opsschool# parted /dev/sda
  GNU Parted 2.3
  Using /dev/sda
  Welcome to GNU Parted! Type 'help' to view a list of commands.
  (parted) print                                                            
  Model: ATA VBOX HARDDISK (scsi)
  Disk /dev/sda: 42.9GB
  Sector size (logical/physical): 512B/512B
  Partition Table: msdos

  Number  Start   End     Size    Type     File system  Flags
   1      8225kB  42.9GB  42.9GB  primary  ext4         boot

   (parted)    

In this example ``parted`` ran against ``/dev/sda``. The user then used the
``print`` command to print out information about the disk and the current
partitioning scheme. A 43GB disk is presented, using the msdos partition table
format. One partition exists and it is flagged as bootable.  Looking at a second
example:

.. code-block:: console

  root@opsschool# parted /dev/sdb
  GNU Parted 2.3
  Using /dev/sdb
  Welcome to GNU Parted! Type 'help' to view a list of commands.
  (parted) print                                                            
  Error: /dev/sdb: unrecognised disk label                                  
  (parted) mklabel msdos                                                    
  (parted) print                                                            
  Model: ATA VBOX HARDDISK (scsi)
  Disk /dev/sdb: 8590MB
  Sector size (logical/physical): 512B/512B
  Partition Table: msdos

  Number  Start  End  Size  Type  File system  Flags

  (parted) mkpart primary 1 1G                                              
  (parted) set 1 boot on                                                    
  (parted) mkpart primary 1G 5G                                             
  (parted) mkpart primary 5G 7G                                             
  (parted) mkpart primary 7G 8G                                             
  (parted) p                                                                
  Model: ATA VBOX HARDDISK (scsi)
  Disk /dev/sdb: 8590MB
  Sector size (logical/physical): 512B/512B
  Partition Table: msdos

  Number  Start   End     Size    Type     File system  Flags
   1      1049kB  1000MB  999MB   primary               boot
   2      1000MB  5000MB  3999MB  primary
   3      5000MB  7000MB  2001MB  primary
   4      7000MB  8590MB  1590MB  primary

  (parted) 

Here we have first found the disk label to be unreadable, so we created a new
disk label of the msdos type. After that we were able to see the disk was 8GB.
We created a primary 1GB partition at the beginning of the disk for ``/boot``
and set the bootable flag on that partition. We created 4GB, 2GB, and 1GB
partitions for root, var, and swap, respectively.



Formatting partitions with new file systems
===========================================
man mkfs

Mounting a filesystem
=====================

Mounting a filesystem is the act of placing the root of one filesystem on
a directory, or mount point, of a currently mounted filesystem. The mount
command allows the user to do this manually. Typically, only the superuser
can perform mounts. The root filesystem, mounted on ``/``, is unique and 
it is mounted at boot. See :doc:`boot_process_101`.

.. code-block:: console

    -root@opsschool # mount -t ext4 -o noatime /dev/sdb1 /mnt

It is common to specify which filesystem type is present on ``/dev/sdb1`` and
which mounting options you would like to use, but if that information is not 
specified then the Linux ``mount`` command is pretty good at picking sane 
defaults. Most administrators would have typed the following instead of the
above:

.. code-block:: console

    -root@opsschool # mount /dev/sdb1 /mnt

The root filesystem, mounted on ``/``, is unique and its mounting 

The type of the filesystem refers to the format of the data structure that is 
the filesystem on disk. Files (generally) do not care what kind of filesystem 
they are on, it is only in initial filesystem creation, automatic 
mounting, and performance tuning that you have to concern yourself with the 
filesystem type. Example filesystem types are ``ext2, ext3, ext4, FAT, NTFS 
HFS, JFS, XFS, ZFS, BtrFS``. On Linux hosts, ext4 is a good default.  For 
maximum compatibility with Windows and Macintosh, use FAT. 

http://en.wikipedia.org/wiki/Comparison_of_file_systems

The fstab, or file system table, is the file that configures automatic mounting
at boot. It tabulates block devices, mount points, type and options for each 
mount.  The dump and pass fields control booting behavior. Dumping is the act 
of creating a backup of the filesystem (often to tape), and is not in common use.
Pass is much more important. When the pass value is nonzero, the filesystem is 
analyzed early in the boot process by fsck, the file system checker, for errors.
The number, fs_passno, indicated priority. The root filesystem should always be
1, other filesystems should be 2 or more. A zero value causes skips to be
checked, an option often used to accelerate the boot process. In ``/etc/fstab``,
there are a number of ways to specify the block device containing the filesystem
. ``UUID``, or universally unique identifier, is one common way in modern Linux
based systems to specify a filesystem.

.. code-block:: console

    -root@opsschool # cat /etc/fstab

    # <file system> <mount point>  <type>  <options>         <dump>  <pass>
    /dev/sda5         /            ext4    errors=remount-ro 0       1
    /dev/sda6         none         swap    sw                0       0
    /dev/sda1         /boot/efi    auto    auto              0       0 

This ``/etc/fstab`` file mounts ``/dev/sda5`` on ``/`` using the ext4 filesystem
. If it encounters a filesystem corruption it will use the ``fsck`` utility
early in the boot process to try to clear the problem. If the physical disk
reports errors in writing while the filesystem is mounted, the os will remount
``/`` readonly. The ``/dev/sda6`` partition will be used as swap. The
``/dev/sda1`` partition will be mounted on ``/boot/efi`` using autodetection, the
partition will not be scanned for filesystem errors.

.. code-block:: console

    -root@opsschool # cat /etc/auto.master

    /home -rw,hard,intr,nosuid,nobrowse bigserver:/exports/home/&
    /stash ldap:ou=auto_stash,ou=Autofs,dc=example,dc=com -rw,hard,intr,nobrowse

The ``auto.master`` file controls the ``autofs`` service. It is another way to
tabulate filesystems for mounting. It is different from the ``/etc/fstab`` 
because the filesystems listed in ``auto.master`` are not mounted at boot. The
automounter allows the system to mount filesystems on demand, then clean up those
filesystems when they are no longer being used. In this case, the system mounts
home directories for each user from a remote NFS server. The filesystem remains
unmounted until the user logs in, and is unmounted a short time after the user
logs out. The automounter is triggered by an attempt to cd into ``/home/<key``,
it will then attempt to find an nfs share on ``/exports/home/<key>`` and mount it
on ``/home/key``, then allow the ``cd`` command to return successfully. The 
``/home`` example above is using the ``&`` expansion syntax, the second line is
using the LDAP syntax to look up a key under ``/stash/<key>`` in LDAP. LDAP will
be covered later in the curriculum. The ``auto.master`` file is known as
``auto_master`` on FreeBSD, Solaris, and Mac OS X.



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
