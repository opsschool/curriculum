Application Components 201
**************************

Message Queue Systems
======================
In a distributed system with the function of sending and receiving messages a Message Queue Systems can provid but is not limited to

 * decoupling processes or hosts from each other
 * parallel processes execution
 * distribute events or information like changes to data, entries from logfiles and statistics for monitoring and notification
 * datastreaming and filetransfers to multiple hosts via multicast and unicast
 * replace polling mechanism with an event system
 * network wide accessible job queues
 * message rerouting on failure / escalation schemes
 * verification of results ( best of three )
 * asynchronous remote procedure calls
 * network wide mutex
 * network or distributed system wide message bus

The two major types are message brokers and routers or brokerless message queuing systems.
   

Message Brokers
===============
Is a class of message queue systems that relies on a central system or cluster to route messages to its destination.
Messages are send to exchanges and received from queues.
Queues resists on the broker, also some protocols implement a clientsite queue for transaction like messaging and prefetching.
Message Brokers tend to become critical system in an otherwise distributed system.


RabbitMQ
--------
RabbitMQ is an open source Message Broker written in erlang.
Its main protocol is the Advance Message Queuing Protocol. (see https://en.wikipedia.org/wiki/AMQP)
It provides a plugin system for different queuing protocols or advanced setups.
Installations scale from single host to multi cluster installations in different networks and bussines critical environments.
Messages are send to different type of exchanges.
While Direct, Fanout, Topic and Header Exchanges are part of the core system, others plugins are rooted in the community.
( https://www.rabbitmq.com/plugins.html , https://www.rabbitmq.com/community-plugins.html )
No prior understanding of the erlang programming language is needed to configure or operate the message broker.
The documentation on the projects homepage reads a fair amount of different configurations with explainations.

To setup a high availability cluster a minimum of two nodes is required.
Also it is possible to send huge messages via AMQP to a RabbitMQ node, the maximum size of a message is limited by the amount of available RAM on the nodes.
Depending on the durability of a message and the policy of queues, the messages are syncronised to other nodes.
A disc-node ensures a message is stored to a disc before delivering.
Memory based nodes do provide a much higher message throughput.
To establish an encrypted communication between nodes, IPSec or TLS can be used.
Authentication is possible via an internal database, LDAP, SASL and PKI client certificates.
With a PKI in place consumers and publisher can share the same autorisation without the need for a prehared password among all processes.
This is accomplished by using the commonname of the dn in the x509 client certificate.
Different applications can be separated by the concept of virtual hosts in a similar way the Apache web server does.
Configuration is provided by webGUI or commandline tools.
Programming libraries and tools for a wide range of environments are available (see https://www.rabbitmq.com/devtools.html and https://www.rabbitmq.com/tutorials/amqp-concepts.html )

To backup and restore configuration data the management plugin must be installed and configured.

.. code-block:: console

  rabbitmqadmin export rabbit.config
  rabbitmqadmin -q import rabbit.config


Show a detailed report:

.. code-block:: console

  rabbitmqctrl report


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

