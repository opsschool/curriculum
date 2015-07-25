Networking 201
**************

VLANs, 802.1q tagging
=====================
Sometimes it is advantageous to break up a single Layer 2 broadcast domain into multiple Layer 2 broadcast domains.
As we had previously discussed, Layer 2 is addressed by MAC addresses, which are flat address space (non-routable) constructs; this creates a Layer 2 broadcast domain, also called a "segment" or "LAN" (local area network.) You can then map one or more IP networks (Layer 3 addressing) onto a Layer 2 segment.
If you map more than one Layer 3 network to a segment, you can get into some headaches, such as having to address multiple IP addr's on various host and router interfaces (one for each network) with the attendant confusion, and also that broadcasts produced by each IP network (i.e. ARP requests) multiply, and bandwidth is consumed.
Therefore, most networking professionals agree that it's best to only map one IP network to one Layer 2 segment.

So, if multiple Layer 3 networks need to be used, instead of using multiple switches for each Layer 2 segment, it is possible to divide a single VLAN-capable switch into two or more "virtual" switches.
This is possible due to the concept of :term:`VLAN`, which stands for "Virtual LAN".

To communicate between the Layer 3 networks mapped to their individual VLANs, packets need to pass through a router or other Layer 3 device.
If one has a switch with many VLANs implemented on it, either of these two methods could work:

1. Devote a switch (and respective router) interface per VLAN, with a single link for each VLAN.
   This is the simplest method, but the most wasteful of (expensive) switch and router interfaces.

