Sysadmin 101
************

*System Administration* is the blanket title for a myriad of job
responsibilites involving the management of information systems and/or
services. We choose to refer to it as a "blanket title" because the
responsiblities are broad and can differ from organization to organization.

System Administration is often refered to as "operations" or "ops" too - you
might have guessed that though, being enrolled in ops school, and all. 

Like almost everything in this world, system administration does not exist in
in a vaccum. Indeed, there are two other groups of people that exist to make
an information system or service come together: developers and customers. 
Customers are often referred to end-users or people or humans - basically, 
they're the folks you're implementing a system or service for.

This section will get into what system administrators (ops) professionals do,
what developers do, where they compare and contrast, what role that ops plays
in the larger business context, a mindset and approach

.. _whats-sysadmin:

What is Systems Administration?
===============================

Like we mentioned above, the role of the System Administrator is diverse.

To illustrate some of the lattitude present in the role, consider how the 
following descriptions compare and contrast:

.. _whats-sysadmin-matt:

* **Matt works for an elementary school.** He's referred to by other employees at
  the school as the system administrator. To the other employees, John ensures
  that the computers "just work." Realistically, Matt spends time making sure
  that the systems that everyone uses are patched. Other employees talk to him
  about software they may need installed -- on their computer, or maybe all of
  them! Matt also has the responsibility of making sure that the school's
  website is available for parents to get information on, and that the classroom
  management software is up and running and accepting requests. Matt also needs
  to be able to look into the future and figure out how much space the employees
  may need on the server, and be able to justify that to his boss.

.. _whats-sysadmin-sarah:

* **Sarah works for a large website that many people use on a day to day basis.**
  Downtime is a big deal for this website: after all, when the site is down,
  people aren't spending money. When these problems occur, Sarah gets paged to
  determine the root cause of the issue and respond. Fortunately for Sarah she's
  not always being paged. During the normal day-to-day operations of the site,
  Sarah writes various scripts and programs that help her other administrators
  manage the site on a day to day basis, and strives to reduce the amount of
  "manual" tasks that the team she's on has to deal with.

.. _whats-sysadmin-will:

* **Will works for a small town.** The town has 35 employees who use computers
  every day, and they need access to email and other collaboration tools. With
  so many services and options for e-mail, which one makes the most amount of
  sense for the town? Daniel collects input from the employees about what
  capabilities are important, input from legal council about what laws towns
  need to follow regarding electronic communication, and evaluates a few
  services or products with his user's needs in mind along with the towns. If he
  ends up using a third- party service, he'll have to figure out how to roll it
  out to the users. If he chooses to host himself, he'll have to coordinate
  getting the appropriate infrastructure set up to make sure that it can meet
  the capacity of the users. He'll also need to set it all up.

.. _whats-sysadmin-karen: 

* **Karen works for a large business.** The business has many groups, but she
  works with one set of users - hardware engineers. She maintains a network that
  they use to run their hardware simulations before the hardware goes out to be
  built. She regularly writes scripts to reduce the number of manual steps
  something takes - like adding new users to the system. She also has to
  coordinate how the network will be patched regularly, what investments will be
  required for the network to continue operating at peek performance, and work
  with the other administrators to create something that works for the engineers
  at the business.

.. _whats-sysadmin-darren:
  
* **Darren works for a large financial company.** The company has many teams that
  do different things, but there's collaboration between the different system
  administrators - the operations team (staffed with administrators) implements
  things that the research team (also administrators) recommends based on the
  needs of the business. They work with developers too - since the financial
  company has a lot of custom software that is written by the developers. The
  developers have different goals when they write their software - so they meet
  with Darren and his team to make sure they understand the context surrounding
  where and how their application will run, and to make sure that Darren's team
  knows what they need as well. Darren's team may have ideas about what
  statistics are important to monitor.
  
.. _whats-sysadmin-comparing-stories:

Comparing Stories 
-----------------
Did you notice any themes throughout the stories? How do they contrast?

Things that we tried to highlight include:

* **System Administrators rarely work on their own.** Some of the stories
  feature System Administrators who are the only "Information Technology" staff
  in their organization - but it's a requirement for them to work with everyone
  - the folks who "sign the paycheck at the end of the day," *and* the people
  who use their technology.
* **System Administration can be difficult if you don't automate!** Imagine if
  Matt had to install a software package on 200 machines by walking to each
  machine and installing it? How could he get any of his other requests done?
* Similarly - yes, **the system administrator is writing code sometimes.**
  Maybe that means a script that can take a list of users and add them to a
  system, maybe that means a script that can check to make sure a service is
  responding appropriately for a monitoring system.
* System Administration means a lot of things to a lot of people. The other
  non-IT employees are looking for something that "just works" - the ops
  person is concerned with proper implementation of a service so it continues 
  to be scalable and reliable for the future.

.. _whats-dev:

What is Development?
====================
As mentioned in the section introduction, system administration doesn't exist
in a vaccum. Computers don't just compute for the sake of it - not yet,
anyways. *Developers* write applications in a variety of programming languages to 
make a computer do *something*.

Historically, developers made a product and operations implemented it. Many
companies still use this model today. Unfortunately, some of the effects of
this style of work are dangerous:

* **Developers and Operations personel may have different goals.** Something 
  important to the operations folk may not be important to the Development
  folk. 
* **Siloing these two organizations means that the most important goal of 
  service availability is compromised.** The person who is using your program
  or service doesn't care who made a mistake, the service is down.
* Speaking of mistakes: when companies don't encourage their development and 
  operations teams to work together, it's possible that 
  
Fortunately, recent companies there has been more effort on ensuring that both
teams are familiar with the concepts of the other - this is often referred to
as "devops" - the combination of Development and Operations. Remember
":ref:`Darren's story <whats-sysadmin-darren>`" - he was an operations
person who worked with the developers to share the context.

If you're working in an environment without developers, that's ok. There are 
other people who you share a common, larger goal with. Businesses may have 
analysts that interpret needs and look to you for assistance. In the example 
featuring Karen above, she supports hardware engineers who have requirements
to deliver a particular sensor. Their ability to work hinges on Karen's ability
to deliver a service for simulation that is available for them to work.

.. _whats-not-sysadmin:

What is System Administration Not?
==================================
* Nick Burns: Your Companys Computer Guy!!! 
* professional roadblock, etc.

The role of the SysAdmin in the organization
============================================
* “Who is that guy? Why is he always muttering about ‘latency’?”
* understanding the greater role of delivering a service for a business


Mindset and approach
====================

Generalists vs Specialists
--------------------------

Problem Solving
===============

Learning Styles - Ways to improve skillset
------------------------------------------

Methodologies for finding solutions
-----------------------------------

Ethics
======
LOPSA ethics statement/SAGE ethics statement?

Where to draw the line
----------------------

Keeping yourself safe from yourself
-----------------------------------

