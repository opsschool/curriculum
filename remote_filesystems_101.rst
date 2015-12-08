Remote Filesystems 101
**********************

NFSv3
=====

iSCSI
=====

SAMBA/CIFS
==========


Samba is an Open Source/Free Software suite that provides seamless file and print services to SMB/CIFS clients. Samba is freely available, unlike other SMB/CIFS implementations, and allows for interoperability between Linux/Unix servers and Windows-based clients.

The Samba source code is distributed via http. View the download area via <a href="https://download.samba.org/pub/samba/">HTTP</a>. The file you probably want is called samba-latest.tar.gz. Old releases are available in the Samba archives.

The Samba distribution GPG public key can be used to verify that current releases have not been tampered with. Using GnuPG, simply download the Samba source distribution, the tarball signature, and the Samba distribution public key. Then run

```
$ gpg --import samba-pubkey.asc
$ gunzip samba-version.tar.gz
$ gpg --verify samba-release.tar.asc
gpg: Signature made Tue 20 Nov 2007 07:12:04 PM CST using \
  DSA key ID 6568B7EA
gpg: Good signature from "Samba Distribution Verification Key \
  ‹samba-bugs@samba.org›
```

You can also install samba by using the command: sudo yum install samba.

The setting can be changes in the smb.conf file, this file is located under /etc/samba.


