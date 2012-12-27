Crontab
*******

You've probably heard of "cron jobs" or "crontabs" before. What are they and what are they good for?

What is "cron"?
===============

    **cron** is the time-based job scheduler in Unix-like computer operating systems. **cron** enables users to schedule jobs (commands or shell scripts) to run periodically at certain times or dates. It is commonly used to automate system maintenance or administration, though its general-purpose nature means that it can be used for other purposes, such as connecting to the Internet and downloading email. [#]_

Put more simply, **cron** is a system that lets a user schedule some command to be run at a future time, instead of right now while that user is sitting in front of the terminal. Commands run by **cron** are triggered by the arrival of a time that fulfills their "cron expression"; at that time, **cron** runs the specified command just as though the user had typed it and pressed Enter right then.

What is a "crontab"?
====================

"crontabs" (or *cron tables*) are the way that a user schedules a job to run under **cron**. They're composed of the following parts:

- *cron expression*: A string that describes the time or times when the command should be run
- *user* (optional): If the crontab is specified in a centralized location instead of the user's personal crontab file, this field contains the username that the command should be run as. This allows **cron** to run a command as you, even if you're not logged in at the time. This field isn't used in a user's personal crontab file.
- *command*: The rest of the line is interpreted as the command that **cron** should run at the specified time.

"crontab" may also refer to a file containing any number of these "crontabs", one per line.

How do I schedule a job with **cron**?
======================================

1. Run ``crontab -e`` to open an editor with your personal "crontab" file (or an empty file if you don't have any crontabs already scheduled).

2. Add a line to the file that looks like this::

    0 * * * * date

3. Save and close the file.

The "crontab" you just created will run the command ``date`` at the top of every hour (e.g. at 10:00, 11:00, 12:00, etc.). Depending on how your system is configured, it may begin to email you the output of that command each time it's run, or it may be saved to a log file somewhere.

To stop running this "crontab", run ``crontab -e`` again, delete that line, and save and close the file.

What are some common "cron expressions"?
========================================

The "cron expression" syntax can be confusing to understand. Here are some common expressions to get you started.

- ``* * * * *``: every minute
- ``0 * * * *``: every hour, on the hour
- ``0 0 * * *``: every day at midnight server time
- ``0 0 * * 0``: every Sunday at midnight server time
- ``*/10 * * * *``: every ten minutes
- ``0 */4 * * *``: every four hours, on the hour

--------

Advanced "crontab"
==================


How do "cron expressions" work?
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Here is the standard "cron expression" cheat sheet [#]_::

    # .---------------- minute (0 - 59)
    # |   .------------- hour (0 - 23)
    # |   |   .---------- day of month (1 - 31)
    # |   |   |   .------- month (1 - 12) OR jan,feb,mar,apr ...
    # |   |   |   |  .----- day of week (0 - 7) (Sunday=0 or 7)  OR sun,mon,tue,wed,thu,fri,sat
    # |   |   |   |  |
    # *   *   *   *  *  command to be executed

Put this template at the top of your crontab file so it'll be easy to remember what the fields do.

Notes on composing good "cron expressions"
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

- If you want to run something every N hours, be sure to specify a minute expression (the first number) also; otherwise, the command will be run once a minute for the entire hour.
- Commands run by **cron** won't have all of the configuration and environment variables that come from your shell initialization files (like ``.bashrc`` or ``.zshrc`` or such). In particular, make sure to specify the full path to your program if it's not in a commonly-used location like ``/usr/bin``.

``crontab`` usage
~~~~~~~~~~~~~~~~~

The ``crontab`` command can be used in several different ways:

- ``crontab -e``: Edit your crontab file (as described above).
- ``crontab -l``: List your crontab file without spawning an editor.
- ``crontab -d``: Wipe out your crontab file entirely.
- ``crontab -u <username>``: Select a user's crontab file to operate on. This option combines with the previous options; for instance, ``crontab -u joe -e`` will open an editor with Joe's crontab file instead of your own.


Modifying "crontab" parameters
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

With some cron implementations [#]_, you can add shell environment variables to the top of a crontab file that affect all of the commands run by those crontabs. For example, you could modify the ``PATH``, ``MAILTO``, ``HOME``, or any other variable.

--------

Footnotes
=========

.. [#] `cron - From Wikipedia, the free encyclopedia <http://en.wikipedia.org/wiki/Cron>`_

.. [#] `"Examples" in cron - Wikipedia, a free encyclopedia <http://en.wikipedia.org/wiki/Cron#Examples_2>`_

.. [#] `Where can I set environment variables that crontab will use?, <http://stackoverflow.com/questions/2229825/where-can-i-set-environment-variables-that-crontab-will-use/10657111#10657111>`_
