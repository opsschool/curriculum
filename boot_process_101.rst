The boot process
****************

The boot processes involves many components:

* List them all - BIOS, POST tests, reading the MBR/GUID, Loading the boot
  loader (grub), the kernel, init, init scripts, getty

The BIOS
========

How an operating system is started
==================================
MBR / GUID partition tables
Maybe a brief mention of ``INT 13H``? A side note about how this used to be a
popular hook for boot loader viruses might be interesting for readers :)

GRUB bootloader
===============

/bin/init
=========

Run levels and Single User Mode
===============================
