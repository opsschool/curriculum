#######
Crontab
#######

Laziness is considered a virtue in system administration, meaning you should
figure out ways to avoid having to perform the same task regularly by hand.
**Cron** to the rescue!

.. epigraph::

    **cron** is the time-based job scheduler in Unix-like computer operating
    systems. **cron** enables users to schedule jobs (commands or shell scripts)
    to run periodically at certain times or dates. It is commonly used to
    automate system maintenance or administration, though its general-purpose
    nature means that it can be used for other purposes, such as connecting
    to the Internet and downloading email.

    -- `cron - From Wikipedia, the free encyclopedia <http://en.wikipedia.org/wiki/Cron>`_

**cron** allows you to run routine jobs on a Unix-based system automatically at
a specific future time or at regular intervals rather than running such jobs
manually. Even knowing just the basics of cron is a huge win, as this tool
increases productivity, saves time and is not prone to human error, i.e.
forgetting to run a job.

What is a "crontab"?
====================

"crontabs" (or *cron tables*) are the way that a user schedules a job to run
under **cron**. They're composed of the following parts:

- *cron expression*: A string that describes the time or times when the command
  should be run
- *user* (optional): If the crontab is specified in a centralized location
  instead of the user's personal crontab file, this field contains the username
  that the command should be run as. This allows **cron** to run a command as
  you, even if you're not logged in at the time. This field isn't used in a
  user's personal crontab file.
- *command*: The rest of the line is interpreted as the command that **cron**
  should run at the specified time.

"crontab" may also refer to a file containing any number of these cron
tables, one per line.

crontab files are commonly found in a couple of different places:

- ``/etc/cron.d``, ``/etc/cron.daily``, ``/etc/cron.hourly``, and so on: System
  crontabs, or crontabs installed by a software package
- ``/var/spool/cron/<username>``: Users' crontabs; these are created
  and managed with the ``crontab`` command.

The ``crontab`` command
=======================

How do I list jobs scheduled to run with **cron**?
--------------------------------------------------

The easiest way to see jobs running regularly on the system (aka "cron jobs")
is to type:

.. code-block:: console

  crontab -l

in the shell. This will show all the cron jobs scheduled to run as the
current user. For example, if you are logged in as ``jdoe`` and there
are no cron jobs running the output will be something like:

.. code-block:: console

  -bash: crontab: no crontab for jdoe

If there are jobs scheduled to run, there will be a list of lines that looks
something like this:

.. code-block:: console

  05 12 * * 0 /home/jdoe /home/jdoe/jobs/copy-to-partition

How do I schedule a job with **cron**?
--------------------------------------

The way to schedule a new job is to type:

.. code-block:: console

  crontab -e

in the shell. This will open an editor with your personal "crontab" file (or
an empty file if you don't have any cron jobs already scheduled).

Then add a line to the file that looks like this:

.. code-block:: console

  0 * * * * date

and save and close the file.

The cron job you just created will run the command ``date`` at the top of
every hour (e.g. at 10:00, 11:00, 12:00, etc.). Depending on how your system
is configured, it may begin to email you the output of that command each time
it's run, or it may be saved to a log file somewhere.

A more in-depth example
-----------------------

Let's look again at the example from before:

.. code-block:: console

  05 12 * * 0 /home/jdoe /home/jdoe/jobs/copy-to-partition

Let's dissect this a bit, as it will help when you're creating your own cron
jobs. What is this output telling you? It is helpful to know what the fields of
a "crontab" are. Here's a table with the order of fields, and their values:

  ====== ==== ========== ===== ========= ================
  MINUTE HOUR DAYOFMONTH MONTH DAYOFWEEK COMMAND
  0-59   0-23 1-31       1-12  0-6       filepath/command
  ====== ==== ========== ===== ========= ================

.. note:: Order matters, and note that the first element is 0 for the minute, hour,
  and day of the week fields, while the day of the month and month
  fields begin at 1.

Knowing this, we can see that this "crontab" means:

  At 12:05 every Sunday, every month, regardless of the day of the month, run the
  command ``copy-to-partition`` in the ``/home/jdoe/jobs`` directory.

Let's take another example and create a cron job that checks disk space
available every minute, every hour, every day of the month, every month, for
every day of the week, and outputs it to a file named :file:``disk_space.txt``.

.. code-block:: console

  * * * * * df -h > disk_space.txt

This would get us what we wanted (``df -h`` is the unix command for checking free
disk space).

Field values can also be ranges. Let's say you want to edit this job to run the
same command (``df -h``), but instead of running every minute, you only want the
job to run it in the first 5 minutes of every hour, every day of the month,
every month, for every day of the week.

Running ``crontab -e`` again and changing the line to:

.. code-block:: console

  0-5 * * * * df -h > disk_space.txt

will get you what you want.

How do I remove a crontab?
--------------------------

Lastly, if you want to remove the command, again type ``crontab -e``, and then
delete the line with that job in it from the file in your editor.

To remove all cron jobs for the current user, type:

.. code-block:: console

  crontab -r

What are some common "cron expressions"?
========================================

The "cron expression" syntax can be confusing to understand. Here are some
common expressions to get you started.

- ``* * * * *``: every minute
- ``0 * * * *``: every hour, on the hour
- ``0 0 * * *``: every day at midnight server time
- ``0 0 * * 0``: every Sunday at midnight server time
- ``*/10 * * * *``: every ten minutes
- ``0 */4 * * *``: every four hours, on the hour

Advanced "crontab"
==================

How do "cron expressions" work?
-------------------------------

Here is the standard "cron expression" cheat sheet [#]_::

    # .---------------- minute (0 - 59)
    # |   .------------- hour (0 - 23)
    # |   |   .---------- day of month (1 - 31)
    # |   |   |   .------- month (1 - 12) OR jan,feb,mar,apr ...
    # |   |   |   |  .----- day of week (0 - 7) (Sunday=0 or 7)  OR sun,mon,tue,wed,thu,fri,sat
    # |   |   |   |  |
    # *   *   *   *  *  command to be executed

Put this template at the top of your crontab file so it'll be easy to remember
what the fields do.

Notes on composing good "cron expressions"
------------------------------------------

- If you want to run something every N hours, be sure to specify a minute
  expression (the first number) also; otherwise, the command will be run once a
  minute for the entire hour.
- Commands run by **cron** won't have all of the configuration and environment
  variables that come from your shell initialization files (like ``.bashrc`` or
  ``.zshrc`` or such). In particular, make sure to specify the full path to your
  program if it's not in a commonly-used location like ``/usr/bin``.

Modify a specific user's crontab
--------------------------------

The ``crontab`` command can be used to view or modify a specific user's crontab
file, instead of the current user's crontab file. For instance, if you
are logged in as ``jdoe`` and you want to edit ``jsmith``'s crontab (and you
have the permissions to do so), type the following in a shell:

.. code-block:: console

  crontab -e -u jsmith

This option also combines with the other options we looked at before (``-l`` for
listing and ``-r`` for removing a user's crontab file).

Modifying crontab parameters
----------------------------

With some cron implementations [#]_, you can add shell environment variables to
the top of a crontab file that affect all of the commands run by those crontabs.
For example, you could modify the ``PATH``, ``MAILTO``, ``HOME``, or any other
variable.

--------

Footnotes
=========

.. [#] `"Examples" in cron - Wikipedia, a free encyclopedia <http://en.wikipedia.org/wiki/Cron#Examples_2>`_

.. [#] `Where can I set environment variables that crontab will use?, <http://stackoverflow.com/questions/2229825/where-can-i-set-environment-variables-that-crontab-will-use/10657111#10657111>`_
