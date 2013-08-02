Useful shell tools
******************

When you work in a unix environment, you will need to make frequent use of the
command line tools available to you in order to complete the tasks you have.

Mention that the philosophy here is to have many tools which perform a few
discrete tasks, and run them together to get what you want.

Working with your system
========================

ps
--

``ps`` shows the processes currently running on the system. ``ps`` takes many
arguments but some good recopies are ``ps aux`` to see all processes from a user
standpoint, whether they have a tty or not.
Also good is ``ps -lfU <username>`` to see all processes owned by ``<username>``
in long format. ``ps`` has output formatting with the ``-o`` switch and works 
well with ``grep``. ``pgrep`` combines ``ps`` and ``grep`` into one command.

.. code-block:: console

  $ ps aux | grep vim
  opsschool      5424  1.1  0.2 247360 10448 pts/0    Sl+  17:01   0:00 vim shell_tools_101.rst
  opsschool      5499  0.0  0.0  13584   936 pts/5    S+   17:01   0:00 grep --color=auto vim

Note that the ``grep`` command also shows up in the proccess table.

top
---

``top`` shows the top most cpu intensive processes in a table form with system
stats summarized at the top.
It refreshes every 1 second. By pressing ``m`` while running, ``top`` will sort 
not by processor use but by memory use.

df
--

``df`` looks at all your mounted filesystems and reports on their status. This
includes the filesystem type, total size, available space, percent used, and
mountpoint. ``df -h`` will show the sizes in human readable form. ``df -h .``
will show only the filesystem the current working directory is on, very useful
when debugging ``nfs`` and ``autofs`` mounts.

.. code-block:: console

    $ df -h .
    Filesystem      Size  Used Avail Use% Mounted on
    /dev/sda5        43G   31G  9.6G  77% /


du
--

``du`` estimates the size on disk of a file or files. ``du -h`` returns the
usage information in human readable format. If the argument to ``du`` is a
directory, ``du`` will run and return a new value for each file in that
directory, recursing into subdirectories as well. ``du -sh`` can be run on a
directory to prevent this behavior and return the sum of the space on disk of
all files under that directory.

.. code-block:: console

    $ du -sh drawing.svg 
    24K drawing.svg
    4.0K    style.css
    20K sitemeter.js

find
----

``find`` is for finding files on the system. ``find .`` recursively walks the
entire filesystem tree below the current working directory and prints out each 
file by its relative path. ``find . -name 'opsschool'`` again recursively walks
the entire filesystem tree, but only prints out files named ``opsschool``.
``find . -name 'opschool' -type d`` prints out only directories and ``-type f``
prints out only regular files. The ``-name`` filter also accepts wildcards.
``find . -name '*.html'`` will print out all files ending in ``.html``. As you 
become more skilled at Unix/Linux, use the find command often, as it becomes 
more and more useful with experience.


kill
----
``kill`` is used to send a signal to a process, which is a way of interrupting
it. The most common use of this is to stop, or kill, processes. For example, to
stop a process with the PID '123':

.. code-block:: console

  $ kill 123

Once the signal is sent the process can decide what to do with it. For instance
the above ``kill 123`` sends a ``TERM`` signal to the process with
PID 123. ``TERM`` (also known as ``SIGTERM``) means "Software Termination
Signal". Some processes will notice that signal, finish what they are doing, and
shut down. Others will just shut down immediately or ignore it entirely.

But there are many more types of signals. (Check out ``man signal`` for a good
list.) For instance, if the above kill command did not succeed in terminating
your process, a more heavy-handed option is to send a ``KILL`` signal to the
process:

.. code-block:: console

  $ kill -KILL 123

Every signal has a name and a number. You can reference them by either
one. Another way of running ``kill -KILL`` is:

.. code-block:: console

  $ kill -9 123

Be careful when using the ``KILL`` signal as it is the one signal that cannot be
caught by the process. It will not have a chance to gracefully shut down. This
can lead to temporary files not being removed, open files not being closed, or
even corruption of database files.

Signals can be used in a wide variety of ways, not just for terminating
processes. One interesting use: if your system is running Unicorn processes, you
can send a ``TTIN`` signal to the master process and it will spawn an additional
worker. Likewise, you can send a ``TTOU`` signal and it will remove one of the
workers. Another example is Apache which accepts ``USR1``, which causes it to
close and re-open log files, which is useful when you need to rotate your log
files.

For more on signals see :doc:`unix_signals`.

ls
--
``ls`` is used to list the contents of a directory.
It's most basic usage would be to list the contents of your shell's current working directory:

