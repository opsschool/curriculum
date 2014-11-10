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

Instead of passing a packet to a specific ip address however, this configuration states that packets intended for the 10.10.20.0/24 network should be forwarded to the host directly connected to Serial 0/0/1.

Static routes do not change when the network changes but can be useful in cases where the network is small enough that it outweighs the cost of dynamic routing.
It is often used in tandem with dynamic routing to specify a default route in case a dynamic route is unavailable.

Dynamic routing protocols (RIP, OSPF, BGP, EIGRP, IS-IS)
------------------------------------------

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
