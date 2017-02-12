Package management
******************

Workflow
========

What is a package manager?
==========================

Modern operating systems use package managers to take care of the installation,
maintenance and removal of software. On Windows this is `Windows Installer
<http://en.wikipedia.org/wiki/Windows_Installer>`_ (formerly Microsoft
Installer). On Linux there are two popular package managers:

* APT (used by Debian, Ubuntu)
* RPM (RedHat, CentOS, Fedora, SuSe)

Specific commands for each system vary, but at their core they all provide the
same functionality:

* Install and uninstall packages
* Upgrade packages
* Install packages from a central repository
* Search for information on installed packages and files

RPM and yum (RedHat, CentOS, Fedora, Scientific Linux)
======================================================

In the following examples, we will be using the package ``dstat``. The
process however applies to any software you may want to install.

Yum provides a wrapper around RPM, which can be used to search for, and install
packages from multiple package repositories. It also resolves dependencies, so
that if a package has prerequisites, they will be installed at the same time.

If your Linux distribution uses RPM and yum, you can search for packages
by running:

.. code-block:: console

   bash-4.0$ yum search dstat
   ======================== N/S Matched: dstat =========================
   dstat.noarch : Versatile resource statistics tool

Installing packages
-------------------

You can install a package using yum, by running:

.. code-block:: console

   bash-4.0$ yum install dstat

   =============================================================================
    Package        Arch            Version              Repository         Size
   =============================================================================
   Installing:
    dstat          noarch          0.7.0-1.el6          CentOS-6          144 k

   Transaction Summary
   =============================================================================
   Install       1 Package(s)

   Total download size: 144 k
   Installed size: 660 k
   Is this ok [y/N]:

If you have a downloaded RPM file, you can also install the package directly
with the ``rpm`` command:

.. code-block:: console

  bash-4.0$ rpm -i dstat-0.7.0-1.el6.noarch.rpm

Upgrading packages
------------------

RPM and yum both make it easy to upgrade existing packages, too.
Over time, new packages may be added to the yum repositories that are configured
on your system, or you may have a newer RPM for an already installed package.

To upgrade a package using yum, when a newer package is available, simply ask yum
to install it again:

.. code-block:: console

   bash-4.0$ yum install dstat

To upgrade all packages that have newer versions available, run:

.. code-block:: console

   bash-4.0$ yum upgrade

To upgrade a package with an RPM file, run:

.. code-block:: console

   bash-4.0$ rpm -Uvh dstat-0.7.1-1.el6.noarch.rpm

Uninstalling packages
---------------------

To uninstall a package using yum, run:

.. code-block:: console

   bash-4.0$ yum remove dstat

Similarly, you can uninstall a package with rpm:

.. code-block:: console

   bash-4.0$ rpm -e dstat

Cleaning the RPM database
-------------------------

You can clean the RPM database, forcing it to refresh package metadata from its
sources on next install or upgrade operation.

.. code-block:: console

   bash-4.0$ yum clean all

Querying the RPM database
-------------------------

Occasionally you will want to find out specific information regarding installed
packages. The ``-q`` option to the ``rpm`` command comes in handy here. Let's
take a look at a few examples:

One common task is to see if you have a package installed. The ``-qa`` option
by itself will list ALL installed packages. You can also ask it to list specific
packages if they are installed:

.. code-block:: console

   bash-4.0$ rpm -qa dstat
   dstat-0.7.0-1.el6.noarch

Now let's say we want to list all of the files installed by a package. The
``-ql`` option is the one to use:

.. code-block:: console

   bash-4.0$ rpm -ql dstat
   /usr/bin/dstat
   /usr/share/doc/dstat-0.7.0
   /usr/share/doc/dstat-0.7.0/AUTHORS
   /usr/share/doc/dstat-0.7.0/COPYING
   /usr/share/doc/dstat-0.7.0/ChangeLog
   ...

We can also do the reverse of the previous operation. If we have a file, and
want to known which package it belongs to:

