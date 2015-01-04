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

As Roy Rapoport states, "change is often the enemey of stability." Frequently,
you will find yourself troubleshooting outages or system issues that were
self-inflicted.

It's probably fairly obvious but when a system problem occurs or performance
degrades, something has changed! If nothing had changed, no one would be
complaining. Often problems can be traced back to recent changes. You
will quickly learn that problems that start around the time of a change,
are not usually a coincidence.

As you troubleshoot an issue, you will find yourself logging in to hosts and
monitoring systems to identify what's different. After all, the system is
responding differently (suddenly) and it's your job to find out what's changed.
Did someone on the team make a change? Or, was the change a degredation in the
operating environment? Such as a service failing, a system failing, a bad disk
or memory that's failed.

Here are a few questions you may ask your servers, yourself, your teammates
or developers who may deploy to assist in identifying a change:

Who was last logged in to the system, and when?

.. code-block:: console

    user@opsschool ~$ last
    brian    pts/0        nosferatu.krs.io Sun Jan  4 00:27   still logged in
    brian    pts/0        nosferatu.krs.io Fri Jan  2 18:04 - 20:53  (02:48)
    sjackson pts/0        192.168.10.2     Fri Jan  2 17:36 - 17:36  (00:00)
    sjackson pts/0        192.168.10.2     Fri Jan  2 09:09 - 09:09  (00:00)

What did they do?

.. code-block:: console

    user@opsschool ~$ sudo less /home/sjackson/.bash_history
    ..
    vi /etc/crontab
    vi /etc/important-app/config
    cd /opt/important-app
    git checkout v1.6.0

What software versions were recently deployed?

What in the runtime environment has changed (configuration, database
schema, etc.)?

What other changes have been made on the network?

What actors are involved from point a to z?

This last question can be used as a very powerful exercise. If you imagine
your service and all of the actors required to make that service work you
can begin to visualize the entire application, platform and network stack
that supports and facilitates your application.  For example, are you
running a Python web service API, behind Apache 2, Hosted inside of a Docker
instance, that running on top of an Amazon EC2 instance that speaks to a
MySQL database? Picturing which actors are involved allows you to quickly
ask which system has had changes recently and allows you to prioritize
your troubleshooting efforts. For example, maybe performance is the current
problem you are troubleshooting and you quickly surmise that delay is commonly
added by a database server or by the Amazon EC2 instance that hosts the docker
image.

Remember, when you find what changed, be sure to add a monitoring test to more
quickly identify this condition in the future. Operations teams are small and
agile and preventing break-fix changes in the future is your first goal when
identifying a change that caused an issue and your second goal is to add
monitoring and notifications around the issue so if, unfortunately, it occurs
again in the future you identify the issue very quickly.

If you are on a medium to larger sized team and you find yourself frequently
troubleshooting breaking changes, many operations teams are having success
with deploying change notification systems. A change notification system is
any lightweight logging tool that servers, automated applications (CI/CD)
and Operations teammates use to post a change that has occurred. These change
notification systems are usually a lightweight REST/Web service API. For
example, when a new version of software is deployed, it could post and
log that change to the web service api or when an operations engineer
adds a new server he could curl and 'post' to the API and share with the
team that 'something has changed.'

Having a change notification platform allows anyone to pull up an Operations
dashboard log that shows recent system changes. This quickly enables anyone
to know very quickly what's recently changed in your environment.

One that was open sourced recently was `Blesk <https://github.com/Netflix/blesk>`

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
    
Permission denied reading to / writing from disk
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
File system permissions, and attributes (chattr)?


Out of disk space
^^^^^^^^^^^^^^^^^

If you have appropriate monitoring in place, hopefully you are troubleshooting
a full disk because you received a proactive alert. If not, take an action item
to dive deep into :doc:`monitoring_101` and configure your notification systems to
alert for any volume that nears 75% utilization and configure a second critical
alert if it exceeds 90% utilization.

A frequent cause of downtime is a volumes disk space being fully utilizes.
Operations engineers often find themselves logged in to a system not knowing what
is specifically wrong but through habit have learned one of the first things
they will want to do is check disk space utilization on all of the volumes.

.. code-block:: console
    
    user@opsschool ~$ df -h
    Filesystem            Size  Used Avail Use% Mounted on
    /dev/sda6             454G  454G    0G 100% /
    tmpfs                  12G     0   12G   0% /lib/init/rw
    udev                   12G  144K   12G   1% /dev
    tmpfs                  12G     0   12G   0% /dev/shm
    /dev/sda1             236M   41M  183M  19% /boot

Here we can see that /dev/sda6 is 100% used. We better quickly find the
culprit.

Find files that are larger than 50M:

.. code-block:: console

    user@opsschool ~$ sudo find / -size +50M
    ..
    user@opsschool ~$ sudo find / -size +50M -exec ls -lh {} \; | awk '{ print $9 ": " $5 }'


Often times, you are going to find a large log file. In fact, you may quickly
find out you are now solving for two or three symptoms. For example,
the application servers disk filled, because it was logging a new critical
event, because database schema changed.

NOTE: Be very careful, just because you can use a few of these tools to
identify large files, it does not mean they are the culprit or that they
should be deleted. For example, it would be disastorous to delete a 1.1G file
named /var/lib/mysql/ibdata1 or /var/lib/mysql/mysql-bin.000971. Use caution
and if you aren't sure what the file is, be sure to Google it and identify the
proper way to perform maintenance.


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


