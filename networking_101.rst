Networking 101
**************

This chapter should provide enough knowledge on networking to enable a systems
administrator to connect a Linux server to a network and troubleshoot basic
network-related problems. First, we will go over the basics of the 7-layer Open
Systems Interconnection (:term:`OSI`) model, which is a standard framework with which to
describe communication system logical layers. Next, we will delve into each layer of the OSI
model in more detail as it applies to the role of systems administration.

Before any discussion of networking, however, it's important to have a
working knowledge of the numbered Request for Comments (:term:`RFC`) documents
and how they apply to computer networking. These documents describe the technical
specifics for every protocol you will run into (e.g., TCP, IP, HTTP, SMTP), and
as such are the authoritative source for how computers communicate
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


OSI 7-layer model (OSI Reference Model)
=======================================

The early history of computer networking was marked by various proprietary
implementations of networking hardware and the protocols that ran on them (such
as for example, IBM's "Systems Network Architecture" (SNA) and Digital Equipment
Corp's "DECnet".) In the mid-1970's, two standards organizations (the
International Organization for Standardization, known as the "ISO", and the
International Telecommunications Union, known as the "ITU") led a push to define
a set of open interconnection standards, which became known as the "Open Systems
Interconnection" (OSI) standard. While the actual protocols they developed did
not become popular and are not in wide use today, the model they came up with
has achieved wide popularity as a way of thinking about the network protocol
stack. It's important to understand that the OSI model is a conceptual model,
and therefore more of a mental framework than a technical framework. It is useful
primarily for troubleshooting and explanation purposes. While the OSI model may
seem esoteric as you're learning it, it will eventually prove invaluable to your
daily work once you have a grasp of the concepts.

The OSI model describes seven layers of abstraction that enable software
programs to communicate with each other on separate systems. The seven layers
are designed to allow communication to occur between systems at a given level,
without concern for how the lower levels are implemented. In this way,
innovation on the protocols in the various layers can be achieved without
modifying any other layer's design and code. The job of each layer is to provide
some service to the layer above by using the services provided by the layer
below.

*  Layer 1 - Physical layer

   The physical layer describes the physical connections between devices. Most
   enterprise networks today implement Ethernet at the physical layer, described
   in IEEE 802.3 for wired connections and IEEE 802.11 for wireless networks.
   Other Layer 1 protocols that are no longer in wide use today are Token-Ring, and FDDI.

*  Layer 2 - Data link layer

   The data link layer defines the basic protocol for communicating between two
   points on a network, that may consist of many intermediate devices and cables,
   possibly spanning a large geographic area. The IEEE 802.3 specification defines
   the data link layer, which includes the Media Access Control (:term:`MAC`)
   addresses that allow hosts to address their data to one or more systems on the
   same Ethernet segment. The MAC address is a flat (non-hierarchical) 48-bit
   address, and therefore when a node sends out a MAC broadcast frame (usually
   written as FF:FF:FF:FF:FF:FF), all stations on the Layer 2 segment receive it.
   Therefore, a layer 2 segment is also known as a "broadcast domain". 

*  Layer 3 - Network layer

   The network layer is what allows many flat Layer 2 segments to be
   interconnected, and separates broadcast domains. It is this layer of
   the OSI model that enables the Internet to exist, using Internet Protocol
   (IP) addressing. Version 4 of the Internet Protocol (IPv4), most commonly found in
   today's production networks, is described in :rfc:`791`. Its successor, IP version
   6 (IPv6) is described in :rfc:`2460`. Both protocols utilize addresses that are
   hierarchical in nature, providing for a *network* portion of the address space,
   and also a *host* portion within the network. 
   
   There are also two control protocols included in this layer; namely, the
   "Internet Control Message Protocol" ("ICMP") as described in :rfc:`792`, and
   "Internet Group Management Protocol" ("IGMP"), as described in :rfc:`1112`. The
   ICMP protocol provides for Layer 3 device control messaging, and among other
   uses, allows small test packets to be sent to a destination for troubleshooting
   purposes, such as those used by the ubiquitous ``ping`` utility. The IGMP
   protocol is used to manage multicast groups, which implements "one-to-many"
   packet sending between interested systems.

*  Layer 4 - Transport layer

   The transport layer is where things really start to get interesting for the
   systems administrator. It is at the transport layer that the Transmission
   Control Protocol (TCP), and the User Datagram Protocol (UDP) are defined.
   The TCP and UDP protocols allow data to be sent from one system to another using
   simple "socket" APIs that make it just as easy to send text across the globe as it
   is to write to a file on a local disk - a technological miracle that is often
   taken for granted.

*  Layer 5 - Session layer

   The purpose of the session layer is to provide a mechanism for ongoing
   conversations between devices using application-layer protocols. Notable
   Layer 5 protocols include Transport Layer Security / Secure Sockets Layer
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


TCP/IP (ARPA) 4-layer model
===========================

When the ARPAnet project spawned TCP/IP, which became the dominant network
protocol stack used on the Internet, it was envisioned with a much simpler
4-layer stack (the ISO layers roughly map to the ARPA model as indicated below.)
The elements of this stack from the lowest to highest are as follows:

*  Link Layer

   Basically a combination of the ISO Layer 1 (physical) and Layer 2 (data link) layers.
   This layer covers the physical network card in the computer (and the transmission
   medium between the computers) as well as the device driver in the operating system
   of the computer. As an example, both Ethernet and MAC are covered in this layer.
   
*  Network Layer

   This layer maps to the ISO's Layer 3, and covers the movement of packets between
   networks. For this reason, it is also called the "Internet Layer", and the main
   protocol used in this layer is named the "Internet Protocol", or "IP" as it is commonly
   referred to. As discussed above, the ICMP and IGMP protocols are also included in
   this layer.
   
*  Transport Layer

   This layer maps to the ISO's Layer 4, and covers the creation, management and teardown
   of "virtual circuits" or "flows" between end hosts. There are two different protocols
   in use at this layer as discussed above in the ISO Layer 4 section, namely, TCP and
   UDP.
   
*  Application Layer

   This layer maps to the ISO's Layer 5 through Layer 7, and covers the application
   processes that use the network to communicate.
   

IP Addressing
=============

IPv4
----

Internet Protocol Version 4 (IPv4) is the fourth version of the Internet
protocol, the first version to be widely deployed. This is the version of the
protocol you're most likely to encounter, and the default version of the IP
protocol in Linux.

IPv4 uses a 32-bit address space most typically represented in 4 dotted decimal
notation, each octet contains a value between 0-255, and is separated by a dot.
An example address is below:

    10.199.0.5

There are several other representations, like dotted hexadecimal, dotted octal,
hexadecimal, decimal, and octal. These are infrequently used, and will be
covered in later sections.



IPv6
----

The Internet has experienced tremendous growth from its beginnings in research
at universities to its use in e-commerce and other applications today.
As a result, more and more IPv4 addresses were given out to enable users to connect
to the Internet.
However, despite the number of IPv4 addresses being a large number (4294967296), they
are slowly running out.

Internet Protocol Version 6 (IPv6) is the latest version of the Internet protocol
which aims to address the IPv4 address exhaustion problem.

Probably the most obvious difference of IPv6 to IPv4 is the representation of addresses.
Unlike IPV4, IPv6 uses a hexadecimal format to represent a 128-bit address. The format is
grouped into 8 groups of four hexadecimal digits separated by a colon.
For example: 

    2001:0DB8:0000:CBAD:4321:0000:0000:1234

Leading zeroes may be left off of a group in the address. Whole groups of zeros may be
left out entirely, but only once. Generally, the longest run of zeros is dropped, but
any run may be. Using these rules, the above address can also be represented in either
of these ways:

    2001:DB8::CBAD:4321:0000:0000:1234
    2001:DB8:0000:CBAD:4321::1234

TCP vs UDP
==========

Both TCP :rfc:`793` and UDP :rfc:`768` provide data transfer between processes 
through ports. These process ports can be on the same computer or separate 
computers connected by a network. TCP provides the following: reliability, 
flow control, and connections (see Example Difference 1 below). UDP is less 
feature-rich, it does its work with a header that only contains a source port, 
destination port, a length, and a checksum. TCP provides its capabilities 
by sending more header data, more packets between ports and performing more 
processing. UDP requires less header data in the individual 
packets and requires fewer packets on the network to do its work. UDP does no 
bookkeeping about the fate of the packets sent from a source. They could be 
dropped because of a full buffer at a random router between the source and 
destination and UDP wouldn't account for it in itself (other monitoring systems
can be put in place to do the accounting, however that is beyond the UDP
protocol). 

