###########
Style Guide
###########

When writing documentation, we should strive to keep the style used consistent, clear and understandable across all topics.

These documents are written in `reStructuredText`_ with `Sphinx`_ directives.

The style used in writing here should follow the guidelines in the following sections.
This document attempts to conform to the guidelines themselves, and may be used as a reference.
You may view this document's source by clicking "Show Source" on the left-hand menu bar.

.. _reStructuredText: http://docutils.sourceforge.net/rst.html
.. _Sphinx: http://sphinx-doc.org/

*******
Editing
*******

This section attempts to guide the author on ways to edit the source documentation and preserve project style so many editors can collaborate efficiently.
Guides on spacing, heading layout and topic structure will be established here.


Spacing
=======

- Do not add two spaces after terminal punctuation, such as periods.

  The extra space will be ignored when rendered to HTML, so strive for consistency while editing.

- Remove trailing whitespace from your lines.

  Almost any modern text editor can do this automatically, you may have to set your editor to do this for you.
  Reasoning is that:

  a. it makes git unhappy
  b. there's no usefulness in those extra characters
  c. when others edit with whitespace-removing editors, their commit can get polluted with the removals instead of the content

  Basically, don't leave a mess for someone else to clean up.

- Start sentences on new lines.

  Start each new sentence of the same paragraph on the immediately following line.
  Historically, lines had been requested to be wrapped at 80 characters.
  In our modern day and age of powerful text editors this may no longer be very useful, rather asking the author to provide concise text is more valuable.
  If a sentence is very long, consider how you might rewrite it to convey the idea across multiple sentences.
  In general, try to consider how long the line would appear in a text editor at 1024 pixels wide, 12pt font (around 140 characters total).

- Sectional spacing should follow these guidelines:

  * 1 blank line after any section heading, before content starts
  * 2 blank lines after an H1/H2 section ends
  * 1 blank line after any other section ends

  When in doubt, leave one blank line.

- Leave one (1) blank line at the end of every file.
