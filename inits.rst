``/bin/init`` and its descendants
*********************************


``init``
========

systemd
=======

systemd is the init system on most current Linux distributions, including Debian, Ubuntu, Red Hat Enterprise Linux and Fedora.
It starts services at boot, restarts them when configured to, and collects their logs in the journal.

Units
-----

systemd manages units.
A service, such as a web server, is a unit with a name ending in ``.service``.
Other kinds of units include mounts (``.mount``), timers (``.timer``) and groups of units (``.target``).

Each unit is described by a unit file.
Packages install unit files in ``/usr/lib/systemd/system`` (``/lib/systemd/system`` on older Debian and Ubuntu releases).
Files in ``/etc/systemd/system`` are local changes, and a file there replaces a package's file of the same name.
To change only some settings, add a drop-in file: a file ending in ``.conf`` in a directory named after the unit, such as ``/etc/systemd/system/nginx.service.d/``.
``systemctl edit nginx`` creates one for you.
``systemctl cat nginx`` shows the unit file and its drop-ins on disk, each with its path.
These contents may differ from systemd's loaded configuration until you run ``systemctl daemon-reload``.
Use ``systemctl show nginx`` to inspect loaded properties.

After you change a unit file or a drop-in yourself, run ``systemctl daemon-reload`` so that systemd reads it again.

A service unit's ``[Service]`` section says how to run the program, for example:

.. code-block:: ini

  [Service]
  EnvironmentFile=/etc/inventory/inventory.env
  ExecStart=/usr/local/bin/inventory
  Restart=always
  RestartSec=2

``EnvironmentFile=`` reads environment variables from a file, which many services use for their settings.
``Restart=always`` restarts the program after both successful and failed exits.
It does not restart the program when systemd stops it, for example through ``systemctl stop``.
``RestartSec=`` sets how long systemd waits before restarting; the default is 100 milliseconds.
``Restart=on-failure`` restarts it only when it fails, for example when it exits with a non-zero status.

Managing services
-----------------

.. code-block:: console

  root@opsschool # systemctl start inventory     # start it now
  root@opsschool # systemctl stop inventory      # stop it now
  root@opsschool # systemctl restart inventory   # stop it, then start it
  root@opsschool # systemctl enable inventory    # start it at boot
  root@opsschool # systemctl disable inventory   # don't start it at boot

``systemctl status`` shows whether a service is running, its main process, and its most recent log lines.
``systemctl --failed`` lists units that have failed.

A service that keeps exiting
----------------------------

With ``Restart=always``, a service that repeatedly exits on its own soon after starting enters a restart loop.
``systemctl status`` then shows it as ``activating (auto-restart)``, and the exit status of the last attempt:

.. code-block:: console

  root@opsschool # systemctl status inventory
  ● inventory.service - Inventory API
       Loaded: loaded (/etc/systemd/system/inventory.service; enabled; preset: enabled)
       Active: activating (auto-restart) (Result: exit-code) since Mon 2026-10-05 02:19:42 UTC; 334ms ago
      Process: 19889 ExecStart=/usr/local/bin/inventory (code=exited, status=1/FAILURE)
     Main PID: 19889 (code=exited, status=1/FAILURE)

Clients of the service fail while this goes on, often with a "connection refused" error, because nothing is listening.
The status line alone doesn't say why the program exits.
The program's own messages are in the journal, between the lines in which systemd starts it and records that it exited:

.. code-block:: console

  root@opsschool # journalctl -u inventory -n 6
  Oct 05 02:19:42 opsschool systemd[1]: inventory.service: Scheduled restart job, restart counter is at 3.
  Oct 05 02:19:42 opsschool systemd[1]: Started inventory.service - Inventory API.
  Oct 05 02:19:42 opsschool inventory[19889]: inventory: invalid configuration: LISTEN_PORT: "80a" is not a number
  Oct 05 02:19:42 opsschool systemd[1]: inventory.service: Main process exited, code=exited, status=1/FAILURE
  Oct 05 02:19:42 opsschool systemd[1]: inventory.service: Failed with result 'exit-code'.

By default, if a unit is started more than 5 times within 10 seconds, systemd stops trying, and ``systemctl status`` shows ``start request repeated too quickly``.
These limits are ``StartLimitBurst=`` and ``StartLimitIntervalSec=`` in the ``[Unit]`` section.
After you fix the cause, ``systemctl reset-failed inventory`` clears the failure, and ``systemctl start inventory`` starts it again.

Useful ``journalctl`` options:

- ``-u <unit>``: only messages from that unit
- ``-n 50``: the last 50 lines
- ``-f``: keep printing new messages as they arrive
- ``--since "10 minutes ago"``: only recent messages
- ``-b``: only messages since the last boot
- ``-p err``: only messages at error priority or higher

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

Control groups
==============

Control groups (cgroups) are a Linux kernel feature that groups processes, measures the resources each group uses, and can limit them.
Current distributions use cgroup version 2, which has a single tree of groups mounted at ``/sys/fs/cgroup``.

systemd puts each service in its own cgroup, so every process that a service starts, including its children, is counted and limited together.
Services are grouped further into slices, such as ``system.slice`` for system services and ``user.slice`` for logged-in users.
A limit on a slice applies to all the services in it together.
``systemd-cgls`` shows the tree, and ``systemd-cgtop`` shows how much CPU, memory and I/O each group uses.
``systemctl status`` shows a service's cgroup on the ``CGroup:`` line.

Limiting CPU
------------

In a unit file, ``CPUQuota=`` limits how much CPU time a service or slice can use, as a percentage of one CPU.
For example, ``CPUQuota=20%`` lets it use 20 milliseconds of CPU time in every 100 millisecond period.
200% would allow two full CPUs.
systemd writes the limit to the group's ``cpu.max`` file, as the quota and the period in microseconds:

.. code-block:: console

  root@opsschool ~# cat /sys/fs/cgroup/system.slice/quotademo.service/cpu.max
  20000 100000

When the processes in the group have used their quota, the kernel doesn't run them again until the next period begins, even if CPUs are idle.
This is called throttling.
A throttled service is slow, but CPU graphs for the host can look healthy: here a busy loop uses only 18% of a CPU, and the rest of the CPU is idle.

The group's ``cpu.stat`` file shows how often it was throttled:

.. code-block:: console

  root@opsschool # cat /sys/fs/cgroup/system.slice/quotademo.service/cpu.stat
  usage_usec 1023799
  user_usec 1023799
  system_usec 0
  nice_usec 0
  nr_periods 50
  nr_throttled 50
  throttled_usec 3938134
  nr_bursts 0
  burst_usec 0

``nr_periods`` is the number of enforcement periods that have elapsed, and ``nr_throttled`` is the number of times the group has been throttled.
``throttled_usec`` is the total time it spent throttled, in microseconds.
If ``nr_throttled`` grows steadily, the quota is limiting the group.

``CPUWeight=`` is a different kind of control.
It sets a group's share of CPU time relative to other groups, and the default is 100.
It only matters when the CPUs are busy: a group with a low weight can still use an idle CPU.

``systemctl show -p CPUQuotaPerSecUSec <unit>`` shows the quota systemd has set, and ``systemctl cat`` shows which unit file or drop-in set it.
Memory has similar controls, such as ``MemoryMax=``; see the ``systemd.resource-control(5)`` manual page.