2. "Trunk" multiple VLANs over one physical link, which only requires a single interface per side (switch to router, or server to switch.)
   This method multiplexes multiple VLANs over the one link by "tagging" the frame with a 4-byte VLAN identifier field which is inserted after the destination MAC field in the Ethernet header.
   This tagging format is officially defined by the IEEE's 802.1q protocol.
   The tag field insertion happens at the sending trunk interface (which is a part of the sending station's VLAN), and the tag is stripped of by the receiving trunk interface, and the frame placed on the proper VLAN on the receiving device.

The VLAN identifier must be between 1 and 4094 inclusive.
The 802.1q tag also provides for a 3-bit priority field and 1-bit discard eligibility field which can be used for QoS applications.
The trunking and priority functions of 802.1q can be enabled and disabled separately.

Spanning Tree
=============
As networks and broadcast domains grow and become more complex, it becomes much easier to intentionally or unintentionally create loops- multiple network paths between one or more switches.
These loops may lead to a infinite forwarding of Ethernet packets, causing switches to be overwhelmed by an ever-increasing amount of packets being forwarded to all other switches in the broadcast domain.
To address this concern, Spanning Tree Protocol (STP) was invented as a standard protocol used by Ethernet switches, designed to prevent such loops in broadcast domains.

To disable loops, each switch that implements Spanning Tree Protocol (STP / IEEE 802.1D) will participate in electing a "root" device where the Spanning Tree is calculated from.
Once the root is chosen, all other participating switches will calculate their least-cost path for each of its ports to the root device.
This is usually a port that traverses the least number of segments between the source and the root.
Once each device has chosen its least-cost port, or Root Port (RP), it calculates the least-cost path for each network segment.
By doing this, each switch obtains a Designated Port (DP), or a port on which it should expect to accept traffic to forward through its RP.
Upon calculating these values, all other paths to the root device are disabled- marking them as Blocked Ports (BP).

A major benefit to using STP to prevent network loops is that network administrators can intentionally connect redundant paths for fault tolerance without worry of causing an infinite loop during normal operation.
Should a path in a broadcast domain fail, each participating switch will recalculate the DP and RP, modifying the status of BPs if necessary.
This has the effect of repairing the broken path by routing traffic through segments around the failure.

Routing
==============
When a network is subdivided into multiple Layer 2 broadcast domains, Layer 3 addressing enables hosts to communicate with another host in another broadcast domain.

The process of forwarding packets from one Layer 2 subdomain to another Layer 2 subdomain using Layer 3 addresses is called routing.

A Layer 3 device or router is the device responsible for this function.
Routers may use many different ways to forward packets but these methods can be categorized into two types.

Static Routing
--------------
The method of using manually configured routes is called static routing.

Typically, there are three pieces of information that are needed to specify a static route:

1. the destination subnet

2. the latter's subnet mask and

3. the next hop host or outgoing interface.

Take for example the following static route configuration in Linux:

.. code-block:: console

   # ip route add 10.10.20.0/24 via 192.168.2.1

The above command means that packets intended for hosts in the 10.10.20.0/24 network must be forwarded to the host with ip address 192.168.2.1.

Below is another example in Cisco IOS:

.. code-block:: console

   Router(config)# ip route 10.10.20.0 255.255.255.0 Serial0/0/1

Instead of passing a packet to a specific ip address however, this configuration states that packets intended for the 10.10.20.0/24 network should be forwarded to the host directly connected to Serial 0/0/1 whose ip address can be anything.

Static routes do not change when the network changes but can be useful in cases where the network is small enough that it outweighs the cost of dynamic routing.
It is often used in tandem with dynamic routing to specify a default route in case a dynamic route is unavailable.

Dynamic routing protocols (RIP, OSPF, BGP, EIGRP, IS-IS)
--------------------------------------------------------
For a small network, manually configuring routes will work just fine.
As a network grows larger, however, doing this can be very arduous if not infeasible.
Dynamic routing solves this problem by programmatically building a routing table.

ACLs
====

Network Bonding (802.3ad / LACP link aggregation)
=================================================

IOS switch configuration
========================

Cisco devices use a command line-based configuration interface which can be accessed via RS-232, Telnet or SSH.
The command line syntax is distinctive and has been loosely copied by many other vendors, so some familiarity with the Cisco-style configuration can go a long way.

When connecting via RS-232, use 9600 baud, 8 data bits, 1 stop bit, no parity, no flow control.

You can abbreviate keywords on the command line without having to press Tab, so long as they are unambiguous.
You can also use Tab as you would in a UNIX shell.

While using the command line, you can type '?' at any time and receive context-sensitive help.

The command line is mode-based. The prompt tells you of your current mode.
Awareness of your current mode is key to efficient operation of the command line.
Here are some examples:

=========================  ===================================================
Prompt                     Meaning
=========================  ===================================================
``hostname>``              'EXEC' mode, non-privileged access
                           (like ``$`` on Linux)
``hostname#``              'EXEC' mode, privileged access
                           (like ``#`` on Linux)
``hostname(config)#``      Global configuration mode
``hostname(config-if)#``   Interface configuration mode
=========================  ===================================================

EXEC mode allows you to execute imperative commands such as ``ping``.
The configuration mode allows you to add and remove configuration statements.

The following commands move you between modes:

=========================  ======  ===========================================
Command                    Mode    Effect
=========================  ======  ===========================================
``enable``                 EXEC    Become privileged (like ``su`` on Linux)
``conf t``                 EXEC    Enter ``(config)`` mode.
``int <interface>``        config  Enter ``(config-if)`` mode for an
                                   interface.
``exit``                   Any     Leave the current mode or logout.
=========================  ======  ===========================================

Other useful commands include:

=========================  ======  ===========================================
Command                    Mode    Effect
=========================  ======  ===========================================
``show ...``               EXEC    The subcommands of this command provide
                                   access to all available information.
``show run``               EXEC    Show the current configuration.
``show ip int brief``      EXEC    Show all interface names, IPv4 addresses,
                                   and their status.
``wr``                     EXEC    Save the current configuration.
``ping <host>``            EXEC    Ping. '!' means a response,
                                   '.' means a timeout.
``no <command>``           config  Delete a configuration statement in config
                                   mode.
``do <command>``           config  Execute a command in EXEC mode from config
                                   mode.
=========================  ======  ===========================================

Note that configuration changes become active as soon as they are made.
The ``show run`` command shows the configuration which is currently in effect.
This is not saved until you execute ``wr``.

GRE and other tunnels
=====================

Multi-homed hosts
=================

Similarities and differences between IPv4 and IPv6 networking
=============================================================

Implications of dual-stack firewalls (especially under Linux)
=============================================================

Multicast uses and limitations
==============================

Latency vs. Bandwidth
=====================

http://www.stuartcheshire.org/rants/Latency.html

VPNs
====

IPSec
-----

SSL
---

Troubleshooting layer 1 problems
================================

Layer one problems (physical) deserve their own section, because of how cryptic
they can seem without a solid understanding of the causes and symptoms.
There are several data points that fall under layer one.

Most layer problems are usually traced back to one of the following causes:

* Bad cabling
* Mismatched duplex/speed
* Electrical interference (EMI)
* Network congestion

To find the metrics below, use either `ip -s link`, or `ethtool -S`.

Example output of `ip -s link show eth0`:

.. code-block:: console

  user@opsschool ~$ ip -s link show eth0
  2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
      link/ether 04:01:1c:dc:82:01 brd ff:ff:ff:ff:ff:ff
      RX: bytes  packets  errors  dropped overrun mcast
      150285334  221431013 0       0       0       0
      TX: bytes  packets  errors  dropped carrier collsns
      999121259  42410636 0       0       0       0

Here we can see that the link is up (`state UP`), the MAC address
(`04:01:1c:dc:82:01`), bytes, packets, errors, drops, overruns, multicast,
carrier, and collisions, for both RX (receive) and TX (transmit).

One thing you may notice is that overrun and mcast are only on RX, and
carrier/collisions are only on TX.
This is because overruns and multicast only affect the receiving side, while carrier
and collisions only affect the transmitting side.
The reason for this will be made clear when described below.

Example output of `ethtool eth0`:

.. code-block:: console

  user@opsschool ~$ sudo ethtool eth0
  Settings for eth0:
    Supported ports: [ TP ]
    Supported link modes:   10baseT/Half 10baseT/Full
                            100baseT/Half 100baseT/Full
                            1000baseT/Full
    Supported pause frame use: No
    Supports auto-negotiation: Yes
    Advertised link modes:  10baseT/Half 10baseT/Full
                            100baseT/Half 100baseT/Full
                            1000baseT/Full
    Advertised pause frame use: No
    Advertised auto-negotiation: Yes
    Speed: 1000Mb/s
    Duplex: Full
    Port: Twisted Pair
    PHYAD: 0
    Transceiver: internal
    Auto-negotiation: on
    MDI-X: Unknown
    Supports Wake-on: umbg
    Wake-on: d
    Current message level: 0x00000007 (7)
               drv probe link
    Link detected: yes

As we can see, ethtool gives a lot of information, and with additional flags,
there's even more.
ethtool is especially useful for looking into layer 1 configuration, due to
how much information it can provide.
A particularly useful option is `-S`, which shows all the same metrics as
the `ip` command mentioned above, plus many, many more:

.. code-block:: console

  user@opsschool ~$ sudo ethtool -S eth0
  NIC statistics:
      rx_packets: 20831
      tx_packets: 11160
      rx_bytes: 14654723
      tx_bytes: 3637509
      rx_broadcast: 0
      tx_broadcast: 9
      rx_multicast: 0
      tx_multicast: 11
      rx_errors: 0
      tx_errors: 0
      tx_dropped: 0
      multicast: 0
      collisions: 0
      rx_length_errors: 0
      rx_over_errors: 0
      rx_crc_errors: 0
      rx_frame_errors: 0
      rx_no_buffer_count: 0
      rx_missed_errors: 0
      tx_aborted_errors: 0
      tx_carrier_errors: 0
      tx_fifo_errors: 0
      tx_heartbeat_errors: 0
      tx_window_errors: 0
      tx_abort_late_coll: 0
      tx_deferred_ok: 0
      tx_single_coll_ok: 0
      tx_multi_coll_ok: 0
      tx_timeout_count: 0
      tx_restart_queue: 0
      rx_long_length_errors: 0
      rx_short_length_errors: 0
      rx_align_errors: 0
      tx_tcp_seg_good: 304
      tx_tcp_seg_failed: 0
      rx_flow_control_xon: 0
      rx_flow_control_xoff: 0
      tx_flow_control_xon: 0
      tx_flow_control_xoff: 0
      rx_long_byte_count: 14654723
      rx_csum_offload_good: 0
      rx_csum_offload_errors: 0
      alloc_rx_buff_failed: 0
      tx_smbus: 0
      rx_smbus: 0
      dropped_smbus: 0

Common Network metrics
----------------------

RX/TX errors
^^^^^^^^^^^^

These counters appear to be aggregates of various other counters, and information
on exactly what's included in them is sparse.
An increase of these counters is an indication that *something* is wrong, but
they're too general to determine the exact cause.
If you see these increasing, it will be more beneficial to check `ethtool -S`
and see exactly which counters are increasing.

Cyclic Redundancy Check (CRC)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

CRC is a sort of checksum.
CRC errors can be a symptom for many different issues, such as a bad cable,
noisy lines, and more.

Drops
^^^^^

Drops occur when the network link is simply too saturated.
The solution is to increase the bandwidth available, perhaps by upgrading the
network link, or implementing some form of traffic limitation such as
Quality-of-Service (QoS).
Traffic limitation/shaping is out of the scope of this section, and should be
considered an advanced topic.
You can verify the saturation by checking your graphs for RX/TX bytes and see how
much data is going through.

Overruns
^^^^^^^^

Overruns occur when the network device is receiving data faster than the kernel
driver can process it.
Overruns sound similar to drops, but overruns have to do with the kernel's ability
to process data, rather than the network link itself.
The solution is either to replace the network device with a more capable one, or
lower the amount of data coming in to the network device.

Carrier
^^^^^^^

Carrier errors occur when the link signal is having issues.
In the absence of collisions, this can usually be attributed to a bad cable
or network device.
If collisions are also increasing, check the duplex settings.

Collisions
^^^^^^^^^^

Collisions occur when there is a duplex or speed mismatch.
Verify duplex and speed at both ends of the cable (at the switch and at the server).
Auto-negotiation being enabled on one end but not on the other is often a
cause of duplex/speed mismatch.
Some devices don't play nicely with each other when auto-negotiation is
enabled on both sides, so be sure to verify what auto-negotiation resulted in.

Electrical interference
-----------------------

The symptoms of electrical interference can be seen by an increase of CRC errors.
Electrical interference can be caused by any number of things: power cables nearby,
mechanical systems (such as HVAC units), etc.
Since a visual inspection will often reveal electrical interferance culprits easier
than checking metrics, it's simpler to do a visual inspection first if you suspect
this.
The problem can sometimes be solved/alleviated by replacing the cable with STP
(Shielded Twisted Pair), but it's usually better to just relocate the network cable.
Only copper cabling is suspectible to EMI.
Fiber is immune to EMI, as signals are transmitted as light instead of current.

Testing copper cabling
----------------------

A link tester is an inexpensive device made by many different vendors, for the
purpose of verifying that a cable's functionality.
Link testers only verify a small amount of criteria, mainly that there are no
opens or shorts.
A link tester, however, is different from a cable certifier, which will verify
full compliance with TIA/EIA and ISO cable standards, including metrics for
crosstalk allowance, cable length, current loss, and many others.

If you need to test cable used in low-impact environments (eg, wall jack to
desktop), a link tester is usually sufficient.
If you need to test cable used in high-impact environments (eg, datacenter),
opt for a cable certifier.
Reputable cable vendors will have already run certification tests on cable,
so this is only useful if you are routinely making your own cable (which
is not advised under most circumstances).

Many vendors sell great testers, with Fluke Networks being one of the primary
vendors.

Fiber errors
------------

Issues with fiber cable generally fall into two categories:

* Dirty/Scratched fiber
* Bad optic/transceiver

The symptoms for both are very similar: intermittent RX/TX and CRC errors
indicate a dirty or scratched fiber, while persistent RX/TX and CRC errors
indicate a bad optic/transceiver.
