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
  ensure traffic is only routed to backends able to service the request.
  These health checks can be simplistic ICMP pings to ensure host availabilty
  to advanced HTTP (layer 7) healtch checks that use HTTP response codes or 
  content matching.
* Can provide a layer abstraction such that end user endpoints remain consistent
  (i.e URLs) as backend application infrastructure expands (or contracts) 
* Can be either software based (aka reverse proxies) or hardware based (physical
  devices optimized for network throughput and ISO layer 2,3 and 4 routing).
* Can be tuned by itself for max conns?
* HTTP path based routing, 1 domain to multiple heterogeneous server pools.

Software
========

Apache
------

Nginx
-----

HAProxy
-------

Hardware
========

BIG-IP
------

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
* Most often used to load balance between geographical disperse data centers.
* Generally has health check mechanisms similar to load balancers which are used
to return an IP address (as part of the DNS lookup) of a host that is currently
able to service the request.
* Concenptually provides coarse grained round robin and affinity balancing
algorithms by setting the time to live (TTL) of the DNS lookup for an 
appropriate duration.

CDN’s
-----

(cparedes: I’d argue that it’s valid in some contexts, depending on what
you’re load balancing)

Application implications
========================

Balancing algorithms
--------------------

Session Affinity, and why it is evil.
-------------------------------------

Local & ISP caching
-------------------

SSL Termination
---------------

Balancing vs. Failover (Active / Passive)
-----------------------------------------

Health Checks. 
---------------

Non-HTTP use cases
==================

Native SMTP, or general TCP for things like RabbitMQ, MongoDB, etc
