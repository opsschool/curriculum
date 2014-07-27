Programming 101
***************

Shell scripting basics
======================

Specifying the interpreter
--------------------------
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

There are two kinds of variables you can use in your scripts: user-defined and built-in.
Quite obviously, user-defined variables are created by the user (you) and built-in variables are automatically set by the shell you use.

* User-defined variables

To create a variable, all you need to do is assign it a value.

.. code-block:: console

  $ ops_var=1

NOTE: When using a BASH shell or writing a BASH script, make sure that you don't leave spaces in the expression; doing so will lead to an error.
Always make sure you don't have spaces in variable assignment statements in BASH.

To echo out the value of this variable, use ``$ops_var`` like so:

.. code-block:: console

  $ echo $ops_var
  1

The same variable can be assigned different values in the same script.
Creating a script called ``ops_script.sh``

.. code-block:: bash

  #!/usr/bin/env bash

  # First an integer value
  ops_var=1
  echo $ops_var

  # Then a string
  ops_var="Hello"
  echo $ops_var

When this script is executed:

.. code-block:: console

  $ ./ops_script.sh
  1
  Hello

NOTE: From this point on, assume that the same ``ops_script.sh`` is going to be used.
I also won't be typing ``#!/usr/bin/env bash`` every time, but know that it's present at the top of the script.

Variables can be used in other strings by calling them with curly braces ``{ }`` around the variable name.

.. code-block:: bash

  ops_var="Yoda"
  echo "${ops_var}, my name is"

.. code-block:: console

  $ ./ops_script.sh
  Yoda, my name is

You can use variables to store user input and use it later on in the script.
For example, if you want to ask the user for their name and echo it back to the screen:

.. code-block:: bash

  print "Hello, what is your name?: "
  read name

  echo "Hello, ${name}"

I'm going to supply the name "Yoda" in this case.

.. code-block:: console

  $ ./ops_script.sh
  Hello, what is your name?: Yoda
  Hello, Yoda

By default, all variable values are treated as strings, unless the operation they are used for explicitly uses them as another data type.
Assume you have two variables that you assign integer values to, and you want to add them together.

.. code-block:: bash

  first_var=1
  second_var=2
  result=${first_var}+${second_var}

  echo ${result}

You would expect the output of this to be 3, but this is not the case with bash.

.. code-block:: console

  $ ./ops_script.sh
  1+2

What happened here was that both values were treated as string and the expression ``${first_var}+${second_var}`` got evaluated as the string "1+2".
To actually add these two numbers, the operation needed is a little different.

.. code-block:: bash

  first_var=1
  second_var=2
  result=$(( ${first_var} + ${second_var}))

  echo ${result}

Here, the ``$(( ))`` tells bash that anything that goes inside these double parentheses needs to be evaluated as an arithmetic operation.

.. code-block:: console

  $ ./ops_script.sh
  3

* Built-in variables

The second kind of variables that you can use in your scripts are the ones that the shell automatically populates for you when it is started up.
These variables store information specific to your shell instance, so if you and a co-worker are both logged on to the same server, both of you might not see the same values for all built-in variables.

Here are some examples of built-in variables:

.. code-block:: bash

  # Home directory of the logged-on user.
  $HOME

  # Name of the computer you are currently logged on to
  $HOSTNAME

  # Present working directory
  $PWD


The above variables contain some information relative to the shell itself.
But there are other built-in variables you can use within your script which store information like exit statuses of scripts/commands and number of arguments passed to a script.

.. code-block:: bash

  # Exit status of the previously executed script/command
  $?

  # Number of arguments passed to the script
  $#

  # Value of the 1st argument passed to the script
  ${1}

These variables let your script take in parameters which can be then used throughout the script.

For example, I will write a script that prints the number of parameters received, and use the first one in a string

.. code-block:: bash

  # Print the number of arguments passed
  echo "Number of arguments: $#"

  # Use the first argument in a string
  echo "First argument passed was: ${1}"

I'll run this script a couple of times with different arguments to show how this works

.. code-block:: console

  $ ./ops_script.sh hello world
  Number of arguments: 2
  First argument passed was: hello

  $ ./ops_script.sh car truck bike scooter
  Number of arguments: 4
  First argument passed was: car


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
