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
The PXE stack by itself is very limited.  It's designed to just be able to retrieve and execute a file, and by itself is not terribly useful.  The two most popular software packages used with PXE are iPXE [#]_ and PXELINUX [#]_.  Of these, PXELINUX is the older one, and is a varient of SysLinux.  SysLinux is most commonly used as the initial menu you see when you boot a Linux CD.  iPXE is newer, and supports booting over many different protocols (HTTP, iSCSI, various SAN types).

Basic Setup Process
===
For both iPXE and SysLinux you will need both a DHCP and TFTP server installed.  For TFTP, just install whatever package your operating system makes available.  All the server software is similar enough that it does not matter which varient you install.  The only important bit is to locate where the 'tftpboot' directory is.  This will be where you copy all the PXE software.

For the DHCP server, ISC DHCPD would be suggested.  There are other DHCP servers available, but this one has the largest amount of documentation.  Aside from setting up a DHCP range, you'd need to add the following line to your dhcpd.conf file:

.. code-block:: 

    next-server <tftp server IP>;

That is the bare minimum needed for both types of PXE software we're going to set up.

iPXE Setup
====
Start by downloading [#]_ iPXE.  Make sure you save this to your tftpboot directory.  Next, add the following to your dhcpd.conf file:

.. code-block:: 

    if exists user-class and option user-class = "iPXE" {
        filename "chainconfig.ipxe";
    } else {
        filename "undionly.kpxe";
    }

What this will do is serve up the iPXE file to any non-iPXE clients.  This would mean that whenever a machine boots off it's internal NIC, it would download a copy of iPXE.  Once iPXE finishes it's startup process, it will do another DHCP request and be told to download the 'chainconfig.ipxe' file.  This file will determine where iPXE goes from there.  A good place to start with this would be a basic menu.  Create a chainconfig.ipxe file in your tftpboot directory with the following:

.. code-block:: 

    #!ipxe

    :start
    # Define a generic title for the menu
    menu iPXE boot menu
    # And now we define some options for our menu
    item localboot    	Boot from local disk
    item centos6_x64	Install CentOS 6 x64
    item centos6_i386	Install CentOS 6 i386

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

    # This demostrates some of the power of iPXE.  We make use of variables to prevent config duplication
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
====

Start by downloading [#]_ SysLinux.  Copy a few files from the archive into your tftpboot directory:

* com32/menu/vesamenu.c32
* core/pxelinux.0

Next, we'll need to create the menu config file.  Create the file tftpboot/pxelinux.cfg/default:

.. code-block:: 

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


Since PXELINUX doesn't support HTTP, we'll need to download the CentOS installer images to the tftpboot directory.  Create two directories and download the initrd.img and vmlinuz files to them:

* Directory: tftpboot/centos/6/x64/ Files: http://mirror.centos.org/centos-/6/os/x86_64/images/pxeboot/
* Directory: tftpboot/centos/6/i386/ Files: http://mirror.centos.org/centos-6/6/os/i386/images/pxeboot/





References
----------
.. [#] http://ipxe.org
.. [#] http://www.syslinux.org/wiki/index.php/PXELINUX
.. [#] http://boot.ipxe.org/undionly.kpxe
.. [#] https://www.kernel.org/pub/linux/utils/boot/syslinux/
