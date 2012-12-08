Tools for productivity
**********************

The work of operations is often helped by using the combinations of tools.
The tools you choose to use will depend heavily on your work style. Over time
you may even change the tools you use.

This is a list of commonly used tools, we recommend trying as many of them as
you have time for. Finding the right tool is often like finding a good pair of
gloves - they have to fit just right!


Terminal emulators
==================

OSX: iTerm2
-----------


SSH
===

Common: ``ssh``
---------------
If your operating system has a good terminal emulator available (eg, Linux, BSD,
OSX), then the ``ssh`` command line utility is by far the best tool to use for
SSH interactions.

Windows: PuTTY
--------------
If you use Windows as your primary work space, PuTTY is possibly one of the best
SSH clients available.

Android: Connectbot
-------------------
Connectbot is a free app available in the Google Play store, and offers most of
the functionality you would need from an SSH client.

iOS: iSSH
---------


Multiplexers
============
Operations work regularly involves connecting to remote servers (there will be
times when you do nothing but work on remote servers - parts of this curriculum
were even typed on remote servers rather than contributors desktops and
laptops!).

There are however two limitations to working this way:

# You'll often need to be connected to more than one remote system at a time.
Opening a whole new terminal each time can result in a lot of windows cluttering
up previous screen space.
# When happens if your internet connection stops working? All of your
connections are reset. Any work you might have been doing on the remote servers
can be lost.

Multiplexers are a good solution to this.
They allow you to run multiple "virtual" windows inside a single windows.
For example:

.. epigraph::
   Bob works on 10 remote servers, all of which run Linux.
   Bob's internet connection at work is questionable.
   To work around this, Bob connects to ``server1`` which is at his data centre.
   It is a reliable server which is close to the other servers Bob works on.
   On ``server1``, but starts a multiplexer. The multiplexer gives Bob a regular
   looking command prompt, and Bob continues his work.
   
   If Bob's internet connection drops, he can reconnect to ``server1``, and then
   re-attach to the multiplexer he started previously. His session is in the
   same state he left it before being disconnected, and he can continue his
   work.

   The multiplexer also lets Bob open more than one command prompt and switch
   between them as he needs to. Bob can now connect to many servers and see them
   all in one window.

.. _gnu-screen:

GNU Screen
----------
``screen`` is one of the longest lived multiplexers. Almost everyone who has
used a multiplexer has used screen, and you can't go far wrong with it.

.. todo::
   Explain how to use ``screen``

.. _tmux:

