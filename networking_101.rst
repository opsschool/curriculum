Networking 101
**************

This chapter should provide enough knowledge on networking and linux, to be able
to connect a linux server to a network and troubleshoot basic network related
problems.

OSI model
=========

The OSI model describes 7 layers of networking that enable users and
applications to communicate with each other on separate systems.

#. Physical layer
   This is the lowest layer and represents the physical wiring between systems.
#. Data link layer
   Ethernet lives at this layer. It provides the basic protocol for communicating
   between two points on the physical layer.
#. Network layer
   It's great that two media devices (ie, network cards) can talk to each other
   on the data link layer, but what about everything else?
   The network layer is where the Internet Protocol lives. The various protocols
   on this layer provide the means for machines to talk to one another.
#. Transport layer
   Finally, we can talk about server-to-server communications.
   When two systems communicate, they send data back and forth. Depending on what
   the data is, there is a choice of which protocols you want to use.
   TCP is common when you need to make sure every packet arrives correctly.
   UDP is best when some packet loss is OK as long as data gets to the other end
   quicky (eg, video streaming or real time games).
   ICMP is used by ``ping`` and other network-activity related things.
   These are all transport later protocols. They use the common Internet Protocol
   and add to it.
#. Session layer
#. Presentation layer
#. Application layer


IP Addressing
=============

IPv4
----

IPv6
----


Static routing
==============


Private address space (RFC1918)
===============================


TCP vs UDP
==========
<discuss 3 way handshake here>


Subnetting, netmasks and CIDR
=============================


NAT
===


Practical networking
====================

Cat5e, Cat6, Cat6a
------------------

Fiber
-----

LC vs SC
^^^^^^^^

SFP, SFP+, X2, QSFP
^^^^^^^^^^^^^^^^^^^

Twinax
------

Multimode vs Single Mode vs OM{3,4}
-----------------------------------

