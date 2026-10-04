Security 201
************

Centralised accounts
====================

LDAP and Kerberos
-----------------

Active Directory
----------------

NIS, NIS+, YP, Hesiod
---------------------


Firewalls and packet filters
============================

host vs network
---------------

"Crunchy outer shell, soft center is bad"

In general terms, implementing defense-in-depth strategies is always a sensible 
practice.  The concept in which multiple layers of security controls (defense) 
are placed throughout an information technology (IT) system.  Its intent is to 
provide redundancy in the event a security control fails or a vulnerability is 
exploited.

Implementing a firewall on the network **and** host-based packet filters 
provides defense-in-depth layers to your infrastructure.  In today's landscape, 
the firewall aspect could be a service like ec2 security groups with host-based 
packet filters such as iptables, Windows Firewall (or WFAS).  Although this can 
add additional complexity to deployment, that is not a reason to not implement 
it where appropriate.

The defense-in-depth concept is mostly regarded in terms of attack and 
compromise, however in ops it also safeguards **us** as everyone makes mistakes.  
Sometimes, we ourselves or our colleagues are the point of failure.

Strange as it may seem, people often make the mistake of disabling "rules" when 
something is not working and they cannot figure out why.  The *just checking* 
test.  This is always the first mistake.  In real world operations these things 
do happen, whether it is a *just checking* mistake, an incorrect configuration 
or action, sometimes we make serious mistakes, to err is human.

In all these situations using a firewall/other and host-based packet filters 
comes to the fore:

- It protects us, the people working on the systems, from ourselves.
- It protects the organisation's assets in the event of a failure or compromise 
  at either layer, whether it be user error or a systematic failure.
- It keeps us proficient and *in-hand* on both the network specific 
  security implementation, the host-based security practices and the 
  applications related to their management.
- It is good practice.

An old skool, real world example:
---------------------------------

- You have a MSSQL server running on Windows Server 2000 (no "firewall" back 
  then & ip filters are not enabled) - it has both private and public network 
  interfaces.
- The MSSQL server has a public NIC because it has run replication with your 
  customer's MSSQL server sometimes for dev purposes and catalog updates.
- You have rules on the diversely routed, mirrored NetScreen NS1000 firewalls 
  that allows port 1433 between the MSSQL servers only on the public interface.
- Your colleague has an issue that cannot be resolved and quickly just sets the
  firewall/s to "Allows from all", the *just checking* test.
- **Immediate** result - network unreachable.
- SQL Slammer had just arrived and proceeded to gobble up 2Gbps of public T1 
  bandwidth.
- All points into the network and the network itself are saturated until you 
  debug the issue via the serial port on the firewall and figure out what 
  happened.

The synopsis is that a practice of disabling rules was implemented, as it had 
always been the last line in debugging.  It was a practice that had been done 
many times by the engineers in the organisation at the time with no "apparent" 
consequences in the past.  This time differed in that the MSSQL server had a 
public interface added to allow for replication with the customer **and** 
SQL Slammer was in the wild.

If the MSSQL server had ip filtering enabled as well the situation would have 
been mitigated.  Needless to say, it was the last time "Allow from all" debug 
strategy was ever implemented.  However it is useful to note, that because this
was done all the time, the engineer in question did not even tie the action of 
"Allow from all" and the network becoming unreachable together at the time, 
because the action previously had never resulted in the outcome that was 
experienced in this instance.

Stateful vs stateless
---------------------

A stateless filter looks at each packet on its own.
It decides by what is in the packet: addresses, ports and flags.
To allow the replies to your outgoing connections, you need rules for the replies too.

A stateful filter remembers connections.
It can use connection state to allow subsequent packets through explicit firewall rules.
On Linux, the kernel's connection tracking (``nf_conntrack``) records this state; firewall rules decide whether to accept packets.
iptables uses it through the ``conntrack`` match:

.. code-block:: console

    root@opsschool # iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT
    root@opsschool # iptables -A INPUT -m conntrack --ctstate INVALID -j DROP

Any rule that uses the ``conntrack`` match turns on connection tracking for all traffic.

The connection tracking table
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Each tracked connection is an entry in a table of fixed size.
Every TCP connection takes an entry, and so does every UDP flow, such as a DNS lookup.
Entries can remain after a connection closes.
For example, conntrack's ``TIME_WAIT`` timeout defaults to 120 seconds; other closing states have different timeouts, and these values are configurable.

To see how full the table is:

.. code-block:: console

    root@opsschool # cat /proc/sys/net/netfilter/nf_conntrack_count
    1342
    root@opsschool # cat /proc/sys/net/netfilter/nf_conntrack_max
    262144

