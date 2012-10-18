Load Balancing
**************

Why do we use load balancers?
=============================

* A fast, HTTP/TCP based load balancer can distribute between slower application
  servers, and can keep each application server running optimally
* They keep application servers busy by buffering responses and serving them to
  slow clients (or keepalives). We want app servers to do real work, not waste
  time waiting on the network.
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
