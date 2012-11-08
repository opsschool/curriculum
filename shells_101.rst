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
- emacs

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
