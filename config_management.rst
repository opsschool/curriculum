Configuration Management 101
****************************

A Brief History of Configuration Management
===========================================

Configuration management as a distinct sub-discipline of operations engineering 
has roots in the mid-1990s. Prior to then, even sites with a large number of 
users like universities and large ISPs had relatively few Unix systems. Each of 
those systems was generally what today's operations community calls a 
"snowflake system" (after the phrase "a precious and unique snowflake"). They 
were carefully hand-built to purpose, rarely replaced, and provided a unique 
set of services to their users.

The rise of free Unix-like Operating Systems and commodity x86 hardware, coupled with the 
increasing demands to scale booming Internet services meant the old paradigms 
of capturing configuration were no longer adequate. Lore kept in text files, 
post-bootstrap shell scripts, and tales told around the proverbial campfire 
just didn't scale.  Administrators needed automation tools which could stamp 
out new machines quickly, plus manage configuration drift as users made changes 
(deliberately or accidentally) that affected the functioning of a running 
system. 

The first such tool to gain prominence was CFEngine, an open-source project 
written in C by Mark Burgess, a CS professor at Oslo University. CFEngine 
popularized the idea of *idempotence* in systems administration tasks, 
encouraging users to describe their system administration tasks in ways that 
would be convergent over time rather than strictly imperative shell or perl 
scripting.

.. todo:: a specific example of convergent over time might help

In the early 2000s, the systems administration community began to focus more
intensely on configuration management as distributed systems became both more 
complex and more common. A series of LISA papers and an explosion in the number 
and sophistication of open-source tools emerged. Some highlights and background 
reading:

* Steve Traugott's isconf3 system and paper `"Bootstrapping an 
  Infrastructure" <http://www.infrastructures.org/papers/bootstrap/bootstrap.html>`_ provided a 
  concrete model for repeatable, scalable provisioning and config management.
* CERN released and wrote about `Quattor <http://quattor.org/index.html>`_ 
  which they used to build and administer high-performance compute clusters at 
  larger scale than most sites at the time had dealt with.
* Alva Couch from Tufts University and Paul Anderson from University of 
  Edinburgh, laid out theoretical underpinnings for configuration management 
  in a `joint session at LISA'04 <http://static.usenix.org/event/lisa04/tech/talks/couch.pdf>`_
* Narayan Desai's `bcfg2 system <http://bcfg2.org>`_ provided a hackable Python 
  CM project with early support for advanced features like templating and 
  encrypted data
* Recapitulating Luke Kanies' `departure from cfengine 
  <http://rootprompt.org/article.php3?article=10981>`_ to start Puppet, Adam 
  Jacob created Chef in 2008 to address `fundamental differences 
  <http://www.akitaonrails.com/2009/11/18/chatting-with-adam-jacob>`_ with 
  Puppet (primarily execution of ordering and writing user code in Ruby vs a 
  DSL).

By 2008, provisioning and configuration management of individual systems were 
well-understood (if not completely `"solved" 
<http://blog.lusis.org/blog/2011/08/22/the-configuration-management-divide/>`_) 
problems, and the community's attention had shifted to the next level of 
complexity: cross-node interactions and orchestration, application deployment, 
and managing ephemeral cloud computing instances rather than (or alongside) 
long-lived physical hardware.

A new crop of CM tools and approaches "born in the cloud" began to emerge in 
the 2010s to address this shift. SaltStack, Ansible, and Chef-v11 built on 
advances in language (Erlang and Clojure vs Ruby and Python), methodology 
(continuous deployment and orchestration vs static policy enforcement), and the 
component stack (ZeroMQ and MongoDB vs MySQL). 

Whatever specific configuration management tooling operations engineers 
encounter as an operations engineer, ultimately the technology exists to enable 
business goals -- short time-to-restore in the face of component failure, 
auditable assurance of control, low ratio of operators per managed system, etc.  
-- in a world whose IT systems are moving, in the words of CERN's Tim Bell, 
"from pets to cattle".

Idempotence
===========

Convergent and Congruent systems
================================

Direct and Indirect systems: ansible, capistrano
================================================

