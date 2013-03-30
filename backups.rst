Backups
*******

Backups are a critical part of preserving an organization's
information. RAID, clustering and other real-time copies of
data do not help when the data is deleted, the buildings containing
these live sources are compromised by weather, human intervention,
or facilities failures.

Could the organization survive if part or all of its information
were lost?

Just remember, RAID is not a backup solution.

.. TODO:: Why is RAID not a backup solution? Experienced ops folks know why, but a junior ops person may not.

Designing a backup policy
=========================

Operations needs to work with organizational leaders to develop a
backup policy that protects the organization's information while
balancing cost and other factors important to the organization.

Operations engineers need to be able to understand organizational
needs in order to evaluate backup solutions, strategies, and current
best practices. They must then translate this knowledge into
recommended solutions which management can use to make a sound
business decision. The ability to assemble a cohesive plan, including
a cost-benefit analysis of available options, can help the organization
make the right decision for the business on an appropriate backup
policy.

Deciding what to backup
-----------------------

Consider cost, time, and staff resources when deciding what to
backup. Does it make sense to backup every client hard drive when
a system can be re-imaged over the network? Does the configuration
management solution maintain client customizations? If so, operations
may not need to backup the operating system on clients and can focus
on user data. Ideally, the goal is not to restore clients starting at
the operating system level.

Where is organizational data stored? Is data local on each client or
does the organization use a file server? Do users actually
store their data on the file server? Some users have had bad experiences
using the network and move all of their data locally, or refuse to
store anything on the server in case the network becomes unavailable.
If users travel frequently, they may store all of their data
on a laptop. This poses additional backup challenges when laptops
are not always on the organization's network to talk to the backup
server.

Some organizations perform incremental backups on all clients to
protect against users who decide where best to store their data.

What about data stored on servers? Do all servers need to be backed
up? What is the available backup window and will the organization's
network push enough data to a backup server to fit within that
window? How will taking the backup affect each application's
performance?

Ultimately, deciding what to backup is a business decision. What
information is critical to the business and how much is the
organization willing to invest to protect that information? There
may also be laws requiring the organization to retain certain data
that will influence this decision.

When not to backup
^^^^^^^^^^^^^^^^^^

.. TODO:: restoring? what not to backup, regulations on same, how to store them (PCI)

Retention periods
-----------------

Is there a risk to having a long retention period?

.. TODO:: The idea of such a risk may not be immediately clear to a beginning ops person.

What is the cost to retain backups for a 30-days, 6-months, or 1 year?

Is there a business requirement for keeping backups for a specific length of time?

Are there laws for maintaining electronic records for a specific period of
time? 


If the organization is required to adhere to Freedom of
Information Act (FOIA) law, sometimes long retention periods are a
detriment. FOIA requests can cover any kind of data including
individual emails on a specific subject. If the information doesn't
benefit the organization, then it might be a liability due to
labor-intensive searches through ancient backups.

.. TODO:: How does FOIA affect what information an organization needs to have available? Assume the reader is a civilian and doesn't know how FOIA affects an organization.

Offsite backups
---------------

How sensitive is the organization's information? Storing a copy
of the organization's crown jewels in an insecure location can
expose the organization to loss, unauthorized modification, or
theft of intellectual property.

Is the organization required to encrypt offsite data?

Does the organization have a second geographically distributed
location that could house an off-site backup of the organization's
information?

Designing a Backup System
=========================

Backup type and frequency
-------------------------

What kind of information does the organization rely on to do
business? Do hourly changes need to be captured or can the
organization survive with backups every 12-hours or once per day?

* Full backups
* Incremental backups
* Replication
* Snapshots
* Bare-metal restore vs data only
* online/offline

.. TODO:: media -- should someone address the state of backup media? Some places are still doing tape. What about orgs who rely on standalone consumer-grade disks for client backups (e.g. Time Machine)? Risks, cost to maintain.

Cost of backups
---------------

What is the cost of not doing backups?


Verification
============

Test backups. If data cannot be restored then what was the
point of backing it up in the first place.

Recovery testing
----------------

How long does it take to restore the largest backup set?

Integrity of backups
--------------------

Completeness of backups
-----------------------

Security implications
=====================

.. TODO:: Using backups to restore to a known "good" state after an incident just serves to put the machine in a known vulnerable state (security hole that was exploited is now back in operation)

.. TODO:: can be used to restore system state that can be useful in a post mortem after an incident (say the attacker covered their tracks but backups were able to capture a rootkit before it was removed or before logs wer etampered with)

Recovery basics
===============

Secure data destruction
=======================

Information Lifecycle Management in relation to backups
========================================================

Main goal of backups is restore system state including data in case of issues and ILM, have data available for functional
reasons other than uptime.

Main items to cover in this chapter are:

Archiving
---------

Data replication
----------------

