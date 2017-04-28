############
Sysadmin 101
############

*System Administration* is the blanket title for a myriad of job
responsibilities involving the management of information systems and/or
services. We choose to refer to it as a "blanket title" because the
responsibilities are broad and can differ from organization to organization.
In some organizations, it even differs from group to group!

System Administration is often referred to as "operations" or "ops" too - you
might have guessed that though, having decided to enroll in Ops School.

Like almost everything in this world, system administration does not exist in
a vacuum. There are two other groups of people that exist to make
an information system or service come together: developers and customers.
Customers are often referred to as end-users or people or humans - basically,
they're the folks you're implementing a system or service for.

A system administrator experiences life in all three of these roles: sometimes
you will be tasked with maintaining a system or service, and sometimes you
may be developing parts of a system. Many times, you'll be someone else's
customer.

This section will get into what system administration (ops) professionals do,
what developers do, where they compare and contrast, what role that ops plays
in the larger business context, a mindset and approach to systems administration,
demonstrate the importance of problem solving and ethics in the every day
life of a system administrator.


.. _whats-sysadmin:

*******************************
What is Systems Administration?
*******************************

Like we mentioned above, the role of the System Administrator is diverse.

To illustrate some of the latitude present in the role, consider how the
following descriptions compare and contrast:

.. _whats-sysadmin-matt:

* **Matt works for an elementary school.** He's referred to by other employees at
  the school as the system administrator. To the other employees, Matt ensures
  that the computers "just work." Realistically, Matt spends time making sure
  that the systems that everyone uses are patched - to avoid security issues.
  Other employees talk to him
  about software they may need installed -- on their computer, or maybe all of
  them! Matt also has the responsibility of making sure that the school's
  website is available for parents to get information on, and that the classroom
  management software is up and running and accepting requests. Matt also needs
  to be able to look into the future and figure out how much space the employees
  may need on the server, and be able to justify that to his boss. When report
  card time runs around, Matt has to make sure that the teachers are able to
  give their report cards out on time.

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
  sense for the town? Will collects input from the employees about what
  capabilities are important, input from legal council about what laws a town
  must follow regarding electronic communication, and evaluates a few
  services or products with his users' needs in mind along with the town's. If he
  ends up using a third-party service, he'll have to figure out how to roll it
  out to the users. If Will chooses to host himself, he'll have to coordinate
  getting the appropriate infrastructure set up to make sure that it can meet
  the capacity of the users. He'll also need to set it all up.

.. _whats-sysadmin-karen:

* **Karen works for a large business.** The business has many groups, but she
  works with one set of users - hardware engineers. She maintains a network that
  they use to run their hardware simulations before the hardware goes out to be
  built. She regularly writes scripts to reduce the number of manual steps
  something takes - like adding new users to the system. She also has to
  coordinate how the network will be patched regularly, what investments will be
  required for the network to continue operating at peak performance, and work
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

.. _whats-sysadmin-chris:

* **Chris works for a consultancy.** The company has many clients for which it
  provides software development and operations support. Every client has their
  own infrastructure, their own requirements - some of them very similar
  to the previous examples, some very different. Therefore, Chris needs to be
  familiar with a wide range of technologies, in addition to being able to
  learn new things quickly. There's a lot of variety in the day-to-day job.
  While it's not his primary responsibility to write code, sometimes the job
  calls for it - where context has already determined the implementation
  language. Chris works closely not only with the clients but also with his
  developer coworkers and the rest of the operations team to meet client goals.
  In addition, Chris also works with the rest of the operations team to support
  the consultancy's own infrastructure.

.. _whats-sysadmin-comparing-stories:

Comparing Stories
=================
Did you notice any themes throughout the stories? How do they contrast?

Highlighted themes include:

* **System Administrators rarely work on their own.** Some of the stories
  feature System Administrators who are the only "Information Technology" staff
  in their organization - but it's a requirement for them to work with everyone:
  the folks who "sign the paycheck at the end of the day," *and* the people
  who use their technology.
* **System Administration can be difficult if you don't automate!** Imagine if
  Matt had to install a software package on 200 machines by walking to each
  machine and installing it? How could he get any of his other requests done?
* **At times, the system administrator may need to write code.**
  Maybe that means a script that can take a list of users and add them to a
  system, maybe that means a script that can check to make sure a service is
  responding appropriately for a monitoring system.
* **System Administration has a wide variety of meanings to different people.**
  The other non-IT employees are looking for something that "just works" - the
  ops person is concerned with proper implementation of a service so it
  continues to be scalable and reliable for the future.


.. _whats-dev:

********************
What is Development?
********************

As mentioned in the section introduction, system administration doesn't exist
in a vacuum. Computers don't just compute for the sake of it; not yet,
anyway.

*Developers* write applications in a variety of programming languages to
make a computer do *something*. Developers created the backend that allows you
to order a book on a site like Amazon.com, post a status on Facebook, or
research a topic in a literary journal database.

These backends can be big, and have many considerations behind their design.

.. _whats-developer-rachel:

* **Rachel works for an ecommerce company.** The company has a large
  team of developers, all working in different applications, usually in Ruby or
  JavaScript, and the applications talk to each other via APIs to support the main
  site that customers interact with. Her obligations to the business are to create
  new features for the website and maintain or fix the old ones. Her obligations to
  her other developers are to keep the code clean and readable, with tests so others
  can confidently refactor her code. She works closely with the Operations team to
  make the app-level changes that help the ops team maintain a robust infrastructure,
  like making pages fully cachable or eliminating unnecessary HTML, CSS, JavaScript
  or images being sent to the browser.

