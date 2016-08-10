Networking 201
**************

VLANs, 802.1q tagging
=====================
Sometimes it is advantageous to break up a single Layer 2 broadcast domain into multiple Layer 2 broadcast domains.
As we had previously discussed, Layer 2 is addressed by MAC addresses, which are flat address space (non-routable) constructs; this creates a Layer 2 broadcast domain, also called a "segment" or "LAN" (local area network.) You can then map one or more IP networks (Layer 3 addressing) onto a Layer 2 segment.
If you map more than one Layer 3 network to a segment, you can get into some headaches, such as having to address multiple IP addr's on various host and router interfaces (one for each network) with the attendant confusion, and also that broadcasts produced by each IP network (i.e. ARP requests) multiply, and bandwidth is consumed.
Therefore, most networking professionals agree that it's best to only map one IP network to one Layer 2 segment.

So, if multiple Layer 3 networks need to be used, instead of using multiple switches for each Layer 2 segment, it is possible to divide a single VLAN-capable switch into two or more "virtual" switches.
This is possible due to the concept of :term:`VLAN`, which stands for "Virtual LAN".

To communicate between the Layer 3 networks mapped to their individual VLANs, packets need to pass through a router or other Layer 3 device.
If one has a switch with many VLANs implemented on it, either of these two methods could work:

1. Devote a switch (and respective router) interface per VLAN, with a single link for each VLAN.
   This is the simplest method, but the most wasteful of (expensive) switch and router interfaces.

