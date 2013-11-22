Networking 201
**************

VLANs, 802.1q tagging
=====================
Sometimes it is advantageous to break up a single Layer 2 broadcast domain into
multiple Layer 2 broadcast domains. As we had previously discussed, Layer 2 is
addressed by MAC addresses, which are flat address space (non-routable)
constructs; this creates a Layer 2 broadcast domain, also called a "segment" or
"LAN" (local area network.) You can then map one or more IP networks (Layer 3
addressing) onto a Layer 2 segment. If you map more than one Layer 3 network to
a segment, you can get into some headaches, such as having to address multiple
IP addr's on various host and router interfaces (one for each network) with the
attendant confusion, and also that broadcasts produced by each IP network (i.e.
ARP requests) multiply, and bandwidth is consumed. Therefore, most networking
professionals agree that it's best to only map one IP network to one Layer 2
segment.

So, if multiple Layer 3 networks need to be used, instead of using multiple
switches for each Layer 2 segment, it is possible to divide a single VLAN-
capable switch into two or more "virtual" switches. This is possible due to the
concept of :term:`VLAN`, which stands for "Virtual LAN".

To communicate between the Layer 3 networks mapped to their individual VLANs,
packets need to pass through a router or other Layer 3 device. If one has a
switch with many VLANs implemented on it, either of these two methods could work:

  1. Devote a switch (and respective router) interface per VLAN, with a single link for
     each VLAN. This is the simplest method, but the most wasteful of (expensive) switch and
     router interfaces.

  2. "Trunk" multiple VLANs over one physical link, which only requires a single interface
     per side (switch to router, or server to switch.) This method multiplexes multiple
     VLANs over the one link by "tagging" the frame with a 4-byte VLAN identifier field
     which is inserted after the destination MAC field in the Ethernet header. This tagging
     format is officially defined by the IEEE's 802.1q protocol. The tag field insertion
     happens at the sending trunk interface (which is a part of the sending station's VLAN),
     and the tag is stripped of by the receiving trunk interface, and the frame placed on
     the proper VLAN on the receiving device. 

Spanning Tree
=============

Static Routing
==============

Dynamic routing protocols (RIP, OSPF, BGP)
==========================================

ACLs
====

Network Bonding (802.3ad / LACP link aggregation)
=================================================

IOS switch configuration
========================

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
