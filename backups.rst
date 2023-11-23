Backups
*******

Backups are a critical part of preserving an organization's
information. RAID, clustering and other real-time copies of
data do not help when the data is deleted, the buildings containing
these live sources are compromised by weather, human intervention,
or facilities failures.

Could the organization survive if part or all of its information
were lost?

Remember, RAID is not a backup solution.
RAID protects against the loss of individual drives, but if data gets corrupted
or deleted, RAID will replicate that corruption across all RAID stripes.
Making backups on separate media is the best way to protect yourself against
both "hard" and "soft" data loss.

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

When should you backup?
=======================

The best bet is to backup your files on a fairly regular basis–daily if possible. If you’re using an online backup solution, they are often configured to immediately start syncing any changed files when your PC is idle for a little while. This can be a great way to keep your files safe without having to wait for the next backup.

What Files Should You Backup?
=============================

The most important files to backup are probably your documents, pictures, music, and other user files, but they are not the only files that you need to backup. Let’s walk through some of them.	
-	Documents: You should backup your entire documents folder all the time. This should be a no-brainer.
-	Music: If you’ve paid lots of money for MP3 downloads, you’ll probably be sad to lose them. Make sure to include this folder. Note: if you’re an iTunes user, you should make sure to backup your iTunes folder, which is thankfully usually inside this directory.
-	Pictures & Videos:  The photos might not have actually costed you anything, but you’ll probably be more sad about losing memories than paying for music downloads again.
-	Desktop Email: If you’re using Outlook or Windows Live Mail, make absolutely certain that you’ve backed up the files from these applications. Outlook stores all your email in a .PST file.
-	Application Settings: If you look within the AppData folders, you’ll see directories for each and every application you’re running. These settings can often be restored from a backup so you don’t have to tweak everything again. Just head into C:\Users\Username\AppData\ to see the Local, Roaming, and LocalLow folders that contain many settings for your applications.
-	Virtual Machines: If you use virtual machines for real work, you should probably create a backup of your virtual machines at some point. We wouldn’t necessarily recommend backing these up every single night, but you should at least consider some type of backup plan.
-	Bookmarks: Most browsers other than Internet Explorer actually make it difficult to backup your bookmarks using Windows Backup, but the much better option is to sync your bookmarks to the cloud. Naturally, we’ve got a full list of all the bookmark syncing services that you can use. If you’d rather use local backup, you can simply backup the application settings folder and restore that—this works especially well for Firefox.
 

Backup These Files More Easily
==============================

Instead of trying to find all those locations, backup your entire Users folder, which is at C:\Users\Username in Windows 7 or Vista, and C:\Documents and Settings\Username for Windows XP. This will include all of those files, unless you’ve stored them somewhere else.

Files You Should Not Bother Backing Up
======================================

There’s simply no reason to backup these directories:
-	Windows: There’s almost never a reason to backup your Windows directory, as you’re going to have to reinstall the whole thing anyway, so this backup will likely do you no good.
-	Program Files: You’re going to have to reinstall your applications if your computer dies and you have to reinstall.








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

.. TODO:: can be used to restore system state that can be useful in a post mortem after an incident (say the attacker covered their tracks but backups were able to capture a rootkit before it was removed or before logs were tampered with)

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

