Useful shell tools
******************

When you work in a unix environment, you will need to make frequent use of the command line tools available to you in order to complete the tasks you have.

The Unix Philosophy
===================

Unix is characterized by its modular design philosophy. 
Unix programs are small, single purpose tools that can easily be chained together with other Unix programs. 
This modular approach allows for a much simpler, more flexible system, since you can combine multiple command line programs to solve a unique problem, without having to write a whole new program to get the job done.

Working with your system
========================

ps
--

``ps`` shows the processes currently running on the system.
``ps`` takes many arguments but some good recipes are ``ps aux`` to see all processes from a user standpoint, whether they have a tty or not.
Also good is ``ps -lfU <username>`` to see all processes owned by ``<username>`` in long format.
``ps`` has output formatting with the ``-o`` switch and works well with ``grep``.
``pgrep`` combines ``ps`` and ``grep`` into one command.

.. code-block:: console

  $ ps aux | grep vim
  opsschool      5424  1.1  0.2 247360 10448 pts/0    Sl+  17:01   0:00 vim shell_tools_101.rst
  opsschool      5499  0.0  0.0  13584   936 pts/5    S+   17:01   0:00 grep --color=auto vim

Note that the ``grep`` command also shows up in the process table.

top
---

``top`` shows the top most cpu intensive processes in a table form with system stats summarized at the top.
It refreshes every 1 second.
By pressing ``m`` while running, ``top`` will sort not by processor use but by memory use.

df
--

``df`` looks at all your mounted filesystems and reports on their status.
This includes the filesystem type, total size, available space, percent used, and mountpoint.
``df -h`` will show the sizes in human readable form.
``df -h .`` will show only the filesystem the current working directory is on, very useful when debugging ``nfs`` and ``autofs`` mounts.

.. code-block:: console

  $ df -h .
  Filesystem      Size  Used Avail Use% Mounted on
  /dev/sda5        43G   31G  9.6G  77% /


du
--

``du`` estimates the size on disk of a file or files.
``du -h`` returns the usage information in human readable format.
If the argument to ``du`` is a directory, ``du`` will run and return a new value for each file in that directory, recursing into subdirectories as well.
``du -sh`` can be run on a directory to prevent this behavior and return the sum of the space on disk of all files under that directory.

.. code-block:: console

  $ du -sh drawing.svg
  24K drawing.svg
  4.0K    style.css
  20K sitemeter.js

find
----

``find`` is for finding files on the system.
``find .`` recursively walks the entire filesystem tree below the current working directory and prints out each file by its relative path.
``find . -name 'opsschool'`` again recursively walks the entire filesystem tree, but only prints out files named ``opsschool``.
``find . -name 'opsschool' -type d`` prints out only directories and ``-type f`` prints out only regular files.
The ``-name`` filter also accepts wildcards.
``find . -name '*.html'`` will print out all files ending in ``.html``.
As you become more skilled at Unix/Linux, use the find command often, as it becomes more and more useful with experience.


kill
----
``kill`` is used to send a signal to a process, which is a way of interrupting it.
The most common use of this is to stop, or kill, processes. For example, to stop a process with the PID '123':

.. code-block:: console

  $ kill 123

Once the signal is sent the process can decide what to do with it.
For instance the above ``kill 123`` sends a ``TERM`` signal to the process with PID 123.
``TERM`` (also known as ``SIGTERM``) means "Software Termination Signal".
Some processes will notice that signal, finish what they are doing, and shut down.
Others will just shut down immediately or ignore it entirely.

There are many more types of signals.
(Check out ``man signal`` for a good list.)
For instance, if the above kill command did not succeed in terminating your process, a more heavy-handed option is to send a ``KILL`` signal to the process:

.. code-block:: console

  $ kill -KILL 123

Every signal has a name and a number. You can reference them by either one.
Another way of running ``kill -KILL`` is:

.. code-block:: console

  $ kill -9 123

Be careful when using the ``KILL`` signal as it is the one signal that cannot be caught by the process.
It will not have a chance to gracefully shut down.
This can lead to temporary files not being removed, open files not being closed, or even corruption of database files.

