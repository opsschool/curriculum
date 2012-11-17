The boot process
****************

A computer exists, in our case, to load an operating system for us to make
use of. In this section, you'll learn the basics of how a standard "PC" 
boots. 

This section may seem old school: after all, modern operating systems have
dealt with making the PC boot in the first place. We encourage you to still
know the basics of this section, though, for a few reasons:

1. Should you find yourself needing to troubleshoot a botched install, 
   knowing where to look will save you time, and make you look like you 
   know what you're doing. (This is likely: hardware fails!)

2. If you ever want to move into more advanced Linux-based development,
   this knowledge will serve you well.

The BIOS
========

At its core, the Basic Input/Output System (BIOS) is a integrated circuit
located on the computer's motherboard that can be programmed with firmware.
This firmware is what facilitates the boot process so that an operating
system can load.

That's a high-level definition, though. Let's go through it in a bit more 
depth.

* Generally, *Firmware* is software that is programmed into Electrically
  Erasable Programmable Read-Only Memory (EEPROM). In this case, the
  firmware faciliates booting an operating system and configuring basic
  hardware settings.

* An *integrated circuit* (IC) is what you would likely think of as a
  stereotypical "computer chip" - a thin wafer that is packaged and has 
  metal traces sticking out from it that can be mounted onto a printed
  circuit board.

Basically, your BIOS is the lowest level interface you'll get to the 
hardware in your computer. The BIOS also performs the Power-On Self Test, 
or POST.

How an operating system is started
==================================

A BIOS can read boot information from a variety of sources: 

* Floppy Disks (even still!)
* CD-ROMs
* USB Flash Drives
* Hard Drives
* A Network

We'll be discussing the first four options. There's another section that
deals with booting over the network.

Once your BIOS has been told (either though setup or through manual 
intervention) which device you'd like to attempt to boot from, the BIOS
attempts to read the master boot record of the device you've selected. This
master boot record (MBR) exists on the very first 512 bytes of the drive.

The MBR has two component parts: the boot loader information block (448 
bytes) and the partition table (64 bytes). The boot loader information 
block is where the first program the computer can run is stored. The 
partition table stores information about how the drive is logically laid
out. 

If you're familliar with windows, perhaps you've seen computers with
"C:" drives and "D:" drives, etc. - these represent different logical 
"partitions" on the drive. These represent partitions defined in that 64-
byte partition table.

The Bootloader
==============
Realistically, any application that can fit in the first 448-bytes of the
Master Boot Record can be a bootloader. The purpose of a bootloader is to
load the initial kernel and supporting modules into memory.

There are a few bootloaders that exist. We'll discuss the GRand Unified
Bootloader (GRUB), a bootloader used by many Linux distributions
today. 

GRUB initializes itself in "stages." These stages are:

* *Stage 1* - This is the very tiny application that can exist in that
  first part of the drive. It exists to load the next, larger part of 
  GRUB. This is the first 208 bytes of the drive.

* *Stage 1.5* - This exists in the next 240 bytes of the MBR. It contains
  the drivers necessary to access the filesystem that stage 2 resides on

* *Stage 2* - This actually loads the menu and configuration options for
  GRUB.

The bootloader loads the kernel and the initram image into memory. We'll
talk about that next.

The Kernel and the Ramdisk
==========================
The kernel is the main component of any operating system. The kernel acts
as the lowest-level intermediary between the hardware on your computer and
the applications running on your computer. The kernel is responsible for 
many other things - memory management, how the processor's time will be used
among some of them.

A kernel, therefore, is what various device drivers interact with to allow
an operating system to effectively "talk" to hardware installed on the 
computer.

So, then, what's an Initial RAM Filesystem, or Ramdisk?

You can imagine that there are tens of thousands of different
devices in the world. Would it be practical for Linux to include drivers
for every storage device under the sun, so that once loaded into memory
it could access a particular hard disk?

It'd be big -- for a boot process file.

So the concept of an Initial RAM disk was created. The Iniital
RAM disk was created to provide a little bit of module support to the kernel
for the boot process. With the kernel and ramdisk loaded into memory, we can
attempt to access the disk drive and continue booting our Linux system.

/bin/init
=========
While there are replacements for the standard "System V init" that has 
existed for years, we'll start with one of the more common ones.

After the initial ramdisk sets the stage for the kernel to access the hard
drive, we now need to execute the first process that will essentially 
"rule them all" - init.

"init" references /etc/inittab to figure out what script should be run to
initialize the system. This is a collection of scripts that vary based on
the desired "runlevel" of the system.

Run levels 
==========
Various run levels have been defined to bring the system up in different
states. In general, the following run levels are consistent in the majority 
of Linux distributions:

* 0: Halt the system
* 1: Single User Mode
* 6: Reboot the machine

Between distributions there can be various meanings for runlevels 2-5.
RedHat-based distributions use runlevel 3 for a multi-user console
environment and 5 for a graphical-based environment. 

.. todo: Finish section. Getty -> Login