.. code-block:: console

    $ ls
    bar  foo

You can also pass a directory name to the ``ls`` command and it will list the contents of that directory:

.. code-block:: console

    $ ls /usr/local
    bin  etc  games  include  lib  libexec  sbin  share  src

There are a number of options that can be passed to the ls command to control both what is output and how it's formatted.
Two of the more useful are the -a option and the -l option.
Files and directories that begin with a '.' are refered to as hidden files.
``ls -a`` will list these hidden files and directories.
``ls -l`` outputs what's called a long listing, where various attributes are given in addition to the filenames:

.. code-block:: console

    $ ls -la
    total 26636
    drwx-----x. 39 luser luser   4096 Jun 28 01:56 .
    drwxr-xr-x.  4 root root     4096 Dec 11  2012 ..
    drwxrwxr-x   7 luser luser   4096 May 23 00:40 Applications
    -rw-------.  1 luser luser  16902 Jun 28 01:33 .bash_history
    -rw-r--r--.  1 luser luser     18 May 10  2012 .bash_logout
    -rw-r--r--.  1 luser luser    176 May 10  2012 .bash_profile

In a long listing the first field lists the file type, it's permissions, and also any special attributes it might have.
The very first letter in the first field indicates the file type.
Notice directories are indicated by a "d" and regular files are indicated by a "-".
After the first field, from left to right the fields are filetype\attributes\permissions, links, owner, group, file size, modification date, and file name.

Notice also the files named "." and "..".
These are the current directory and the directory up one level respectively.

lsof
----
``lsof`` lists open files.
This command can be very useful in examining what a particular process or user happens to be doing on a system.
For each open file information is listed such as the process id that holds the file open, the command that started the process, and the name of the user running the process.

``lsof`` doesn't just list regular files.
Of particular use is examing what network activity is currently going on.
This can be viewed with by issuing ``lsof -i``.

mount
-----
The ``mount`` command is used to mount filesystems.
For example, mounting an ext4 filesystem that resides on the :file:`/dev/sda2` partiton could be done as follows: ``mount -t ext4 /dev/sda2 /mnt``
In this example the "-t" switch tells the ``mount`` command that the filesystem type is ext4.
    
When passed no arguments the ``mount`` command lists the filesystems that are currently mounted:

.. code-block:: console

   $ mount
   /dev/sda2 on / type ext4 (rw)
   proc on /proc type proc (rw)
   sysfs on /sys type sysfs (rw)
   devpts on /dev/pts type devpts (rw,gid=5,mode=620)
   tmpfs on /dev/shm type tmpfs (rw)
   /dev/sda1 on /boot type ext4 (rw)


The ``mount`` command will also consult :file:`/etc/fstab` and if it's able to and use the entries and options it finds there.
If an entry for /home exists in /etc/fstab one would be able to simply issue the command ``mount /home``.
This command would mount whatever parition is found that is associated with the :file:`/home` entry in /etc/fstab, and use any options that happen to be present.

Things don't necessarily have to be a partition on a disk to be mounted either.
Mounting an iso file, an image of a optical disk, is especially handy: ``mount -o loop -t iso9660 /home/luser/installer.iso /mnt/cdrom``

``mount`` can also operate on currently mounted filesystems.
Of particular use is switching a currently mounted filesystem from read-write to read-only so that a filesystem check can be performed: ``mount -o remount,ro /``
This particular command tells mount to remount the currently mounted / filesystem as read-only.

There are quite a number of options that can be passed to the ``mount`` command's "-o" switch.
Some are filesystem independent, while others depend entirely on the type of filesystem that's being mounted.
Further documentation on either can be found in the man pages.

stat
----

vmstat
------

strace
------

ulimit
------

Extracting and manipulating data
================================

A very common pattern in unix is to take some data (a text file, a directory
listing, the output from a command) and either extract specific data from it,
change some of the data, or both. These tools help you when you do this.

cat
---

``cat`` outputs the contents of a file either to the shell, another file that already exists, or a file that does not yet exist.    

Perhaps most frequently, ``cat`` is used to print the contents of a file to the shell.  For example, if file foo.txt contains the word 'foo': ::

  $ cat /tmp/foo.txt
  -bash: bar

When cat is called on multiple files, the output is in the same order as the files.  If we have another file bar.txt that contains 'bar' and run: ::

  $ cat /tmp/foo.txt /home/jdoe/bar.txt
  -bash: foo bar

