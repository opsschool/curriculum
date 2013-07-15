``/bin/init`` and its descendants
*********************************


``init``
========

systemd
=======

upstart
=======

Upstart is a project that aims to replace to old init system by providing one
standard way of starting and stopping daemons with the correct environment.
A second goal is to speed up a computer's boot time. It achieves this by
removing the slow init shell scripts, and also by parallelizing as much of the
startup as possible. Where old init daemons start daemons in successive order,
upstart issues "events" on which "jobs" can listen.

Such an event can be e.g.: ``filesystem`` - indicates that the system has mounted
all its filesystems and we can proceed to start any jobs that would depend
on a filesystem. Each job then becomes an event of its own, upon which others
can depend. These events can be broken up into stages: ``starting(7)``, and
``started(7)``; and ``stopping(7)``, and ``stopped(7)`` respectively.

A good starting point for learning how different jobs on a system are interconnected
is ``initctl(8)``'s ``show-config`` command:

.. code-block:: console

    igalic@tynix ~ % initctl show-config
    avahi-daemon
      start on (filesystem and started dbus)
      stop on stopping dbus
    cgroup-lite
      start on mounted MOUNTPOINT=/sys
    elasticsearch
      start on (filesystem or runlevel [2345])
      stop on runlevel [!2345]
    mountall-net
      start on net-device-up
    ...

This snippet reveals that upstart will ``stop`` the ``avahi-daemon`` at the same
time as dbus. Unlike many other daemons that depend on the whole filesystem, upstart
will ``start`` ``cgroup-lite`` as soon as the ``/sys`` filesystem is mounted.

Upstart is also able to "supervise" programs: that is, to restart a program
after it crashed, or was killed. To achieve this, upstart needs to "follow" a
programs progression. It uses the ``ptrace(2)`` system call to do so. However,
following a daemons forks is *complex*, because not all daemons are written alike.
The upstart documentation recommends to avoid
this whenever possible and force a to remain in the foreground. That
makes upstart's job a lot easier.

Finally, upstart can also switch to a predefined user (and group) before
starting the program. Unlike systemd_ and SMF_, however, it cannot drop to a
limited set of ``capabilities(7)`` before doing so.

Putting it all together in an example:

.. code-block:: console

    # httpd-examplecom - Apache HTTP Server for example.com
    #

    # description     "Apache HTTP Server for example.com"

    start on filesystems
    stop on runlevel [06]

    respawn
    respawn limit 5 30

    setuid examplecom
    setgid examplecom
    console log

    script
        exec /opt/httpd/bin/httpd -f /etc/httpds/example.com/httpd.conf -DNO_DETACH -k start
    end script

In this example we define an upstart job for serving ``example.com`` from
the Apache HTTP Server. We switch to the user/group ``examplecom`` and start
``httpd`` in the foreground, by passing the option ``-DNO_DETACH``.

To activate this job, we simply place it in a file in ``/etc/init/``, e.g.
``/etc/init/httpd-examplecom.conf``. We can then start/stop the job by issuing:

.. code-block:: console

    % sudo start httpd-examplecom

Note that this job definition alone already will guarantee that the system will
start the job on reboot. If this is not what we want, we can add the stanza
``manual`` to the job definition.


SMF
===

daemontools
===========
