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

IPv4
====

Internet Protocol Version 4 (IPv4) is the fourth version of the Internet protocol, the first
version to be widely deployed. This is the version of the protocol you're most likely to
encounter, and the default version of the IP protocol in Linux.

Addressing
----------

IPv4 uses a 32-bit address space most typically represented in 4 dotted decimal notation,
each octet contains a value between 0-255, and is separated by a dot. An example
address is below:

    10.199.0.5

There are several other representations, like dotted hexadecimal, dotted octal, hexadecimal,
decimal, and octal. These are infrequently used, and will be covered in later sections.

Subnetting, netmasks and CIDR
-----------------------------

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
-----------------------------------

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

NAT
---

Network Address Translation is a process of a router modifiying IPv4 traffic to include additional information in the packet headers.

NAT is commonly deployed in one of two configurations:

* One To Many

  This method allows many private computers to utilise a single public IP address for accessing the Internet. This can be combined with other methods to allow access to internal services from external devices.
  
* One To One

  This method creates an explicit routing relationship between two addresses, commonly one internal and one external. This allows for traffic destined for one (e.g. a public address) to be routed to the other (e.g. a private address), and also allows for the reverse traffic (e.g. traffic to the Internet from a private network device) to appear to be from that second specified address

Because computers deployed behind a NAT are not directly accessible, several methods have been developed to allow services inside a private network to be accessed from external sites. These include:

* Port Forwarding

  Using Port Forwarding means a router will forward connections in-bound on a particular port to another device inside its network. 

.. todo: Add in STUN / ICE / UPnP / NAT-PMP


IPv6
====

IPv6 is the latest revision of the Internet Protocol, and is intended to replace IPv4. Although 
less commonly deployed, new systems should aspire to be deployed with IPv6 support. IPv6 is now
supported in all modern Linux distributions, Windows, OS X, and most newer networking hardware.

Addressing
----------

IPv6 uses a 128-bit address space, giving significantly more addresses than IPv4. Addresses are represented as eight colon separated groups of four hexadecimal digits, for example:

    2001:41c8:0020:60e:0000:004c:0010

In the IPv6 addressing scheme, leading zeros can be omitted, and sections of consecutive zeros can be omitted, leaving the above address as a shortened:

    2001:41c8:20:60e::4c:10

IPv6 has been designed to avoid the requirement for NAT (as above). As part of this, the idea of "private ranges" has a different meaning within IPv6 - :rfc:`4193` covers the definition of IPv6 unicast addresses for use within a private network which are not intended to be globally routed.

Subnetting, netmasks and CIDR
-----------------------------

CIDR was originally introduced into IPv4 to attmept to improve the efficiency of utilisation of the small address space available. As this is less of a concern in IPv6, subnetting is handled differently.

The practical implementations of this are that:

 * An (:rfc:`4291`) compliant subnet uses IPv6 addresses with 64 bits for the host portion, meaning it will have a /64 routing prefix
 * All addresses within a range are useable as there are no reserved addresses for broadcast traffic as there is with IPv4
 * Customer site allocations are recommended to be 48-bit at a minimum


TCP vs UDP
==========

Both TCP :rfc:`793` and UDP :rfc:`768` provide data transfer between processes 
through ports. These process ports can be on the same computer or separate 
computers connected by a network. TCP and UDP are the two most commonly referenced
Transport Layer protocols used by the Internet Protocol. They have very different 
use cases - the choice of protocols to use is often based on whether the risk of losing 
packets in real-time without immediate alerting is acceptable. In some cases 
UDP may be acceptable, such as video or audio streaming where programs can 
interpolate over missing packets. However, TCP will be required due to its 
reliable delivery guarantee in systems that support banking or healthcare.


UDP is a less feature-rich, it does its work with a header that only contains a source port, 
destination port, a length, and a checksum. TCP provides its capabilities 
by sending more header data, more packets between ports and performing more 
processing


TCP
---

TCP - Transmission Control Protocol - is a connection oriented protocol. It provides reliability, flow control and connection ensurance by implementing error checking, checksumming and packet ordering to ensure that data sent is received, and in the order with the same data with which it was sent. (See Example 1 below)

Because of these extra checks, TCP is generally considered slower than UDP, especially over higher latency connections. Because of this it can be considered be unsuitable for file transfers over long slow links. 

.. todo: Add in TCP handshake

UDP
---

UDP - User Datagram Protocol - is a simple, connectionless protocol. It is used commonly in media and voice delivery where low latency is preferable to guarantee of delivery. It has lower overheads because there is no error checking [#checksum]_, however this means that applications must tolerate this. 

UDP requires less header data in the individual 
packets and requires fewer packets on the network to do its work. UDP does no 
bookkeeping about the fate of the packets sent from a source. They could be 
dropped because of a full buffer at a random router between the source and 
destination and UDP wouldn't account for it in itself (other monitoring systems
can be put in place to do the accounting, however that is beyond the UDP
protocol). 


UDP is commonly used for DNS, SNMP, and VoIP.

.. [#checksum] UDP does implement checksumming - but this is only relevant if the packet reaches its destination.



Examples
---

* Example 1
 
  The TCP protocol requires upfront communication and the UDP protocol does 
  not.  TCP requires an initial connection, known as the "three way handshake",
  in order to begin sending data. That amounts to one initial packet sent 
  between ports from initiator of the communication to the receiver, then 
  another packet sent back, and then a final packet sent from the initiator 
  to the receiver again. All that happens before sending the first byte of
  data. In UDP the first packet sent contains the first byte of data.

* Example 2 
  
  TCP and UDP differ in the size of their packet headers. The TCP header is 
  20 bytes and the UDP header is 8 bytes. For programs that send a lot of 
  packets with very little data, the header length can be a large percentage of
  overhead data (e.g. games that send small packets about player position and state). 

Static routing
==============


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