(mpdehaan: Are we talking about deployment here?  Then let's start a deployment section.  What does direct/indirect mean? How about not addressing tools in 101 and talking about concepts, so as to make a better tools section? Ansible operates in both push and pull topologies, so I'm guessing that is not what is meant about direct/indirect?)

Chef
====

Chef (adam: I'm biased here, but I would do Chef in 101, puppet and cfengine in
201, but it's because I want junior admins to get better at scripting, not just
because I'm a dick.)
(Magnus: this goes back to why Ruby will be so much more for new guys coming in
today like Perl was for a lot of us in the 90's)

Configuration Management 201
****************************

Ansible
=======

`Ansible <http://ansible.com>`_ is a configuration management, deployment, and remote execution tool that uses SSH to address remote machines (though it offers other connection types, including 0mq).  It requires no server software nor any remote programs, and works by shipping small modules to remote machines that provide idempotent resource management.  While implemented in Python, Ansible uses a basic YAML data language to describe how to orchestrate operations on remote systems.  

Ansible can be extended by writing modules in any language you want, though there is some accelerated module writing ability that makes it easier to do them in Python.

To prevent documentation drift, see `Ansible documentation site <http://docs.ansible.com>`_.

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
time you will need to work through all the content is about 10 minutes.

Prerequisites:

* access to 2 Linux/Solaris/FreeBSD/Windows machines in the same network
* basic understanding of command line instructions
* basic understanding of YAML file format

Installation
------------

Salt has a `dedicated page <https://salt.readthedocs.org/en/latest/topics/installation/index.html>`_
on how to get it installed and ready to use, please refer to it after deciding
what OS you will be using. These examples are shown on an Ubuntu installation
with Salt installed from a `project personal package archive
<https://salt.readthedocs.org/en/latest/topics/installation/ubuntu.html>`_.

To set-up the environment you can use virtual machines or real boxes, in the
examples we will be using hostnames **master** and **slave** to refer to each
one.

At this point, you should install the latest version on both machines with the
directions provided above, and have a command line session open on both your
**master** and **slave** machines.
You can check what version are you using on master with:

.. code-block:: console

  root@master:~# salt --version
  salt 0.10.3

and on slave with:

.. code-block:: console

  root@slave:~# salt-minion --version
  salt-minion 0.10.3

Configuration
-------------

A minimum configuration is required to get the slave server to
communicate with master. You will need to tell it what IP address and port
master uses.
The configuration file can typically be found at :file:`/etc/salt/minion`.

You will need to edit the configuration file directive ``master: salt`` replacing
``salt`` with master IP address or its hostname/FQDN.

Once done, you will need to restart the service: **salt-minion**. On most
Linux distributions you can execute ``service salt-minion restart`` to restart
the service.

Authentication keys for master/slave are generated during installation so
you don't need to manage those manually, except in case when you want to
`preseed minions <https://salt.readthedocs.org/en/latest/topics/tutorials/preseed_key.html>`_.

To add the slave to minions list, you will have to use the command ``salt-key``
on master. Execute ``salt-key -L`` to list available minions:

.. code-block:: console

  root@master:~# salt-key -L
  Unaccepted Keys:
  slave
  Accepted Keys:
  Rejected:

To accept a minion, execute ``salt-key -a <minion-name>``:

.. code-block:: console

  root@master:~# salt-key -a slave
  Key for slave accepted.

  root@master:~# salt-key -L
  Unaccepted Keys:
  Accepted Keys:
  slave
  Rejected:

Once the minion is added, you can start managing it by using command ``salt``.
For example, to check the communication with slave, you can ping the slave from the master:

.. code-block:: console

  root@master:~# salt 'slave*' test.ping
  slave: True

Remote execution
----------------

In order to understand how Salt does its configuration management on minions,
we'll take look at the ``salt`` command line tool. Let's take our
previous command and inspect the parts of the command:

.. code-block:: console

  root@master:~# salt 'slave*' test.ping
                             ^ ^
                       ______| |__________________
                       target  function to execute

**target** is the minion(s) name. It can represent the exact name or only
a part of it followed by a wildcard. For more details on how to match minions
please take a look at `Salt Globbing <http://docs.saltstack.org/en/latest/topics/targeting/globbing.html>`_.

  In order to run target matching by OS, architecture or other identifiers
  take a look at `Salt Grains <https://salt.readthedocs.org/en/latest/topics/targeting/grains.html>`_.

Functions that can be executed are called Salt Modules.
These modules are Python or Cython code written to abstract access to CLI or
other minion resources. For the full list of modules please take a look
`this page <https://salt.readthedocs.org/en/latest/ref/modules/all/index.html>`_.

One of the modules provided by Salt, is the **cmd** module. It has the **run**
method, which accepts a string as an argument. The string is the exact
command line which will be executed on the minions and contains both
the command name and command's arguments. The result of the command execution
will be listed on master with the minion name as prefix.

For example, to run command ``uname -a`` on our slave we will execute:

.. code-block:: console

  root@master:~# salt slave cmd.run 'uname -a'
  slave: Linux slave 2.6.24-27-openvz #1 SMP Fri Mar 12 04:18:54 UTC 2010 i686 GNU/Linux

Writing configuration files
---------------------------

One of the Salt modules is called ``state``. Its purpose is to manage minions
state.

  Salt configuration management is fully managed by states, which purpose is
  to describe a machine behaviour: from what services are running to what
  software is installed and how it is configured. Salt configuration management
  files (``.sls`` extension) contain collections of such states written in YAML
  format.

Salt states make use of modules and represent different module calls organised
to achieve a specific purpose/result.

Below you can find an example of such a **SLS** file, whose purpose is to get
Apache Web server installed and running:

.. code-block:: yaml

  apache2:
    pkg:
      - installed
    service.running:
      - require:
        - pkg: apache2

To understand the snippet above, you will need to refer to documentation on
states: pkg and service. Basically our state calls methods ``pkg.installed``
and ``service.running`` with argument ``apache2``. ``require`` directive is
available for most of the states and describe dependencies if any.

Back to ``state`` module, it has a couple of methods to manage these states. In
a nutshell the state file form above can be executed using ``state.sls``
function. Before we do that, let's take a look where state files reside on
the master server.

Salt master server configuration file has a directive named ``file_roots``,
it accepts an YAML hash/dictionary as a value, where keys will represent the
environment (the default value is ``base``) and values represent a set/array
of paths on the file system (the default value is :file:`/srv/salt`).

Now, lets save our state file and try to deploy it.

Ideally you would split state files in directories (so that if there
are also other files, say certificates or assets, we keep those organised). The
directory layout we will use in our example will look like this: ::

  /srv/salt/
  |-- apache
  |   `-- init.sls
  `-- top.sls

When creating new states, there is a file naming convention.
Look at ``init.sls``, it is the default filename to be searched when loading
a state. This is similar to Python or default web page name ``index.html``.

So when you create a new directory for a state with an ``init.sls`` file in it
it translates as the Salt state name and you can refer to it as that. For example,
you do not write ``pkg: new_state.init``, write just ``pkg: new_state``.

Now to deploy it, we will use the function ``state.sls`` and indicate the state
name:

.. code-block:: console

  root@master:~# salt slave state.sls apache
  slave:
  ----------
      State: - pkg
      Name:      apache2
      Function:  installed
          Result:    True
          Comment:   Package apache2 installed
          Changes:   apache2.2-bin: {'new': '2.2.14-5ubuntu8.10', 'old': ''}
                     libapr1: {'new': '1.3.8-1ubuntu0.3', 'old': ''}
                     perl-modules: {'new': '5.10.1-8ubuntu2.1', 'old': ''}
                     ssl-cert: {'new': '1.0.23ubuntu2', 'old': ''}
                     apache2-utils: {'new': '2.2.14-5ubuntu8.10', 'old': ''}
                     libaprutil1-ldap: {'new': '1.3.9+dfsg-3ubuntu0.10.04.1', 'old': ''}
                     apache2-mpm-worker: {'new': '2.2.14-5ubuntu8.10', 'old': ''}
                     make: {'new': '3.81-7ubuntu1', 'old': ''}
                     libaprutil1: {'new': '1.3.9+dfsg-3ubuntu0.10.04.1', 'old': ''}
                     apache2: {'new': '2.2.14-5ubuntu8.10', 'old': ''}
                     libcap2: {'new': '1:2.17-2ubuntu1', 'old': ''}
                     libaprutil1-dbd-sqlite3: {'new': '1.3.9+dfsg-3ubuntu0.10.04.1', 'old': ''}
                     libgdbm3: {'new': '1.8.3-9', 'old': ''}
                     perl: {'new': '5.10.1-8ubuntu2.1', 'old': ''}
                     apache2.2-common: {'new': '2.2.14-5ubuntu8.10', 'old': ''}
                     libexpat1: {'new': '2.0.1-7ubuntu1.1', 'old': ''}

  ----------
      State: - service
      Name:      apache2
      Function:  running
          Result:    True
          Comment:   The service apache2 is already running
          Changes:

You can see from the above that Salt deployed our state to **slave** and reported changes.

In our state file we indicated that our service requires that the package must
be installed. Following the same approach, we can add other requirements like
files, other packages or services.

Let's add a new virtual host to our server now using the ``file`` state. We
can do this by creating a separate state file or re-using the existing one.
Since creating a new file will keep code better organised, we will take that approach.

We will create a new ``sls`` file with a relevant name, say ``www_opsschool_org.sls``
with the content below:

.. code-block:: yaml

  include:
    - apache

  extend:
    apache2:
      service:
        - require:
          - file: www_opsschool_org
        - watch:
          - file: www_opsschool_org

  www_opsschool_org:
    file.managed:
    - name: /etc/apache2/sites-enabled/www.opsschool.org
    - source: salt://vhosts/conf/www.opsschool.org

Above, we include already described state of the Apache service and extend it
to include our configuration file. Notice we use a new directive ``watch``
to describe our state as being dependent on what changes the configuration
file triggers. This way, if a newer version of the same file is deployed, it
should restart the Apache service.

Below is the directory listing of the changes we did: ::

  /srv/salt/
  |-- apache
  |   `-- init.sls
  |-- top.sls
  `-- vhosts
      |-- conf
      |   `-- www.opsschool.org
      `-- www_opsschool_org.sls

Using the newly created state file, we can try and deploy our brand new
virtual host:

.. code-block:: console

  root@master:~# salt slave state.sls vhosts.www_opsschool_org
  slave:
  ----------
      State: - file
      Name:      /etc/apache2/sites-enabled/www.opsschool.org
      Function:  managed
          Result:    True
          Comment:   File /etc/apache2/sites-enabled/www.opsschool.org updated
          Changes:   diff: New file

  ----------
      State: - pkg
      Name:      apache2
      Function:  installed
          Result:    True
          Comment:   Package apache2 is already installed
          Changes:
  ----------
      State: - service
      Name:      apache2
      Function:  running
          Result:    True
          Comment:   Started Service apache2
          Changes:   apache2: True

Salt reports another successful deploy and lists the changes as in the example
above.

All this time, you were probably wondering why there is a file ``top.sls`` and
it was never used?! Salt master will search for this file as indicated in the
configuration of your install. This file is used to describe the state of all
the servers that are being managed and is deployed across all the machines
using the function ``state.highstate``.

Let's add our state files to it to describe the high state of the ``slave``.

.. code-block:: yaml

  base:
    'slave*':
      - vhosts.www_opsschool_org

Where ``base`` is the default environment containing minion matchers followed
by a list of states to be deployed on the matched host.

Now you can execute:

.. code-block:: console

  root@master:~# salt slave state.highstate

Salt should output the same results, as nothing changed since the last run. In order to
add more services to your slave, feel free to create new states or extend the
existing one. A good collection of states that can be used as examples can be
found on Github:

* https://github.com/saltstack/salt-states -- Community contributed states
* https://github.com/AppThemes/salt-config-example -- WordPress stack
  with deployments using Git

.. seealso:: For the full documentation on available states, please see `Salt States documentation <http://salt.readthedocs.org/en/latest/ref/states/all/index.html>`_.
