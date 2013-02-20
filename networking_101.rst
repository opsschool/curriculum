Networking 101
**************

This chapter should provide enough knowledge on networking to enable a systems
administrator to connect a Linux server to a network and troubleshoot basic
network-related problems. First, we will go over the basics of the 7-layer Open
Systems Interconnection (:term:`OSI`) model, which is a standard framework with which to
implement communication systems. Next, we will delve into each layer of the OSI
model in more detail as it applies to the role of systems administration.

Before any discussion of networking, however, it's important to have a
working knowledge of the numbered Request for Comments (:term:`RFC`) documents
and how they apply to computer networking. These documents describe the
mechanisms for the OSI layer implementations (e.g. TCP, IP, HTTP, SMTP)
and as such are the authoritative source for how computers communicate
with one another.

The RFC Documents
=================

Starting in 1969, the RFC document series describes standards for how computers
communicate. The series gets its name from the RFC process, wherein industry
experts publish documents for the community at large and solicit comments on
them. If the Internet community finds errors in a document, a new, revised
version is published. This new version obsoletes the prior versions. Some
documents, such as the document specifying email messages, have had several
revisions.

The `RFC Editor <http://www.rfc-editor.org/>`_ manages the RFC archive, as well
as associated standards. New documents go to the RFC Editor for publication
[#notquite]_, whether revision or new standard. These documents go through a
standardization process, eventually becoming Internet-wide standards. Many of
the networking protocols discussed in later chapters have RFCs governing their
behavior, and each section should provide information on the relevant RFCs.

.. [#notquite] This is a simplification, as there are actually many standards
  bodies involved in the process. The `RFC Editor Publication Process
  <http://www.rfc-editor.org/pubprocess.html>`_ document explains in full detail.

Important RFCs
--------------

There are a number of RFCs which don't pertain to any specific technology but
which are nevertheless seminal. These documents establish procedure or policy
which have shaped everything after, and as such have a timeless quality.
In some cases, later documents make references to them. This list is given in
increasing numerical order, though is not exhaustive.

* :rfc:`1796`: Not All RFCs are Standards

  This document describes the different kinds of documents in the RFC series.

* :rfc:`2026`: The Internet Standards Process

  This document (and those that update it) describes in detail how RFCs are
  published and how they become Internet standards.

* :rfc:`2119`: Key words for use in RFCs to Indicate Requirement Levels

  This document, referenced in many following RFCs, presents a common vocabulary
  for specifying the relationship between a standard and implementations of that
  standard. It provides keywords that specify how closely an implementation
  needs to follow the standard for it to be compliant.

* :rfc:`5000`: Internet Official Protocol Standards

  This document provides an overview of the current standards documented by the
  RFCs and which RFC is the most recent for each standard. This document is
  regularly updated with the current standards document status.


OSI model
=========

The OSI model describes seven layers of abstraction that enable software
programs to communicate with each other on separate systems. The seven layers
are designed to allow communication to occur between systems at a given level of
abstraction without concern for how the lower levels are implemented. In this
way, more complex protocols can be built on top of simpler ones that can be used
interchangeably without modifying the higher-level code. The job of each layer
is to provide some service to the layer above by using the services provided by
the layer below.

*  Layer 1 - Physical layer

   The physical layer describes the physical connections between devices. Most
   enterprise networks today implement Ethernet at the physical layer, described
   in IEEE 802.3 for wired connections and IEEE 802.11 for wireless networks.

*  Layer 2 - Data link layer

   The data link layer defines the basic protocol for communicating between two
   points on a network that may consist of many intermediate devices and cables,
   possibly spanning a large geographic area. Ethernet defines the data link
   layer in addition to the physical layer, including (Media Access Control
   (:term:`MAC`) addresses that allow hosts to address their data as being
   relevant to one or more other hosts in particular.

*  Layer 3 - Network layer

   The network layer is what allows many "Layer 2" networks to be
   interconnected, forming much larger "Layer 3" networks. It is this layer of
   the OSI model that enables the Internet to exist, using Internet Protocol
   (IP) addressing. IP addressing allows for a logical taxonomy of systems and
   networks built on top of the MAC addresses provided by Ethernet, which are
   more closely tied to the physical hardware. Version 4 of the Internet
   Protocol, most commonly found in production networks, is described in
   :rfc:`791`.

*  Layer 4 - Transport layer

   The transport layer is where things really start to get interesting for the
   systems administrator. It is at the transport layer that the Transmission
   Control Protocol (TCP), User Datagram Protocol (UDP), and Internet Control
   Message Protocol (ICMP) are defined. The TCP and UDP protocols allow data to
   be sent from one system to another using simple "socket" APIs that make it
   just as easy to send text across the globe as it is to write to a file on a
   local disk - a technological miracle that is often taken for granted. The
   ICMP protocol, used by the ubiquitous ``ping`` utility, allows small test
   packets to be sent to a destination for troubleshooting purposes.

*  Layer 5 - Session layer

   The purpose of the session layer is to provide a mechanism for ongoing
   conversations between devices using application-layer protocols. Notable
   "Layer 5" protocols include Transport Layer Security / Secure Sockets Layer
   (TLS/SSL) and, more recently, Google's SPDY protocol.

*  Layer 6 - Presentation layer

   The job of the presentation layer is to handle data encoding and decoding as
   required by the application. An example of this function is the Multipurpose
   Internet Mail Extensions (MIME) protocol, used to encode things other than
   unformatted ASCII text into email messages. Both the session layer and the
   presentation layer are often neglected when discussing TCP/IP because many
   application-layer protocols implement the functionality of these layers
   internally.

*  Layer 7 - Application layer

   The application layer is where most of the interesting work gets done,
   standing on the shoulders of the layers below. It is at the application layer
   that we see protocols such as Domain Name System (DNS), HyperText Transfer
   Protocol (HTTP), Simple Mail Transfer Protocol (SMTP), and Secure SHell
   (SSH). The various application-layer protocols are at the core of a good
   systems administrator's knowledge base.

IP Addressing
=============

IPv4
----

Internet Protocol Version 4 (IPv4) is the fourth version of the Internet protocol, the first
version to be widely deployed. This is the version of the protocol you're most likely to
encounter, and the default version of the IP protocol in Linux.

IPv4 uses a 32-bit address space most typically represented in 4 dotted decimal notation,
each octet contains a value between 0-255, and is separated by a dot. An example
address is below:

    10.199.0.5

There are several other representations, like dotted hexadecimal, dotted octal, hexadecimal,
decimal, and octal. These are infrequently used, and will be covered in later sections.



IPv6
----



TCP vs UDP
==========
<discuss 3 way handshake here>


Subnetting, netmasks and CIDR
=============================
A subnet is a logical division of an IP network, and allows the host system to identify which
other hosts can be reached on the local network. The host system determines
this by the application of a routing prefix. There are two typical representations of this
prefix: a netmask and CIDR.

Netmasks typically appear in the dotted decimal notation, with values between 0-255 in each
octet. These are applied as bitmasks, and numbers at 255 mean that this host is not reachable.
Netmask can also be referred to as a Subnet Mask and these terms are often used interchangeably. An
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

Private address space (:rfc:`1918`)
===================================

Certain ranges of addresses were reserved for private networks. Using this address space
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
the type of connector used on SFPs or XFPs. Typically the connector is packaged in a duplex configuration
with each cable side by side.

SC connectors are also know as Subscriber Connector, Square Connector, or Standard Connector. This is the type
of connector typically used in the telecom industry. They have a larger form factor than the LC connectors, and
are often found in single and duplex configurations.


SFP, SFP+, X2, QSFP
^^^^^^^^^^^^^^^^^^^

Twinax
------


