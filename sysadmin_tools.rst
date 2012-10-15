Tools for productivity
*******************************

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


GNU Screen
----------
``screen`` is one of the longest lived multiplexers. Almost everyone who has
used a multiplexer has used screen, and you can't go far wrong with it.

.. todo::
   Explain how to use ``screen``


``tmux``
--------
``tmux`` is relatively new compared to ``screen``. It covers the same basic
feature set and has added a few more advanced features. It is recommended you
get comfortable with ``screen`` first before attempting to use ``tmux``. 

.. todo::
   Explain how to use ``tmux``