Signals can be used in a wide variety of ways, not just for terminating processes.
One interesting use: if your system is running Unicorn processes, you can send a ``TTIN`` signal to the master process and it will spawn an additional worker.
Likewise, you can send a ``TTOU`` signal and it will remove one of the workers.
Another example is Apache HTTPD Web Server which accepts ``USR1``, which causes it to close and re-open log files, which is useful when you need to rotate your log files.

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
Files and directories that begin with a '.' are referred to as hidden files.
Two of the more useful options are: ``-a`` and ``-l``:

- ``ls -a`` will list these hidden files and directories.
- ``ls -l`` outputs what's called a long listing, where various attributes are given in addition to the filenames.

Example of using both:

.. code-block:: console

  $ ls -la
  total 26636
  drwx-----x. 39 luser luser   4096 Jun 28 01:56 .
  drwxr-xr-x.  4 root root     4096 Dec 11  2012 ..
  drwxrwxr-x   7 luser luser   4096 May 23 00:40 Applications
  -rw-------.  1 luser luser  16902 Jun 28 01:33 .bash_history
  -rw-r--r--.  1 luser luser     18 May 10  2012 .bash_logout
  -rw-r--r--.  1 luser luser    176 May 10  2012 .bash_profile

In a long listing the first field lists the file type, its permissions, and also any special attributes it might have.
The very first letter in the first field indicates the file type.
Notice directories are indicated by a "d" and regular files are indicated by a "-".
After the first field, from left to right the fields are filetype\attributes\permissions, links, owner, group, file size, modification date, and file name.

Notice also the files named "." and "..".
These are the current directory and the directory up one level, respectively.

lsof
----
``lsof`` lists open files.
This command can be very useful in examining what a particular process or user happens to be doing on a system.
For each open file information is listed such as the process id that holds the file open, the command that started the process, and the name of the user running the process.

``lsof`` doesn't just list regular files.
Of particular use is examing what network activity is currently going on.
This can be viewed with by issuing ``lsof -i``.

man
---
The ``man`` command is used to access the ``man`` pages.
A ``man`` page, short for manual page, is documentation on a particular aspect of your operating system, be it a command, a configuration file, or even functions from a library for a programming language.
To access a ``man`` page, simply type the ``man`` command followed by the name of the command, file, etc. that you want to view documentation on.
In the old days these manuals were hardcopy and on some systems (e.g. Solaris) you will still see evidence of their page layout.
There are ``man`` pages for most, if not all, programs on your system.
If you install new software from a package manager, usually new ``man`` pages will be installed.
When man is invoked, it will display the ``man`` page to you, when you press 'q', the page will disappear.

The man pages are split up into different sections based on their types.
For example if you access the ``bash`` ``man`` page, at the very top you will see "BASH(1)", indicating that the ``bash`` manual is in section 1: general commands.
Depending on what you're trying to access, you may have to include a section number when you run man.
For example ``man printf`` will show you the ``printf`` commands man page.
But if instead you were wanting to view documentation on the C printf function you would type ``man 3 printf`` as section 3 contains documentation on library functions.

The ``man`` page sections are as follows:

- Section 1: General commands
- Section 2: System calls (functions provided by the kernel)
- Section 3: Library functions (functions within program libraries)
- Section 4: Special files (usually found in /dev)
- Section 5: File formats and conventions (eg /etc/passwd)
- Section 6: Games and screensavers
- Section 7: Miscellaneous
- Section 8: System administration commands

To search through the ``man`` pages run either ``man -k`` or ``apropos`` followed by your search term.
This will return a list of man pages who's descriptions match your search term.

The ``info`` command is another way to find out information about the system and its utilities.
Most system administrators are comfortable with the 'basics' of their core command set, but are frequently checking the ``man`` pages to look up odd flags and functionality.

mount
-----
The ``mount`` command is used to mount filesystems.
For example, mounting an ext4 filesystem that resides on the :file:`/dev/sda2` partition could be done as follows: ``mount -t ext4 /dev/sda2 /mnt``
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
This command would mount whatever partition is found that is associated with the :file:`/home` entry in /etc/fstab, and use any options that happen to be present.

