Text Editing 201
****************

Vim
===
Vim is an insertion mode editor just like VI.
The differences between Vi and Vim aren’t terribly significant. Vim is simply an improved version of Vi. It pretty much has a ton of functionality that Vi doesn’t.

Within Vim you can see the differences between Vi and Vim by running the following command: 

  :h vi-differences

Window Splitting
----------------
In Vim, you can view several buffers at once by loading them into multiple windows. 
Here we discuss the essentials of working with windows: opening, closing, resizing, moving between and rearranging them.

### Opening split windows
 
 +------------+------------+
| Action  | Command  |
+============+============+
| Split the current window horizontally, loading the same file in the new window | `ctrl-w s`   |
+------------+------------+
| Split the current window vertically, loading the same file in the new window | `ctrl-w v`  |
+------------+------------+
| Split the current window horizontally, loading filename in the new window  | `:sp[lit] filename`  |
+------------+------------+
| Split the current window vertically, loading filename in the new window | `:vsp[lit] filename`    |
+------------+------------+

 
| Action                                      | Command                   |
| :---                                        | :---                      |
| Split the current window horizontally, loading the same file in the new window                  | `ctrl-w s`         |
| Split the current window vertically, loading the same file in the new window | `ctrl-w v`         |
| Split the current window horizontally, loading filename in the new window                      | `:sp[lit] filename` |
| Split the current window vertically, loading filename in the new window               | `:vsp[lit] filename`         |

### Closing split windows

| Action                                      | Command                   |
| :---                                        | :---                      |
| Close the currently active window                  | `:q[uit]`              |
| Close all windows except the currently active window | `:on[ly]`         |

### Changing focus between windows
 
| Action                                      | Command                   |
| :---                                        | :---                      |
| Cycle between the open windows                  | `ctrl-w w`|
| Focus the window at the left | `ctrl-w h`         |
| Focus the window at the bottom                  | `ctrl-w j` |
| Focus the window at the top        | `ctrl-w k` |
| Focus the window at the right             | `ctrl-w l` |

### Resizing windows

| Action                                      | Command                   |
| :---                                        | :---                      |
| Increase height of current window by 1 line                 | `ctrl-w -`|
| Decrease height of current window by 1 line | `ctrl-w +`         |
| Maximise height of current window                  | `ctrl-w _` |
| Maximise width of current window        | `ctrl-w |` |

### Moving windows

| Action                                      | Command                   |
| :---                                        | :---                      |
| Rotate all windows                 | `ctrl-w r`|
| Exchange current window with its neighbour | `ctrl-w x`         |
| Move current window to far left                | `ctrl-w shift-H` |
| Move current window to bottom        | `ctrl-w shift-J` |
| Move current window to top      | `ctrl-w shift-K` |
| Move current window to far right    | `ctrl-w shift-L` |

Plugins
-------

Emacs
=====
Emacs is an extensible system for editing text built on top of a powerful programming language. With a typical text editor you can only do what the author thought to include commands for. Emacs can be made to do pretty much anything you can think of. Unlike VI and VIM, Emacs is not an insertion mode editor, meaning that any character typed in Emacs is automatically inserted into the file, unless it includes a command prefix.

Both VIM and Emacs support many fundamental virtues of text editors such as extensive syntax highlighting, collapsible functions, spell checking, macros, undo-redo, multiple document editing, and a large support community. They are both free, Open Source, mature and well developed pieces of software. But...while both Emacs and Vim are considered powerful text editors, both capable of providing many of the same feature lists for general editing commands, the primary difference between these editors is the fundamental philosophy behind their design, and more to the point, the types of workflow that they were originally designed to handle.

[Intro to emacs, what it is, how it's different to Vi]

Edit/Open/Close Files
---------------------

Edit/Open/Close Buffers
-----------------------

Directory Navigation
--------------------

Syntax Highlighting
-------------------

Line numbers
------------

Window Splitting
----------------

Buffers
-------
