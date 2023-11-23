Remote Filesystems 101
**********************

NFSv3
=====

iSCSI
=====

SAMBA/CIFS
==========

Intro
-----
Samba is an Open Source/Free Software suite that provides seamless file and print services to SMB/CIFS clients. Samba is freely available, unlike other SMB/CIFS implementations, and allows for interoperability between Linux/Unix servers and Windows-based clients.

For more information, you can visit the `Samba Wiki
<https://wiki.samba.org/index.php/Using_Git_for_Samba_Development>`_.

Install
-------
The Samba source code is distributed via http. View the download area via https://download.samba.org/pub/samba/. The file you probably want is called samba-latest.tar.gz. Old releases are available in the Samba archives.

The Samba distribution GPG public key can be used to verify that current releases have not been tampered with. Using GnuPG, simply download the Samba source distribution, the tarball signature, and the Samba distribution public key. 

You can also install samba on Linux distributions by using the command: "sudo yum install samba".

Roles
-----
There is many code and configuration about samba which you can use from experienced users. Many Github projects are available. 

There is also a specific platform for Linux server roles:
`Ansible Galaxy
<https://galaxy.ansible.com/>`_.

On this platform are roles which you can install on you server with Vagrant and Ansible. 

When you search for samba on Ansible Galaxy you get 26 results for a configuration of samba. An example of a good documented role is that from 
`Bertvv
<https://galaxy.ansible.com/detail#/role/3118>`_.

Configuration
-------------
The settings can be changed in the smb.conf file, this file is located under /etc/samba.
In this file you can configure shares, users ...

You can combine the samba role with a vsftpd role so you can access your files from a Windows host.