Items to you wish to mount don't necessarily have to be a partition on a disk to be mounted either.
Mounting an ISO file, an image of a optical disk, is especially handy: ``mount -o loop -t iso9660 /home/luser/installer.iso /mnt/cdrom``

``mount`` can also operate on currently mounted filesystems.
Of particular use is switching a currently mounted filesystem from read-write to read-only so that a filesystem check can be performed: ``mount -o remount,ro /``
This particular command tells mount to remount the currently mounted ``/`` filesystem as read-only.

There are quite a number of options that can be passed to the ``mount`` command's "-o" switch.
Some are filesystem independent, while others depend entirely on the type of filesystem that's being mounted.
Further documentation on either can be found in the ``man`` pages.

stat
----

.. todo:: stat command

vmstat
------

.. todo:: vmstat command

strace
------

.. todo:: strace command

ulimit
------

.. todo:: ulimit command


Extracting and manipulating data
================================

A very common pattern in unix is to take some data (a text file, a directory listing, the output from a command) and either extract specific data from it, change some of the data, or both.
These tools help you when you do this.

cat
---

``cat`` outputs the contents of a file either to the shell, another file that already exists, or a file that does not yet exist.

Perhaps most frequently, ``cat`` is used to print the contents of a file to the shell.
For example, if file :file:`foo.txt` contains the word 'foo':

.. code-block:: console

  $ cat /tmp/foo.txt
  foo

When ``cat`` is called on multiple files, the output is in the same order as the files.
If we have another file :file:`bar.txt` that contains 'bar' and run:

.. code-block:: console

  $ cat /tmp/foo.txt /home/jdoe/bar.txt
  foo bar

If you want to combine the contents of the two files:

.. code-block:: console

  $ cat /tmp/foo.txt /home/jdoe/bar.txt > /home/jdoe/foobar.txt
  $ cat /home/jdoe/foobar.txt
  foo
  bar

It is important to note that :file:`foobar.txt` did not exist before running this command.
For this particular usage, ``cat`` can create a file "on the fly".

``cat`` can also be used to output the contents of one file to another file.

.. WARNING:: You should be careful when using ``cat`` this way since it will overwrite the contents of the receiving file.

  .. code-block:: console

    $ cat /tmp/foo.txt > /home/jdoe/bar.txt
    $ cat /home/jdoe/bar.txt
    foo

There are many tools that you may use to parse the output of files, and since ``cat`` can provide a comfortable input method for other tools, it is not always necessary.
Read more on `Useless Use of cat <http://en.wikipedia.org/wiki/Cat_(Unix)#Useless_use_of_cat>`_.

cut
---

The ``cut`` utility cuts out selected portions of each line and writes them to the standard output.

As an example, let's take a file ``students.txt`` that stores a list of student names, ages and email addresses in columns separated by a tab:

.. code-block:: console

  $ cat students.txt
  John	Doe	25	john@example.com
  Jack	Smith	26	jack@example.com
  Jane	Doe	24	jane@example.com

Here, you can see that the first two columns contain the student's name, the third has an age and the fourth, an email address.
You can use ``cut`` to extract just the student's first name and email like this:

.. code-block:: console

  $ cut -f1,4 students.txt
  John john@example.com
  Jack jack@example.com
  Jane jane@example.com

The flag, ``-f`` is used to select which fields we want to output.

``cut``, by default, uses tab as a delimiter, but we can change that by using the ``-d`` flag.

Suppose the ``students.txt`` instead stores data like this:

.. code-block:: console

  $ cat students.txt
  John Doe|25|john@example.com
  Jack Smith|26|jack@example.com
  Jane Doe|24|jane@example.com

Now, if the ``|`` character is used as a delimiter, the first column would be the student's full name:

.. code-block:: console

  $ cut -f1 -d| students.txt
  John Doe
  Jack Smith
  Jane Doe

If you want to use the space to delimit strings, you would do it like so:

