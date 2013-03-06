Booting over the network
************************

Typically, a machine will boot off of some type of local device.  However, there are some cases where it makes more sense to boot a machine off a remote host.  For example:

* You have a large number of machines to install, and carrying a CD/USB drive to each of them is time consuming.
* You have a machine where the hard drive is failing, and you need to run some diagnostic tools
* You have a machine in a remote location, but you need to reinstall the operating system

In situations like this, it's possible to boot a machine entirely over the network.  This relies on a little piece of software built into pretty much every modern NIC.  This software is called Preboot eXecution Envionment (PXE).

The network boot process works like this:

* Machine boots off the NIC, which has PXE firmware built in
* The PXE firmware intializes the NIC and sends out a DHCP request
* The DHCP response contains instructions to contact another server and download a file, as well as the normal network information
* The PXE firmware downloads the file via TFTP, and begins executing it
* From here, many different things can happen depending on what file was actually downloaded


TFTP
====
TFTP is a very basic file transfer protocol.  The PXE roms in most network cards do not support anything more complicated then TFTP.  This is okay, as the built in PXE stack is typically only used to retrieve another, more fully featured program.

As the server itself is pretty simple, virtually every operating system has some type of TFTP server available.  These are typically quite easy to setup, as there's very little complexity in TFTP.  You usually just point the server at a directory, and start it.  TFTP has no support for authentication, or any other complex features.  You are advised to make sure your TFTP server only allows read access to the files.  This is setup by default for most TFTP servers.

PXE
===
The PXE stack by itself is very limited.  It's designed to just be able to retrieve and execute a file, and by itself is not terribly useful.  The two most popular software packages used with PXE are iPXE [#]_ and PXELINUX [#]_.  Of these, PXELINUX is the older one, and is a varient of SysLinux.  SysLinux is most commonly used as the initial menu you see when you boot a Linux CD.  iPXE is newer, and supports booting over many different protocols (HTTP, )

bootp
=====


References
----------
.. [#] http://ipxe.org
.. [#] http://www.syslinux.org/wiki/index.php/PXELINUX
