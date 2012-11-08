Shells
******

What is a shell?
================

Introduction to Bash
====================

Shell fundamentals
==================

Command-line Editing Modes
--------------------------
- vi
 + set -o vi
- emacs
 + set -o emacs

See `Text Editing 101`_ for details on appropriate edit commands to use on the command line.

.. _`Text Editing 101`: /text_editing_101.html

Environment variables
---------------------
``$PATH``, ``$HOME``, ``$USER``, etc

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

For information on ensuring running jobs continue, even when terminal connectivity is lost, see the sections on `GNU screen`_ and tmux_.

.. _`GNU screen`: /sysadmin_tools.html#gnu-screen
.. _tmux: /sysadmin_tools.html#tmux

Customizing the Prompt (Or Fun with the PS variables)
-----------------------------------------------------
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
