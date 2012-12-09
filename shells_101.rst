Shells
******

What is a shell?
================
A [Li|U]nix shell is the command-line interface between the user and the system.
It is used to perform some action, specifically, typing commands and displaying
output, requested by the user.

Introduction to Bash
====================
**Bash** is known as the Bourne-again shell, written by Brian Fox, and is a play
on the name of the previously named Bourne shell (/bin/sh) written by Steve
Bourne. [#f1]_

Shell fundamentals
==================

Command-line Editing Modes
--------------------------
- vi

  ``set -o vi``

- emacs

  ``set -o emacs``

See :doc:`text_editing_101` for details on appropriate edit commands to use on
the command line.

.. todo: Explain setting the command line editing mode - what does it do, who does it affect?

Environment variables
---------------------
Environment variables are used to define values for often-used attributes of a
user's shell. In total, these variables define the user's environment. Some
environment variables provide a simple value describing some basic attribute,
such the user's current directory (``$PWD``). Others define the behavior of a
command, such as whether or not the ``history`` command should log repeated
commands individually or log the repeated command once (``$HISTCONTROL``).

$PATH
~~~~~
The most common, or most recognized, environment variable is the ``$PATH``
variable. It defines the set of directories that the shell can search to find a
command. Without an explicit path provided when calling a command (i.e. ``/bin/ps``),
the shell will search the directories listed in the ``$PATH`` variable until it
finds the command. If the command is not found in any of the defined directories
in ``$PATH``, the shell will produce an error explaining as much. ::

  $ foobar -V
  -bash: foobar: command not found


To view the contents of the ``$PATH`` variable, use ``echo`` to print the variable's value: ::

  $ echo $PATH
  /usr/local/sbin:/sbin:/bin:/usr/sbin:/usr/bin:/usr/local/bin

The order of the directories in the ``$PATH`` variable, from left to right, is
important; when searching directories for a command, the shell will stop looking
after it finds its first match.
In other words, using our example ``$PATH`` variable above, if there is a
version of ``ps`` that exists in ``/usr/local/bin`` that is preferred (by the sysadmin)
over the version that exists in ``/bin``, the shell will still execute ``/bin/ps``
due to the precedence of the directories defined in the ``$PATH`` variable.

To list all of the shell's environment variables, use the ``env`` command: ::

  $ env
  HOSTNAME=foobar
  SHELL=/bin/bash
  TERM=xterm
  HISTSIZE=1000
  USER=root
  PATH=/usr/local/sbin:/sbin:/bin:/usr/sbin:/usr/bin:/root/bin:/usr/local/bin
  MAIL=/var/spool/mail/root
  PWD=/root/curriculum
  PS1=[\[\e[33;1m\]\t \[\e[31;1m\]\u\[\e[0m\]@\[\e[31;1m\]\h\[\e[0m\] \W\[\e[0m\]]# 
  AWS_IAM_HOME=/opt/aws/apitools/iam
  HISTCONTROL=ignoredups
  SHLVL=1
  SUDO_COMMAND=/bin/bash
  HOME=/root
  HISTTIMEFORMAT=[%Y-%m-%d %H:%M:%S] 
  OLDPWD=/tmp

Global vs. User Profiles
------------------------
- ``/etc/profile``, ``~/.bash_profile``
- Profile precedence

Special environment variables
-----------------------------
``$$``, ``$!``, ``$?``, etc

History
-------
``history``, ``.bash_history``, ``!!``, ``!#`` (i.e. ``!1`` to repeat the first command in the history), ``^R``, etc

Job control
-----------
- ``^Z``, ``bg``, ``fg``, ``%1/2/3..``, ``jobs``

For information on ensuring running jobs continue, even when terminal
connectivity is lost, see the sections on :ref:`gnu-screen` and :ref:`tmux`.

Customizing the Prompt for Fun or Profit
----------------------------------------
``$PS1``, ``$PS2``, ``$PS3``, ``$PS4``

Example (needs explanation)::

  # if I'm root, set my terminal colors to alert me!
  if [ "$EUID" = "0" ]
  then
    PS1="[\[\e[33;1m\]\t \[\e[31;1m\]\u\[\e[0m\]@\[\e[31;1m\]\h\[\e[0m\] \W\[\e[0m\]]\$ "
    export PS1
  else
    PS1="[\t \[\e[34;1m\]\u\[\e[0m\]@\[\e[34;1m\]\h\[\e[0m\] \[\e[33;1m\]\W\[\e[0m\]]\$ "
    export PS1
  fi

.. todo::
   - Link to any content describing profiles (global, user-level) as the above example should be placed in a profile
   - Link to content describing terminal color codes/ANSI escape codes
   - Determine if it's important to discuss such an esoteric topic as terminal color/escape codes or if I'm really just showing off...

.. rubric:: Footnotes

.. [#f1] `C Programming by Al Stevens <http://www.drdobbs.com/i-almost-get-a-linux-editor-and-compiler/184404693>`_, Dr. Dobb's Journal, July 1, 2001
