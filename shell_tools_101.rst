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
---
Finding free disk space is critical.

du
---
Finding disk space used is critical.

find
---
Finding specific files is critical.

kill
----

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
---
Definitely need to discuss sort and how to use it in basic activities.
