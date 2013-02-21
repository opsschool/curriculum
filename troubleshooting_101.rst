Troubleshooting
***************

A key skill for anyone doing operations, is the ability to successfully
troubleshoot problems.

.. contents::
   :depth: 4
   :local:

Methodologies
=============

Here we will got over a few steps you can take to help quickly narrow down
problems to their causes.


Triaging with differential diagnosis
------------------------------------

What is broken? First think about how it works in most basic terms.
Then build on that the things which can break.
Then from those, pick the ones that could cause the symptoms you see.

Example:
    You cannot ping a server.


Preparing your toolbelt
-----------------------
You have a variety of tools at your fingertips to help work out the cause of a
problem. Over time you will expand what is in your toolbelt, but to start with
you must know how to use each of these:

* ``top``, ``vmstat``, ``iostat``, ``systat``, ``sar``, ``mpstat``
  These help you see the current state of the system - what is running, what is
  using cpu, memory? Is the disk being heavily used? There is a lot of
  information, and knowing how these tools work will help you pick out the bits
  you should focus on.
* ``tcpdump``, ``ngrep``
  If you suspect you have a network-related problem, ``tcpdump`` and ``ngrep``
  can help you confirm it.

Walk through of a diagnosis
---------------------------

* Eliminating variables

 * What changed recently?
 * Could any of the symptoms be red herrings?

* Common culprits (is it plugged in?)
* Look through your logs
* Communicating during an outage
* 'Talking Out-Loud' (IRC/GroupChat)
* Communicating after an outage (postmortems)


Recent changes
--------------

Often problems can be traced back to recent changes.
Problems that start around the time of a change aren't usually coincidence.

Learning common errors
----------------------

Over time you may find that a small set of errors cause a large portion of the
problems you have to fix. Let's cause some of these problems and see how we
identify and fix them.

Cannot bind to socket
^^^^^^^^^^^^^^^^^^^^^

Permission denied reading to / writing from disk
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Out of disk space
^^^^^^^^^^^^^^^^^
(finding large files, and also finding deleted-but-open files)

Mystery problems and SELinux
^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Out of inodes
^^^^^^^^^^^^^
Manifests as "disk full" when ``df`` claims you have disk space free.


Working effectively during a crisis
===================================

Being able to work successfully through a crisis is crucial to being a good
operations person. For some it is a personality trait, but it can certainly be
learned and is almost a requirement for many employers.

The most important skill to learn is the ability to remain calm in the face of
disaster. It's not always easy, especially with a client on the phone, but
panicking will only make a situation worse. Yes, the most critical server in
the infrastructure may have just completely failed without a backup. Instead of
focusing on what will happen as a result of the crisis, focus on what needs to
be done to bring the system back up. Deal with the results later, after fixing
the immediate failure. The fallout of the crisis might be terrible, but it will
almost certainly be worse if the immediate problem isn't fixed.

Different people will adapt to handling crisis situations in different ways.
Some will adapt the detached, analytical calm of a surgeon. Others will
take a few deep breaths to calm themselves before digging in to analyze
the problem. The ability to stay calm in the face of disaster is more
important than the method by which calm is achieved. It will take
practice to reach the point of reacting to a disaster calmly.

The Importance of Procedure
---------------------------



Non-technical skills
--------------------

Situational Awareness (Mica Endsley)
Decision Making (NDM and RPD) - Klein
Communication (Common ground, Basic Compact, Assertiveness)
Team Working (Joint Activity, fundamentals of coordination and collaboration)
Leadership (before, during, after incidents) (Weick, Sutcliffe work on HROs)
Managing Stress
Coping with Fatigue
Training and Assessment Methods
Cognitive Psychology concerns (escalating scenarios, team-based troubleshooting)


