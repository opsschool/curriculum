Configuration Management 101
****************************

Idempotency
===========

Convergent and Congruent systems
================================

Direct and Indirect systems: ansible, capistrano
================================================

Chef
====

Chef (adam: I’m biased here, but I would do Chef in 101, puppet and cfengine in
201, but it’s because I want junior admins to get better at scripting, not just
because I’m a dick.)
(Magnus: this goes back to why Ruby will be so much more for new guys coming in
today like Perl was for a lot of us in the 90’s)

Configuration Management 201
****************************

Puppet
======

Cfengine 3
==========

SaltStack
=========

SaltStack or just **Salt**, is a configuration management and remote
execution tool written in Python. Salt uses ZeroMQ to manage communication 
between master and minions, and RSA keys to handle authentication. 
This chapter will explain the basics on how to get started with it.

Salt is a centralized system, which means there is a main server (also referred
here as *master*) which manages other machines connected to it or itself (also
referred here as *minions*). This topology can be further split using
`Salt Syndic <http://docs.saltstack.org/en/latest/ref/syndic.html>`_, 
please refer to Salt documentation for more details on this topic.

In examples below we will be using the master + 1 minion setup. The approximate 
time you will need to walk through all the content is about 10 minutes.

Prerequisites:

* access to 2 Linux/Solaris/FreeBSD/Windows machines in the same network
* basic understanding of command line instructions
* basic understanding of YAML file format

Installation
------------

Salt has a `dedicated page <https://salt.readthedocs.org/en/latest/topics/installation/index.html>`_ 
on how to get it installed and ready to use, please refer to it after deciding
what OS you will be using.

To set-up the environment you can use virtual machines or real boxes, in the 
examples we will be using hostnames **master** and **slave** to refer to each
one.

We will continue considering you already have the latest version installed
and available from command line on your machines.
You can check what version are you using on master with:

::

  ~# salt --version
  salt 0.10.3

and on slave with:

::

  ~# salt-minion --version
  salt-minion 0.10.3

