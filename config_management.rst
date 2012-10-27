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

Configuration
-------------

Minimum configuration has to be done in order to get the slave server to
communicate with master. You will need to tell it what IP address and port
master uses.
The configuration file usually can be found at ``/etc/salt/minion``.

You will need to edit the file where it says ``master: salt`` replacing
``salt`` with master IP address.

Once done, you will need to restart the service: **salt-minion**. On most
Linux OSes you can use ``service salt-minion restart`` for that.

Authentication keys for master/slave are generated during installation so
you don't need to manage those manually, except in case when you want to
`preseed minions <https://salt.readthedocs.org/en/latest/topics/tutorials/preseed_key.html>`_.

To add the slave to minions list, you will have to use the command ``salt-key``
on master. Run ``salt-key -L`` to list available minions:

::

  root@master:~# salt-key -L
  Unaccepted Keys:
  slave
  Accepted Keys:
  Rejected:

To accept a minion, run ``salt-key -a``:

::

  root@master:~# salt-key -a slave
  Key for slave accepted.

  root@master:~# salt-key -L
  Unaccepted Keys:
  Accepted Keys:
  slave
  Rejected:

Once the minion was added, you can start managing it by using command ``salt``.
For example to check the communication with slave, you can ping it:

::

  root@master:~# salt 'slave*' test.ping
  slave: True

Remote execution
----------------

In order to understand how Salt manages minions and configuration files syntax,
lets have a look at ``salt`` command line tool. Lets take our previous command
and inspect it:

::
  
  root@master:~# salt 'slave*' test.ping
                             ^ ^
                       ______| |__________________
                       target  function to execute

**target** is the minion(s) name. It can represent the exact name or just
a part of it followed by a wildcard.

  In order to run target matching by OS, architecture or other identifiers
  take a look at `Salt Grains <https://salt.readthedocs.org/en/latest/topics/targeting/grains.html>`_.

Functions that can be executed represent `Salt Modules <https://salt.readthedocs.org/en/latest/ref/modules/index.html>`_.
These modules are Python or Cython code written to abstract access to CLI or
other minion resources. For the full list of modules please take a look
`this page <https://salt.readthedocs.org/en/latest/ref/modules/all/index.html>`_.

One of the modules Salt can use is called **cmd**, it has the method **run**
which accepts as arguments a string which will be executed as a command on
minions and the resulted data will be outputted prefixed with minion name.

For example, to run command ``uptime`` on our slave we will fire:

::
  
  root@master:~# salt 'slave*' cmd.run 'uptime'
  slave: 22:45:52 up 96 days, 10:42,  0 users,  load average: 0.00, 0.00, 0.00

