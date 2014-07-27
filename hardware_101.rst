Hardware 101
************

Hardware Types
==============

Rackmount
---------
Rackmount refers to servers mounted in 19-inch relay racks. The server's height is measured in U or rack units (1.75 inches) and servers typically range in height from 1U to 4U(though they can be larger).
The rack they are installed into is typically 42U high, but you rarely see installations with more that 40 servers in a rack.
Servers are secured to the rack via specialty rails.

Rackmount servers are the most common physical hardware installation in a datacenter.
Individual rackmount systems can usually meet most density requirements, but in large installations physical cable management can become an issue.
In addition systems management can be a problem in large installations.
Rackmount servers usually are less expensive than other solutions, and modern rackmount servers typically include some type of remote-management.

Blades
------
A blade is a server on a removable card which connects to a chassis for power, network, video, and other services.
The blade chassis is typically a rackmount unit, but sizes of this are usually much larger, using 8U or more depending on the number of blades the chassis supports.

Blades were originally hoped to solve server density issues, and some products have the ability to house four servers in only 2U of space.
In recent years blades have been used to solve physical cable management issues in complex environments, since there are very few cables required to connect the systems to other pieces of infrastructure.

SANs
----
A Storage Area Network or SAN is a network that provides access to centralized block storage.
Common choices for the network fabric are Fiber Channel and Ethernet.
Solutions built around Ethernet usually rely on an additional protocols to implement access, and may require special configurations for high throughput (thing like large frames).
The most common Ethernet based approach uses iSCSI.
These networks are typically separate and dedicated to providing access to storage.

SAN Array
---------
A SAN Array is a dedicated array of disks that has been designed to connect to a SAN.
SAN Arrays will often support Ethernet or Fiber Channel, and some vendors will support both network types within the same unit. 

Most SAN Arrays have two main components, a head unit and an array.
The head unit controls a collection of arrays, and the SAN configuration.
Most SANs come with two head units for redundancy.
The Array unit holds some number of disks, and typically connects directly to the head units.
The head unit is usually connected to the network fabric.

SANs installations vary from vendor to vendor, but they generally can be installed in standard 19-inch relay racks, or come in a custom cabinet.
Power requirements also vary from vendor to vendor, but it its not atypical that they require dedicated power circuit due to the large number of spinning disks.

NAS
---
A Network-Attached Storage device or NAS is a specialty server designed to provide storage to other servers.
Unlike SANs, a NASs will typically implement an standard network protocol like SMB or NFS to expose the storage to the client.
The client will present these to its users as standard network share.

NASs and SANs are almost indistinguishable from each other physically.
They will typically have head units and array units, and there will typically be two head units for redundancy.
They will typically install into a 19-inch rack, or come with a custom cabinet.
Lastly, they will typically need a dedicated circuit due to the large number of spinning disks.

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


