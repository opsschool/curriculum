Booting over the network
************************

Typically, a machine will boot off of some type of local device.
However, there are some cases where it makes more sense to boot a machine off a remote host.
For example:

* You have a large number of machines to install, and carrying a CD/USB drive to each of them is time consuming.
* You have a machine where the hard drive is failing, and you need to run some diagnostic tools
* You have a machine in a remote location, but you need to reinstall the operating system

In situations like this, it's possible to boot a machine entirely over the network.
This relies on a little piece of firmware built into pretty much every modern NIC.
This software is called Preboot eXecution Environment (PXE).

The network boot process works like this:

* The machine boots off the NIC, which has PXE firmware built in
* The PXE firmware initializes the NIC and sends out a DHCP request
* The DHCP response contains instructions to contact another server and download a file
* The PXE firmware downloads the file via TFTP, and begins executing it
* From here, many different things can happen depending on the type of software downloaded

TFTP
====
Trivial File Transfer Protocol (TFTP) is a very basic file transfer protocol that's used for bootstrapping the PXE firmware into something more useful.
TFTP is implemented as a UDP service, typically running on port 69.
TFTP does not support any type of authentication.

Given the low complexity of the TFTP protocol, servers are available for virtually every operating system.
Configuration is typically just pointing the server at a directory full of files to serve up.
Unless you have very specific reasons to do so (and fully understand the implications), you should make sure that your TFTP server is setup in read only mode.

PXE
===
The PXE stack by itself is very limited.
It's designed to just be able to retrieve and execute a file, and by itself is not terribly useful.
The two most popular software packages used with PXE are iPXE [#]_ and PXELINUX [#]_.
Of these, PXELINUX is the older one, and is a variant of SysLinux.
SysLinux is most commonly used as the initial menu you see when you boot a Linux CD.
iPXE is newer, and supports booting over many different protocols (HTTP, iSCSI, various SAN types, among others).

Basic Setup Process
===================
For both iPXE and SysLinux you will need both a DHCP and TFTP server installed.  Your operating system will likely have a tftp server package available.
Don't worry too much about the exact one you use, there are only minor differences between them.
After this is installed, you will need to locate the directory that it is serving files from.
For the rest of this document we will refer to this directory as the 'tftpboot' directory.

For the DHCP server, we'll be using the ISC DHCPD server.
This is available for many platforms, and has a large amount of documentation available.
Aside from setting up a DHCP range, you'd need to add the following line to your dhcpd.conf file:

.. code-block:: none

    # Any clients that are PXE booting should attempt to retrieve files from this server
    next-server <tftp server IP>;

That is the bare minimum needed for both types of PXE software we're going to set up.

iPXE Setup
==========
Start by downloading iPXE [#]_.
Make sure you save this to your tftpboot directory.
Next, add the following to your dhcpd.conf file:

.. code-block:: none

    if exists user-class and option user-class = "iPXE" {
        filename "chainconfig.ipxe";
    } else {
        filename "undionly.kpxe";
    }

This will cause the DHCP server to first tell clients to download iPXE.
Once iPXE starts up and does another DHCP request, it will be told the actual location of the configuration file to download.
Without this, we would end up with a continuous loop of iPXE downloading itself.
Create a chainconfig.ipxe file in your tftpboot directory with the following:

.. code-block:: none

    #!ipxe

    :start
    # Define a generic title for the menu
    menu iPXE boot menu
    # And now we define some options for our menu
    item localboot      Boot from local disk
    item centos6_x64    Install CentOS 6 x64
    item centos6_i386   Install CentOS 6 i386

    # Now we prompt the user to choose what they want to do.  If they don't select an option
    # within 60s we default to booting from the local disk.  If they somehow choose an invalid
    # option, we drop them to the iPXE shell for further debugging
    choose --default localboot 60000 bootoption && goto ${bootoption} ||
    echo Invalid option selected!
    shell

    # Here we define our two network installation options.  iPXE supports basic variable operations,
    # so we can reuse much the same code for booting the two different architectures.  We'll define
    # some options specifying the version and architecture we want, then jump to some common installer
    # code
    :centos6_x64
    set centos-version 6
    set arch x86_64
    goto centos_installer

    :centos6_i386
    set centos-version 6
    set arch i386
    goto centos_installer

    # This demonstrates some of the power of iPXE.  We make use of variables to prevent config duplication
    # and we load the installer files directly off the CentOS mirror.  There's no need to copy everything
    # to a local TFTP server.  We also fallback to a shell if the boot fails so any issues can be debugged
    :centos_installer
    kernel http://mirror.centos.org/centos-${centos-version}/${centos-version}/os/${arch}/images/pxeboot/vmlinuz ramdisk_size=65535 noipv6 network
    initrd http://mirror.centos.org/centos-${centos-version}/${centos-version}/os/${arch}/images/pxeboot/initrd.img
    boot ||
    goto shell

    # This just exits iPXE entirely, and allows the rest of the boot process to proceed
    :localboot
    exit


PXELINUX setup
==============

Start by downloading SysLinux [#]_.
Copy a few files from the archive into your tftpboot directory:

* com32/menu/vesamenu.c32
* core/pxelinux.0

Next, we'll need to create the menu config file.
Create the file tftpboot/pxelinux.cfg/default:

.. code-block:: none

    # We want to load the vesamenu module, which generates GUI menus
    UI vesamenu.c32
    # Don't display a prompt for the user to type in a boot option (they'll be selecting one instead)
    PROMPT 0
    # Our default option is to boot from the local drive
    DEFAULT localboot
    # Wait 60s before booting to the default option
    TIMEOUT 600

    # Define a title for our menu
    MENU TITLE SysLinux Boot Menu

    # This is the internal name for this option.
    LABEL centos6_x64
        # And a human readable description
        MENU LABEL Install CentOS 6 x64
        # This is the kernel file to download (via TFTP) and boot
        KERNEL centos/6/x64/vmlinuz
        # And any command line options to pass to the kernel
        APPEND initrd=centos/6/x64/initrd.img ramdisk_size=65535 noipv6 network

    # Now for the i386 version.  As SysLinux doesn't support variables, we end up duplicating
    # the majority of the config from the x64 version
    LABEL centos6_i386
        MENU LABEL Install CentOS 6 i386
        KERNEL centos/6/i386/vmlinuz
        APPEND initrd=centos/6/i386/initrd.img ramdisk_size=65535 noipv6 network

    LABEL localboot
        # Proceed through the rest of the normal boot process
        LOCALBOOT 0


Since PXELINUX doesn't support HTTP, we'll need to download the CentOS installer images to the tftpboot directory.
Create two directories and download the initrd.img and vmlinuz files to them:

* Directory: tftpboot/centos/6/x64/ Files: http://mirror.centos.org/centos-/6/os/x86_64/images/pxeboot/
* Directory: tftpboot/centos/6/i386/ Files: http://mirror.centos.org/centos-6/6/os/i386/images/pxeboot/





References
----------
.. [#] http://ipxe.org
.. [#] http://www.syslinux.org/wiki/index.php/PXELINUX
.. [#] http://boot.ipxe.org/undionly.kpxe
.. [#] https://www.kernel.org/pub/linux/utils/boot/syslinux/
