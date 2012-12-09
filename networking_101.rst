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

Internet Protocol Version 4 (IPv4) is the fourth version of the Internet protocol, the first
version to be widely deployed. This is the version of the protocol you're most likely to
encounter, and the default version of the IP protocol in Linux.

IPv4 uses a 32-bit address space most typically represented in 4 dotted decimal notation,
each octect contains a value between 0-255, and is seperated by a dot. An example 
address is below:

    10.199.0.5 

There are several other representations, like dotted hexidecimal, dotted octal, hexidecimal, 
decimal, and octal. These are infrequently used, and will be covered in later sections. 



IPv6
----



TCP vs UDP
==========
<discuss 3 way handshake here>


Subnetting, netmasks and CIDR
=============================
A subnet is a logical devision of an IP network, and allows the host system to identify which 
other hosts can be reached on the local network. The host system determines
this is by the application of a routing prefix. There are two typical representations of this
prefix: a netmask and CIDR. 

Netmasks typically appear in the dotted decimal notation, with values between 0-255 in each 
octet. These are applied as bitmasks, and numbers at 255 mean that this host is not reachable.
Netmask can also be refered to as a Subnet Mask and these terms are often used interchangeably. An 
example IP Address with a typical netmask is below:

============= ===============
IP Address    Netmask   
============= ===============
192.168.1.1   255.255.255.0 
============= ===============

CIDR notation is a two-digit representation of this routing prefix. Its value can range
between 0 and 32. This representation is typically used for networking equipment. Below
is the same example as above with CIDR notation:

============= ===============
IP Address    CIDR   
============= ===============
192.168.1.1   /24 
============= ===============

Private address space (:rfc: 1918)
===============================

Certian ranges of addresses were reserved for private networks. Using this address space
you cannot communicate with public machines without a NAT gateway or proxy. There are 
three reserved blocks:

============== ===================== =============== ==============
First Address  Last Address          Netmask         CIDR
============== ===================== =============== ==============
10.0.0.0       10.255.255.255        255.0.0.0       /8
172.16.0.0     172.31.255.255        255.240.0.0     /12
192.168.0.0    192.168.255.255       255.255.0.0     /16
============== ===================== =============== ==============


Static routing
==============


NAT
===


Practical networking
====================

Cat5e, Cat6, Cat6a
------------------

Cat5e, Cat6, and Cat6a are all coper transport mediums. They use twisted pair 
wiring, relying on the twist with differential signaling to prevent noise. This is the most
common form of cabling for connecting computers in a network. 

Fiber
-----
Fiber is a generic term that refers to optical transport mediums. It comes in several types,
all of which look identical but are generally incompatible.

Multimode vs Single Mode vs OM{3,4}
-----------------------------------
Multimode fiber is a less expensive fiber optic cable, that is typically useable with lower
cost optical components. Depending on the application and bandwidth required, multimode fiber
can have a range up to 2000 meters, but as low as 33 meters. It is very common to see it
used for building backbones, and system to switch applications. 

LC vs SC
^^^^^^^^

LC and SC connectors are the two most common type of fiber connectors. 

LC is also known as a Lucent Connector. They are typically used for high-density applications, and are
the type of connector used on SFPs or XFPs. Typcially the connector is packaged in a duplex configuration
with each cable side by side. 

SC connectors are also know as Subscriber Connector, Square Connector, or Standard Connector. This is the type
of connector typically used in the telcom industry. They have a larger form factor than the LC connectors, and 
are often found in single and duplex configurations.


SFP, SFP+, X2, QSFP
^^^^^^^^^^^^^^^^^^^

Twinax
------


