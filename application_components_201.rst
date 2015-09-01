Application Components 201
**************************

Message Queue Systems
======================
In a distributed system a Message Queue Systems can provide a basic infrastructure to higher level functions like

 * decoupling processes or hosts from each other
 * parallel processes execution
 * distribute events and information like changes to data, entries from logfiles and statistics for monitoring and notification
 * data-streaming and file-transfers to multiple hosts via multicast and unicast
 * replace polling mechanism with an event orientated system
 * network wide accessible job queues
 * message rerouting on failure / escalation schemes
 * verification of results ( best of three )
 * asynchronous remote procedure calls
 * network wide mutex
 * network or distributed system wide message bus

Message Queue Systems are more like peer2peer networks then client-server applications.
They can be split up into message brokers and routers or brokerless message queuing systems. [0MQ]_
A message as understood by the system is everything that can be represented as a bytestream.
Properties like type and timestamp may are added to message. [Bok]_

.. [Bok] http://sardes.inrialpes.fr/papers/files/Bouchenak08a.pdf
.. [0MQ] https://en.wikipedia.org/wiki/%C3%98MQ


Further Readings
----------------

 * http://c2.com/cgi/wiki?MessageQueuingArchitectures
 * https://en.wikipedia.org/wiki/Message-oriented_middleware and https://en.wikipedia.org/wiki/Message_oriented_middleware
 * https://en.wikipedia.org/wiki/Queueing_theory
 * https://en.wikipedia.org/wiki/Message_queue
 * http://www.slideshare.net/mwillbanks/art-of-message-queues

Message Brokers
===============
Message brokers represent a message queue systems that relies on a central system to route messages to its destination.
Messages are send to exchanges and received from queues.
Queues resists on the broker, also some protocols implement a client site queue for multipart transactions and prefetching.
Message Brokers tend to become a critical system in an otherwise distributed environment.


RabbitMQ
--------
RabbitMQ_ is an open source Message Broker written in erlang.
Its queuing protocol is the Advance Message Queuing Protocol. [wiki]_  [specs]_ 
It provides plugins for different queuing protocols like STOMP and advanced setups.
Installations scale from single host as an applicationmessagebus, to multi cluster installations in different networks.
Messages are send to different type of exchanges.
While Direct, Fanout, Topic and Header Exchanges are part of the core system, others plugins are rooted in the community. 
No prior understanding of the erlang programming language is needed to setup, configure and operate the message broker.
The documentation on the projects homepage reads a fair amount of different configurations with explanations.

To setup a cluster a minimum of two nodes is required.
All nodes of a cluster must be in the same IP network segment.
It is possible to send huge messages via AMQP to a RabbitMQ node.
The maximum message size is limited by the amount of available RAM on the node.
Depending on the durability of a message and the policy of queues, the messages are synchronized to other nodes.
A disc-node ensures a message is stored to a disc before delivering.
Memory based nodes provide a much higher message throughput.
To establish an encrypted communication between nodes, IPSec or TLS can be used.
Authentication is possible via an internal database, LDAP, SASL and PKI client certificates.
With a PKI in place consumers and publisher can share the same authorization.
No preshared password among all processes is needed, but different private keys are mandatory.
This is accomplished by using the commonname of the dn in the x509 client certificate as username.
Different applications can be separated by the concept of virtual hosts in a similar way the Apache web server does.
Configuration is provided by webGUI or commandline tools.
Programming libraries and tools for a wide range of environments are available.

The installation from the Debian repository leads to a configuration with a single host.
Similar installation instructions are provided for Windows, Solaris and other operation systems.

.. code-block:: bash

  curl 'https://www.rabbitmq.com/rabbitmq-signing-key-public.asc' | sudo apt-key add -
  echo 'deb http://www.rabbitmq.com/debian/ testing main' | sudo tee /etc/sources.d/rabbitmq.list
  sudo apt-get update
  sudo apt-get upgrade
  sudo apt-get install rabbitmq-server

To backup and restore configuration data the management plugin should be configured.

.. code-block:: none

  { rabbitmq_management, [
      {  listener, [ { port, 15672 },
                     { ssl, true },
                     { ssl_opts, [ { cacertfile, "/etc/ssl/certs/cacert.crt" },
                                   { certfile,   "/etc/ssl/certs/node1.crt" },
                                   { keyfile,    "/etc/ssl/private/node1.key" } ]}
                   ]} // configured listener
  ]} // configured rabbitmq_management
  ]. // EOC

.. code-block:: console

  rabbitmqadmin export rabbit.config
  rabbitmqadmin -q import rabbit.config

Show a detailed report about queues, users and connections

.. code-block:: console

  rabbitmqctl report

.. _RabbitMQ: https://www.rabbitmq.com
.. [wiki] https://en.wikipedia.org/wiki/AMQP
.. [specs] http://www.amqp.org/resources/download

Apache ActiveMQ
---------------

Memory Caches
=============

Memcached
---------

Redis
-----

Specialized Caches
==================

Varnish
-------

nginx+memcached
---------------