.. code-block:: console

  $ cut -f1 -d' ' students.txt
  John
  Jack
  Jane

``cut`` also has some other options. If you have some input with fixed width columns, you can use ``-c`` to break them apart.
For example, to show the login names and times of the currently logged in users:

.. code-block:: console

  $ who | cut -c 1-9,19-30
  mike     Aug  1 23:42
  mike     Aug  5 20:58
  mike     Aug 22 10:34
  mike     Aug  6 19:18

You might have to change some of the character positions to make it work on your system.

grep
----

``grep`` matches patterns in files. 
Its name comes from the ``ed`` command g/re/p (globally search a regular expression and print the results).
``grep`` looks for the given pattern in the specified file or files and prints all lines which match. 

Grep is an excellent tool for when you know the approximate location of the information you want and can describe its structure using a regular expression.  

As you can learn from grep's man page, it takes some options, a pattern, and a file or list of files. 
The files can be specified by ordinary shell globbing_, such as using ``*.log`` to refer to all files in the current directory whose names end in .log. 

Grep's options allow you to customize what type of regular expressions you're using, modify which matches are printed (such as ignoring capitalization or printing only the lines which don't match), and control what's printed as output.
`This post`_ explains some of the optimizations which allow GNU grep to search through large files so quickly, if you're interested in implementation details. 

.. _globbing: http://tldp.org/LDP/abs/html/globbingref.html
.. _`This post`: http://lists.freebsd.org/pipermail/freebsd-current/2010-August/019310.html

**Intro to Regular Expressions**

If you're looking for a specific string, either a word or part of a word, the pattern is just that string.
For example, let's say that I vaguely recall someone telling me about OpsSchool on IRC, and I'd like to find the article that they linked me to.
I think that it was mentioned in a channel called #cschat:

.. code-block:: console

    user@opsschool ~$ grep opsschool \#cschat.log 
     23:02 < nibalizer> http://www.opsschool.org/en/latest/

That's the only place that 'opsschool' was mentioned.
Since grep is case-sensitive by default, 'OpsSchool' would not have shown up in that search.
To ignore case, use the -i flag: 

.. code-block:: console

    user@opsschool ~$ grep -i opsschool \#cschat.log 
     23:02 < nibalizer> http://www.opsschool.org/en/latest/
     15:33 < edunham> hmm, I wonder what I should use as an example in the OpsSchool writeup on grep...

That's better.
But what if I can't remember whether there was a space in 'ops school'? 
I could grep twice, once with the space and once without, but that starts to get ugly very fast.
The correct solution is to describe the pattern I'm trying to match using a regular expression. 

There are a variety of regex tutorials available.
The most important thing to remember about regular expressions is that some characters are special, and not taken literally. 
The special characters are:

==========      =======
Characters      Meaning
==========      =======
``$``           End of a line
``^``           Start of a line
``[]``          Character class
``?``           The preceding item is optional and matched at most once
``*``           Zero or more of the preceding item
``+``           One or more of the preceding item
``{}``          Match some number of the preceding item
``|``           Alternation (true if either of the regexes it separates match)
``.``           Any one character
``\``           Escape (take the following character literally)
==========      =======

Note that there's almost always more than one way to express a particular pattern. 
When you're developing a regular expression, it can be helpful to test it on simplified input to see what it catches.
To test a regex for various spellings of opsschool, you might put a variety of spellings in a file and then grep it to see which regex catches which spellings.

.. code-block:: console

    user@opsschool ~$ cat ops.txt 
     OpsSchool
     Ops School
     opsschool
     ops school
     ops School
     ops  School
     Ops   school
    user@opsschool ~$ grep -i "ops *school" ops.txt
    user@opsschool ~$ grep "[oO]ps[ ?][sS]chool" ops.txt

Try it yourself.
Part of developing a regular expression is clarifying your ideas about exactly what pattern you're looking for.
Will you get any false positives?
For example, the first example above will return true if there are 2 or more spaces between the words of ops school.
It's up to you to decide whether that behavior is correct.
If your regular expression catches strings that it shouldn't, be sure to include some of those possible false positives when testing.

Think about which of the regular expressions above you'd prefer.
Which is easier to read?
Which is better at matching your idea of the correct output?

For more information about regular expressions, try ``man 7 regex``, `regularexpressions.info`_, and the `Advanced Bash Scripting Guide`_ chapter on the subject.

.. _`regularexpressions.info`: http://www.regular-expressions.info/tutorial.html
.. _`Advanced Bash Scripting Guide`: http://tldp.org/LDP/abs/html/x17046.html

**Single vs. Double quotes in the shell**

When grepping for bash variables in scripts, you'll probably want the name of the variable. 
However, there might be times when you want its value.
Below is a quick exercise to explore the difference:

.. code-block:: console

    user@opsschool ~$ echo "$HOME has my username in it" >> home.txt
    user@opsschool ~$ echo '$HOME has a dollar sign in it' >> home.txt
    user@opsschool ~$ cat home.txt
     /home/username has my username in it
     $HOME has a dollar sign in it
    user@opsschool ~$ grep $HOME home.txt
     /home/username has my username in it
    user@opsschool ~$ grep "$HOME" home.txt
     /home/username has my username in it
    user@opsschool ~$ grep '$HOME' home.txt
     $HOME has a dollar sign in it
    user@opsschool ~$ grep \$HOME home.txt
     $HOME has a dollar sign in it                                      

awk
---

``awk`` is a very powerful utility that lets you extract and manipulate data from files.

For example, if you had a file ``students.txt`` similar to the one above, but with the fields separated by a space:

.. code-block:: console

  $ cat students.txt
  John Doe 25 john@example.com
  Jack Smith 26 jack@example.com
  Jane Doe 24 jane@example.com

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

.. NOTE:: The ``-F`` option was not used here since rows are being manipulated, and the ``-F`` option specifies a delimiter for column manipulation


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

In addition to the ``END`` pattern, ``awk`` also provides a ``BEGIN`` pattern, which describes an action that needs to be taken before the first line of the file is processed.

For example:

.. code-block:: console

  $ awk 'BEGIN {print "This is the second line of the file"} NR==2' students.txt
  This is the second line of the file
  Jack Smith - 26 - jack@example.com

sed
---
.. todo:: Only talk about replacing text for now? It's the most common / needed piece of sed at this level.

sort
----
``sort`` can be used to sort lines of text.

For example, if you had a file ``coffee.txt`` that listed different types of coffee drinks:

.. code-block:: console

  $ cat coffee.txt
  Mocha
  Cappuccino
  Espresso
  Americano

Running ``sort`` would sort these in alphabetical order:

.. code-block:: console

  $ sort coffee.txt
  Americano
  Cappuccino
  Espresso
  Mocha

You can also reverse the order by passing in ``-r`` to the command:

.. code-block:: console

  $ sort -r coffee.txt
  Mocha
  Espresso
  Cappuccino
  Americano

All very easy so far.
But, say we have another file ``orders.txt`` that is a list of how many of each drink has been bought in a day:

.. code-block:: console

  $ cat orders.txt
  100 Mocha
  25 Cappuccino
  63 Espresso
  1003 Americano

What happens when we run ``sort`` on this file?

.. code-block:: console

  $ sort orders.txt
  1003 Americano
  100 Mocha
  25 Cappuccino
  63 Espresso

This isn't what we want at all.
Luckily, ``sort`` has some more flags, ``-n`` is what we want here:

.. code-block:: console

  $ sort -n orders.txt
  25 Cappuccino
  63 Espresso
  100 Mocha
  1003 Americano

What if we want to sort the new list by name? We will have to sort by the second column, not the first one. ``sort`` assumes that columns are space separated by default. ``sort`` has the flag ``-k`` that let us specify what key we want to use.

.. code-block:: console

  $ sort -k2 orders.txt
  1003 Americano
  25 Cappuccino
  63 Espresso
  100 Mocha

There are many more flags available, ``man sort`` will show you them all.
There is probably already something there for whatever you can throw at it.