Tmux
----
``tmux`` [#]_  is relatively new compared to
``screen``. It covers the same basic feature set and has added a few
more advanced features. It is recommended you get comfortable with
``screen`` first before attempting to use ``tmux``.

In this chapter you will learn to start a tmux session, get to know a
few first keyboard shortcuts and detach from and re-attach to the
session.

Installation
~~~~~~~~~~~~
tmux is available on Debian and its descendants like Ubuntu or Mint
with the command:

.. code-block:: bash

  aptitude install tmux

On RedHat-style distributions you will have to use the :term:`EPEL` repo to
get a pre-built package, and install with the command:

.. code-block:: bash

  yum install tmux

On MacOS you can use Homebrew to install via:

.. code-block:: bash

  brew install tmux

tmux basics
~~~~~~~~~~~
``tmux`` is usually started with the command ``tmux`` in a
terminal window. Depending of your version of tmux you will see either
a line at the bottom of the screen or nothing at all. ``tmux`` is
controlled with keyboard shortcuts, the default shortcut usually is
``ctrl-b``. If you press ``ctrl-b`` and then a ``t`` in the newly
started tmux window you should see the local time displayed as a large
digital clock. If you hit ``ctrl-b`` and ``c`` you should see a new
empty window with an empty input prompt.

If you want to detach from the session you have to hit ``ctrl-b`` and
``d``. The ``tmux`` window will disappear and you will see a message
``[detached]`` in your terminal window. All the shells and processes
you started onside the ``tmux`` session continue to run, you can see
this with a simple

.. code-block:: bash

  ps -ef | grep tmux

You should see something like the following:

.. code-block:: bash

  cdrexler 13751     1  0 Nov30 ?        00:00:41 tmux

You will notice that the ``tmux`` process has a parent process id of 1
which means that it is not a child process of the shell you started it
in anymore. Accordingly you can leave your working shell, start a new
one and attach to the running tmux process again which is very handy
if your connectivity is flaky or you have to work from different
locations. If you check the process table for the process id of the
tmux process

.. code-block:: bash

  ps -ef | grep 13751

you will find that is the parent process of the two shells you created
in the beginning of the chapter:

.. code-block:: bash

   cdrexler  4525 13751  0 17:54 pts/2    00:00:00 -zsh
   cdrexler  4533 13751  0 17:54 pts/5    00:00:00 -zsh

If you want to get an overview of the running tmux processes on your
system you can use the command

.. code-block:: bash

  tmux ls

It will list all available ``tmux`` sessions on your system [#]_. If there
is only one you can attach to it with the command:

.. code-block:: bash

  tmux att

If there is more than one session the output of ``tmux ls`` will look like this:

.. code-block:: bash

   0: 3 windows (created Fri Nov 30 18:32:37 2012) [80x38]
   4: 1 windows (created Sun Dec  2 17:44:15 2012) [150x39] (attached) 

You will then have to select the right session with the ``-t`` command line switch:

.. code-block:: bash

  tmux att -t 4

``tmux`` runs as a server process that can handle several sessions so
you should only see one tmux process per user per system.

You should see the original session with the two shells again after
running this command.

tmux configuration 
~~~~~~~~~~~~~~~~~~~
 
``tmux`` is configured via a
config file which is usually called ``.tmux.conf`` that should live in
your ``$HOME`` directory.

A typical ``.tmux.conf`` looks like this:

.. code-block:: ini

   #set keyboard shortcut to ctrl-g
   unbind C-b
   set -g prefix C-g
   bind C-g send-prefix
   bind g send-prefix
   #end of keybord shortcut setting
   # Highlight active window
   set-window-option -g window-status-current-bg red
   # Set window notifications
   setw -g monitor-activity on
   set -g visual-activity on
   #automatically rename windows according to the running program
   setw -g automatic-rename
   #set scroll back buffer
   set -g history-limit 10000
   set -g default-terminal "xterm-256color"
   set -g base-index 1
   set -g status-left ‘#[fg=green]#H

This illustrates a method to change the default keybinding and some
useful settings.

Please note that you can force ``tmux`` to use another configfile with
the ``-f`` command line switch like so:

.. code-block:: bash

  tmux -f mytmuxconf.conf

There is a nifty cheat sheet [#]_ for the most important
``screen`` and ``tmux`` keybindings or even a whole book about tmux [#]_.



byobu
-----
.. todo::

   - describe advantages of meta-multiplexers like ``byobu`` [#]_ that can use different backends.
   - describe scrollback and copy and paste

References
----------
.. [#] http://tmux.sourceforge.net/
.. [#] Please note that ``tmux ls`` will *only* list tmux sessions that belong to your userid!
.. [#] http://www.dayid.org/os/notes/tm.html
.. [#] http://pragprog.com/book/bhtmux/tmux
.. [#] https://launchpad.net/byobu


Shell customisations
====================

As you read in :doc:`shells_101`, your shell is your primary tool during the
work day. It's also incredibly customisable to suit your needs. Let's look at
some changes you can make.

How to customise your shell
---------------------------

Your shell's configuration is stored in its ``rc`` file. For bash, this file is
``~/.bashrc``. Each time you edit this, you can reload the configuration by
typing:

.. epigraph::
   ``source ~/.bashrc``

Changing your prompt
--------------------

Your default prompt probably looks something like this:

.. epigraph::
   ``bash-3.2$``

That's pretty plain and doesn't tell you much. In fact, all it does tell you is
that you're using Bash version 3.2, and that you are not the root user (the
``$`` at the end signifies a regular user, whereas if you were root, you would
see a ``#`` instead).

Let's change this up a little. Edit your ``~/.bashrc`` file, and add this line
to the end:

.. epigraph::
   ``PS1="\u@\h \w> "``

Save, quit, and then reload your ``.bashrc`` file. Your prompt should change to
something like this:

.. epigraph::
   ``avleen@laptop ~>``

Much better! Now your know your username, the name of the machine you're on (in
this case "``laptop``"), and the directory you're in ("``~``" is your home
directory).

The ``PS1`` variable has a lot of different options you can use to customise it
further.


Mosh
====


Ticketing systems
=================


Note-taking
===========

Wiki
----

EverNote
--------

OneNote
-------
