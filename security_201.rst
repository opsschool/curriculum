Security 201
************

Centralised accounts
====================

LDAP and Kerberos
-----------------

Active Directory
----------------

NIS, NIS+, YP, Hesiod
---------------------


Firewalls and packet filters
============================

host vs network
---------------

"Crunchy outer shell, soft center is bad"

In general terms, implementing defense-in-depth strategies is always a sensible 
practice.  The concept in which multiple layers of security controls (defense) 
are placed throughout an information technology (IT) system.  Its intent is to 
provide redundancy in the event a security control fails or a vulnerability is 
exploited.

Implementing a firewall on the network **and** host-based packet filters 
provides defense-in-depth layers to your infrastructure.  In today's landscape, 
the firewall aspect could be a service like ec2 security groups with host-based 
packet filters such as iptables, Windows Firewall (or WFAS).  Although this can 
add additional complexity to deployment, that is not a reason to not implement 
it where appropriate.

The defense-in-depth concept is mostly regarded in terms of attack and 
compromise, however in ops it also safeguards **us** as everyone makes mistakes.  
Sometimes, we ourselves or our colleagues are the point of failure.

Strange as it may seem, people often make the mistake of disabling "rules" when 
something is not working and they cannot figure out why.  The *just checking* 
test.  This is always the first mistake.  In real world operations these things 
do happen, whether it is a *just checking* mistake, an incorrect configuration 
or action, sometimes we make serious mistakes, to err is human.

In all these situations using a firewall/other and host-based packet filters 
comes to the fore:

- It protects us, the people working on the systems, from ourselves.
- It protects the organisation's assets in the event of a failure or compromise 
  at either layer, whether it be user error or a systematic failure.
- It keeps us proficient and *in-hand* on both the network specific 
  security implementation, the host-based security practices and the 
  applications related to their management.
- It is good practice.

An old skool, real world example:
---------------------------------

- You have a MSSQL server running on Windows Server 2000 (no "firewall" back 
  then & ip filters are not enabled) - it has both private and public network 
  interfaces.
- The MSSQL server has a public NIC because it has run replication with your 
  customer's MSSQL server sometimes for dev purposes and catalog updates.
- You have rules on the diversely routed, mirrored NetScreen NS1000 firewalls 
  that allows port 1433 between the MSSQL servers only on the public interface.
- Your colleague has an issue that cannot be resolved and quickly just sets the
  firewall/s to "Allows from all", the *just checking* test.
- **Immediate** result - network unreachable.
- SQL Slammer had just arrived and proceeded to gobble up 2Gbps of public T1 
  bandwidth.
- All points into the network and the network itself are saturated until you 
  debug the issue via the serial port on the firewall and figure out what 
  happened.

The synopsis is that a practice of disabling rules was implemented, as it had 
always been the last line in debugging.  It was a practice that had been done 
many times by the engineers in the organisation at the time with no "apparent" 
consequences in the past.  This time differed in that the MSSQL server had a 
public interface added to allow for replication with the customer **and** 
SQL Slammer was in the wild.

If the MSSQL server had ip filtering enabled as well the situation would have 
been mitigated.  Needless to say, it was the last time "Allow from all" debug 
strategy was ever implemented.  However it is useful to note, that because this
was done all the time, the engineer in question did not even tie the action of 
"Allow from all" and the network becoming unreachable together at the time, 
because the action previously had never resulted in the outcome that was 
experienced in this instance.

Stateful vs stateless
---------------------

IPTables: Adding and deleting rules
-----------------------------------

pf: Adding and deleting rules
-----------------------------


Public Key Cryptography
=======================

A public key infrastructure (PKI) supports the distribution and identification of public encryption keys, enabling users and computers to both securely exchange data over networks such as the Internet and verify the identity of the other party. When PKI is not being used, sensitive information can be easily encrypted. Any form of sensitive data exchanged over the Internet is reliant on PKI for its security. 

A PKI consists out of 4 elements. First there's the trusted party, also known as the certificate authority (CA). This certificate acts as the root of trust and provides services that authenticate the identity of individuals, computers and others. Next there's the policies that govern the operation of the PKI. The third element are the digital certificates themselves and the last element are the applications that are written to use the PKI. 

When we ask the question, where is PKI being used? Basically, it's very often being used when there's sensitive information that has to cross between two places. For example, customers expect their bank to verify their identity before allowing them to view account balances or transfer funds via an online application.  This authentication is often accomplished using user names and passwords, but PKI provides an attractive alternative. SSL has the optional ability to require a client-side certificate before establishing a secure and encrypted connection. A big advantage of authentication using PKI is that the server does not need to have access to a database of user passwords.  Users may still need to supply a password to unlock their private key in their certificate, but that password is stored and managed locally on the client. This is inherently more secure and easier to manage than server password databases.

Now we know the most important things about PKI its concept, but we might still be wondering, why are PKI's so important? Well, the answer is that almost all security controls come down to authentication and access controls. Encryption is a powerful tool, but unless that data can be decrypted, it might not be very useful. If a company cares about the integrity of its data and systems, it must either deploy a PKI with an appropriate set of checks and balances or use a third party service it can trust. Failure to do so leaves an organization exposed and increasingly vulnerable compared to other potential victims.


Using public and private keys for SSH authentication
----------------------------------------------------


Two factor authentication
=========================

Two-factor authentication adds a second level of authentication to an account log-in. When you have to enter only your username and one password, that's considered a single-factor authentication. 2FA requires the user to have two out of three types of credentials before being able to access an account. The three types are:

- Something you know, such as a personal identification number (PIN), password or a pattern
- Something you have, such as an ATM card, phone, or fob
- Something you are, such as a biometric like a fingerprint or voice print

Although many may think that 2FA is something new, it's actually not new at all. For example when you use your credit card and you must enter in your ZIP code to confirm a charge (online banking), that's 2FA in action. There's the physical factor (the card), and the knowledge factor (the ZIP code). But just because it's been around for a long time doesn't mean that it's easy to set up and use. It adds an extra step to your log-in process, meaning it can take some time to set up. Of course, alot depends on your patience and your willingness to spend some time on ensuring a higher level of security. 

2FA isn't protected against hackers just because 2FA requires an extra credential. But then again, hackers will definitely have a harder time of intercepting or receiving your credentials.

Building systems to be auditable
================================

Data retention
--------------

Log aggregation
---------------

Log and event reviews
---------------------

Role accounts vs individual accounts
------------------------------------


Network Intrusion Detection
============================


Host Intrusion Detection
=========================


Defense practices
=================


Risk and risk management
========================


Compliance: The bare minimum
============================

What is compliance and why do you need it?

What kinds of data can't you store without it?

Legal obligations


Dealing with security incidents
===============================


ACLs and extended attributes (xattrs)
=====================================


SELinux
=======


Data placement
==============
Eg, local vs cloud, the implications, etc


Additional reading
==================
Ken Thompson, Reflections on Trusting Trust:
http://dl.acm.org/citation.cfm?id=358210
