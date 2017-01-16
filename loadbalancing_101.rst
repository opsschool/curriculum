Load Balancing
**************

Why do we use load balancers?
=============================

Scalability
-----------

The most basic scalability need is to allow one address handle a load that's too much for one physical server.
A secondary scalability factor is offloading or reducing underlying tasks, such as SSL termination and re-using already-established TCP links between the load balancer and the servers to reduce the time taken to send data to the user.
Smart load balancers that are able to poll and react to client load can seamlessly distribute work among heterogeneous machines.

Reliability
------------

Simple health checking alone increases reliability.
A load balancer that simply polls an endpoint like status/health can detect a machine or application that has failed and remove it from the pool of active servers.
More advanced health checking can remove partially failed systems that are exhibiting poor or intermittent performance.

Agility
-------

A load balancer can assist in making application development faster too.
Load balancers could be configured to route requests using higher level information about the connection,
such as certain URLs served by certain types of servers,
or sending everybody who logs in as an employee to an internal staging server instead of production.
Finally, a load balancer permits various different strategies for introducing new deployments in stages.


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

Nginx reverse proxy implementation supports load balancing for HTTP, HTTPS, FastCGI, uwsgi, SCGI, 
and memcached. The default method is round-robin, but session persistence/stickiness (``ip_hash``) 
and least connections (``least_conn``) methods are available. Flags for weighted load balancing and passive 
server health checks can also be set on any of the backends.

`Nginx load balancing documentation`_

.. _Nginx load balancing documentation: http://nginx.org/en/docs/http/load_balancing.html


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

