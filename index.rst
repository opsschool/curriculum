####################################
Learn reliability-minded engineering
####################################

Ops School is a free, openly licensed curriculum for learning reliability-minded engineering.
It is for SREs, software engineers, operations engineers, and anyone who builds or runs systems that people depend on.
These include the computer systems of businesses big and small.
They also include the systems that allow websites, networks, payment systems and other Internet services to function.
The curriculum covers a wide variety of topics, including systems administration, security, networking and beyond.
Ops School will guide you through all of these skill sets from beginner to expert.

.. rst-class:: actions

* :doc:`Start with Critical Thinking 101 <critical_thinking_101>`
* :ref:`Browse the curriculum <curriculum>`

.. container:: terminal

   .. code-block:: console

      user@opsschool ~$ curl -s -o /dev/null -w '%{time_total}\n' http://web-1/health
      0.042
      user@opsschool ~$ curl -s -o /dev/null -w '%{time_total}\n' http://web-3/health
      1.874


.. rst-class:: start

**************
Where to start
**************

Since the early 1990s, operations engineers have been in high demand, and the more recent SRE role is just as sought after.
The work suits people who enjoy diving into the inner workings of computer systems.

.. rst-class:: cards

* :doc:`Critical Thinking 101 <critical_thinking_101>`

  Check what you believe, and why, before you act on it.
  These habits come before troubleshooting, design, and everything else in the curriculum.

* :doc:`Careers <careers>`

  If you are reading about this career for the first time and want to know if it is for you, start here.

* :ref:`How to start <how-to-become-an-operations-engineer>`

  If you already know about the profession and want to know how to start, read what employers look for.


.. _curriculum:

.. rst-class:: curriculum

**************
The curriculum
**************

Many topics start at 101 and continue at 201, and some at 301.
Work through a track, or go straight to the topic you need.

.. rst-class:: legend

* 101
* 201
* 301

.. container:: tracks

   .. toctree::
      :maxdepth: 1
      :caption: Foundations

      introduction
      critical_thinking_101
      meta/guidelines
      careers
      sysadmin_101

   .. toctree::
      :maxdepth: 1
      :caption: Operating systems

      unix_101
      unix_201
      windows_101
      text_editing_101
      text_editing_201
      sysadmin_tools
      hardware_101

   .. toctree::
      :maxdepth: 1
      :caption: Networking

      networking_101
      networking_201
      common_services
      loadbalancing_101

   .. toctree::
      :maxdepth: 1
      :caption: Security and identity

      security_101
      security_201
      identity_management
      active_directory_101
      active_directory_201

   .. toctree::
      :maxdepth: 1
      :caption: Storage and data

      remote_filesystems_101
      remote_filesystems_201
      databases_101
      databases_201
      logs_101
      logs_201

   .. toctree::
      :maxdepth: 1
      :caption: Running services

      troubleshooting_101
      monitoring_101
      monitoring_201
      deployment_101
      deployment_201
      config_management
      capacity_planning
      statistics
      bcp

   .. toctree::
      :maxdepth: 1
      :caption: Design and build

      programming_101
      programming_201
      architecture_101
      architecture_201
      application_components_201
      virtualization_101
      virtualization_201
      datacenters/datacenters_101
      datacenters/datacenters_201
      datacenters/datacenters_301

   .. toctree::
      :maxdepth: 1
      :caption: People and practice

      soft_skills_101
      soft_skills_201
      labs
      learning
      reading_list
      seealso

   .. toctree::
      :maxdepth: 1
      :caption: Contributing and reference

      meta/contributions
      meta/conventions
      meta/style_guide
      meta/todo
      glossary


.. rst-class:: contribute

*********************************
Written by people who do the work
*********************************

Ops School is open source and written by volunteers.
Many chapters are still incomplete.
If you know a topic well, read the guidelines and the style guide, then open a pull request.

.. rst-class:: actions

* `Contribute on GitHub <https://github.com/opsschool/curriculum>`_
* :doc:`Read the style guide <meta/style_guide>`
