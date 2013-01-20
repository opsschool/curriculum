Soft Skills
***********

The term soft skills seems to imply that these skills are somehow less
important than technical skills. In reality, soft skills often separate
the technical-level operations person from the senior engineer. Soft
skills encompass general business acuity, communications skills, time
management, project management, and people management. General
business skills include positioning, understanding of organizational
finances, risk management, managing customer preference, and thinking
strategically.

Communication
=============

Communication 101 would tell you that the first step to effective
communication is audience analysis. A basic audience analysis consists
of answering some simple questions:

* Is your audience technical or non-technical?
* How much do they know about the topic you are trying to communicate?
* How much do they care about the topic?
* What communication mechanism will work best to deliver your message: face-to-face meeting or class, email, social media, internal poster or bulletin board, internal corporate web site?

Before sending one email or updating that web site, think about the
answers to these questions. People are innundated with communication
from email, voicemail, twitter, facebook/Google+, internal web/wikis,
text, IM. How can you deliver your message in the best way possible to
get the appropriate level of attention for your topic while not overly
burden the recipient(s)?

Communicating to internal customers
-----------------------------------

For those who are still communicating with internal customers via email,
consider how many emails you send and how long they are. 

Here's a hint with email: shorter is better. If you have to include
a long technical writeup, consider posting that on an internal
wiki/site and including a link for those interested.

Put the most important info at the top of the email. Many people will
skim the first few lines to see if an email applies to them. Consider
addressing it directly "All users of Macintosh systems" at the top of
the email. 

If the email requires an action, put that with a deadline in the very
first line: 
"The file cluster will be taken off-line for maintenance at 7:00pm
Friday, January 10th. Please save before that time so you do not lose
any of your work."

For major outages like the above, it is also good to give more than a
week notice for those people who are out of the office for a week or
more. You should also consider a followup message closer to the outage
to remind people that it is coming up. Forward the original email and
repeat the first line as a reminder.

Remember that we all get too much email. If you send an email every day,
or even every week, people will start to ignore/filter you. Consider how
critical it is to communicate with all of your customers at once. If you
limit it to topics that are important to them (outages, new service
announcements), they will be more likely to read and respond to your
attempts.

Communicating to external customers
-----------------------------------

Communicating with external customers can offer additional challenges.
If the external customers are customers of your organization, there 
is the possibility that your dealings with them could result in a
complaint to your upper management.

You can't please everyone, all the time. Every operations person has
dealt with complaints and customer unhappiness. 

You can reduce this by considering how you communicate with these
external customers. When communicating about a service outage, consider
timing of the outage, duration, and impact of the outage on these
external customers. Are most of your customers in the same time zone? If
so, then your maintenance window could be outside of traditional working
hours. If your external customers include international people in
varying timezones, you will have to choose an outage window that impacts
your core customers the least. 

Communicate maintenance windows decisions with your line management
for communication up the chain. It is best if your management knows
that their customers are about to be impacted by operations. You
may also need to include a justification for the maintenance: why
is it necessary, why did you choose this time window, how did you
come up with the duration, what happens if the outage goes beyond
the advertised window, how will you communicate to the external
customers? All of these pieces of information may not be necessary
if you have already established yourself as someone who supports
external customers on a regular basis.

Time Management
===============

Tom Limoncelli's Time Management Wiki:
* http://code.google.com/p/tomontime/wiki/Main

Project Management (Does time management go here too?)
======================================================

The project management triangle (good, cheap, fast: pick two)

* http://en.wikipedia.org/wiki/Project_management_triangle

Waterfall
---------

Waterfall is a sequential form of project management that was adapted
from other industries for the software development world. Waterfall
breaks down in practice because it requires a promise of delivery that
may be several years out. Instead of delivering a working product in a
short timeframe, waterfall assumes that each phase of the project will be
complete before moving to the next phase with the entirety of the
product delivered at the end. 

Detractors of the waterfall method point to its rigidity and
inflexibility. One of the issues in operations and development
work is that stakeholders may not have a solid grasp of requirements
until they see a working prototype, or iterations of working
prototypes during the implementation of the product. It is common
for stakeholders in a project not to know what technology can deliver
until they see it. Many operations teams are moving to Agile methods
for several reasons and one of them is because agile development
allows stakeholders to see working bits of the product before the
end and to modify requirements before it's too late.

Agile
-----

Agile is a methodology. Agile started in 2001 when a group of software
developers created the Agile Manifesto. The agile manifesto outlines the
12 principles of agile: http://agilemanifesto.org/. There are some
common implementations of Agile: Scrum and Kanban and even a hybrid
Scrumban that was created more for operational work.

