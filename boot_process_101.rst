The Boot Process
****************

The boot process of a modern system involves multiple phases.
Here we will discuss each step, how it contributes to the boot process, what can
go wrong and things you can do to diagnose problems during booting.

You should read this document first, and then power on a computer.
Note each of the phases described in this section. Some of them will last many
seconds, other will be by very quickly!

The following components are involved in the boot process. They are each
executed in this order:

.. contents::
   :depth: 2
   :local:


Power Supply Unit
=================

When the power button is pressed, an electric circuit is closed which causes the
power supply unit to perform a self test. In order for the boot process to
continue, this self test has to complete successfully. If the power supply
cannot confirm the self test, there will usually be no output at all.

Most modern x86 computers, especially those using the `ATX
<http://en.wikipedia.org/wiki/ATX>`_ standard will have two main connectors to
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

Once the CPU has powered in, the first call made is to the BIOS (Basic Input /
Output System). The BIOS contains the code required to get the system started
and to begin loading the operating system.

The first step taken by the BIOS is to ensure that the minimum required hardware
exists:

* CPU
* Memory
* Video card

Once the existance of the hardware has been confirmed, it must be configured.

The BIOS has its own memory storage known as the CMOS (Complimentary Metal Oxide
Semiconductor). The CMOS contains all of the settings the BIOS needs to save
about a system. Amongst others, these include the memory speed and the CPU
frequency multipler and the location and configuration of the hard drives and
other devices.

The BIOS first takes the memory frequency and attempts to set that on the
memory controller.

Next the BIOS multiplies the memory frequency by the CPU frequency multiplier.
This is the speed at which the CPU is set to run. Sometimes it is possible to
"overclock" a CPU, but telling it to run at a higher multiplier than it was
designed to, effectively making it run faster. There can be benefits and risks
to doing this, including the potential for damaging your CPU.


POST tests
==========

Once the memory and CPU frequencies have been set, the BIOS begins the Power On
Self Test (POST). The POST will perform basic checks on many system components,
including:

* Check that the memory is working
* Check that hard drives and other devices are all responding
* Check that the keyboard and mouse are connected (this check usually be
  disabled)
* Initialise any additional BIOSes which may be installed (eg, RAID cards)

Possible failures:

In the event that a POST test fails, the BIOS will normally indicate failure
through a series of beeps on the internal computer speaker. The pattern of the
beeps incidates which specific test failed. A few beep codes are common across
systems:

* One beep: All tests passed successfully (Have you noticed that your computer
  beeps once when you press the power button? This is why!)
* Three beeps: Often a memory error
* One long, two short beeps: Video card or display problem

Your BIOS manual should document what its specific beep codes mean.


Master Boot Record (the old way)
================================

Once the BIOS has identified which drive it should attempt to boot from, it
looks at the first sector on that drive. These sectors should contain the
Master Boot Record.

The MBR should contain at least the following information:

* The partition table
* Instructions on where to find code, to continue booting

The MBR has been heavily limited in its design, as it can only occupy the first
512 bytes of space on the drive (which is the size of one physical sector).
This limits how much the MBR is able to store. As the complexity of systems
grew, it became necssary to add "chain boot loading". This allows the MBR to
load an another program from elsewhere on the drive into memory. The new program
is then executed and continues the boot process.


GUID Partition Table (the new way)
==================================


Boot Loader
===========


OS Kernel
=========


Init
====


Run levels and Single User Mode
===============================