.. _whats-developer-tyler:

* **Tyler is a systems developer at a small technology company.** He takes
  complex processes, like developing a search engine or collecting statistics, and
  creates several pieces of software to accomplish those tasks. He works primarily
  in C, Perl and Python and has to have a deep understanding of the operating system
  his code will run on. He works closely with his Operations engineers to make sure his
  code is performant and on capacity planning.

  .. _whats-developer-francesca:

* **Francesca is one of five students working on the website for her school's
  internationally famous hackathon.** She and her friends run a Tornado web server
  with a Postgres database. For the front end, they use GitHub pages and React.
  Their website is hosted on Amazon Web Services. The team uses New Relic and
  Pingdom to profile the code to see what time, CPU, and memory it consumes, as
  well as to alert them if any aspect of the site is acting abnormally.

.. _whats-developer-jean:

* **Jean works at a large technology firm.** The company has many teams of
  developers, all working on different aspects of their product. Jean is a data
  engineer who manages the collection and display of certain statistics that are
  useful for other employees of the company. If Jean has any issues with their
  computer or development environment, they ask the people running the
  information technology desk. Core parts of the technology at this company are
  maintained by site reliability engineers. Jean also has an embedded site
  reliability engineer on their team who makes sure that the features they
  create are scalable and maintainable.

.. todo:: "What is Development" Section needs more developer perspective.


.. _constrasting-devandops:

**************************************
Contrasting Development and Operations
**************************************

At the end of the day, both groups have an important shared goal: to ensure that
a system or service remains as *available* (think: accessible, usable, works how
people expect it to) as a customer expects it to be. You'll see references to
the idea of how "available" a system or service is later when Service Level
Agreements (SLAs) are discussed.

That being said, Development and Operations have different day-to-day thoughts.

Operations thoughts include:

* How are we going to install (or deploy) the servers that run this application?
* How will we monitor the system/service to make sure it's working as we expect?
* Can we deploy this system or service in a way that is easy to maintain?
* What are the pros/cons of implementing this application in this way?

Development thoughts include:

* How will I implement message passing between two parts of this application?
* What's the best algorithm to use for searching through this amount of data?
* Should I be thinking of a key-value store for this data vs. using a relational
  database?
* What language will I implement this in?

Again, this is by no means an exhaustive list - entire books could be written
about the subject. It's just to give a feel for the different considerations.


*************************************
History of Development and Operations
*************************************

Historically, Developers made a product or application and Operations
implemented it. Some companies still use this model today. Unfortunately, the
effects of this style of work can be dangerous:

* **Developers and Operations personnel may have different goals.** Something
  important to the operations folk may not be important to the Development
  folk.
* **Siloing these two organizations means that the most important goal of
  service availability is compromised.** The person who is using your program
  or service doesn't care who made a mistake, the service is down.
* **Speaking of mistakes:** when companies don't encourage their development and
  operations teams to work together, it's possible that groups get too invested
  in assigning blame, instead of working together to fix an issue

Fortunately, in recent years, many organizations have made more effort to ensure
that both teams are familiar with the concepts of the other; this is often 
referred to as *DevOps*, the combination of Development and Operations. Recall
:ref:`Darren's story <whats-sysadmin-darren>` - he's an operations person who
worked with the developers to ensure that the developers understand the
environment that their application will run on. The street went both ways,
though: the developers need to share how they plan to implement various product
features so that the operations team can figure out how to best support the
developers' needs.

If you're working in an environment without developers, that's OK. There are
other people who you share a common, larger goal with. Businesses may have
analysts that interpret needs and look to you for assistance. In
:ref:`Karen's story <whats-sysadmin-karen>`, she supports hardware engineers who
have requirements to deliver a particular sensor. Their ability to work hinges
on Karen's ability to deliver a service for simulation that is available for
them to work, which requires an understanding of their requirements and needs
as well.

********************************
What System Administration Isn't
********************************

System Administration, like many things in life, can suffer from a cultural
perception issue.  Primarily, some outside of the operations field that believe
that the role of IT is not to innovate; rather that IT exists to enforce rules
decided on by others.

We challenge that mindset. System administration is not about saying "no" to 
requests but about finding ways to intelligently fulfill the needs of a business
in a way that increases maintainability, usability, and security for a group of
people and enhances their ability to do useful work.

Another perception issue is that of the `BOFH
<http://www.theregister.co.uk/data_centre/bofh/>`_. While there's something of
a grain of truth to the meme, most system administrators who adopt such an
attitude quickly find themselves out of work. Beyond the social implications,
this sort of obstructionist, adversarial behavior is the least effective way of
addressing problems. The best path to success and satisfaction is collaboration,
not obstinacy. This requires cooperation from all involved, whether their role
is technical or not. An administrator's job is to create order from chaos, to
control the technology -- not the people. As with all social interactions, 
everyone has their own needs and goals, and it can be complex to navigate safely, 
but it is critical to the success of the organization.

Problem Solving
===============

Learning Styles - Ways to improve skill set
-------------------------------------------

Methodologies for finding solutions
-----------------------------------

Ethics
======
* LOPSA ethics statement/SAGE(now LISA) ethics statement?

Where to draw the line
----------------------

Keeping yourself safe from yourself
-----------------------------------