Some benefits of agile include the following:

* reduced process overhead
* improved team and stakeholder communication and collaboration
* errors and bugs are fixed in development instead of waiting till the product is "complete" to address them.
* stakeholders see the product as it is shaped and have the ability to adjust requirements during development
* project teams are empowered
* can easily be combined with DevOps methodology to improve effectiveness of development-into-operations
* if done well, can increase work output of teams (increased velocity)
* everyone on the project can easily see where the project stands (e.g. scrumboard or kanban wall)

Kanban
^^^^^^

Scrum
^^^^^

Scrumban
^^^^^^^^

Agile Toolkit
^^^^^^^^^^^^^

jira
http://www.atlassian.com/software/jira/overview


The Tao of DevOps
=================

What is DevOps
--------------

DevOps seeks to include the IT operations team as an important
stakeholder in the development process. Instead of developers solely
coding to meet the stakeholder's requirements on time and on budget,
they are also held responsible for how easily it deploys, how few
bugs turn up in production, and how well it runs. Basically, how
easily can operations support the product once it rolls into
production. Instead of bringing operations into the conversation
after the product is complete, the DevOps methodology includes
operations in the development stream.

Development's view: 

* Roll a product out to meet customer specifications within a certain timeframe
* Continuous delivery means recurring change as bugs are fixed and features added
* fast changing environments are needed to support dev
* agility is key

Operation's view:

* supporting the product for customers
* keeping a handle on IT security
* planning for deployment to production state 
* changes are slow/incremental
* consistent environments are needed to support ops
* stability is key

By combining these two, often competing mindsets, both sides can be
satisfied and the result is a product that potentially has fewer bugs, higher
availability, increased security, and a process for improved development
over the life of the product that works for both the developers and the
operations people.

What isn't DevOps
-----------------

Why devops is important
-----------------------

The importance of Business Acumen in Operations
===============================================

What is business acumen? Business acumen a leadership competency simply
defined as a general understanding of business principles that leads
to an organization's success. We aren't trying to turn every operations
person into a senior executive, but development of business
acumen as applied to operations can sure help to bridge the gap
between your organization's senior leadership and the operations
team. Business acumen as applied to operations works on multiple
levels. In many organizations, operations is a service unit within
the larger organization but it also serves the needs of the
organization as a whole. The savvy operations person will look at
operations within that context, applying the following skills to
appropriately position operations and act with the best interests of the
greater organization in mind.

