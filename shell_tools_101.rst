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

kill
----

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

Estracting and manipulating data
================================

A very common pattern in unix is to take some data (a text file, a directory
listing, the output from a command) and either extract specific data from it,
change some of the data, or both. These tools help you when you do this.

cat
---

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