The choice of protocols to use is often based on whether the risk of losing 
packets in real-time without immediate alerting is acceptable. In some cases 
UDP may be acceptable, such as video or audio streaming where programs can 
interpolate over missing packets. However, TCP will be required due to its 
reliable delivery guarantee in systems that support banking or healthcare.

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

Subnetting, netmasks and CIDR
=============================
A subnet is a logical division of an IP network, and allows the host system to
identify which other hosts can be reached on the local network. The host system
determines this by the application of a routing prefix. There are two typical
representations of this prefix: a netmask and Classless Inter-Domain Routing (CIDR).

Netmasks typically appear in the dotted decimal notation, with values between
0-255 in each octet. These are applied as bitmasks, and numbers at 255 mean that
this host is not reachable. Netmask can also be referred to as a Subnet Mask and
these terms are often used interchangeably. An example IPv4 address with a typical
netmask is below:

============= ===============
IP Address    Netmask
============= ===============
192.168.1.1   255.255.255.0
============= ===============

CIDR notation is a two-digit representation of this routing prefix. Its value can range
between 0 and 32. While having long been a staple in network equipment and used mainly by
network engineers, CIDR notation can now be found in Linux (iproute's `ip` command uses
CIDR notation). In addition, CIDR notation is often encountered in discussions, as it's
quicker and simpler to reference than a netmask. Below is the same example as above
with CIDR notation:

