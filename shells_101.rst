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
- Maybe a brief mention of tmux/screen here, or at least a link to such as section as it relates to continuing jobs even if terminal connectivity is lost.

Customizing the Prompt (Or Fun with the PS variables)
-----------------------------------------------------
``$PS1``, ``$PS2``, ``$PS3``, ``$PS4``
