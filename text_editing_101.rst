Text Editing 101
****************

As a sysadmin, and *especially* as a Linux/Unix sysadmin, you're going to have to get used to
not only making your way around the command line, but also editing text files.
Both of these tasks require you to get familiarized and comfortable with the two most common
text editors in this space: ``vi`` and ``emacs``

Both of these editors, ``vi`` and ``emacs``, are command-line utilities that you run from a
terminal window.  This means that you can also run them on remote machines over an ``ssh`` connection.
While there are graphical versions of these tools available (``gvim`` and ``xemacs``), many
administrators prefer the text-only versions, as it creates a uniform experience whether they
are working on their own local machine, or a remote one being accessed over the network.
Additionally, avoiding use of the mouse means your hands stay on the home row of the keyboard,
meaning you work faster and with less fatigue.

A little history
================

Back when UNIX was just starting to appear on computers that had actual input/output systems
(e.g. teletype), interaction with the computer was strictly a text-based affair.  The computer
would send some data in the form of `ASCII <http://en.wikipedia.org/wiki/ASCII/>`_\* characters over a slow
telephone link.  The user would then hit some keys on the keyboard, which would send back more
ASCII characters to the computer, again over the slow telephone link.

Because communication was so slow, UNIX programmers and users wrote their internal tools to
be as efficient as possible.  Early editors, such as ``ed``, would only let you see one
line of text in the file at a time, since the commonly used interface at the time, a teletype,
printed on paper and thus there was no concept of "redrawing" or even of a "screen".  These
early tools later evolved into the more advanced editors like ``vi`` and ``emacs`` in use
today.  But these tools retain their UNIX heritage of lean efficiency.  Commands are sent
usually in single characters, there's no "toolbar", and the editor provides powerful tools
to quickly shuttle the cursor around the file quickly, without having to display unrelated
text/context in the file.

\* Note: you can also run ``man ascii`` to quickly get an ASCII table

``vi`` basics
=============

``vi`` was created as an extension to ``ex`` (a successor to ``ed``).  When
"modern" terminals (e.g. electronic displays, not teletypes) became common, it
made sense to use a "visual" editing mode, where you could see the context of
the line you were editing.  From ``ed``, you would enter "visual" mode by
running ``:vi`` -- thus VI was born.

Since then, ``vi`` has been written and re-written many, many times.  But inside, you
will still find good ol' ``ed``.  In ``Text Editing 201``, you will see how you can
still access the very powerful ``ed`` features by using the colon (``:``) command.

There are many versions of ``vi`` out there, but the most popular is ``vim`` (VI iMproved).
Virtually every Linux distribution out there provides ``vim`` either by default, or
as an easily installable package.  ``vim`` extends the feature set provided by the 
POSIX-compliant ``vi`` utility in many ways.  In the examples below, features that
are only present in ``vim`` will be noted.


A Note About Vimtutor
---------------------

Vim, the modern, extended version of ``vi``, includes a tutorial of its own.  On
a system with Vim installed, you should be able to run ``vimtutor`` to get an interactive
session that walks you through most of the same content you'll get on this page.  You
don't have to use ``vimtutor`` if you don't want to, though -- just know that it is
there if you wish to try it out.

``vi``  Modes
-----------------------

The first thing to understand about how ``vi`` works is that it operates in
different modes.

**Normal mode** is ``vi``'s default mode and should be the mode you are in most
of the time. When in normal mode, the keys you type are interpreted by ``vi``
as commands, not as text you wish to put into the file.  For example, moving
the cursor around the screen is a series of commands, as is deleting text.

**Insert mode** is when you are actually typing in text.  The keys you type
while in insert mode are simply dumped into the buffer (``vi``'s name for the
file you are editing).  You exit insert mode by pressing the ``Esc`` key.

**Command-line mode** is when you are typing a command on the cmd line at the
bottom of the window. This is for the Ex commands, ":", the pattern search
commands, "?" and "/", and the filter command, "!".

Hello, world!
-------------

Before learning anything else, you'll need to know how to do two basic operations,
which are best described by working through these steps.

Start ``vi`` and edit a new file::

  vi helloworld.txt