2. "Trunk" multiple VLANs over one physical link, which only requires a single interface per side (switch to router, or server to switch.)
   This method multiplexes multiple VLANs over the one link by "tagging" the frame with a 4-byte VLAN identifier field which is inserted after the destination MAC field in the Ethernet header.
   This tagging format is officially defined by the IEEE's 802.1q protocol.
   The tag field insertion happens at the sending trunk interface (which is a part of the sending station's VLAN), and the tag is stripped of by the receiving trunk interface, and the frame placed on the proper VLAN on the receiving device.

The VLAN identifier must be between 1 and 4094 inclusive.
The 802.1q tag also provides for a 3-bit priority field and 1-bit discard eligibility field which can be used for QoS applications.
The trunking and priority functions of 802.1q can be enabled and disabled separately.

Spanning Tree
=============
As networks and broadcast domains grow and become more complex, it becomes much easier to intentionally or unintentionally create loops- multiple network paths between one or more switches.
These loops may lead to a infinite forwarding of Ethernet packets, causing switches to be overwhelmed by an ever-increasing amount of packets being forwarded to all other switches in the broadcast domain.
To address this concern, Spanning Tree Protocol (STP) was invented as a standard protocol used by Ethernet switches, designed to prevent such loops in broadcast domains.

To disable loops, each switch that implements Spanning Tree Protocol (STP / IEEE 802.1D) will participate in electing a "root" device where the Spanning Tree is calculated from.
Once the root is chosen, all other participating switches will calculate their least-cost path for each of its ports to the root device.
This is usually a port that traverses the least number of segments between the source and the root.
Once each device has chosen its least-cost port, or Root Port (RP), it calculates the least-cost path for each network segment.
By doing this, each switch obtains a Designated Port (DP), or a port on which it should expect to accept traffic to forward through its RP.
Upon calculating these values, all other paths to the root device are disabled- marking them as Blocked Ports (BP).

A major benefit to using STP to prevent network loops is that network administrators can intentionally connect redundant paths for fault tolerance without worry of causing an infinite loop during normal operation.
Should a path in a broadcast domain fail, each participating switch will recalculate the DP and RP, modifying the status of BPs if necessary.
This has the effect of repairing the broken path by routing traffic through segments around the failure.

Routing
==============
When a network is subdivided into multiple Layer 2 broadcast domains, Layer 3 addressing enables hosts to communicate with another host in another broadcast domain.

The process of forwarding packets from one Layer 2 subdomain to another Layer 2 subdomain using Layer 3 addresses is called routing.

A Layer 3 device or router is the device responsible for this function.
Routers may use many different ways to forward packets but these methods can be categorized into two types.

Static Routing
--------------
The method of using manually configured routes is called static routing.

Typically, there are three pieces of information that are needed to specify a static route:

1. the destination subnet

2. the latter's subnet mask and

3. the next hop host or outgoing interface.

Take for example the following static route configuration in Linux:

.. code-block:: console

   # ip route add 10.10.20.0/24 via 192.168.2.1

The above command means that packets intended for hosts in the 10.10.20.0/24 network must be forwarded to the host with ip address 192.168.2.1.

Below is another example in Cisco IOS:

.. code-block:: console

   Router(config)# ip route 10.10.20.0 255.255.255.0 Serial0/0/1

Instead of passing a packet to a specific ip address however, this configuration states that packets intended for the 10.10.20.0/24 network should be forwarded to the host directly connected to Serial 0/0/1 whose ip address can be anything.

Static routes do not change when the network changes but can be useful in cases where the network is small enough that it outweighs the cost of dynamic routing.
It is often used in tandem with dynamic routing to specify a default route in case a dynamic route is unavailable.

Dynamic routing protocols (RIP, OSPF, BGP)
--------------------------------------------------------
For a small network, manually configuring routes will work just fine.
As a network grows larger, however, doing this can be very arduous if not infeasible.
Dynamic routing solves this problem by programmatically building a routing table.

Here are some examples of dynamic routing protocols:

1. Routing Information Protocol (RIP) - RIP uses the number of routers between a router and the destination subnet as its path selection metric.
When a RIP router first comes online, it sends a broadcast message and then all its neighbor routers respond back with their own routing table.
This new RIP router then compares the routes in its own routing table with the routes from the other routing tables it has received.
If there are routes to other subnets that are currently not in its own routing table, that route is added.
More importantly, it updates routes in its own routing table if there is a route with less hops to a destination subnet.
RIP is part of a class of routing protocols called distance vector routing protocols.
They are characterized by their sharing of routing tables with their neighbors.
Similar protocols to RIP include Cisco's Enhanced Interior Gateway Routing Protocol (EIGRP).
EIGRP is similar to RIP but uses a mix of path selection metrics such as bandwidth and delay to determine the route to a destination subnet.
Unlike RIP, it only sends updates incrementally to reduce overhead.

2. Open Shortest Path First (OSPF) - OSPF is a routing protocol that calculates the shortest path to each destination subnet in the network.
Routing protocols like OSPF that work this way are said to be called Link State routing protocols.
To be able to know the shortest path to each destination subnet, each router should know everything about the network from links, subnets and even other routers.
To do this in OSPF, each router packages routing information they know in a data structure called Link State Advertisements (LSA) and floods the whole network with this information.
Fellow OSPF routers listen for LSAs, collect this data and store them in a Link State Database (LSDB).
The point of all this sharing is for each router to have identical LSDBs.
Once this is achieved, each OSPF router calculates routes to all subnets from its own point of view using a shortest path algorithm.
These routes are added then added to the router's routing table.

3. Border Gateway Protocol (BGP) -

ACLs
====

Network Bonding (802.3ad / LACP link aggregation)
=================================================

IOS switch configuration
========================

Cisco devices use a command line-based configuration interface which can be accessed via RS-232, Telnet or SSH.
The command line syntax is distinctive and has been loosely copied by many other vendors, so some familiarity with the Cisco-style configuration can go a long way.

When connecting via RS-232, use 9600 baud, 8 data bits, 1 stop bit, no parity, no flow control.

You can abbreviate keywords on the command line without having to press Tab, so long as they are unambiguous.
You can also use Tab as you would in a UNIX shell.

While using the command line, you can type '?' at any time and receive context-sensitive help.

The command line is mode-based. The prompt tells you of your current mode.
Awareness of your current mode is key to efficient operation of the command line.
Here are some examples:

=========================  ===================================================
Prompt                     Meaning
=========================  ===================================================
``hostname>``              'EXEC' mode, non-privileged access
                           (like ``$`` on Linux)
``hostname#``              'EXEC' mode, privileged access
                           (like ``#`` on Linux)
``hostname(config)#``      Global configuration mode
``hostname(config-if)#``   Interface configuration mode
=========================  ===================================================

EXEC mode allows you to execute imperative commands such as ``ping``.
The configuration mode allows you to add and remove configuration statements.

The following commands move you between modes:

=========================  ======  ===========================================
Command                    Mode    Effect
=========================  ======  ===========================================
``enable``                 EXEC    Become privileged (like ``su`` on Linux)
``conf t``                 EXEC    Enter ``(config)`` mode.
``int <interface>``        config  Enter ``(config-if)`` mode for an
                                   interface.
``exit``                   Any     Leave the current mode or logout.
=========================  ======  ===========================================

Other useful commands include:

=========================  ======  ===========================================
Command                    Mode    Effect
=========================  ======  ===========================================
``show ...``               EXEC    The subcommands of this command provide
                                   access to all available information.
``show run``               EXEC    Show the current configuration.
``show ip int brief``      EXEC    Show all interface names, IPv4 addresses,
                                   and their status.
``wr``                     EXEC    Save the current configuration.
``ping <host>``            EXEC    Ping. '!' means a response,
                                   '.' means a timeout.
``no <command>``           config  Delete a configuration statement in config
                                   mode.
``do <command>``           config  Execute a command in EXEC mode from config
                                   mode.
=========================  ======  ===========================================

Note that configuration changes become active as soon as they are made.
The ``show run`` command shows the configuration which is currently in effect.
This is not saved until you execute ``wr``.

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


Network Troubleshooting
=======================

ping
----
ping should always be the first step in any network-related troubleshooting session, due to
the simple manner in which it works, and the information it returns.
Many times, the resulting information from a ping will point you in the next direction.
For example, if you notice jitter (ping responses varying wildly), you will know to start
looking at layer 1 problems somewhere on the path, and that available bandwidth probably
isn't an issue.

You're probably already familiar with the basic usage of ping, but there are some really handy options
that can used (these are on Linux, but also exist in most versions of ping, but under different flags).
Here are a couple of my most-often-used options:

-I - Change the source IP address you're pinging from.
For example, there might be multiple IP addresses on your box, and you want to verify that a particular
IP address can ping another.
This comes in useful when there's more than just default routes on a box.

-s - Set the packet size.
This is useful when debugging MTU mismatches, by increasing the packet size.
Use in conjuction with the -M flag to set MTU Path Discovery hint.

An example of using ping to test MTU:

.. code-block:: console

  user@opsschool ~$ ping -M do -s 1473 upstream-host
  PING local-host (local-host) 1473(1501) bytes of data.
  From upstream-host icmp_seq=1 Frag needed and DF set (mtu = 1500)
  From upstream-host icmp_seq=1 Frag needed and DF set (mtu = 1500)
  From upstream-host icmp_seq=1 Frag needed and DF set (mtu = 1500)

I've used the `-M do` option to set the 'Don't Fragment' (DF) flag on the packet, then
set the packet size to 1473.
With the 28 bytes of overhead for Ethernet, you can see the total packet size becomes
1501--just one byte over the MTU of the remote end.
As you can see from the example, since the DF flag is set and the packet needs to fragment,
it spits back an error, and helpfully tells us what the MTU size is on the other end.
ping can be used to determine Path MTU (the smallest MTU size along a path), but other
tools are better for that (see below).

telnet
------

While telnet daemons are a big no-no in the wild (unencrypted traffic), the telnet client
utility can be used to test whether TCP connections can be made on the specified port.
For example, you can verify a TCP connection to port 80 by connecting via telnet:

.. code-block:: console

  user@opsschool ~$ telnet yahoo.com 80
  Trying 98.138.253.109...
  Connected to yahoo.com.
  Escape character is '^]'.

A connection failure would look like this (using port 8000, since nothing is listening on that port):

.. code-block:: console

  user@opsschool ~$ telnet yahoo.com 8000
  Trying 98.138.253.109...
  telnet: connect to address 98.138.253.109: Connection timed out

You can also send raw data via telnet, allowing you to verify operation of the
daemon at the other end.
For example, we can send HTTP headers by hand:

.. code-block:: console

  user@opsschool ~$ telnet opsschool.org 80
  Trying 208.88.16.54...
  Connected to opsschool.org.
  Escape character is '^]'.
  GET / HTTP/1.1
  host: www.opsschool.org

The last two lines are the commands sent to the remote host.

Here's the response from the web server:

.. code-block:: console

  HTTP/1.1 302 Found
  Date: Fri, 26 Dec 2014 14:55:45 GMT
  Server: Apache
  Location: http://www.opsschool.org/
  Vary: Accept-Encoding
  Content-Length: 276
  Content-Type: text/html; charset=iso-8859-1

  <!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
  <html><head>
  <title>302 Found</title>
  </head><body>
  <h1>Found</h1>
  <p>The document has moved <a href="http://www.opsschool.org/">here</a>.</p>
  <hr>
  <address>Apache Server at www.opsschool.org Port 80</address>
  </body></html>
  Connection closed by foreign host.

Here we passed the bare minimum required to initiate an HTTP session to a remote web server,
and it responded with HTTP data, in this case, telling us that the page we requested is located
elsewhere.

Note that the port you're connecting to might the port for HTTP, but it could be something
other than an HTTP daemon running.
Nothing prevents a service from running on a port other than it's usual one.
In effect, you could run a SMTP daemon on port 80, and an HTTP daemon on port 25.
Testing TCP connections with telnet would verify TCP operation, but you still would not
have a working web server on port 80.
Since the scope of this section is focused only on the networking aspect, see the other
sections of OpsSchool for troubleshooting daemon operation and Linux troubleshooting.

iproute / ifconfig
------------------

ifconfig is ubiquitous and a mainstay of any network-related work on Linux, but it's actually
`deprecated in RHEL7 <https://bugzilla.redhat.com/show_bug.cgi?id=1119297>`_ (the net-tools
package which contains `ifconfig` isn't included in RHEL 7/CentOS 7 by default) and many major
distributions include iproute by default.
The `ifconfig <http://linux.die.net/man/8/ifconfig>`_ man page also recommends using the iproute package.
All examples used below will use iproute, and will cover only the basics of troubleshooting.
It's highly recommended to play around and see what you can find.
The `ip <http://linux.die.net/man/8/ip>`_ man page contains a wealth of knowledge on the tool.

ip addr show
^^^^^^^^^^^^^^

Show all IP addresses on all interfaces.
Many options can be passed to filter out information.
This will show several important pieces of information, such as MAC address, IP address, MTU, and link state.

.. code-block:: console

  user@opsschool ~$ ip addr show
  1: lo: <LOOPBACK,UP,LOWER_UP> mtu 16436 qdisc noqueue state UNKNOWN
      link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
      inet 127.0.0.1/8 scope host lo
      inet6 ::1/128 scope host
         valid_lft forever preferred_lft forever
  2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
      link/ether 08:00:27:8a:6d:07 brd ff:ff:ff:ff:ff:ff
      inet 10.0.2.15/24 brd 10.0.2.255 scope global eth0
      inet6 fe80::a00:27ff:fe8a:6d07/64 scope link
         valid_lft forever preferred_lft forever

ip route
^^^^^^^^

Show all routes on the box.

.. code-block:: console

  user@opsschool ~$ ip route
  10.0.2.0/24 dev eth0  proto kernel  scope link  src 10.0.2.15
  169.254.0.0/16 dev eth0  scope link  metric 1002
  default via 10.0.2.2 dev eth0


ss
--

`ss` is the replacement for `netstat`, which is obsolete according to the `netstat man
page <http://linux.die.net/man/8/netstat>`_.
While most distributions will probably have netstat available for some time, it is
worthwhile to get used to using ss instead, which is already included in the iproute
package.

ss is very useful for checking connections on a box.
ss will show SOCKET, TCP, and UDP connections, in various connection states.
For example, here's ss showing all TCP and UDP connections in the LISTEN state, with
numeric representation.
In other words, this shows all daemons listening on UDP or TCP with DNS and port lookup
disabled.

.. code-block:: console

  user@opsschool ~$ ss -tuln
  Netid  State      Recv-Q Send-Q         Local Address:Port         Peer Address:Port
  tcp    LISTEN     0      128                        *:80                      *:*
  tcp    LISTEN     0      50                         *:4242                    *:*
  tcp    LISTEN     0      50                        :::4242                   :::*
  tcp    LISTEN     0      50                         *:2003                    *:*
  tcp    LISTEN     0      50                         *:2004                    *:*
  tcp    LISTEN     0      128                       :::22                     :::*
  tcp    LISTEN     0      128                        *:22                      *:*
  tcp    LISTEN     0      100                        *:3000                    *:*
  tcp    LISTEN     0      100                      ::1:25                     :::*
  tcp    LISTEN     0      100                127.0.0.1:25                      *:*
  tcp    LISTEN     0      50                         *:7002                    *:*

There are a few things to note in the output.
Local address of * means the daemon is listening on all IP addresses the server might have.
Local address of 127.0.0.1 means the daemon is listening only to the loopback interface, and therefore
won't accept connections from outside of the server itself.
Local address of ::: is the same thing as \*, but for IPv6.
Likewise, ::1 is the same as 127.0.0.1, but for IPv6.
In this example, we used the flags `-tuln`, which just happens to be one of my more-often
used sets of flags.

By default, ss shows only non-listening TCP connections:

.. code-block:: console

  user@opsschool ~$ ss

  State      Recv-Q Send-Q                                    Local Address:Port                                        Peer Address:Port
  ESTAB      0      0                                             10.0.2.15:ssh                                             10.0.2.2:64667

ss has many more useful flags than just these, which you can find in the `ss man page <http://linux.die.net/man/8/ss>`_.

traceroute
----------

If you've familiarized yourself with the basics of networking, you'll know that networks are comprised of
many different routers.
The internet is mostly a messy jumble of routers, with multiple paths to end points.
traceroute is useful for finding connection problems along the path.

traceroute works by a very clever mechanism, using UDP packets on Linux or ICMP packets on Windows.
traceroute can also use TCP, if so configured.
traceroute sends packets with an increasing TTL value, starting the TTL value at 1.
The first router (hop) receives the packet, then decrements the TTL value, resulting in the packet
getting dropped since the TTL has reached zero.
The router then sends an ICMP Time Exceeded back to the source.
This response indicates to the source the identity of the hop.
The source sends another packet, this time with TTL value 2.
The first router decrements it as usual, then sends it to the second router, which decrements to
zero and sends a Time Exceeded back.
This continues until the final destination is reached.

An example:

.. code-block:: console

  user@opsschool ~$ traceroute google.com
  traceroute to google.com (173.194.123.39), 30 hops max, 60 byte packets
   1  (redacted) (redacted)  1.153 ms  1.114 ms  1.096 ms
   2  192.241.164.253 (192.241.164.253)  0.226 ms 192.241.164.241 (192.241.164.241)  3.267 ms 192.241.164.253 (192.241.164.253)  0.222 ms
   3  core1-0-2-0.lga.net.google.com (198.32.160.130)  0.291 ms  0.322 ms 192.241.164.250 (192.241.164.250)  0.201 ms
   4  core1-0-2-0.lga.net.google.com (198.32.160.130)  0.290 ms 216.239.50.108 (216.239.50.108)  0.980 ms  1.172 ms
   5  216.239.50.108 (216.239.50.108)  1.166 ms 209.85.240.113 (209.85.240.113)  1.143 ms  1.358 ms
   6  209.85.240.113 (209.85.240.113)  1.631 ms lga15s47-in-f7.1e100.net (173.194.123.39)  0.593 ms  0.554 ms


mtr
---

mtr is a program that combines the functionality of `ping` and `traceroute` into one utility.

.. code-block:: console

  user@opsschool ~$ mtr -r google.com
  HOST: opsschool                 Loss%   Snt   Last   Avg  Best  Wrst StDev
  1. (redacted)                    0.0%    10    0.3   0.4   0.3   0.5   0.1
  2. 192.241.164.237               0.0%    10    0.3   0.4   0.3   0.4   0.0
  3. core1-0-2-0.lga.net.google.c  0.0%    10    0.4   0.7   0.4   2.8   0.8
  4. 209.85.248.178                0.0%    10    0.5   1.1   0.4   6.2   1.8
  5. 72.14.239.245                 0.0%    10    0.7   0.8   0.7   1.4   0.2
  6. lga15s46-in-f5.1e100.net      0.0%    10    0.5   0.4   0.4   0.5   0.0

mtr can be run continuously or in report mode (-r).
The columns are self-explanatory, as they are the same columns seen when running traceroute
or ping independently.

Reading mtr reports can be a skill in itself, since there's so much information packed into them.
There are many excellent in-depth guides to mtr that can be found online.

iftop
-----

iftop displays bandwidth usage on a specific interface, broken down by remote host.
You can use filters to filter out data you don't care about, such as DNS traffic.
iftop is not available in the base reposities for RHEL/CentOS or Ubuntu, but is available in 
`EPEL <https://fedoraproject.org/wiki/EPEL>`_, and the Universe repository, respectively.

In this example, iftop is listening only to the eth0 interface, and for purposes of this document,
is also using the -t option, which disables the ncurses interface (for your use, you won't need -t).
This box is a very low-traffic VM, so there's not much here, but it does give a sense of what
information is available via the tool.

.. code-block:: console

  user@opsschool ~$ sudo iftop -i eth0 -t
  interface: eth0
  IP address is: 10.0.2.15
  MAC address is: 08:00:27:ffffff8a:6d:07
  Listening on eth0
    # Host name (port/service if enabled)            last 2s   last 10s   last 40s cumulative
  --------------------------------------------------------------------------------------------
    1 10.0.2.15                                =>       804b       804b       804b       201B
      google-public-dns-a.google.com           <=       980b       980b       980b       245B
    2 10.0.2.15                                =>       352b       352b       352b        88B
      10.0.2.2                                 <=       320b       320b       320b        80B
  --------------------------------------------------------------------------------------------
  Total send rate:                                     1.13Kb     1.13Kb     1.13Kb
  Total receive rate:                                  1.27Kb     1.27Kb     1.27Kb
  Total send and receive rate:                         2.40Kb     2.40Kb     2.40Kb
  --------------------------------------------------------------------------------------------
  Peak rate (sent/received/total):                     1.12Kb     1.27Kb     2.40Kb
  Cumulative (sent/received/total):                      289B       325B       614B
  ============================================================================================

iperf
-----

iperf is a bandwidth testing utility.
It consists of a daemon and client, running on separate machines.

This output shows from the client's side:

.. code-block:: console

  user@opsschool ~$ sudo iperf3 -c remote-host
  Connecting to host remote-host, port 5201
  [  4] local 10.0.2.15 port 45687 connected to x.x.x.x port 5201
  [ ID] Interval           Transfer     Bandwidth       Retr  Cwnd
  [  4]   0.00-1.00   sec   548 KBytes  4.48 Mbits/sec    0   21.4 KBytes
  [  4]   1.00-2.00   sec   503 KBytes  4.12 Mbits/sec    0   24.2 KBytes
  [  4]   2.00-3.00   sec   157 KBytes  1.28 Mbits/sec    0   14.3 KBytes
  [  4]   3.00-4.00   sec  0.00 Bytes  0.00 bits/sec    0   14.3 KBytes
  [  4]   4.00-5.00   sec   472 KBytes  3.88 Mbits/sec    0   20.0 KBytes
  [  4]   5.00-6.00   sec   701 KBytes  5.74 Mbits/sec    0   45.6 KBytes
  [  4]   6.00-7.00   sec   177 KBytes  1.45 Mbits/sec    0   14.3 KBytes
  (snip)

Some of the really handy options:

-m - Use the maximum segment size (the largest amount of data, in bytes, that a system can support
in an unfragmented TCP segment).
This option will use the default size for the particular network media in use (eg, Ethernet is 1500 bytes).

-M - Set MSS, used in conjuction with the previous -m option to set the MSS to a different value than default.
Useful for testing performance at various MTU settings.

-u - Use UDP instead of TCP.
Since UDP is connectionless, this will give great information about jitter and packet loss.

tcpdump
-------

For occasions where you need to get into the nitty-gritty and look at actual network behavior,
tcpdump is the go-to tool.
tcpdump will show raw connection details and packet contents.

Since one could devote entire pages to the usage of tcpdump, it is recommended to search online
for any one of the many great guides on tcpdump usage.


Troubleshooting layer 1 problems
================================

Layer one problems (physical) deserve their own section, because of how cryptic
they can seem without a solid understanding of the causes and symptoms.
There are several data points that fall under layer one.

Most layer problems are usually traced back to one of the following causes:

* Bad cabling
* Mismatched duplex/speed
* Electromagnetic interference (EMI)
* Network congestion

To find the metrics below, use either `ip -s link`, or `ethtool -S`.

Example output of `ip -s link show eth0`:

.. code-block:: console

  user@opsschool ~$ ip -s link show eth0
  2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
      link/ether 04:01:1c:dc:82:01 brd ff:ff:ff:ff:ff:ff
      RX: bytes  packets  errors  dropped overrun mcast
      150285334  221431013 0       0       0       0
      TX: bytes  packets  errors  dropped carrier collsns
      999121259  42410636 0       0       0       0

Here we can see that the link is up (`state UP`), the MAC address
(`04:01:1c:dc:82:01`), bytes, packets, errors, drops, overruns, multicast,
carrier, and collisions, for both RX (receive) and TX (transmit).

One thing you may notice is that overrun and mcast are only on RX, and
carrier/collisions are only on TX.
This is because overruns and multicast only affect the receiving side, while carrier
and collisions only affect the transmitting side.
The reason for this will be made clear when described below.

Example output of `ethtool eth0`:

.. code-block:: console

  user@opsschool ~$ sudo ethtool eth0
  Settings for eth0:
    Supported ports: [ TP ]
    Supported link modes:   10baseT/Half 10baseT/Full
                            100baseT/Half 100baseT/Full
                            1000baseT/Full
    Supported pause frame use: No
    Supports auto-negotiation: Yes
    Advertised link modes:  10baseT/Half 10baseT/Full
                            100baseT/Half 100baseT/Full
                            1000baseT/Full
    Advertised pause frame use: No
    Advertised auto-negotiation: Yes
    Speed: 1000Mb/s
    Duplex: Full
    Port: Twisted Pair
    PHYAD: 0
    Transceiver: internal
    Auto-negotiation: on
    MDI-X: Unknown
    Supports Wake-on: umbg
    Wake-on: d
    Current message level: 0x00000007 (7)
               drv probe link
    Link detected: yes

As we can see, ethtool gives a lot of information, and with additional flags,
there's even more.
ethtool is especially useful for looking into layer 1 configuration, due to
how much information it can provide.
A particularly useful option is `-S`, which shows all the same metrics as
the `ip` command mentioned above, plus many, many more:

.. code-block:: console

  user@opsschool ~$ sudo ethtool -S eth0
  NIC statistics:
      rx_packets: 20831
      tx_packets: 11160
      rx_bytes: 14654723
      tx_bytes: 3637509
      rx_broadcast: 0
      tx_broadcast: 9
      rx_multicast: 0
      tx_multicast: 11
      rx_errors: 0
      tx_errors: 0
      tx_dropped: 0
      multicast: 0
      collisions: 0
      rx_length_errors: 0
      rx_over_errors: 0
      rx_crc_errors: 0
      rx_frame_errors: 0
      rx_no_buffer_count: 0
      rx_missed_errors: 0
      tx_aborted_errors: 0
      tx_carrier_errors: 0
      tx_fifo_errors: 0
      tx_heartbeat_errors: 0
      tx_window_errors: 0
      tx_abort_late_coll: 0
      tx_deferred_ok: 0
      tx_single_coll_ok: 0
      tx_multi_coll_ok: 0
      tx_timeout_count: 0
      tx_restart_queue: 0
      rx_long_length_errors: 0
      rx_short_length_errors: 0
      rx_align_errors: 0
      tx_tcp_seg_good: 304
      tx_tcp_seg_failed: 0
      rx_flow_control_xon: 0
      rx_flow_control_xoff: 0
      tx_flow_control_xon: 0
      tx_flow_control_xoff: 0
      rx_long_byte_count: 14654723
      rx_csum_offload_good: 0
      rx_csum_offload_errors: 0
      alloc_rx_buff_failed: 0
      tx_smbus: 0
      rx_smbus: 0
      dropped_smbus: 0

Common Network metrics
----------------------

RX/TX errors
^^^^^^^^^^^^

These counters appear to be aggregates of various other counters, and information
on exactly what's included in them is sparse.
An increase of these counters is an indication that *something* is wrong, but
they're too general to determine the exact cause.
If you see these increasing, it will be more beneficial to check `ethtool -S`
and see exactly which counters are increasing.

Cyclic Redundancy Check (CRC)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

CRC is a sort of checksum.
CRC errors can be a symptom for many different issues, such as a bad cable,
noisy lines, and more.

Drops
^^^^^

Drops occur when the network link is simply too saturated.
The solution is to increase the bandwidth available, perhaps by upgrading the
network link, or implementing some form of traffic limitation such as
Quality-of-Service (QoS).
Traffic limitation/shaping is out of the scope of this section, and should be
considered an advanced topic.
You can verify the saturation by checking your graphs for RX/TX bytes and see how
much data is going through.

Overruns
^^^^^^^^

Overruns occur when the network device is receiving data faster than the kernel
driver can process it.
Overruns sound similar to drops, but overruns have to do with the kernel's ability
to process data, rather than the network link itself.
The solution is either to replace the network device with a more capable one, or
lower the amount of data coming in to the network device.

Carrier
^^^^^^^

Carrier errors occur when the link signal is having issues.
In the absence of collisions, this can usually be attributed to a bad cable
or network device.
If collisions are also increasing, check the duplex settings.

Collisions
^^^^^^^^^^

Collisions occur when there is a duplex or speed mismatch.
Verify duplex and speed at both ends of the cable (at the switch and at the server).
Auto-negotiation being enabled on one end but not on the other is often a
cause of duplex/speed mismatch.
Some devices don't play nicely with each other when auto-negotiation is
enabled on both sides, so be sure to verify what auto-negotiation resulted in.

Electromagnetic interference (EMI)
----------------------------------

The symptoms of electromagnetic interference can be seen by an increase of CRC errors.
Electromagnetic interference can be caused by any number of things: power cables nearby,
mechanical systems (such as HVAC units), etc.
Since a visual inspection will often reveal electromagnetic interference culprits easier
than checking metrics, it's simpler to do a visual inspection first if you suspect
this.
The problem can sometimes be solved/alleviated by replacing the cable with STP
(Shielded Twisted Pair), but it's usually better to just relocate the network cable.
Only copper cabling is suspectible to EMI.
Fiber is immune to EMI, as signals are transmitted as light instead of current.

Testing copper cabling
----------------------

A link tester is an inexpensive device made by many different vendors, for the
purpose of verifying that a cable's functionality.
Link testers only verify a small amount of criteria, mainly that there are no
opens or shorts.
A link tester, however, is different from a cable certifier, which will verify
full compliance with TIA/EIA and ISO cable standards, including metrics for
crosstalk allowance, cable length, current loss, and many others.

If you need to test cable used in low-impact environments (e.g., wall jack to
desktop), a link tester is usually sufficient.
If you need to test cable used in high-impact environments (e.g., datacenter),
opt for a cable certifier.
Reputable cable vendors will have already run certification tests on cable,
so this is only useful if you are routinely making your own cable (which
is not advised under most circumstances).

Many vendors sell great testers, with Fluke Networks being one of the primary
vendors.

Fiber errors
------------

Issues with fiber cable generally fall into two categories:

* Dirty/Scratched fiber
* Bad optic/transceiver

The symptoms for both are very similar: intermittent RX/TX and CRC errors
indicate a dirty or scratched fiber, while persistent RX/TX and CRC errors
indicate a bad optic/transceiver.


Differences in perspective: network engineering and systems administration
==========================================================================

Network engineering and systems administration have a tendency to speak different languages,
due to the divide in skillsets and lack of overlap.

A good way to view how network engineering sees the technology differently is consider
the OSI model: network engineering is focused primarily on layers 1, 2, and 3, with the
occasional venture into the higher layers when certain routing protocols are involved (eg,
BGP peering sessions operate over TCP).
System administrators, on the other hand, are typically more concerned with layers 4 through 7,
with the occasional venture into layer 3 for IP addressing.
If one considers the perspective of the other, empathy is understanding comes easier, and
anticipating what the other side expects becomes straightforward.

As such, here are a few tips on how the two specializations see the same technology:

1. Network engineers output in bits-per-second (bps), while many server-specific utilities
   output in bytes-per-second (Bps).
   As such, be sure when you're sending throughput data to network engineering that it's in
   bits-per-second.
   Your monitoring tools will usually do this for you (look for a config option).
   In the occasion you need to do it by hand, and you're working in bytes, simply multiply by eight
   to get bits.
   For more information on conversions, `wikipedia <http://en.wikipedia.org/wiki/Data_rate_units>`_
   has a good article on unit measurements.
   Alternatively, use an online calculator.
2. Systems administrators often don't worry about network topology since it so often "just works".
   However, in some cases, especially situations where you're troubleshooting hosts across the
   open internet, you may run into something called an 'asymmetrical path', that is, the routing
   is using a different path out than it does coming back in.
   In such situations, one path might have issues, while the other path is perfectly fine.
   For this reason, when sending issue reports to a network engineer, be sure to send
   a `traceroute`/`mtr` report from *both* directions.
   This situation is not common on internal networks, but can be on the open Internet.
   You may also run into this situation when you have more complex routing set up on the local server.
3. A simple `ping` sometimes isn't enough for an issue report, for the simple reason that it contains
   so little information.
   A `traceroute`/`mtr` report is better.
   If there's suspected throughput issues, try to get an `iperf` report as well.
   Also include an interface configuration report, showing MTU, IP address, and MAC address
   of the relevant interface(s).

