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
  help you see the current state of the system - what is running, what is
  using cpu, memory? Is the disk being heavily used? There is a great deal of
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

* Common culprits (is it plugged in? is it network accessible?)
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

There are two common reasons that you can't bind to a socket: the port is
already in use, or you don't have permission.
As an example, you can see what happens when I try to start a Python 
SimpleHTTPServer on a port that is already in use:

.. code-block:: console
    
    user@opsschool ~$ python -m SimpleHTTPServer 8080
    ...
    socket.error: [Errno 98] Address already in use

Here's an example of what happens when I try to bind to a privileged port 
without proper permissions (in Linux, ports < 1024 are privileged):

.. code-block:: console

    user@opsschool ~$ python -m SimpleHTTPServer 80
    ...
    socket.error: [Errno 13] Permission denied
    
Permission denied reading from / writing to disk
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

A very important skill to learn is the ability to remain calm in the face of
disaster. It's not always easy, especially with a client on the phone, but
panicking will only make a situation worse. Yes, the most critical server in
the infrastructure may have just completely failed without a backup. Instead of
focusing on what will happen as a result of the crisis, focus on what needs to
be done to bring the system back up. Deal with the results later, after fixing
the immediate failure. The fallout of the crisis might be terrible, but it will
almost certainly be worse if the immediate problem isn't fixed. A calm
mind can carefully analyze a situation to determine the best solution.
Panic responses do not benefit from the same calculating rationality.

Different people will adapt to handling crisis situations in different ways.
Some will adopt the detached, analytical calm of a surgeon. Others will
take a few deep breaths to calm themselves before digging in to analyze
the problem. The ability to stay calm in the face of disaster is more
important than the method by which calm is achieved. It will take
practice to reach the point of reacting to a disaster calmly.

Avoid placing blame. It doesn't accomplish anything beyond creating
animosity and tension when a team most needs cohesion and efficiency.
While a good practice in general, it is even more important to resist
the urge to point fingers during a crisis. It doesn't assist in solving
the problem, which is the top priority. Everything else is secondary.

The Importance of Procedure
---------------------------

Creating procedures for responding to disasters provides both a
checklist of things to do in the given situation as well as a structured
way to practice responding to the situation. The practice serves to
solidify understanding of how to react, while the procedure itself
provides a target of mental focus during an actual disaster. Adhering to
the procedure ensures the steps taken to resolve a crisis are well-known
and tested. Focus on the procedure to the exclusion of everything else.

That said, not every situation will have an associated procedure. These
situations call for their own procedures. Try to create a procedure for
every situation that doesn't already have one. This diligence pays off
over time, as history tends to repeat itself. In addition to this, a
procedure for situations lacking a procedure provides a safety net when
everything else fails. This will differ from one organization to the
next, but the value is constant.

Like backups, no disaster recovery procedure is useful unless and until it is
tested. Thorough testing and practicing--in a real environment if
possible--quickly finds problems that will happen in the real world. Beyond
having procedures for known possible failures, a procedure for situations other
procedures do not cover provides a fallback for what to do in the inevitable
unpredictable crisis.

In addition to the technical sector, other industries deal regularly with
crisis response--fire fighters, law enforcement, paramedics. These organizations
have their own procedures. These industries all predate technology, offering
much to learn.

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


