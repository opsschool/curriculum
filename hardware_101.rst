Hardware 101
************

Hardware Types
==============

Rackmount
---------
Rackmount refers to servers mounted in 19-inch relay racks. The servers height
is measured in U or rack units(1.75 inches) and servers typically range in
height from 1U to 4U(though they can be larger). The rack they are installed
into is typically 42U high, but you rarely see installations with more that 40
servers in a rack. Servers are secured to the rack via specialty rails.

Rackmount servers are the most common physical hardware installation in a
datacenter. Individual rackmount systems can usually meet most density
requirements, but in large installations physical cable management can become
an issue. In addition systems management can be a problem in large installations.
Rackmount servers usually are less expensive than other solutions, and modern
rackmount servers typically include some type of remote-management.

Blades
------
A blade is a server on a removable card which connects to a chassis for power,
network, video, and other services. The blade chassis is typically a rackmount
unit, but sizes of this are usually much larger, using 8U or more depending on
the number of blades the chassis supports.

Blades were originally hoped to solve server density issues, and the some have
the ability to house 4 servers in only 2U of space. The primary marketing
message these days is the ease of physical cable management in complex
environments.

SANs
----

A Storage Area Network or SAN is a large array of storage connected to some type
of network fabric. Common choices for the network are Fiber Channel and Ethernet.
Solutions built around Ethernet usually rely on an additional protocol implemented
in IP, though there are some pure Ethernet solutions. The client using the SAN
will expose the SAN to its users as a typical block device, like a hard drive.

Most SANs have two main components, a head unit and an array. The head unit
controls a collection of arrays, and the SAN configuration. Most SANs come with
two head units for redundancy. The Array unit holds some number of disks, and
typically connects directly to the head units. 

SANs installations vary from vendor to vendor, but they generally can be
installed in standard 19-inch relay racks, or come in a custom cabinet. Power
requirements also vary from vendor to vendor, but it its not atypical that they
require dedicated power circuit due to the large number of spinning disks.

NAS
---

A Network-Attache Storage device or NAS is a specialty server designed to
provide storage to other servers. Unlike SANs, a NASs will typically implement
an standard network protocol like SMB or NFS to expose the storage to the 
client. The client will present these to its users as standard network share.

NASs and SANs are almost indistinguishable from each other physically. They
will typically have head units and array units, and there will typically be
two head units for redundancy. They will typically install into a 19-inch rack,
or come with a custom cabinit. Lastly, they will typically need a dedicated
curicuit due to the large number of spinning disks.

Basic server architecture
=========================

CPU
---

RAM
---

Motherboard
-----------

Bus
---

Disk controllers
----------------

Network devices
---------------

BIOS/UEFI
---------

PCI-E Cards
-----------

Disk management
===============

Arrays
------

RAID/ZFS
--------

Logical Volume Management
-------------------------

Performance/Redundancy
======================

RAID/ZFS
--------

Power
-----

Electrical overview
^^^^^^^^^^^^^^^^^^^

110v/220v, amperage, 3 phase power
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

UPS
^^^

common plug types
^^^^^^^^^^^^^^^^^

common circuit types
^^^^^^^^^^^^^^^^^^^^

Troubleshooting
===============

Remote management (IPMI, iLO, DRAC, etc)
----------------------------------------

Common MCE log errors
---------------------

System burn-in
--------------


