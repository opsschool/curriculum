.. _conventions:

###########
Conventions
###########

This set of documents uses a number of conventions throughout for ease
of understanding. Contributors should be familiar with everything
described here, and optimally contributions will be compared to what's
here for uniformity.


Style Guide
===========

See :doc:`style_guide` for more.

.. _sample-network:

Sample Network
==============

A sample network makes explaining networking topics easier for both the
teacher and the student. The teacher doesn't need to spend time creating
an example network to explain a topic, and the student doesn't need to
learn a new network for every section. To this end, the curriculum has
adopted a default sample network. The goal is to construct a
hypothetical network that resembles what one is likely to find in a
well-constructed production environment. This means firewalls, load
balancers, separation of concerns between different machine roles, and
so on.

Topography
----------

The network is hosted in a datacenter, with the subnet of 10.10.10.0 and a
subnet mask of 255.255.255.0. This yields an IP address range of 10.10.10.0
through 10.10.10.255. In front of everything else on the network are two
firewall machines, fw-1 and fw-2. These dual-homed machines have the LAN
addresses 10.10.10.1 and 10.10.10.2, respectively. The network also has two
DNS servers, dns-1 and dns-2, occupying 10.10.10.3 and 10.10.10.4. In between
the firewalls and the rest of the LAN are two load balancers, lb-1 and lb-2.
These machines share the rest of the public address space and have the internal
LAN addresses 10.10.10.5 and 10.10.10.6.

Beyond this, there are two machines that function as jumphosts into the
network (independent of any VPN or the like). These machines also host
the source repositories, configuration management, and similar services.
Their LAN addresses are 10.10.10.7 and 10.10.10.8, and their hostnames
jump-1 and jump-2. In addition, there is an LDAP server, ldap-1, at
10.10.10.9.

In terms of the application stack, there are four Web servers
(10.10.10.10-10.10.10.13), two app servers (10.10.10.14-10.10.10.15), two
database servers (10.10.10.16-10.10.10.17), two DNS servers
(10.10.10.18-10.10.10.19), and a message queue (10.10.10.20). The Web server
hostnames are web-1 through web-4; the app servers, app-1 and app-2; the
databases, db-1 and db-2; the message queue, mq-1.

.. todo:: Mainly for @apeiron: Simplify the topography description, possibly with use of a table to describe IP assingments

This is a lot to conceptualize all at once; a diagram will elucidate the
architecture.

.. graphviz::

    digraph network {
      "firewall x2" [shape = box];
      "load balancer x2" [shape = box];
      "app stack" [shape = box];
      "jumphost x2" [shape = box];
      "config mgmt" [shape = box];
      "dns" [shape = box];
      "mq" [shape = box];
      "internet" [shape = ellipse];
      "internet" -> "dns";
      "internet" -> "firewall x2" -> "load balancer x2" -> "app stack";
      "firewall x2" -> "jumphost x2" [dir = both];
      "jumphost x2" -> {"app stack"; "config mgmt"; "load balancer x2"; "mq"; "ldap"};
      "app stack" -> "config mgmt" [dir = both];
      "app stack" -> {"ldap"; "mq"};
      "firewall x2" -> "dns";
      "load balancer x2" -> "dns";
      "app stack" -> "dns";
      "jumphost x2" -> "dns";
      "config mgmt" -> "dns";
      "mq" -> "dns";
    }