============= ===============
IP Address    CIDR
============= ===============
192.168.1.1   /24
============= ===============

As a server administrator, it is helpful to understand subnetting so as to communicate more
effectively with network engineers, and to design networks yourself. There are many guides
online for learning how to subnet, each with different approaches and tricks. A full subnetting
guide is out of the scope of this document.

Classful addressing
===================

You may hear people refer to networks as "Class A", "Class B", or "Class C". This is known
as "Classful Addressing" and has been deprecated for decades, thanks to the introduction
of CIDR in 1993 (:rfc:`1519`). Classful addressing has the drawback of assuming that a
network is drawn on very large boundaries. For example, in the table in the next section,
each block would be a single network. One can see how using an entire /16 (65534 hosts)
when only five hosts are needed would be wasteful. As such, CIDR was created, allowing
people to create subnets only as large as required. For these reasons, one should not refer
to any particular network block as a "Class X" (eg, 10.0.0.0/28 should not be called a Class A
network) as it is misleading at best, and incorrect at worst. You should always use CIDR
notation to ensure accuracy.

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

Network Address Translation, or NAT, is a technology that allows multiple
internet-connected devices to share one common **public** IP address, while
still retaining unique, individual **private** IP addresses. The distinction
between public and private is integral to the understanding of the service
that NAT provides and how it works.

