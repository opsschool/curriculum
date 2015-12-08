Text Editing 201
****************

Vim
===
Vim is an insertion mode editor just like VI.
The differences between Vi and Vim aren’t terribly significant. Vim is simply an improved version of Vi. It pretty much has a ton of functionality that Vi doesn’t.

Within Vim you can see the differences between Vi and Vim by running the following command: 
::
  :h vi-differences

Window Splitting
----------------
In Vim, you can view several buffers at once by loading them into multiple windows. 
Here we discuss the essentials of working with windows: opening, closing, resizing, moving between and rearranging them.

Opening split windows
~~~~~~~~~~~~~~~~~~~~~
 
+--------------------------------------------------------------------------------------+----------------------------+
| Action                                                                               |  Command                   |
+======================================================================================+============================+
| Split the current window horizontally, loading the same file in the new window       | `ctrl-w s`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Split the current window vertically, loading the same file in the new window         |  `ctrl-w v`                |
+--------------------------------------------------------------------------------------+----------------------------+
| Split the current window horizontally, loading filename in the new window            | `:sp[lit] filename`        |
+--------------------------------------------------------------------------------------+----------------------------+
| Split the current window vertically, loading filename in the new window              | `:vsp[lit] filename`       |
+--------------------------------------------------------------------------------------+----------------------------+
 


Closing split windows
~~~~~~~~~~~~~~~~~~~~~

+--------------------------------------------------------------------------------------+----------------------------+
| Action                                                                               |  Command                   |
+======================================================================================+============================+
| Close the currently active window                                                    | `:q[uit]`                  |
+--------------------------------------------------------------------------------------+----------------------------+
| Close all windows except the currently active window                                 |  `:on[ly]`                 |
+--------------------------------------------------------------------------------------+----------------------------+


Changing focus between windows
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

+--------------------------------------------------------------------------------------+----------------------------+
| Action                                                                               |  Command                   |
+======================================================================================+============================+
| Cycle between the open windows                                                       | `ctrl-w w`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Focus the window at the left                                                         |  `ctrl-w h`                |
+--------------------------------------------------------------------------------------+----------------------------+
| Focus the window at the bottom                                                       | `ctrl-w j`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Focus the window at the top                                                          | `ctrl-w k`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Focus the window at the right                                                        | `ctrl-w l`                 |
+--------------------------------------------------------------------------------------+----------------------------+

Resizing windows
~~~~~~~~~~~~~~~~

+--------------------------------------------------------------------------------------+----------------------------+
| Action                                                                               |  Command                   |
+======================================================================================+============================+
| Increase height of current window by 1 line                                          | `ctrl-w +`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Decrease height of current window by 1 line                                          | `ctrl-w -`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Maximise height of current window                                                    | `ctrl-w _`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Maximise width of current window                                                     | `ctrl-w |`                 |
+--------------------------------------------------------------------------------------+----------------------------+

Moving windows
~~~~~~~~~~~~~~

+--------------------------------------------------------------------------------------+----------------------------+
| Action                                                                               |  Command                   |
+======================================================================================+============================+
| Rotate all windows                                                                   | `ctrl-w r`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Exchange current window with its neighbour                                           |  `ctrl-w x`                |
+--------------------------------------------------------------------------------------+----------------------------+
| Move current window to far left                                                      | `ctrl-w shift-H`           |
+--------------------------------------------------------------------------------------+----------------------------+
| Move current window to bottom                                                        | `ctrl-w shift-J`           |
+--------------------------------------------------------------------------------------+----------------------------+
| Move current window to top                                                           | `ctrl-w shift-K`           |
+--------------------------------------------------------------------------------------+----------------------------+
| Move current window to far right                                                     | `ctrl-w shift-L`           |
+--------------------------------------------------------------------------------------+----------------------------+

Plugins
-------
Vim has a plugin system which can be used to add more functionality. These plugins are implemented as simple configuration files, located in subdirectories inside the `$HOME/.vim directory`. If you are using a lot of plugins, managing them can be a hassle, so a plugin manager can help take care of them effectively.

Emacs
=====
Emacs is an extensible system for editing text built on top of a powerful programming language. With a typical text editor you can only do what the author thought to include commands for. Emacs can be made to do pretty much anything you can think of. Unlike VI and VIM, Emacs is not an insertion mode editor, meaning that any character typed in Emacs is automatically inserted into the file, unless it includes a command prefix.