Breaking down business acumen further for operations yields the
following:
* Understanding the role of operations within the context of your organization. This leads to correct positioning of operations within the organization.
* Thinking broadly about decisions and acting decisively (how IT decisions can impact the internal and external workings of the organization.
* See change as a constant and support or promote change as needed

Understanding the role of Operations
------------------------------------
Under any of the Operations professions, the most fundamental role
of the Operations person is to deliver services to a set of customers.
To build upon this further, the Operations person maintains existing IT
infrastructures, translates customer requirements into tangible and
actionable solutions, assists in the protection of customer information
and services, and advises stakeholders on application of technology
under existing limitations of time, money, or capabilities.

By thinking of operations as a business unit instead of a forgotten
office within the organization, the operations engineer is already
thinking at the correct level to assess how to support the needs
of the organization.

Understand how your organization competes within its industry.
Commercial entities, non-profits, educational institutions, government
agencies all measure success in some way. For commerce, it will be sales
and profit. For educational institutions, it might be numbers of
incoming students and retention rate of students. For a non-profit it
might be the number of people willing to give to support the work of the
organization and the number of people who use its services.

All of this leads into correct positioning of operations within your
organization.

* What are the core competencies of operations and how do they serve the internal business units and the organization as a whole?

* What core competencies are you missing and should develop in order to better support your orgnization's mission?

Maintaining Existing IT Infrastructures
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

The most visible role of Operations is to maintain the status quo.
For the system administrator this means maintaining servers and
processes such as logging, monitoring, backups, authentication, or
naming services. For the network administrator it means maintaining
routers, switches, the edge network, gateways, or the relationship
with the corporate Internet Service Provider (ISP). A security
engineer might be responsible for maintaining a vulnerability
scanning capability, incident response policy and processes, intrusion
detection systems, firewalls, and a customer security awareness
training program. Operations may also be responsible for maintaining
access to internal services (e.g. financial systems, corporate content
management systems, procurement systems, etc.) that may impact the
various business units within the organization. These roles are
distinct but there is sometimes overlap between them in smaller
organizations where fewer people server in multiple roles.

Translating Customer Requirements
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
Operations roles are customer service positions. These careers
require a level of customer interaction because the services delivered
by the Operations professional must be driven by customer needs.
In this case, customer is used to mean the business, organization,
or other entity that is employing the Operations professional. Some
questions to ask to help the Operations person understand requirements
from the customer perspective:

* What is the core mission of this organization?
* How does Operations support, hinder, or allow your organization to innovate for the mission?
* Who are your core customers (internal, external, or both)?
* What does the organization need from the Operations professionals?
* Why should this organization come to these Operations people for this service or solution? (What is the value proposition for Operations within this organization?)?
* How could Operations provide more value: higher level of competitiveness, faster service delivery, stronger security, or other benefit that aligns with the mission?

Translating customer requirements is key to focusing the efforts
of Operations. Operations work can be a slippery slope where the
professionals are spreading themselves too thin on projects and
deliverables that do not serve the organization's mission. One way
to focus the efforts of Operations is to answer these questions and
to ensure that the Operations organization, whether insourced or
outsourced, is delivering services that provide the most value.

Protection of Information and Services
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Often the Operations professionals in an organization are the people
who most completely understand the technical risk to organizational
assets from an IT perspective. Senior management within an organization
will usually understand risks related to financials, competition,
manufacturing, etc. but they often do not understand IT enough to make
an informed decision. Operations professionals are the ones with the
deep-dive technical expertise required to comprehend risks, threats,
vulnerabilities, and countermeasures then translate them into
language senior management can understand.

This is another area where the Operations professional is communicating
with the organization's leaders to advise on appropriate actions
to address IT security where it makes sense for the organization.

Areas where organizations need the Operations professional
to advice on IT security could include threats to data from internal
and external sources, hardware failure, site availability or
resilience, data preservation, and information integrity. Again,
these areas are dependent on the organization's mission.

For example: an ecommerce organization will most likely want strong
site availability and protection of customer personal information.
The Operations professionals might build a site with high resilience
and availability including use of Content Delivery Networks (CDNs),
strong encryption not only for the ecommerce session but also data
at rest, role-based access for internal employees accessing customer
information to reduce access to only those people who need access
to that information. Organizational leaders often do not understand
how these solutions are implemented so it is up to the Operations
professional to communicate the threat, solution, cost, impact to
the organization of implementing the solution.

Advising within Current Limitations
+++++++++++++++++++++++++++++++++++

The Operations professional who advises an organization must also
consider limitations that impact the potential solution. Cost,
timing, expertise within the organization, available time of the
people who would implement the solution, or IT security issues may
be considerations. For example, decision makers within the
organization will need to know what is possible and for what cost
so they can make the decision how to spend the organization's money.
Good, fast, or cheap (pick two): it may be the Operations professional's
responsibility to explain this concept from an IT perspective.

 * http://en.wikipedia.org/wiki/Project_management_triangle

Thinking broadly about decisions and acting decisively 
------------------------------------------------------

Sometimes written as "mindful of the implications of a choice for all
the affected parties"

These people can look at a problem from the viewpoint of other
people and business units within the organization. Instead of insular
thinking, they come at a problem with a broad-minded perspective.
How do decisions impact other areas of the organization and,
alternatively, how does the organization view this particular issue?
Those with strong acuity for business will see the big picture and
be able to understand the implications of a decision on more than
just operations.

In some cases it may not be a problem, but an opportunity that injects
potential life into an organization or recalibrates it. Business
leaders, stakeholders, customers or whatever you call them often don't
understand what technology can do for them. Operations should understand
the organization well enough to see where technology can support
innovation. This leads into change as a constant.

What would it take to make this happen? What are the missing ingredients
for success?

See change as a constant and support or promote change as needed
----------------------------------------------------------------


Specific Examples
=================

Below are some specific examples to demonstrate the importance of soft
skills in operations. In each example, soft skills closed the deal
because they enabled the operations person to see the situation from
other perspectives and communicate the needs of operations in terms of
the organization as a whole.

Selling system changes and new proposals
----------------------------------------

Negotiating budgetary constraints vs. need/want requirements
------------------------------------------------------------

Evaluating a product offering
-----------------------------

The importance of Documentation
-------------------------------

What to document
^^^^^^^^^^^^^^^^

* Runbooks? SOP? (cparedes: might be worthwhile even though we want to automate SOP's away as much as possible - what should we check at 2 AM? What do folks typically do in this situation if automation fails?)

* Architecture and design (cparedes: also maybe talk about *why* we choose that design - what problems did we try to solve? Why is this a good solution?) How to manage documentation

Documentation through Diagrams
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Developing the Trusted Advisor Relationship
-------------------------------------------
Be enablers and problem solvers, not a hinderance. The antithesis of the BOFH.


Working with other teams
========================