In our :ref:`Sample Network <sample-network>` we can see that two firewall
machines sit between the Internet and the other hosts in the network; traffic
going in and out of the network will pass through these firewall machines. The
addresses assigned to the firewalls (10.10.10.1 and 10.10.10.2) are private IPs
and are visible to just the hosts within the network. A device on the external
Internet will, instead, see the public IP address for each firewall. It's
important to note that none of the hosts within the network will have a public
IP address, except for the firewalls and the DNS servers, since they are the
only parts of the network that directly communicate with the external Internet.

When a device behind a NAT-enabled router sends a packet, the source IP address
on the packet is the device's local, private IP address. If the packet is going
outside the local network, it will pass through a router, which will modify the
source IP address to its own public IP address. When a response for the packet
is received by the router, it needs to ensure that the response can be forwarded
to the host that sent the packet in the first place. To do this, the router
maintains a **Translation Table**. This table maps a device's IP address and
port to a port on the router itself. The router's public IP address and
the correct port number from the table are used as the source IP and port on
the packet and then sent to the destination.

These maps are temporary and exist on a per-connection basis. This means that each
connection opened by a device will have a unique port number on the device and a
unique port number on the router as well. This port on the router is used as the
public port for that connection. Once the connection terminates, the router is
free to assign that port to another connection. However, the total number of
available ports is limited to 65,536, so it is entirely possible that a router
has no more free ports and cannot assign a new NAT address. This is commonly
referred to as port exhaustion.

Similar to port exhaustion, timeouts can also affect the router's ability to
assign new NAT addresses. Each entry in the translation table has a timeout
value which refers to the amount of time for which that entry can remain
inactive and still keep its place in the table. An entry that has remained
inactive for a period of time longer than the timeout will automatically be
removed, freeing up space for a new one.

Networking cable
================
There are two main types of network cable in use today, namely copper and fiber-optic.

Copper
------
The most common type of network cables are what is known as "unshielded twisted
pair" cables. They use 4 sets of twisted pairs of copper, relying on the twist
with differential signaling to prevent noise and signal propagation between the
pairs. The four pairs of twisted copper wires are encased in a plastic sheath.

There are different standards for copper network cables set by the
Telecommunications Industry Association (TIA) and the International Organization
for Standardization (ISO). Both organizations use the same naming convention
("Category __") for the components, but unfortunately differ on the naming for
the cable standards. The most common reference is the TIA's, and the category
designation is usually shortened to "Cat", so you'll hear references to "Cat5"
or "Cat6" cable.

Copper Cable Standards
^^^^^^^^^^^^^^^^^^^^^^

- Category 5e ("Cat5", ISO class D)

- Category 6 ("Cat6", ISO class E)

- Category 6A ("Cat6A", ISO class Ea)

Fiber
-----
Fiber is a generic term that refers to optical transport mediums. It comes in
several types, all of which look identical but are generally incompatible.

Multimode vs Single Mode
^^^^^^^^^^^^^^^^^^^^^^^^
Single-mode fiber has a small core diameter, which only allows one (a single)
mode of light to be transmitted through the fiber. Using a single mode of light
completely eliminates the possibility of light dispersion and associated signal
loss, and so is used mainly for long-haul runs, such as the cables that run
between buildings and cities. However, since single-mode fiber can only transmit
one wavelength of light at a time, it typically involves much more expensive
light generation sources (i.e., laser diode transmitters) and is very expensive
to produce.

Multimode fiber has a larger core diameter (either 50u or 62.5u) and can
therefore carry multiple modes ("multimode") of light, which can be used to
transmit much more information during a given timeslice. The drawback is that
carrying multimode lightwaves causes light dispersion and associated signal
loss, which limits its effective distance. Multimode is a less expensive fiber
optic cable, that is typically useable with lower cost optical components. It is
very common to see it used for building intra-building backbones, and
system/switch to switch applications.

Multimode Fiber Standards
^^^^^^^^^^^^^^^^^^^^^^^^^
Multimode cables have classifications much like the copper cables discussed above; these
are known as "Optical Multimode" (OM) classes. The four designations are:

- OM1 - a "legacy" fiber class, the core being 62.5u, and cladding being 125u.
  The bandwidth that can be carried ranges from 160 to 500 MHz.
  
- OM2 - a "legacy" fiber class, the core being 50u, and cladding being 125u.
  The bandwidth that can be carried is 500 MHz.
  
- OM3 - a "modern" fiber class, the core being 50u, and cladding being 125u.
  The bandwidth that can be carried ranges from 1500 to 2000 MHz.
  
- OM4 - a "modern" fiber class, the core being 50u, and cladding being 125u.
  The bandwidth that can be carried ranges from 3500 to 4700 MHz.

Optical Connector Types
^^^^^^^^^^^^^^^^^^^^^^^

LC and SC connectors are the two most common type of fiber connectors you will
use.

LC stands for "Lucent Connector", but is also referred to as "Little Connector".
They are typically used for high-density applications, and are the type of
connector used on SFPs or XFPs. Typically the connector is packaged in a duplex
configuration with each cable side by side, and have a latch mechanism for
locking.

SC stands for "Subscriber Connector", but are also known as "Square Connector",
or "Standard Connector". This is the type of connector typically used in the
telecom industry. They have a larger form factor than the LC connectors, and can
be found in single and duplex configurations. SC connectors have a push/pull
locking mechanism, and because of this, are also colloquially known as
"Stab-and-Click" connectors.

Transceivers
^^^^^^^^^^^^

The variety in optical fiber makes for a correspondingly large variety in
optical fiber interface standards. Different interface types will impose
different requirements on the fiber used and the length of the connection.

If optical fiber interfaces were incorporated directly into network equipment,
the number of models made by the manufacturer would have to be multiplied by
the number of interface standards in existence. For this reason, modern network
hardware rarely incorporates such interfaces directly. Instead, pluggable
transceiver modules are used as a layer of indirection between medium-dependent
and medium-independent interfaces. This allows a transceiver slot to be
provided supporting any desired interface standard, whether copper or fiber.
There are some limitations to this, detailed below.

Various module types have been introduced over the years:

============ ============ =========== =======
  Name        Introduced   Speed       Size
============ ============ =========== =======
  GBIC        1995         1 Gb/s      Large
  SFP         2001         1 Gb/s      Small
  XENPAK      2001         10 Gb/s     Large
  XFP         2002         10 Gb/s     Small
  SFP+        2006         10 Gb/s     Small
  QSFP        2006         40 Gb/s     Small
  CFP         2009         100 Gb/s    Large
============ ============ =========== =======

There are a large number of compatibility issues with such modules. Some
compatibility issues cause problems between two ends of a link; others cause
problems between a module and its host device.

- Transceivers are not generally compatible with lower speed versions of the
  same standard. A 1000BASE-T Ethernet card can interface with a 10BASE-T card,
  but a 1 Gb/s fiber transceiver cannot interface with a 10 or 100 Mb/s
  transceiver. In the case of fiber, this is generally due to the different
  wavelengths used; but even many copper transceivers do not support lower
  speeds, although some do. You should assume that any transceiver will only
  support the exact interface for which it is designed unless specified
  otherwise.

- Modules are only made for the speed targeted by a format. For example, SFP+
  modules are only made for 10 Gb/s standards, and not for lower speeds.

- Some equipment may accept SFP modules in SFP+ slots, but this is not
  universal.

- Vendor lock-in is widely practiced. Equipment may refuse to operate with the
  modules made by a different manufacturer. Workarounds are generally
  available, but this may complicate support or warranty arrangements.

These issues can create pathological cases. Suppose you have two switches which
should be connected to one another. One is connected via a 1 Gb/s transceiver
to fiber. The other only has SFP+ slots. If these slots also support SFP
modules, then a 1 Gb/s SFP transceiver can be used, but if they do not,
interconnection is impossible: all SFP+ modules target 10 Gb/s, and fiber
transceivers do not support lower speeds.

Twinax
^^^^^^

