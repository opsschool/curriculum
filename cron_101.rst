#######
Crontab
#######

Laziness is considered a virtue in system administration, meaning you should 
figure out ways to avoid having to perform the same task regularly. Crontab 
allows you to run routine jobs on a user's system automatically at regular 
intervals rather than running such jobs manually. Even knowing just the basics 
of crontab is a huge win, as this tool increases productivity, saves time and is
not prone to human error, i.e. forgetting to run a job.

The ``crontab`` command
=======================

The easiest way to see jobs running regularly on the current user's system
(aka "cron jobs") is to type:

.. code-block:: bash

  crontab -l

in the shell. This will show all the cron jobs that currently run on the current
user's system. For example, if you are logged in as ``jdoe``, the current user
is ``jdoe``, and if there are no cron jobs running, the output will be something
like:

.. code-block:: console

  -bash: crontab: no crontab for jdoe

If there are jobs scheduled to run, there will be a list of lines that looks
something like this:

.. code-block:: console

  05 12 * * 0 /home/jdoe /home/jdoe/jobs/copy-to-partition

Let's dissect this a bit, as it will help when you're creating your own cron
jobs. What is this output telling you? It is helpful to know that the fields of
a cron job are as follows:

  **MINUTE HOUR DAYOFMONTH MONTH DAYOFWEEK COMMAND**

and that the acceptable values for each field are:

  **0-59 0-23 1-31 1-12 0-6 filepath/command**

.. note:: Order matters, and that the first element is 0 for the minute, hour,
  and day of the week fields, while the day of the month and month
  field begin at 1.

Knowing this, we can see that the output of ``crontab -l`` from before means:

  At 12:05 every Sunday, every month, regardless of the day of the month, run the
  command ``copy-to-partition`` in the :file:`/home/jdoe/jobs` directory.

Let's take another example and create a cron job that checks disk space
available every minute, every hour, every day of the month, every month, for
every day of the week, and outputs it to a file named :file:`disk_space.txt`.

.. code-block:: bash

  * * * * * df -h > disk_space.txt

would get us what we wanted (``df -h`` is the unix command for checking free
disk space).

Field values can also be ranges. Let's say you want to edit this job to run the
same command (``df -h``), but instead of running every minute, you only want the
job to run it in the first 5 minutes of every hour, every day of the month,
every month, for every day of the week.

.. code-block:: bash

  crontab -e

will open up your default shell editor, where you will see a list of your cron
jobs. Editing the one we just wrote to:

.. code-block:: console

  0-5 * * * * df -h > disk_space.txt

will get you what you want.

Lastly, if you want to remove the command, again type ``crontab -e``, and then
delete the line with that job in it from the file in your editor.

To remove all cron jobs, type:

.. code-block:: bash

  crontab -r

