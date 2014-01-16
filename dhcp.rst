DHCP
****
DHCP, or, Dynamic Host Control Protocol is a standard that allows one central 
source, ie. a server or router, to automatically assign requesting hosts an IP
address.

The most common DHCP servers on Linux platforms are the ISC's
(Internet System Consortium) ``dhcpd``, and ``dnsmasq``. On Windows platforms
there are a variety of DHCP servers, though the Microsoft DHCP server is
arguably the best in terms of support and integration.


Tools: ISC dhcpd
================
``dhcpd`` is the most common DHCP server, and it offers excellent integration
with BIND (DNS server).

Protocol
========

dhcp helper
===========
DHCP helpers are sometimes referred to as DHCP relayers. The basic idea is that
a relay agent will forward DHCP requests to the appropriate server. This is
necessary because a host that comes online on a subnet with no DHCP server
has no way of finding the correct server; it needs a DHCP assigned address to 
find a route to the correct DHCP server, but can't get there because it has no
IP address! DHCP relaying solves this chicken and egg problem by acting as an
intermediary. See this wikipedia [#]_ article for more details.

Defining classes, leases
========================

Options
=======

Default gateway
---------------

DNS server(s)
-------------
(Tie in previous chapters re: TFTP, PXE with related options?)

.. [#] `DHCP relaying explained <http://en.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol#DHCP_relaying>`_

