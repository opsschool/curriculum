The Boot Process
****************

The boot process of a modern system involves multiple phases.
Here we will discuss each step, how it contributes to the boot process, what can
go wrong and things you can do to diagnose problems during booting.

Learning the details of the boot process will give you a strong understanding of
how to troubleshoot issues that occur during boot - either at the hardware level
or at the operating system level.

You should read this document first, and then power on a computer.
Note each of the phases described in this section. Some of them will last many
seconds, others will fly by very quickly!

The following components are involved in the boot process. They are each
executed in this order:

.. contents::
   :depth: 1
   :local:


Power Supply Unit
=================

When the power button is pressed, an electric circuit is closed which causes the
power supply unit to perform a self test. In order for the boot process to
continue, this self test has to complete successfully. If the power supply
cannot confirm the self test, there will usually be no output at all.

Most modern x86 computers, especially those using the `ATX
<http://en.wikipedia.org/wiki/ATX>`_ standard, will have two main connectors to
the motherboard: a 4-pin connector to power the CPU, and a 24-pin connector to
power other motherboard components. If the self test passes successfully, the
PSU will send a signal to the CPU on the 4-pin connector to indicate that it
should power on.

Possible failures:

* If the PSU is unable to perform a good self test, the PSU may be damaged. This
  could be caused by a blown fuse, or other damage caused by over-/under-current
  on the power line. Using a UPS or a good surge protector is always
  recommended.


BIOS and CMOS
=============

At its core, the Basic Input/Output System (:term:`BIOS`) is an integrated circuit
located on the computer's motherboard that can be programmed with firmware.
This firmware facilitates the boot process so that an operating system
can load.

Let's examine each of these in more detail:

* *Firmware* is the software that is programmed into Electrically Erasable
  Programmable Read-Only Memory (EEPROM). In this case, the firmware facilitates
  booting an operating system and configuring basic hardware settings.

* An *integrated circuit* (IC) is what you would likely think of as a
  stereotypical "computer chip" - a thin wafer that is packaged and has metal
  traces sticking out from it that can be mounted onto a printed circuit board.

Your BIOS is the lowest level interface you'll get to the hardware in your
computer. The BIOS also performs the Power-On Self Test, or POST.

Once the CPU has powered up, the first call made is to the BIOS.
The first step then taken by the BIOS is to ensure that the minimum required
hardware exists:

* CPU
* Memory
* Video card

Once the existence of the hardware has been confirmed, it must be configured.

The BIOS has its own memory storage known as the CMOS (Complementary
Metal Oxide Semiconductor). The CMOS contains all of the settings the
BIOS needs to save, such as the memory speed, CPU frequency
multiplier, and the location and configuration of the hard drives and
other devices.

The BIOS first takes the memory frequency and attempts to set that on the memory
controller.

Next the BIOS multiplies the memory frequency by the CPU frequency multiplier.
This is the speed at which the CPU is set to run. Sometimes it is possible to
"overclock" a CPU, by telling it to run at a higher multiplier than it was
designed to, effectively making it run faster. There can be benefits and risks
to doing this, including the potential for damaging your CPU.


POST tests
==========

Once the memory and CPU frequencies have been set, the BIOS begins the Power-On
Self Test (POST). The POST will perform basic checks on many system components,
including:

* Check that the memory is working
* Check that hard drives and other devices are all responding
* Check that the keyboard and mouse are connected (this check can usually be
  disabled)
* Initialise any additional BIOSes which may be installed (e.g. RAID cards)

Possible failures:

In the event that a POST test fails, the BIOS will normally indicate failure
through a series of beeps on the internal computer speaker. The pattern of the
beeps indicates which specific test failed. A few beep codes are common across
systems:

* One beep: All tests passed successfully (Have you noticed that your computer
  beeps once when you press the power button? This is why!)
* Three beeps: Often a memory error
* One long, two short beeps: Video card or display problem

Your BIOS manual should document what its specific beep codes mean.

Reading the Partition Table
===========================

The next major function of the BIOS is to determine which device to use to
start an operating system.

A typical BIOS can read boot information from the devices below, and will
boot from the first device that provides a successful response. The order of
devices to scan can be set in the BIOS:

* Floppy disks
* CD-ROMs
* USB flash drives
* Hard drives
* A network

We'll cover the first four options here. There's another section that
deals with booting over the network.

There are two separate partition table formats: Master Boot Record (MBR) and
the GUID Partition Table (GPT). We'll illustrate how both store data about
what's on the drive, and how they're used to boot the operating system.

Master Boot Record (the old way)
--------------------------------

Once the BIOS has identified which drive it should attempt to boot from, it
looks at the first sector on that drive. These sectors should contain the Master
Boot Record.

The MBR has two component parts:

* The boot loader information block (448 bytes)
* The partition table (64 bytes)

The boot loader information block is where the first program the computer can
run is stored. The partition table stores information about how the drive is
logically laid out.

The MBR has been heavily limited in its design, as it can only occupy the first
512 bytes of space on the drive (which is the size of one physical sector).
This limits the tasks the boot loader program is able to do. The execution of 
the boot loader literally starts from the first byte. As the complexity of 
systems grew, it became necessary to add "chain boot loading". This allows the
MBR to load an another program from elsewhere on the drive into memory. The new
program is then executed and continues the boot process.

If you're familiar with Windows, you may have seen drives labelled as "C:" and
"D:" - these represent different logical "partitions" on the drive. These
represent partitions defined in that 64-byte partition table.


