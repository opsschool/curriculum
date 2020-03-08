Security 101
************

Authentication in unix
======================

.. todo::
   Discuss how authentication works.
   Touch on ``/etc/(passwd|group|shadow)``, hashing.
   What are groups? Lead in to the users/groups permissions model and how
   permissions are based on the user/group/other bits.

Adding and deleting users and groups
====================================

Standard unix filesystem permissions
====================================
The simplest way of displaying filesystem permissions is by typing:

.. code-block:: console

  $ ls -l
  drwxr-xr-x   2 john  company  68  3 Oct 10:34 files
  -rwxrwxrwx   1 john  company    0  3 Oct 10:29 hello_world.txt

The left column is a 10-character string that indicates the permissions for a file. It consists of the symbols d, r, w, x, -.

- **Directory (d)** - This is the first character in the permissions string. 
  This indicates a *directory*. 
  Otherwise, the first character is a - to indicate that it is not a directory.
- **Read (r)** - The *read* permission allows the user to read the contents of the file or list the files in the directory.
- **Write (w)**- The *write* permission allows the user to write or modify a file. 
  In the case of directories, the use may delete files from the directory or move files into the directory.
- **Execute (x)** -The *execute* permission allows the user to execute a file or access the contents of a directory. 
  In the case of directories, this indicated that the user may read files in the directory, provided that the user has read permission on an individual file.

The 9 remaining characters are split into 3 sets to represent the access rights based on 3 groups of users. 
Take the "files" directory above as an example, we can split the characters like this: ``[d][rwx][r-x][r-x]``

- The first character, as explained above, indicates a directory or a file
- The first group gives the file permissions for the *owner* of the file or directory. 
  This means that the user "john" has read/write/execute permissions to the directory.
- The second group gives the file permissions for the *group* of users to whom the file or directory belongs to. 
  This means that anyone who is under the group "company" has read/execute permissions to the directory.
- The third group gives the file permissions for *other* users. 
  Basically anyone who are not the owner or a part of the user group. 
  This means that everyone else has read/execute permissions to the directory.

Some more examples of permissions:

- ``-rwxrwxrwx`` is a file everyone can read, modify (including delete), and execute.
- ``-rw-------`` is a file only the user can read and modify.


PAM
===

Chroot, jails and containers
============================

Sudo (or, "Why you should not log in as root")
==============================================

History and Lore
================

The Morris Worm
---------------
http://www.snowplow.org/tom/worm/worm.html

/bin/false is not security
--------------------------
https://web.archive.org/web/20150907095805/http://www.semicomplete.com/articles/ssh-security

