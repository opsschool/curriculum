Load Balancing
**************

Why do we use load balancers?
=============================

* A fast, HTTP/TCP based load balancer can distribute between slower application
  servers, and can keep each application server running optimally
* They keep application servers busy by buffering responses and serving them to
  slow clients (or keepalives). We want app servers to do real work, not waste
  time waiting on the network.
* Load balancers provide a mechanism to verify the health of backend servers to
  ensure traffic is only routed to backends available to service the request.
  These health checks can be simplistic ICMP pings to ensure host availability
  to advanced HTTP (layer 7) health checks that use HTTP response codes or
  content matching.
* Can provide a layer of abstraction such that end user endpoints remain
  consistent (i.e URLs) as backend application infrastructure expands
  (or contracts)
* Can be either software based (aka reverse proxies) or hardware based (physical
  devices optimized for network throughput and ISO layer 2,3 and 4 routing).
* Can be tuned by itself for max conns?
* HTTP path based routing, 1 domain to multiple heterogeneous server pools.

Application implications
========================

Backend
-------

An application server behind a load balancer is typically referred to as a
"backend" or "backend server".

Pools
-----

A collection of backends is referred to as a pool.

Balancing algorithms
--------------------

Session Affinity
----------------

Session affinity or session stickiness, is when the load balancer applies an
algorithm to the incoming connection to direct it to a single server. This
is typically done for HTTP applications by setting cookies. In general TCP
applications often use IP addresses to determine the server to direct
the connection to.

Local & ISP caching
-------------------

SSL Termination and Offloading
------------------------------

SSL Termination is when the load balancer established an SSL connection
between the client and the load balancer and a non-encrypted connection between
the load balancer and backend server.

Terminating SSL on the load balancer eliminates the need to distribute your
certificate and key amongst all the servers. When using hardware load balancers
they typically have special hardware acceleration which is much more performant
compared to terminating connections on the backend server.

Balancing vs. Failover (Active / Passive)
-----------------------------------------

Often load balancing is used as a high-availability technique, by allowing
multiple backends to service a request if one node should become unavailable. It
differs from failover configuration because all nodes generally participate in
servicing clients. In failover configurations a single (active) node handles all
requests until an issue arises, and the secondary (passive) node takes over all
of the incoming traffic. Failover configurations are usually not configured for
scaling purposes.

Health Checks
---------------

Most load balancers have some ability to test backend servers for availability,
these are called health checks. They can be as simple as whether hosts are
responding on the port the application is bound to, or complex configurations
that test special URIs, or response times. If a server fails a health check the
load balancer will temporarily remove the node from the pool, until it
successfully passes the health check.

Non-HTTP use cases
==================

Native SMTP, or general TCP for things like RabbitMQ, MongoDB, etc

Software
========

Apache
------

Apache has the ability to load balance using ``mod_proxy_balancer``, and ``mod_jk``.

Mod_proxy_balancer is purpose-built load balancing solution. It supports the HTTP, FTP,
and AJP protocols. There is basic support for sessions stickiness's, weighted round-robin,
and can remove unhealthy backends from the balancing pool. It lacks support for customizable
health checks, other TCP protocols. Support for AJP is provided by ``mod_proxy_ajp`` and support
for FTP is provided by ``mod_proxy_ftp``.

``Mod_jk`` supports the AJPv1.2 and AJPv1.3 protocols. It's a purpose build connector to support
the Java application servers, like Apache Tomcat. It supports weighted round-robin, session
stickiness, session counting, traffic based routing, and backend busyness (server with the least
connections). It supports health checks using the method built into AJP protocol.


Nginx
-----

HAProxy
-------

HAProxy is a general TCP load balancing server that is highly configurable. It
will generally support any TCP based protocol, and has special modes for HTTP,
RDP, MySQL, and Postgresql protocols. It has support for multiple types of
health check including URL based, traffic-based health, and external checks via
the ``httpchk`` options. It has several load balancing algorithms: round robin,
static round-robin, least connections, source hashing, URI hashing, URI
parameter, and RDP-cookie.


Hardware
========

BIG-IP
------

BIG-IP has purpose-built hardware load balancers. They support protocols in layers
2, 4, and 7 of the OSI model. They allow for very complex configurations, and
support writing special TCL programs to modify the load balancing behavior. The
product supports SSL termination and offloading, with additional licensing.

Netscaler
---------

Multi-dc
========

Anycast
-------

DNS GSLB
--------
* A GSLB (Global Site Load Balancer) at the most simplistic level is a health
  checking DNS server.
* Most often used to load balance between geographically dispersed data centers.
* Generally has health check mechanisms similar to load balancers which are used
  to return an IP address (as part of the DNS lookup) of a host that is currently
  available to service the request.
* Conceptually provides coarse-grained round robin and affinity balancing
  algorithms by setting the time to live (TTL) of the DNS lookup for an
  appropriate duration.

CDN's
-----

(cparedes: I'd argue that it's valid in some contexts, depending on what
you're load balancing)

