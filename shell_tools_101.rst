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

top
---

df
--
Finding free disk space is critical.

du
--
Finding disk space used is critical.

find
----
Finding specific files is critical.

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
Knowing what exists on your system is critical.

mount
-----

stat
----

vmstat
------

lsof
----

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

