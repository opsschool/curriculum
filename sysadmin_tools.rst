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

``tmux``
--------
``tmux`` is relatively new compared to ``screen``. It covers the same basic
feature set and has added a few more advanced features. It is recommended you
get comfortable with ``screen`` first before attempting to use ``tmux``. 

In this chapter we will learn to start a tmux session, get to know a
few first keyboard shortcuts and detach from and re-attach to the
session.

``tmux`` is usually started with the command ``tmux`` in a
terminal window. Depending of your version of tmux you will see either
a line at the bottom of the screen or nothing at all. ``tmux`` is
controlled with keyboard shortcuts, the default shortcut usually is
``Ctrl-a``. If you press ``ctrl-a`` and then a ``t``in the newly
started tmux window you should see the local time displayed as a large
digital clock. If you hit ``ctrl-a`` and ``c`` you should see a new
empty window with an empty input prompt.

If you want to  detach from the session you have to hit ``ctrl-a`` and ``d``. The ``tmux`` window will disappear ynd you will see a message ``[detached]`` in your terminal window. All the shells and processes you started onside the ``tmux``session continue to run, you can see this with a simple 

.. epigraph::
   ``ps -ef | grep tmux``

You should see something like the following:

.. epigraph::
   ``cdrexler 13751     1  0 Nov30 ?        00:00:41 tmux``

You will notice that the ``tmux process has a parent process id of 1
which means that it is not a child process of the shell you started it
in anymore. Accordingly you can leave your working shell, start a new
one and attach to the running tmux process again.

If you want to get an overview of the running tmux processes on your system you can use the command

.. epigraph::
   ``tmux ls``

It will list all available ``tmux`` sessions on your system. If there is only one you can attach to it with the command

.. epigraph::
   ``tmux att``

You should see the original session with the two shells again after running this command.

.. _meta_multiplexers:

byobu
-----
.. todo::
   describe advantages of meta-multiplexers like ``byobu`` that can use different backends.


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