GPT - The GUID Partition Table (the new way)
--------------------------------------------

The design of the IBM-Compatible BIOS is an old design and has
limitations in today's world of hardware. To address this, the United
Extensible Firmware Interface (UEFI) was created, along with GPT, a
new partition format.

There are a few advantages to the GPT format, specifically:

* A Globally-Unique ID that references a partition, rather than a partition
  number. The MBR only has 64 bytes to store partition information - and each
  partition definition is 16 bytes. This design allows for unlimited partitions.

* The ability to boot from storage devices that are greater than 2 TBs, due to
  a larger address space to identify sectors on the disk. The MBR simply had no
  way to address disks greater than 2 TB.

* A backup copy of the table that can be used in the event that the primary
  copy is corrupted. This copy is stored at the 'end' of the disk.

There is some compatibility maintained to allow standard PCs that are using
an old BIOS to boot from a drive that has a GPT on it.

The Bootloader
==============

The purpose of a bootloader is to load the initial kernel and supporting modules
into memory.

There are a few common bootloaders. We'll discuss the GRand Unified
Bootloader (GRUB), a bootloader used by many Linux distributions today.

GRUB is a "chain bootloader," meaning it initializes itself in stages. These stages are:

* *Stage 1* - This is the very tiny application that can exist in that first
  part of the drive. It exists to load the next, larger part of GRUB.

* *Stage 1.5* - Contains the drivers necessary to access the filesystem with
  stage 2.

* *Stage 2* - This stage loads the menu and configuration options for GRUB.

On an MBR-formatted drive and standard BIOS
-------------------------------------------

These stages must fit in that first 448 bytes of the boot loader information 
block table. Generally, Stage 1 and Stage 1.5 are small enough to exist in that 
first 448 bytes. They contain the appropriate logic that allow the loader to 
read the filesystem that Stage 2 is located on.

On a GPT-formatted drive and UEFI
---------------------------------

UEFI motherboards are able to read FAT32 filesystems and execute code.
The system firmware looks for an image file that contains the boot code
for Stages 1 and 1.5, so that Stage 2 can be managed by the operating
system.

The Kernel and the Ramdisk
==========================

The kernel is the main component of any operating system. The kernel
acts as the lowest-level intermediary between the hardware on your
computer and the applications running on your computer. The kernel
abstracts away such resource management tasks as memory and
processor allocation.

The kernel and other software can access peripherals such as disk
drives by way of device drivers.

So what, then, is this Initial RAM Filesystem, or Ramdisk?

You can imagine there are tens of thousands of different devices in the world.
Hard drives made with different connectors, video cards made by different
manufacturers, network cards with special interfaces. Each of these needs its
own device driver to bridge the hardware and software.

For our small and efficient little boot process, trying to keep every possible
device driver in the kernel wouldn't work very well.

This lead to the creation of the Initial RAM disk as a way to provide module
support to the kernel for the boot process. It allows the kernel to load just
enough drivers to read from the filesystem, where it can find other specific
device drivers as needed.

With the kernel and ramdisk loaded into memory, we can attempt to access the
disk drive and continue booting our Linux system.


OS Kernel and Init
==================

The organizational scheme for determining the load order for system
services during the boot process is referred to as an init system.
The traditional and still most common init system in Linux is called
"System V init".

After the initial ramdisk sets the stage for the kernel to access the hard
drive, we now need to execute the first process that will essentially
"rule them all" - ``/bin/init``.

The init process reads ``/etc/inittab`` to figure out what script should be run to
initialize the system. This is a collection of scripts that vary based on the
desired "runlevel" of the system.


Runlevels
==========

Various runlevels have been defined to bring the system up in different
states. In general, the following runlevels are consistent in most Linux
distributions:

* 0: Halt the system
* 1: Single User Mode
* 6: Reboot the machine

Across distributions there can be various meanings for runlevels 2-5.
RedHat-based distributions use runlevel 3 for a multiuser console
environment and 5 for a graphical-based environment.

Multiuser vs. Single user runlevels
-----------------------------------

As the name implies, in some runlevels multiple users can use the machine, versus
one user in single user mode. So why does single user mode exist, anyways?

In multiuser runlevels, the system boots as normal. All standard
services such as SSH and HTTP daemons load in the order defined in the
init system. The network interfaces, if configured, are enabled. It's
business as usual if you're booting to a multiuser runlevel.

Conversely, single user mode has the bare minimum of services enabled
(notably there is no networking enabled), making it useful for
troubleshooting (and not much else).

You will need (or involuntarily find yourself in) single user mode
when something breaks: something you configured interferes with the
boot process and you need to turn it off, or perhaps a key filesystem
is corrupt and you need to run a disk check.

In single user mode, the only available access is via the console,
although that need not be limited to physical presence. Remote console
access by way of serial consoles and similar devices is a common
management tool for data centers.

Getty
=====

.. todo:: Check this section. I think i've got it down, but I'm not super
         familiar with this part.

After all the system initialization scripts have run, we're ready to present the
user with a prompt to login. The method of doing this is to provide a login prompt
on a "TTY" which is short for teletype. This is a holdover from the days that a
user running a Unix-based operating system sat at a serially-connected teletype
machine. A TTY can be a physical or virtual serial console, such as the
various terminals you'd be presented with if you used ALT+F# on the console of a
Linux machine.

Getty is often used to continuously spawn ``/bin/login``, which reads
the username and password of the user and, if authentication succeeds,
spawn the user's preferred shell. At this point, the boot and login
process has completed.