``conntrack -L`` lists the entries.
``conntrack -S`` shows counters for each CPU, including ``drop`` and ``insert_failed``.

When the table is full, the kernel drops packets that need a new entry, and logs this:

.. code-block:: none

    nf_conntrack: table full, dropping packet

Existing connections keep working.
New connections fail, but only some of them, and only while the table is full.
Clients send again, wait, and time out, so it looks like a slow or unreliable network.
The service sees nothing, because the dropped packets never reach it.
It often happens only at busy times, when there are the most connections.

To raise the limit at once, choose a value above the current limit; for example:

.. code-block:: console

    root@opsschool # sysctl -w net.netfilter.nf_conntrack_max=524288

To keep the new limit after a reboot, set it in a file in ``/etc/sysctl.d/``.
Each entry uses a few hundred bytes of kernel memory, so a large table is cheap.
Size it for your busiest time, with room to spare.

IPTables: Adding and deleting rules
-----------------------------------

pf: Adding and deleting rules
-----------------------------


Public Key Cryptography
=======================

.. todo::
   What is PKI? What uses it? Why is it important?

Using public and private keys for SSH authentication
----------------------------------------------------


Two factor authentication
=========================


Building systems to be auditable
================================

Data retention
--------------

Log aggregation
---------------

Log and event reviews
---------------------

Role accounts vs individual accounts
------------------------------------


Network Intrusion Detection
============================


Host Intrusion Detection
=========================


Defense practices
=================


Risk and risk management
========================


Compliance: The bare minimum
============================

What is compliance and why do you need it?

What kinds of data can't you store without it?

Legal obligations


Dealing with security incidents
===============================


ACLs and extended attributes (xattrs)
=====================================


SELinux
=======


AppArmor
========

AppArmor limits what a program can do, even when it runs as root.
Each confined program has a profile: a list of the files it may read and write, and what else it may do, such as use the network.
Anything that is not in the profile is denied.
These rules apply as well as the normal file permissions.
So a file can have the right owner and mode, and the program still can't open it.
This is called mandatory access control.

Ubuntu and Debian use AppArmor.
Red Hat and similar systems use SELinux, which does the same job in a different way.

To see which programs are confined:

.. code-block:: console

    root@opsschool ~# aa-status
    apparmor module is loaded.
    31 profiles are loaded.
    29 profiles are in enforce mode.
       /usr/sbin/rsyslogd
       tcpdump
    ...

Profiles are in ``/etc/apparmor.d``.
Most are named after the program's path, with ``.`` in place of ``/``: the profile for ``/usr/bin/tcpdump`` is ``/etc/apparmor.d/usr.bin.tcpdump``.
A profile is in one of two modes.
In ``enforce`` mode, AppArmor denies anything the profile doesn't allow, and logs it.
In ``complain`` mode, it only logs it.

When AppArmor denies something
------------------------------

The program usually gets "Permission denied", the same error as for a normal permission problem.
The kernel log shows the real reason.
For example, Ubuntu's profile for ``tcpdump`` lets it write capture files whose names end in ``.pcap``, but not other names:

.. code-block:: console

    root@opsschool # tcpdump -i eth0 -w /srv/captures/web-01
    tcpdump: /srv/captures/web-01: Permission denied
    root@opsschool # journalctl -k | grep DENIED
    audit: type=1400 audit(1759581234.123:42): apparmor="DENIED" operation="mknod" profile="tcpdump" name="/srv/captures/web-01" pid=2817 comm="tcpdump" requested_mask="c" denied_mask="c" fsuid=0 ouid=0

The line names the profile, the file, and what the program asked for: ``c`` is create, ``r`` is read, ``w`` is write.

Changing a profile
------------------

Don't edit the main profile file, because a package upgrade will replace it.
Most profiles include a file for local additions, with the same name, in ``/etc/apparmor.d/local/``.
Add your rules there, then reload the profile:

.. code-block:: console

    root@opsschool # echo '/srv/captures/** rw,' >> /etc/apparmor.d/local/usr.bin.tcpdump
    root@opsschool # apparmor_parser -r /etc/apparmor.d/usr.bin.tcpdump

``aa-complain`` puts a profile in complain mode, where violations of allow rules are logged instead of blocked, although explicit ``deny`` rules remain enforced.
``apparmor_parser -R`` removes the profile and turns off its protection.
They are useful to check whether AppArmor causes a problem, but they are not a fix.


Data placement
==============
Eg, local vs cloud, the implications, etc


Additional reading
==================
Ken Thompson, Reflections on Trusting Trust:
http://dl.acm.org/citation.cfm?id=358210
