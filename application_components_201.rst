Application Components 201
**************************

Message Queues Systems
======================
Messages Queue Systems can be used in many different ways.
This includes but is not limited to
 * decouple processes or hosts from each other
 * parallize processes execution
 * distribute events or informations
 * replace polling mechanisment with an event system
 * asyncron remote procedure calls
 * job queues
 * network wide mutex
 * data transfer to multiple hosts
 * message rerouting on failure / escalation schemas
 * datastreaming
 * network or systemwide message bus


Message Brokers
===============
Is one class of message queue systems that relies on a central system or cluster to route messages to its destination.


RabbitMQ
--------
RabbitMQ is an erlang based implementation of the Advance Message Queueing Protocoll.
It also provides plugins for different to queuing protocols like STOMP or MQTT.

To setup a high availability cluster a minimum of two nodes are requiered.
A disc-node ensures messages are stored to a disc before delivering it.
Memory based nodes do provide a much higher message thougtput but messages in queues are lost durring shutdown of the node.
To establish an encrypted communication between two nodes, ipsec or TLS can be used.
For authentification it is possible to use an internal database, LDAP, SASL or PKI client certificates.
Different applications can be seperated by the concept of virtual hosts in a similar way apache webserver does.
For configuration a webgui or the commandline tool can be used.
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

