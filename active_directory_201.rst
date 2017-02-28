####################
Active Directory 201
####################

Detailed Breakdown of Active Directory Components/Services
==========================================================

NTDS, the database
------------------

Kerberos KDC
------------

DNS
---

DHCP
----

Global Catalog
--------------

The Global Catalog (GC) holds a subset of attributes for all objects in a multi-domain forest.
It can be used to search all objects in the forest.
For example, when searching for the a specific user, the GC can be searched to find users not only in the current domain but also in other domains in the forest.
The attributes included in the GC's subset of object attributes are those required by the schema and those most commonly used in searches.
Attributes can be added and removed from the Global Catalog using the Active Directory Schema snap-in. Alternatively, they can be added with PowerShell using the following command:

.. code-block:: powershell

	Set-ADObject "yourAttribute" -Replace @{"isMemberOfPartialAttributeSet"=$true}

When Active Directory Domain Services is installed, the GC is added to the first domain controller in the forest.
In order to avoid traversing a WAN to gather GC information, it is recommended that each Active Directory site have at least one Global Catalog server.
It can be added to additional domain controllers with Active Directory Sites and Services by following the steps in this link:
https://technet.microsoft.com/en-us/library/cc755257.aspx

It can also be done with a PowerShell command:

.. code-block:: powershell

	Set-ADObject "CN=NTDS Settings,CN=yourDC,CN=Servers,CN=Default-First-Site-Name,CN=Sites,CN=Configuration,DC=yourDomain,DC=COM" -Replace@{options='1'}


Universal groups can contain objects from any domain in the forest.
The GC holds the membership for all universal groups.
As part of the authorization process, the authenticating DC retrieves the security identifier for all the member's security groups, which includes universal groups.
Therefore, if the environment has more than one domain, the global catalog server must be queried during authentication in order to find objects from the universal groups in other domains.

FSMO Roles
----------

FSMO stands for Flexible Single Master Operation.
FSMO roles are also referred to as Operations Master roles.
They are various functions performed by Domain Controllers (DCs) in Active Directory.

There are 5 FSMO roles.
Three of the roles exist for every domain in the forest and two apply to the entire forest.
While it is possible to have 5 domain controllers each with a separate FSMO role, it is common to have multiple roles on one server.
The forest-wide roles are Schema Master and Domain Naming Master.
The domain-wide roles are PDC Emulator, RID Master, and Infrastructure Master.

Schema Master
^^^^^^^^^^^^^

This role is forest-wide.
This domain controller can update the schema.
Given that the schema is not frequently updated, availability of this DC is less important than it is for other roles.
By default, the role is held on the first server promoted to a DC in the forest.

Domain Naming Master
^^^^^^^^^^^^^^^^^^^^

This role is also forest-wide.
It controls changes to the forest namespace.
It can add, remove, rename, and move domains.
Since domain name changes are infrequent, this role is rarely used.
By default, the role is held on the first server promoted to a DC in the forest.

PDC Emulator
^^^^^^^^^^^^

This role is domain-wide.
PDC stands for primary domain controller.
It holds the latest password hashes and lockout information for all accounts.
Password changes are immediately forwarded to the PDC from other DCs.
Also, it is the authoritative time source for the domain.
This role most directly affects users.
Thus, the availability of the DC with the role is critically important.

The name emulator came from its ability to emulate the functionality of a Windows NT 4.0 PDC, which was useful in a mixed environment with Windows NT 4.0 BDCs. Extended support for Windows NT 4.0 server ended in 2004 and as a result, it is unlikely to be seen in modern environments.

RID Master
^^^^^^^^^^

This role is domain-wide.
A RID is a relative identifier used as part of the SID.
The SID is used to uniquely identify objects.
It is similar in function to the GUID but used for security.
As SIDs must be unique in the domain, the RID master maintains the global pool of unique RID values for the domain, a subset of which are given out to other DCs.
In past versions of Active Directory, rolling back a DC on a virtual machine could cause RID reuse.
However, newer versions of Active Directory are able to detect and prevent this event when a virtual machine is restored.

Infrastructure Master
^^^^^^^^^^^^^^^^^^^^^

This role is domain-wide.
It is used to help manage cross-domain references.
Thus, if the forest only holds one domain, the role is not relevant as there are no cross-domain references.
In order to update the cross-domain references, it compares the data to the Global Catalog (GC).
If all DCs have the GC, then the placement of this role is irrelevant as it will already have the updated data.

Advanced Active Directory Maintenance
=====================================
