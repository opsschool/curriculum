Identity Management 101
***********************

LDAP
====

LDAP (Lightweight Directory Access Protocol) is an application protocol used for accessing and maintaining directory services. Directory Services uses a database-styled system to correlate information about objects in that directory with metadata about that object.

Directory services provide information about a unique ID, in much the same way that an ISBN number identifies a book in a library. That unique ID can be a user, a computer, or a group, or any number of other objects, depending on how the directory index (or schema, as it is referred to) has been specified.

The metadata for that object can include things such as a Display Name, Address, Date Created, Date Modified, Home Directory, etc... Since LDAP can be extended by modifying or adding to the schema, it is a very flexible format and serves as the base of most modern directory services.


Active Directory
----------------

Active Directory is Microsoft's implementation of LDAP, coupled with _Kerberos: http://en.wikipedia.org/wiki/Kerberos_(protocol) encrypted authentication security. One benefit of Active Directory is _Group Policy:http://en.wikipedia.org/wiki/Group_Policy , which is a flexible and granular system for controlling a great many user and computer settings based on flexible criteria, such as Organizational Unit or Group membership.

_Active Directory Domain Services Port Requirements:  http://technet.microsoft.com/en-us/library/dd772723(v=ws.10).aspx  can be found on Technet.

Active Directory works best with Windows Domain Controllers serving as both DHCP and DNS servers, in an AD-integrated DNS zone. While DNS or DHCP can be handled by alternative systems such as BIND, this is an advanced configuration, and should not be attempted in most scenarios.

The Active Directory schema has been modified with each release of Windows Server, and will often be modified when deploying core Microsoft server applications such as Exchange or Certificate Services. The overall state of the domain is referred to as the _Domain Functional Level: http://support.microsoft.com/kb/322692 - this may require consideration when determining requirements of a project or feature implementation.

Active Directory is typically managed through a variety of tools, including:

* Gui Tools
- Active Directory Users & Computers (to manage the users, computers, and groups in a domain)
- Active Directory Sites & Services (to manage the replication of AD to different sites/servers)
- adsiedit (a useful tool for viewing the attributes of objects within the domain)

* Command Line Tools
- Powershell (through the Active Directory cmdlets)
- ldp (an extremely low-level tool for interacting with LDAP directly - not recommended for most uses)

OpenLDAP
--------

Directory389
------------

NIS
===


