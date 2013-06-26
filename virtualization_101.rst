Virtualization 101
******************

Intro to virtualization technologies
====================================

What problems is virtualization good at solving, what isn't it good at.
-----------------------------------------------------------------------

Hardware assisted vs Paravirtualization
---------------------------------------

Hypervisors - overview, some comparison
---------------------------------------

Container-based virtualization (Jails, Containers/Zones, OpenVZ)
----------------------------------------------------------------

The Cloud
=========

The different types of "cloud"
------------------------------

Unfortunately "cloud" is a term which is overloaded in the industry. This course refers to cloud in the Infrastructure as a Service (IaaS) sense, not the hosted applications sense.

There are currently two types of IaaS cloud being deployed. These are "public clouds", which are services such as Amazon's AWS, Rackspace's OpenStack based cloud offering or HP's OpenStack based cloud. There are a variety of companies of various sizes offering these public cloud services. These services are run in the hosting company's datacenter, and an administrator's access to the infrastruture is limited to launching and configuring virtual machines, normally through a web-based dashboard or an API.

The second type of cloud is "private clouds". These are conceptually similar to public clouds, but run in your own datacenter. The point here is that you retain complete control over the environment -- if some of your virtual machines need access to unusual hardware, or store data which you're unwilling to move offsite, then private cloud might be a good choice. The open source OpenStack system is often used to build these private clouds, and there are various vendors offering private cloud solutions, including RackSpace, Piston, and Canonical.

Cloud providers comparison. Pros & Cons
---------------------------------------

Cloud management tools and APIs
-------------------------------
