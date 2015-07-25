Application Components 201
**************************

Message Queues Systems
======================
Messages Queue Systems can be used in many different ways.
This includes but is not limited to
 * decouple processes or hosts from each other
 * parallel processes execution
 * distribute events or information
 * replace polling mechanism with an event system
 * asynchronous remote procedure calls
 * job queues
 * network wide mutex
 * data transfer to multiple hosts
 * message rerouting on failure / escalation schemes
 * data-streaming
 * network or system wide message bus


Message Brokers
===============
Is one class of message queue systems that relies on a central system or cluster to route messages to its destination.


RabbitMQ
--------
RabbitMQ is an erlang based implementation of the Advance Message Queuing Protocol.
It also provides plugins for different to queuing protocols like STOMP or MQTT.

To setup a high availability cluster a minimum of two nodes are required.
A disc-node ensures messages are stored to a disc before delivering it.
Memory based nodes do provide a much higher message throughput but messages in queues are lost during shutdown of the node.
To establish an encrypted communication between two nodes, IP-Sec or TLS can be used.
For authentication it is possible to use an internal database, LDAP, SASL or PKI client certificates.
Different applications can be separated by the concept of virtual hosts in a similar way Apache web server does.
For configuration a webGUI or the command line tool can be used.
Programming libraries and tools are provided for a wide range of environments (see https://www.rabbitmq.com/devtools.html)

backup configuration with

.. code-block:: console

  curl http://guest:guest@localhost:15671/api/definitions > /var/backups/rabbitmq_config

full report:

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

