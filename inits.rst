``/bin/init`` and its descendants
*********************************


``init``
========

systemd
=======

upstart
=======

Upstart is a development that arose from the need to speed up the boot time
of a computer. The strategy it uses is to parallelize as much of the startup
as possible. It achieves this by issuing "events" on which the "jobs" can listen.

Such an event can be e.g.: ``filesystem`` - indicating that (all) file-systems
have now been mounted and we can proceed to start any jobs that would depend
on a file-system. Each job then becomes an event of its own, upon which others
can depend. These events can be broken up into stages, ``starting(7)``, and
``started(7)``; and ``stopping(7)``, and ``stopped(7)`` respectively.

A good starting point for learning how different jobs on a system are inter
connected is ``initctl(8)``'s ``show-config`` command:

.. code-blocks:: console

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

Upstart is also able to "supervise" programs: That is, to restart a program
after it crashed, or was killed. To achieve this, upstart needs to "follow" a
programs progression. It uses the ``ptrace(2)`` system call to do so. However,
doing so can be rather complex. The upstart documentation recommends to avoid
this whenever possible and simply start a program in the foreground. That
makes upstart's job a lot easier.

Finally, upstart can also switch to a predefined user (and group), before
starting the program. Unlike systemd and SMF however, it cannot drop to a
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

Note that this job definition alone already guarantees that this job will
be started on reboot. If this is not intended, we can add the stanza
``manual`` to the job definition.


SMF
===

daemontools
===========
