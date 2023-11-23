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
Sudo allows users to run a command or program using the security privileges of another user, this will be the superuser or root by default. As the name implies the superuser is an user at the most powerful level granting the command or programs run without restrictions. A security policy can be used to determine what priviliges a user has to run sudo. Usually a policy will require users to authenticate themselves such as entering a password.

Sudo can be used by simply placing it infront of the command you want to run. Alternatively the command 'su' can be used, this allows you to log in as another user or the root. Whereas sudo will only run one command as another user, su will run all following commands as the specified user.

So why not just make my life easy, run su at the start of my session and get rid of those peksy security restrictions?
Well there are several reasons why you shouldn't log in as root:

- **Security** - It destroys the built-in security model that has been put there to protect the underlying system from being messed with.
- **Programs** - Running an program as root gives a program freedom. Total freedom. Running a program on a standard user account will give it write access to your homefolder only. Running it as root removes these restrictions, even giving the program write access to the system files. A malicious or buggy program could decide to delete all the files it can access, and since you elevated its rights to the root it can.
- **Yourself** - Su is trap for the inexperienced linux user. You are no longer a 'mere mortal' one click, a tap on the enter button and its done. The system won't interrupt you or ask if you are sure you want to run a heavy impact command. You might be exploring some new commands and before you know it you're formatting your disks.

It's good practice to be wary of what permissions you give your commands or programs. Educate yourself before you go and give away root powers.

For more information on how to use the 'sudo' command please refer to the man page:
https://www.sudo.ws/man/1.8.15/sudo.man.html

History and Lore
================

The Morris Worm
---------------
http://www.snowplow.org/tom/worm/worm.html

/bin/false is not security
--------------------------
https://web.archive.org/web/20150907095805/http://www.semicomplete.com/articles/ssh-security