If you need to combine the contents of two files: ::

  $ cat /tmp/foo.txt /home/jdoe/bar.txt > /home/jdoe/foobar.txt
  $ cat /home/jdoe/foobar.txt
  -bash: foo
         bar

It is important to note that foobar.txt did not exist before running this command.  For this particular usage, ``cat`` can create a file "on the fly"

``cat`` can also be used to output the contents of one file to another file.  WARNING!  You should be careful when using ``cat`` this way since it will overwrite the contents of the receiving file. ::

  $ cat /tmp/foo.txt /home/jdoe/bar.txt
  $ cat /home/jdoe/bar.txt
  -bash: foo












  




cut
---
This is a very useful command which should be covered.

grep
----

awk
---

``awk`` is a very powerful utility that lets you extract and manipulate data from files.

For example, if you had a file ``students.txt`` that stored a list of student names, ages and email addresses in columns separated by a space: 

.. code-block:: console

  $ cat students.txt
  John Doe 25 john@example.com
  Jack Smith 26 jack@example.com
  Jane Doe 24 jane@example.com

Here, you can see that the first two columns have contain the student's name, the third has an age and the fourth, an email address.
You can use awk to extract just the student's first name and email like this: 

.. code-block:: console

  $ awk '{print $1, $4}' students.txt
  John john@example.com
  Jack jack@example.com
  Jane jane@example.com

By default, ``awk`` uses the space character to differentiate between columns. 
Using this, ``$1`` and ``$4`` told ``awk`` to only show the 1st and 4th columns of the file. 

The order in which the columns is specified is important, because ``awk`` will print them out to the screen in exactly that order.
So if you wanted to print the email column before the first name, here's how you would do it: 

.. code-block:: console

  $ awk '{print $4, $1}' students.txt
  john@example.com John
  jack@example.com Jack
  jane@example.com Jane

You can also specify a custom delimiter for ``awk`` and override the default one (the space character) by using the ``-F`` option.
Suppose the ``students.txt`` instead stored data like this:

.. code-block:: console

  $ cat students.txt
  John Doe - 25 - john@example.com
  Jack Smith - 26 - jack@example.com
  Jane Doe - 24 - jane@example.com

Now, if the ``-`` character is used as a delimiter, the first column would be the student's full name: 

.. code-block:: console
  
  $ awk -F '-' '{print $1}' students.txt
  John Doe
  Jack Smith
  Jane Doe

Using this same logic, the second column would be the student's age, and the third their email address. 

.. code-block:: console

  $ awk -F '-' '{print $1, $3}' students.txt
  John Doe john@example.com
  Jack Smith jack@example.com
  Jane Doe jane@example.com


``awk`` functionality doesn't stop at printing specific columns out to the screen; you can use it to:


* extract a specific row from the file using the ``NR`` command

.. code-block:: console

  $ awk 'NR==2' students.txt
  Jack Smith - 26 - jack@example.com

NOTE: I didn't have to use the ``-F`` option here since rows are being manipulated, and the ``-F`` option specifies a delimiter for column manipulation 


* extract lines longer than a specific length by using the ``length($0)`` command

.. code-block:: console

  $ awk 'length($0) > 30' students.txt
  John Doe - 25 - john@example.com
  Jack Smith - 26 - jack@example.com
  Jane Doe - 24 - jane@example.com

  $ awk 'length($0) > 32' students.txt
  Jack Smith - 26 - jack@example.com


* find the average of numbers in a column

.. code-block:: console

  $ awk -F '-' '{sum+=$2} END {print "Average age = ",sum/NR}' students.txt
  Average age =  25

In the last example, with the average age, ``{sum+=$2}`` tells awk to take each value in the second column and add it to the existing value of the variable ``sum``.
It's important to note here that the variable ``sum`` didn't have to be declared or initialised anywhere, ``awk`` creates it on-the-fly.
The ``END`` pattern tells ``awk`` what to do after all lines in the file have been processed.
In our case, that involves printing out the average age of all students.
To get the average age, the sum of all ages (stored in variable ``sum``) was divided by the total number of lines in the file, represented by ``NR``.

In addition to the ``END`` pattern, ``awk`` also provides a ``BEGIN`` pattern, which describes an action that needs to be taken before a the first line of the file is processed.

For example:

.. code-block:: console

  $ awk 'BEGIN {print "This is the second line of the file"} NR==2' students.txt
  This is the second line of the file
  Jack Smith - 26 - jack@example.com

sed
---
Only talk about replacing text for now? It's the most common / needed piece of
sed at this level.

sort
----
Definitely need to discuss sort and how to use it in basic activities.

man
___
This needs to be here.
