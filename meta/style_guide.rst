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

Showing Examples
================

Very frequently, we may want to show the reader an example of code or console output.
Sphinx provides `language-specific highlighting <http://sphinx-doc.org/markup/code.html>`_, so instead of using a standard literal block (``::``), use the more explicit ``code-block::`` syntax.

An example::

  .. code-block:: ruby

  require 'date'
  today = Date.today
  puts "Today's date is #{today}"

This would render so:

.. code-block:: ruby

  require 'date'
  today = Date.today
  puts "Today's date is #{today}"

This highlighted version is superior, and more easily understood.

Please use the language being displayed for the correct syntax.
For examples of console output, use ``console``.
When showing pseudo-code or items that do not fit a language or console, use ``none``.

When showing shell commands, please use the following standard prompts for user and root:

.. code-block:: console

  user@opsschool ~$ uptime
   14:17:34 up 12:05, 16 users,  load average: 0.34, 0.16, 0.05

  root@opsschool ~# uptime
   14:18:15 up 12:05, 16 users,  load average: 0.22, 0.15, 0.06

These prompts can be set with the following PS1 entries:

.. code-block:: console

  export PS1='user@opsschool \w$ '
  export PS1='root@opsschool \w# '


*******
Writing
*******

This is educational material.
Readers come from many countries, and many of them do not speak English as their first language.
Write so that they can follow the text easily, and so that what you write stays correct.

Language
========

- Use plain, simple English.
  Prefer common words and complete sentences.
  Explain a term the first time you use it.

- Use precise terms.
  For example, write "IP packets are carried inside frames addressed to MAC addresses", not "packets go to a MAC address".

- Say which part of the system does what.
  For example: "connection tracking records the state of each connection; firewall rules decide whether to accept packets."

Accuracy
========

- Check every absolute statement.
  If something is true only by default, or only in some cases, say so.
  Words such as "can", "by default", "usually" and "for example" are often more accurate than "always" or "never".

- Give real values with their context.
  For example: "conntrack's ``TIME_WAIT`` timeout defaults to 120 seconds; other closing states have different timeouts, and these values are configurable."

- Mention the exceptions that matter.
  For example, AppArmor's complain mode logs violations of allow rules instead of blocking them, but explicit ``deny`` rules are still enforced.

- Make examples realistic.
  For example, to show how to raise a limit, use a value that is higher than the current limit.

Tone
====

- Write in a calm, neutral tone.
  This is not a blog post or a news article.

- Avoid "headline style": short, dramatic, absolute sentences written for effect, such as "The service sees nothing."
  Write the full, accurate statement instead.

- Don't over-explain or justify.
  Say what something is, how to see it, what can go wrong, and how to fix it.
  Leave out text that only tells the reader why the topic is important.
