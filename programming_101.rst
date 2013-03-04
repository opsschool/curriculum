Programming 101
***************

Shell scripting basics
======================

Specifying the interpeter
-------------------------
Shell scripts will start with ``#!/bin/sh`` or ``#!/bin/bash``.  This line, affectionally by various names including hashbang and shebang, 
is treated as a directive to run the rest of the script with the given interpreter. This directive, combined with setting
the execute bit via ``chmod`` tells the system that the file is an executable, rather than simply a file containing text.

The same line may be used instead to specify another shell on the system, such as ``ksh``, ``tcsh``, ``zsh``, or another interpreter entirely, such as Perl, Python,
or Ruby.

Portability Considerations
--------------------------

If you only target Linux systems, you may never run into much in the way of shell portability concerns, and you can usually
specify the intrepeter to run your script as ``#!/bin/bash`` safely. However, many systems other than Linux (and some limited 
Linux enviroments such as recovery modes and embedded systems) will only have the standard Bourne shell installed 
as ``/bin/sh``, and if bash is installed at all, it will be in another location.

Bash is backwards-compatabile with scripts written for the Bourne shell, but adds new extensions, which will not work
in the original /bin/sh that's found on BSD, Solaris, and other Unix systems.  Anything that needs bash should call
it specificly, anything which does not should simply refer to ``#!/bin/sh``

Caution must be used in writing scripts to specify the correct interpeter. ``#!/bin/sh`` is virtually guranteed to be a
POSIX-compatable shell - that is, a shell which is compatable with the original bourne shell - but there is no gurantee
that this shell will bash. The term "bashism" commonly refers to features of bash not contained in the POSIX ``/bin/sh``
specifications. The `bash manual <http://www.gnu.org/software/bash/manual/html_node/Bash-POSIX-Mode.html#Bash-POSIX-Mode>`_ contains a list of these functions.

Aside from the ``/bin/bash`` vs ``/bin/sh`` considerations, it's not guranteed that bash will be located at ``/bin/bash`` - it may
commonly be found at ``/usr/bin/bash``, ``/usr/local/bin/bash``, and many other locations on systems other than Linux - since you
won't know beforehand where it's located, a common trick is to specify the intrepeter as ``#!/usr/bin/env bash``, using the env
utility as a way to search the path for bash. This technique is also highly recommended for calling other shells and interpreters,
as this allows them to be found anywhere in the path, rather than at some fixed locatation which may change between systems.

A less common portability consideration, but once still worth mentioning, is that of limited enviroments, such as
systems on which /usr has not yet been mounted, or recovery enviroments. If a script must run under such conditions,
it is best to aim for strict POSIX compatability, and specify your intrepeter as ``#!/bin/sh``, so as not to rely on
utilities which may be on unmounted filesystems,


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


