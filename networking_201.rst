Networking 201
**************

VLANs, 802.1q tagging
=====================
Sometimes it is advantageous to break up a single layer-2 broadcast domain into multiple
layer-2 broadcast domains. The typical use case would be to divide a single VLAN-
capable switch into two or more "virtual" switches. This is possible due to the concept 
of :term:`VLAN`, which stands for "Virtual LAN". Best practices dictate the mapping of
a single IP subnet to a VLAN.

To communicate between VLANs, packets need to pass through a router or other Layer-3
device. If one has a switch with many VLANs implemented on it, either of these two
methods could work:

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
