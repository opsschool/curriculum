Labs exercises
**************

Bare-Metal Provisioning 101
===========================

Install CentOS by hand
----------------------

Install the same node using kickstart
-------------------------------------

Bare-Metal Provisioning 201
===========================

Setup a basic cobbler server
----------------------------

Build a profile
---------------

Kickstart a node
----------------

Change the profile, re-kickstart the node
-----------------------------------------

Cloud Provisioning 101
======================

Setup a free-tier account with AWS
----------------------------------

Spawn a Linux node from the AWS Console
---------------------------------------

Cloud Provisioning 201
======================

Spawn a Linux node using the AWS API
------------------------------------

Attach an Elastic IP and EBS to it
----------------------------------


Database 101
============

Install and start up MySQL
--------------------------

Create basic relational database / tables, using a variety of field types
-------------------------------------------------------------------------

Grant and revoke privileges
---------------------------

Install Riak
------------

Write (or provide, probably depends on where this fits in relation
to scripting tutorial?) basic tool to insert and retrieve some data.

Database 201
============

Spawn up second VM/MySQL install
--------------------------------

Set up Master->Slave replication
--------------------------------

Deliberately break and then fix replication
-------------------------------------------

Set up Master<->Master replication
----------------------------------

Introduction to percona toolkit
-------------------------------

Set up Riak Cluster, modify the tool from 101 to demonstrate replication.

Database 301
============

Galera cluster
--------------

* Introduction to variables and their meaning. Tuning MySQL configuration (use
  mysqltuner.pl as a launch point?), pros and cons of various options.
* Introducing EXPLAIN and how to analyse and improve queries and schema.
* Backup options, mysqldump, LVM Snapshotting, Xtrabackup.

Automation 101
==============

Do you need it? How much and why?
---------------------------------

http://www.kitchensoap.com/2012/09/21/a-mature-role-for-automation-part-i/

* Talk basic theory and approach in terms of Idempotency and Convergence.
* Write a bash script to install something as idempotently as possible.
* Discuss Chef and Puppet and while reflecting on the bash script.

Automation - Chef 201
=====================

Setup an Opscode account
------------------------

Setup your workstation as a client to your Opscode account
----------------------------------------------------------

Download the build-essential cookbook, and apply it to your workstation
-----------------------------------------------------------------------

Automation - Chef 301
=====================

Setup a chef-repo
-----------------

Write your own cookbook
-----------------------

Automation - Chef 302
=====================

Setup your own Chef Server
--------------------------

Write your own resources/providers
----------------------------------

Write sanity tests for your code
--------------------------------

Automation - Puppet 201
=======================

Install Puppet
--------------

Install a Forge module using the module tool
--------------------------------------------

Apply it to your local machine
------------------------------

Automation - Puppet 301
=======================

Install Puppet
--------------

Create your own module
----------------------

Apply it to your local machine
------------------------------

Package Management 101
======================

Setup a basic YUM or APT repository and put some packages in it
---------------------------------------------------------------

Setup a local mirror of CentOS (or what have you)
-------------------------------------------------

Setup a client to install from it
---------------------------------

Package Management 201
======================

Build a simple RPM or deb
-------------------------

FPM
---

Build automation fleets
=======================

koji
----

D
-

Version Control with Git 101
============================

Open a GitHub account
---------------------

Create a new repository called 'scripts'
----------------------------------------

Place a useful shell script in it
---------------------------------

Commit and push
---------------

Make a change, commit and push
------------------------------

Create a branch, make a change, commit, and push
------------------------------------------------

Create a pull request and merge the branch into the master branch
-----------------------------------------------------------------

* Read Chapters 1-3 of the `Pro Git <http://git-scm.com/book>`_ book online
* Work through Code School's `Try Git <http://try.github.com/>`_ online

DNS 101
=======

Install Bind
------------

Configure one zone
------------------

Show DNS resolution for an A and a CNAME record in the configured zone
----------------------------------------------------------------------

HTTP 101
========

Install Apache
--------------

Configure a virtual host
------------------------

Display a simple web page
--------------------------
