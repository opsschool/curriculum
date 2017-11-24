Text Editing 201
****************

Vim
===
[What is Vim, compared to Vi?]

Window Splitting
----------------

Plugins
-------

Emacs
=====
[Intro to emacs, what it is, how it's different to Vi]

Vi was first developed on computers where the keyboards looked like this:

https://en.wikipedia.org/wiki/Vi#/media/File:KB_Terminal_ADM3A.svg

Emacs was first developed on computers where the keyboards looked like this:

https://en.wikipedia.org/wiki/Symbolics#/media/File:Symbolics-keyboard.jpg

The important thing to note in the Symbolics keyboard above is the keys on the bottom row:

``[HYPER]`` ``[SUPER]`` ``[META]`` ``[CTRL]`` ``[   SPACE   ]`` ``[CTRL]`` ``[META]`` ``[SUPER]`` ``[HYPER]``

Where Vi has modes for different editing tasks with composable commands, Emacs relies on key combinations or "chords", particularly the ``CTRL`` and ``META`` keys, abbreviated ``C`` and ``M``. 
The convention for canonically specifying a keyboard combination in Emacs is based on those abbreviations using a connecting dash for keys pressed simultanously and a space for follow-up key "arguments".
For example, the most important key combination in Emacs is this one: ``C-h ?`` meaning, press and hold ``Ctrl`` then press ``h``, release and press ``?``. This takes you to Emacs' "help" options.

Modern keyboards lack a ``META`` key, but Emacs has long recognized ``Alt`` as ``META``. 
For example, the second most important key combination in Emacs is ``M-x``, meaning press and hold ``Alt`` and then press ``x``. 
This brings up what amounts to Emacs' "command line" where you can enter any Emacs command to run.

It bears mentioning the "zeroth" most important key command: ``C-g`` meaning press and hold ``Ctrl`` and press ``g``. 
This tells Emacs to cancel whatever it's doing. 
It almost always works, although sometimes after a short wait. 
``C-g`` combines with every partial keyboard combination to cancel it.

What Emacs lacks in command composability, it makes up for in programmability. 
Nearly every aspect of the editing experience is modifiable using Emacs' native language, Emacs Lisp. 
You can use ``M-:`` to open another Emacs' command-line, this one takes any Emacs Lisp expression and evaluates it. 
You can also go to the "scratch" buffer, enter an Emacs lisp expression and press ``C-x C-e`` to evaluate it.

You can read Emacs' online manuals at ``C-h i``. 
Going through the Emacs tutorial---``C-h t``---is highly recommended.

To close Emacs ``C-x C-c``.

Edit/Open/Close Files
---------------------

* ``C-x C-f`` ``find-file``   open a file in a buffer to edit
* ``C-x C-r`` ``find-file-read-only``
* ``C-x k``   ``kill-buffer`` close a buffer
* ``C-x C-f /sudo:localhost:/etc/hosts`` use Tramp-mode to edit a privileged file (of course, you must be privileged)

Edit/Open/Close Buffers
-----------------------

* ``C-x b`` ``switch-to-buffer`` switch to another open buffer
* ``M-x ibuffer`` see a list of open buffers, to navigate or run other commands on 

Directory Navigation
--------------------

* ``C-x d`` see directory contents as a simple list
* ``C-x D`` use Dired to operate on a directory's contents

Syntax Highlighting
-------------------

Emacs autodetects the kind of file one is operating on and applies syntax highlighting as part of the "mode" used for that buffer. 
Emacs comes with built-in modes for almost everything. 
You can explicitly run any mode with ``M-x``.

Line numbers
------------

Emacs typically shows the cursor's position in a buffer on the "mode-line" near the bottom of the screen.
You can configure it to show the position in a "row and column" format running ``M-x line-number-mode`` (usually on by default) and ``M-x column-number-mode``.
To see line numbers in buffer use ``linum-mode``.
To highlight the current line use ``hl-line-mode``.

Window Splitting
----------------

* ``C-x 2`` ``split-window-below`` splits the window horizonally (especially useful for looking at two widely separated parts of the same buffer).
* ``C-x 3`` ``split-window-right`` splits the window vertically
* ``C-x o`` ``other-window`` switches focus to other window
* ``C-x 0`` ``delete-window`` closes this window
* ``C-x 1`` ``delete-other-windows`` closes the other window
* ``C-x 4 f`` ``find-file-other-window`` opens file in other window
* ``C-x 4 r`` ``find-file-read-only-other-window``
* ``C-x 4 b`` ``switch-to-buffer-other-window``

Buffers
-------

``M-x ibuffer``

References
----------

* https://www.gnu.org/software/emacs/manual/ (the web version of ``C-h i``)
* https://www.emacswiki.org
* http://planet.emacsen.org
* http://hyperpolyglot.org/text-mode-editors
