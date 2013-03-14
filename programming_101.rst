Programming 101
***************

Shell scripting basics
======================

Specifying the interpeter
-------------------------
Shell scripts will typically start with ``#!/bin/sh`` or ``#!/bin/bash``.  This line, `affectionally known by various names including hashbang and shebang <http://en.wikipedia.org/wiki/Shebang_(Unix)>`_, is treated as a directive to run the rest of the script with the given interpreter.
This directive, combined with setting the execute bit via ``chmod`` tells the system that the file is meant to be executed, rather than simply a file containing text.

The same line may be used instead to specify another shell on the system, such as ``ksh``, ``tcsh``, ``zsh``, or another interpreter entirely, such as Perl, Python, or Ruby.

Portability Considerations
--------------------------

If you only target Linux systems, you may never run into much in the way of shell portability concerns, and you can usually
specify the interpreter to run your script as ``#!/bin/bash`` safely. However, many systems other than Linux (and some limited
Linux environments such as recovery modes and embedded systems) will only have the standard Bourne shell installed
as ``/bin/sh``, and if bash is installed at all, it will be in another location.

Bash is backwards-compatible with scripts written for the Bourne shell, but adds new extensions, which will not work
in the original ``/bin/sh`` that's found on BSD, Solaris, and other Unix systems.  Anything that needs bash should call
it specifically, anything which does not should simply refer to ``#!/bin/sh``

Caution must be used in writing scripts to specify the correct interpreter. ``#!/bin/sh`` is virtually guaranteed to be a
:term:`POSIX`-compatible shell - that is, a shell which is compatible with the original Bourne shell - but there is no guarantee
that this shell will bash. The term "bashism" commonly refers to features of bash not contained in the :term:`POSIX` ``/bin/sh``
specifications. The `bash manual <http://www.gnu.org/software/bash/manual/html_node/Bash-POSIX-Mode.html#Bash-POSIX-Mode>`_ contains a list of these functions.

Aside from the ``/bin/bash`` vs ``/bin/sh`` considerations, it's not guaranteed that bash will be located at ``/bin/bash`` - it may
commonly be found at ``/usr/bin/bash``, ``/usr/local/bin/bash``, and many other locations on systems other than Linux - since you
won't know beforehand where it's located, a common trick is to specify the interpreter as ``#!/usr/bin/env bash``, using the ``env``
utility as a way to search the path for bash. This technique is also highly recommended for calling other shells and interpreters,
as this allows them to be found anywhere in the path, rather than at some fixed location which may change between systems.

A less common portability consideration, but once still worth mentioning, is that of limited environments, such as
systems on which ``/usr`` has not yet been mounted, or recovery environments. If a script must run under such conditions,
it is best to aim for strict :term:`POSIX` compatibility, and specify your interpreter as ``#!/bin/sh``, so as not to rely on
utilities which may be on unmounted filesystems.


Variables
---------

user-defined
built-in

Control Statements
------------------

tests / conditionals
loops

functions
---------

arrays
------

style
-----

Redirection
-----------

I/O
---

Pipes
-----

stderr vs. stdout
------------------

/dev/null and /dev/zero
-----------------------

Regular Expressions
===================

Sed & awk
=========

GIGO
====

Validating input
----------------

Validating output
-----------------

Trapping & handling exceptions with grace
-----------------------------------------


