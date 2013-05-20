System daemons 101
******************

NTP
===

The Network Time Protocol, or NTP, is a protocol that networked
entities (routers, switches, servers, firewalls and the like) can use
to synchronise their clocks.

There are two entities in an NTP conversation: servers and clients.

NTP servers
-----------

NTP servers make available readings from a reference clock to NTP
clients. This reference clock can take many forms:

* Global Positioning System (GPS) sensors
* Atomic clocks that measure time through observing the decay of
  radioactive material
* Ordinary server clocks

Along with a clock reading, NTP servers also provide a number to NTP
clients that indicates how trustworthy their reference clock might be.
This number is known as the stratum.  Lower stratum numbers are
better.  Typically GPS and atomic clock sources are considered to be
stratum zero: the most reliable. An NTP server providing access to a
stratum zero reference clock typically tells its clients that it has a
stratum one clock.

.. TODO:: Verify NTP stratum purpose and update!

NTP clients
-----------
NTP clients can take many forms but the most common implementations
manage the clock in their device continually. Typically, an NTP client
will be configured to use multiple NTP servers. The software will
monitor the NTP servers and attempt to synchronise the system clock
with whichever server appears to be the most reliable. NTP server
stratum is a major factor here but is not the only factor.

If adjustments are necessary in order to synchronise the local clock
with an NTP server, two methods are available, depending on the amount
of adjustment required. These methods are slewing and stepping.

If the required adjustment is small, the clock can be slewed. With
this method, the clock continues to monotonically tick, but the length
of each tick is altered: shorter ticks if the clock needs to be
advanced, or longer ticks if it needs to go backward. No ticks are
missed or repeated. This last point is very important - it means that
scheduled tasks (eg. UNIX ``cron`` or ``at`` jobs) will happen as per
their normal schedules, and audit data will appear in the correct
order. Slewing does not result in an immediate "fix" and so is
typically used for smaller discrepancies.

For larger adjustments, the clock may need to be stepped instead. This
results in the clock being instantaneously set to the desired time.
Unlike slewing, this means that clock ticks will be missed or
repeated, depending on the direction of adjustment. This may result in
scheduled tasks being repeated, and audit data may have incorrect
timestamps.

UNIX NTP clients will not generally not step the clock unless
specifically directed to do so. Many systems are configured to step
the clock if required during the system boot process, when the NTP
client daemon is started, but to never do so at any other time.

Some NTP clients will stop synchronising the clock altogether if clock
skew exceeds a defined value. Often, an audit log event will be
generated when this happens.

Kernel daemons
==============
``klogd``, etc

Filesystem daemons
==================
``xfslogd``, etc

udevd
=====

cron
====

UNIX systems provide a facility to schedule and execute recurring
background processing jobs. This facility is known as ``cron``. Users
add entries to their own ``crontab`` in order to schedule their jobs.

On most UNIX systems the ``cron`` daemon is simply named ``cron``. On
some Linux systems, however, it is named ``crond``. Apple's Mac OS X
operating system has subsumed the functionality of ``cron`` (along
with ``atd`` and ``init``) into a single daemon named ``launchd``. The
same basic functionality is available in all cases.

Job execution
-------------

UNIX systems will have a service running that knows about all
scheduled ``cron`` jobs, executes them under the job owner's user
account, and, if any output is generated, emails that output to the
job owner. The scheduler runs under the ``root`` user account.

The exact process of executing cron jobs varies from system to system,
but typically involves at least the following:

1. Checking if the user is allowed to execute ``cron`` jobs at all,
   and aborting if not, probably with an audit event generated. For
   example, users whose accounts are locked will not be allowed to
   execute jobs
2. Creating a new process (from which all further tasks for this job
   execution happen)
3. Changing to the correct user if the job is not to be executed as
   the ``root`` user
4. Setting some environment variables (eg. a user may have specified a
   desired setting for ``$PATH`` in their ``crontab``)
5. Executing the command to start the job
6. If any output is generated, send it in an email to the MAILTO
   address in the ``crontab``, if specified, otherwise in an email to
   the owner of the ``crontab``

Job specification
-----------------

A ``crontab`` consists of up to three types of entries.

* comments
* environment variables
* jobs

The job specification format is best learned by example and
experimentation. Most systems will have some built-in jobs used for
system maintenance. Try entering ``crontab -l`` as the root user on
commercial UNIX systems such as Solaris or HP-UX, or looking in the
various ``/etc/cron.*`` directories on Linux systems. For definitive
documentation, see the manual page for the ``crontab`` file format. On
many systems this can be viewed with the command ``man 5 crontab``.

A user's ``crontab`` can be edited with the command ``crontab -e``.

Typically ``crontab`` data is kept under ``/var/spool/cron`` but this
should be used only for investigative purposes. Editing a ``crontab``
should always involve the ``crontab`` command.

atd
===

The ``atd`` daemon is also a job scheduler, but unlike ``cron`` it is
used to execute each job once only. Several commands are used to
manage jobs:

* ``at`` adds new jobs
* ``atq`` shows the queue of ``at`` jobs to be executed
* ``atrm`` removes jobs that have not yet been executed

The ``at`` command is exceedingly flexible in how execution times may
be specified. Words such as ``today``, ``tomorrow`` and ``teatime``
may all be used. These are documented in the ``at`` manual page.

sshd
====

Automounter
===========

Automounters are a class of system daemons that mount filesystems,
often of the networked variety such as NFS or AFS, on demand (that is,
when a UNIX process tries to access it), and optionally also unmount
them after a period of no usage.

More complex environments may provide a central directory of available
networked filesystems to automounter daemons. A common application of
this arrangement may be found in some university environments where
users' ``$HOME`` directories are shared between many servers and are
not necessarily all located on the same server. A central map can
inform automounters of the location of any given user's ``$HOME`` and
so allow it to be automatically mounted upon login.

There are three main varieties of automounter daemon - Linux, BSD and
SysV.  Their names and configuration files are different, but the
basic behaviour and purpose are the same.

Quick Linux automounter example
-------------------------------
If your operating system is Red Hat Enterprise Linux or a derivative
such as CentOS, Scientific Linux or Oracle Linux, and your system has
an optical drive, you can get a quick demonstration very easily by
following the below steps (NOT on a production system!):

1. Install the ``autofs`` daemon if it isn't already installed:
   ``yum -y install autofs``
2. Ensure the line beginning with ``/misc`` in ``/etc/auto.master``
   ends with ``-t 180`` (this enables automatic unmounting)
3. Enable and start the service: ``chkconfig autofs on`` and then
   ``service autofs restart``
4. Ensure that your optical drive is empty, then try to access the
   ``/misc/cd`` directory: ``ls /misc/cd``
5. Insert any data CD or DVD into the optical drive
6. Try again to list the contents of the ``/misc/cd`` directory.
7. Type ``mount`` and observe that a filesystem is now mounted on
   ``/misc/cd``

If everything is working as it should, the optical media should be
automatically mounted on ``/misc/cd`` and the access to that directory
should succeed. Note that this is wholly transparent to the
application accessing the directory and/or any files in it. This is
critical to the usefulness of the automounter - applications do not
need to have any special logic to handle automounting, nor in fact
even be aware of whether the system uses an automounter or not. It is
entirely transparent.

Now if you wait a few minutes (per the ``auto.master`` setting in step
#2 above) and then type ``mount`` again; if nothing has accessed
``/misc/cd`` in the last 180 seconds, it will have been automatically
unmounted.