Both VIM and Emacs support many fundamental virtues of text editors such as extensive syntax highlighting, collapsible functions, spell checking, macros, undo-redo, multiple document editing, and a large support community. They are both free, Open Source, mature and well developed pieces of software. But...while both Emacs and Vim are considered powerful text editors, both capable of providing many of the same feature lists for general editing commands, the primary difference between these editors is the fundamental philosophy behind their design, and more to the point, the types of workflow that they were originally designed to handle.


Edit/Open/Close Files
---------------------

+--------------------------------------------------------------------------------------+----------------------------+
| Action                                                                               |  Command                   |
+======================================================================================+============================+
| Open/Create File                                                                     | `ctrl-x ctrl-f`            |
+--------------------------------------------------------------------------------------+----------------------------+
| Close file                                                                           |  `ctrl-x k`                |
+--------------------------------------------------------------------------------------+----------------------------+
| Save file                                                                            | `ctrl-w ctrl-s`            |
+--------------------------------------------------------------------------------------+----------------------------+
| Undo                                                                                 | `ctrl-x u`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Cut                                                                                  | `shift-del`                |
+--------------------------------------------------------------------------------------+----------------------------+
| Copy                                                                                 | `ctrl-ins`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Paste                                                                                | `shift-ins`                |
+--------------------------------------------------------------------------------------+----------------------------+
| Interactive search                                                                   | `ctrl-s`                   |
+--------------------------------------------------------------------------------------+----------------------------+

Edit/Open/Close Buffers
-----------------------

+--------------------------------------------------------------------------------------+----------------------------+
| Action                                                                               |  Command                   |
+======================================================================================+============================+
| Open/Create buffer                                                                   | `ctrl-x b name`            |
+--------------------------------------------------------------------------------------+----------------------------+
| Close buffer                                                                         | `ctrl-x k name`            |
+--------------------------------------------------------------------------------------+----------------------------+
| Save buffer                                                                          | `ctrl-x ctrl-s`            |
+--------------------------------------------------------------------------------------+----------------------------+
| Get a list of buffers                                                                | `ctrl-x ctrl-b`            |
+--------------------------------------------------------------------------------------+----------------------------+

Directory Navigation
--------------------

Syntax Highlighting
-------------------
The syntax highlighting in Emacs is enabled by default. If this is not the case you can enable the syntax highlighting for the current buffer with the following command.
:: 
  M-x font-lock-mode
  
If you want to enable syntax highlighting for all buffers then use following command.
::
  M-x global-font-lock-mode
  
If you want to enable the syntax highlighting permanently, you can also add next line to the .emacs file. 
::
  (global-font-lock-mode 1)
  
With font-lock-mode turned on, different types of text will appear in different colors. For instance, in a programming mode, variables will appear in one face, keywords in a second, and comments in a third. With the syntax highlighting the user experience will be a lot better. 

Line numbers
------------
Line numbers are always a must when you are using a texteditor, especially when you are writing a script.
To enable line numbers for the current buffer of Emacs, use the following command.
::
  M-x linum-mode

To enable line numbers globally.
::
  M-x global-linum-mode
  
If you want to enable line numbers permanently, you can also add next line to the .emacs file. 
::
  (global-linum-mode 1)

Window Splitting
----------------

+--------------------------------------------------------------------------------------+----------------------------+
| Action                                                                               |  Command                   |
+======================================================================================+============================+
| Split window vertically                                                              | `ctrl-x 2`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Split window horizontally                                                            | `ctrl-x 3`                 |
+--------------------------------------------------------------------------------------+----------------------------+
| Select another window                                                                | `ctrl-x o`                 |
+--------------------------------------------------------------------------------------+----------------------------+

Buffers
-------
The text you are editing in Emacs resides in an object called a buffer. Each time you visit a file, a buffer is used to hold the file’s text. Each time you invoke Dired, a buffer is used to hold the directory listing. If you send a message with C-x m, a buffer is used to hold the text of the message. When you ask for a command’s documentation, that appears in a buffer named *Help*.

Each buffer has a unique name, which can be of any length. When a buffer is displayed in a window, its name is shown in the mode line. The distinction between upper and lower case matters in buffer names. Most buffers are made by visiting files, and their names are derived from the files’ names; however, you can also create an empty buffer with any name you want. A newly started Emacs has several buffers, including one named *scratch*, which can be used for evaluating Lisp expressions and is not associated with any file.

At any time, one and only one buffer is selected; we call it the current buffer. We sometimes say that a command operates on “the buffer”; this really means that it operates on the current buffer. When there is only one Emacs window, the buffer displayed in that window is current. When there are multiple windows, the buffer displayed in the selected window is current.
