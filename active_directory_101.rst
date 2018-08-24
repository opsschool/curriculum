####################
Active Directory 101
####################

What is Active Directory?
=========================

**Active Directory** is a Directory Service created by Microsoft. It
is included in Windows Server operating systems since Windows 2000
Server, which was launched in 2000.

Almost all Active Directory installations actually include several
separate but related components; although the term "Active Directory"
technically refers only to the directory service, in general use it
refers to the entire constellation of parts.

What is Active Directory used for?
==================================

Active Directory is primarily used to store directory objects (like
users and groups) and their attributes and relationships to one
another. These objects are most commonly used to control access to
various resources; for instance, an Active Directory might contain a
group which grants its members permission to log into a certain
server, or to print to a specific printer, or even to perform
administrative tasks on the directory itself.

Active Directory also provides a useful configuration management
service called *Group Policy*, which can be used to manage computers
which connect to the domain in order to install packages, configure
software, and much more.

You mention "separate components"; what is Active Directory composed of?
========================================================================

Most Active Directory installations have a few distinct parts which
all work together:

- a *directory database* which stores the actual directory information
- a *Kerberos key distribution center* which helps manage user
  passwords and other security resources
- a *DNS server*, which maps IP addresses to hostnames and back
- a *DHCP server*, which grants dynamic IP addresses to hosts as they
  join the network
- one or more *Global Catalogs*, which cache parts of the directory
  database and help speed up some common queries

Also, Active Directory is designed in such a way that it can be run on
multiple computers at the same time, which coordinate between
themselves to ensure that their data is always consistent; this
process is called "replication".

What specific services does Active Directory provide?
=====================================================

Group Policy
------------

DNS
---

Kerberos Key Distribution Center
--------------------------------

DHCP
----


Best Practices for managing an Active Directory installation
============================================================

Cleaning up Unneeded Objects
----------------------------


Making Backups
--------------


Replication Health
------------------