.. code-block:: console

   bash-4.0$ rpm -qf /usr/bin/dstat
   dstat-0.7.0-1.el6.noarch

Creating packages
-----------------
.. todo:: Mention spec files and roughly how RPMs are put together.
.. todo:: Then introduce FPM and tell them not to bother with spec files yet.

There are two todos here.

dpkg and APT (Debian, Ubuntu)
=============================

In the following examples, we will be using the package ``dstat`` in our examples. The
process however applies to any software you may want to install.

apt provides a wrapper around dpkg, which can be used to search for, and install
packages from multiple package repositories.

If your Linux distribution uses dpkg and apt, you can search for packages
by running:

.. code-block:: console

   bash-4.0$ apt-cache search dstat
   dstat - versatile resource statistics tool

Installing packages
-------------------

You can install a package through apt, by running:

.. code-block:: console

   bash-4.0# apt-get install dstat

   The following NEW packages will be installed:
     dstat
   0 upgraded, 1 newly installed, 0 to remove and 0 not upgraded.
   Need to get 0 B/79.3 kB of archives.
   After this operation, 351 kB of additional disk space will be used.
   Selecting previously unselected package dstat.
   (Reading database ... 124189 files and directories currently installed.)
   Unpacking dstat (from .../archives/dstat_0.7.2-3_all.deb) ...
   Processing triggers for man-db ...
   Setting up dstat (0.7.2-3) ...

If you have a downloaded DEB file, you can also install the package directly
with the ``dpkg`` command:

.. code-block:: console

  bash-4.0# dpkg -i dstat_0.7.2-3_all.deb

Upgrading packages
------------------

dpkg and APT both make it easy to upgrade existing packages, too.
Over time, new packages may be added to the apt repositories that are configured
on your system, or you may have a newer deb for an already installed package.

To upgrade using apt, when a newer package is available, simply ask apt to
install it again:

.. code-block:: console

   bash-4.0# apt-get install dstat

To upgrade a package with an deb file, run:

.. code-block:: console

   bash-4.0# dpkg -i dstat_0.7.2-3_all.deb

Uninstalling packages
---------------------

To uninstall a package using apt, run:

.. code-block:: console

   bash-4.0# apt-get remove dstat

Similarly, you can uninstall a package with dpkg:

.. code-block:: console

   bash-4.0# dpkg -r dstat

With APT and dpkg, removing a package still leaves behind any configuration
files, in case you wish to reinstall the package again later. To fully delete
packages and their configuration files, you need to ``purge``:

.. code-block:: console

   bash-4.0# apt-get purge dstat

or:

.. code-block:: console

   bash-4.0# apt-get --purge remove dstat

or:

.. code-block:: console

   bash-4.0# dpkg -P dstat


Querying the dpkg database
--------------------------

Occasionally you will want to find out specific information regarding installed
packages. The ``dpkg-query`` command has many options to help. Let's
take a look at a few examples:

One common task is to see if you have a package installed. The ``-l`` option
by itself will list ALL installed packages. You can also ask it to list specific
packages if they are installed:

.. code-block:: console

   bash-4.0$ dpkg-query -l dstat
   Desired=Unknown/Install/Remove/Purge/Hold
   | Status=Not/Inst/Conf-files/Unpacked/halF-conf/Half-inst/trig-aWait/Trig-pend
   |/ Err?=(none)/Reinst-required (Status,Err: uppercase=bad)
   ||/ Name           Version      Architecture Description
   +++-==============-============-============-==================================
   ii  dstat          0.7.2-3      all          versatile resource statistics tool

Now let's say we want to list all of the files installed by a package. The
``-L`` option is the one to use:

.. code-block:: console

   bash-4.0$ dpkg-query -L dstat
   /.
   /usr
   /usr/bin
   /usr/bin/dstat
   /usr/share
   ...

We can also do the reverse of the previous operation. If we have a file, and
want to know to which package it belongs:

.. code-block:: console

   bash-4.0$ dpkg-query -S /usr/bin/dstat
   dstat: /usr/bin/dstat

