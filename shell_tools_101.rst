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

.. code-block:: bash

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

.. code-block:: bash

  $ kill -KILL 123

Every signal has a name and a number. You can reference them by either
one. Another way of running ``kill -KILL`` is:

.. code-block:: bash

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
Only talk about column extraction for now? It's the most common / needed piece
of awk at this level.

sed
---
Only talk about replacing text for now? It's the most common / needed piece of
sed at this level.

sort
----
Definitely need to discuss sort and how to use it in basic activities.