The editor will start up in normal mode.  The first thing we want to do is add
some text to the file.  Press ``i`` to enter "insert" mode.  You will notice
that the editor begins to let you start typing text into the file.  Enter
some text into the file::

  Hello, World!

Now you want to save the contents, so follow these steps:

1. Press ``Esc`` to exit insert mode and enter command mode
2. Press ``:`` and then ``w``, then ``Enter``.  This will write out to the file.
3. The editor remains in command mode, waiting for your next command.

Now you want to exit the editor:

1. Press ``Esc`` to exit insert mode (hitting ``Esc`` while in command mode does nothing, so it's safe to just hit it for good measure)
2. Press ``:`` and then ``q`` (quit), then ``Enter``.  You will exit ``vi``.

Cursor motion
-------------

Now that you know how to open a file, put some text into it, and then exit the editor.  The next most important thing to learn
is how to move around in the file.  While using the arrow keys might work, it's far from optimal, and may not be implemented in
strictly POSIX-compliant versions of ``vi``.

When in **normal mode**, you use the ``h``, ``j``, ``k``, and ``l`` (ell) keys to move the cursor around:

* ``h`` - left
* ``j`` - down
* ``k`` - up
* ``l`` - right

Seems confusing and awkward at first, right?  It's actually not.  You will quickly learn the motions without
having to think about it, and you will be able to move around very fast in the file, without ever having
to move your hands from the home row of the keyboard.

In addition to the simple cursor motion commands, there are a few other cursor motion commands that are
extremely helpful to know:

* ``w`` - move forward one word
* ``b`` - move back one word
* ``0`` - move to the beginning of the line
* ``^`` - move to the first non-blank character of the line
* ``$`` - move to the end of the line
* ``/some text`` - search forward to the next match of ``some text``, and place the cursor there
* ``?some text`` - search backward to the previous instance of ``some text``, and place the cursor there
* ``n`` - repeat the most recent search
* ``N`` - repeat the most recent search, but in the opposite direction
* ``gg`` - Go directly to the top of the file
* ``G`` - Go directly to the bottom of the file
* ``numberG`` - Go directly to line ``number`` (example: ``20G`` goes to line 20)

Text insertion commands
-----------------------

``vi`` gives you several options for how you actually want to insert text when you enter insert mode.

* ``i`` - insert text at the position of the cursor.  The first character you type will appear to the left of where the cursor is
* ``a`` - append text at the position of the cursor.  The first character you type will appear to the right of where the cursor is
* ``I`` - Same as insert, but first moves the cursor to the beginning of the line (equivalent to ``^i``.
* ``A`` - Same as append, but first moves the cursor to the end of the line (equivalent to ``$a``.
* ``o`` - Open a new line under the cursor and begin inserting text there
* ``O`` - Open a new line above the cursor and begin inserting text there

Text removal commands
---------------------

You will notice that while in **insert mode**, you can use the backspace and delete keys as expected.  This
makes insert mode easiy to use, but it's not particularly efficient if you're trying to eliminate a whole paragraph or something
from your document.  When in **normal mode**, you can issue some commands that remove whole chunks of text:

* ``x`` - Delete the character under the cursor
* ``dw`` - Delete from the character under the cursor to the beginning of the next word
* ``dd`` - Delete the line under the cursor
* ``NUMdd`` - Delete NUM lines of text, ex: ``10dd`` deletes 10 lines
* ``dgg`` - Delete the current line, and everything else to the top of the file
* ``dG`` - Delete the current line, and everything else to the bottom of the file

Undo and Redo
-------------

Now that you know how to add and remove text, you'll inevitably end up making a mistake.
Luckily, ``vi`` lets you undo the last command or insertion, by going back to **normal mode** and hitting
the ``u`` key.

In ``vim`` (but not strict POSIX ``vi``), you can also press ``Ctrl-R`` to redo the last thing
you un-did.

Professional Fonts (vim?)
-------------------------

Syntax Highlighting (vim)
-------------------------

Directory navigation (NERDtree)
-------------------------------

Edit/Open/Close Files
---------------------

Edit/Open/Close Buffers (vim)
-----------------------------
