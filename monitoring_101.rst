Monitoring, Notifications, and Metrics 101
******************************************

History: How we used to monitor, and how we got better (monitors as tests)
==========================================================================

Perspective (end-to-end) vs Introspective monitoring
====================================================

Metrics: what to collect, what to do with them
==============================================

Common tools
============

Syslog (basics!)
----------------

Syslog-ng
---------

Nagios
------

Graphite
--------

Ganglia
-------

Munin
-----

RRDTool / cacti
---------------

Icinga
------

SNMP
----

Collectd
--------

`Collectd <https://collectd.org>`_ collects system-level metrics on
each machine.  It works by loading a list of plugins, and polls data
from various sources.  The data are sent to different backend
(Graphite, Riemann) and can be used to trigger alerts with Nagios.

Sensu
-----
`Sensu <https://github.com/sensu>`_ was written as a highly
configurable, Nagios replacement. Sensu can be described as a
"monitoring router", since it connects check scripts across any number
of systems with handler scripts run on one or more Sensu servers. It
is compatible with existing Nagios checks and additional checks can be
written in any language similar to writing Nagios checks. Check
scripts can send alert data to one or more handlers for flexible
notifications. Sensu provides the server, client, api and dashboard
needed to build a complete monitoring system.

Diamond
-------
`Diamond <https://github.com/BrightcoveOS/Diamond>`_ is a python daemon
that collects system metrics and publishes them to Graphite
(and others). It is capable of collecting cpu, memory, network, i/o,
load and disk metrics. Additionally, it features an API for implementing
custom collectors for gathering metrics from almost any source.
