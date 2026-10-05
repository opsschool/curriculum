Looking at system metrics
*************************

``vmstat``
==========

``vmstat`` reports processes, memory, swap, disk I/O and CPU use for the whole system.
Give it an interval in seconds, and optionally a count.
The first line shows averages since the system started; the following lines each cover one interval:

.. code-block:: console

  root@opsschool # vmstat 1 3
  procs -----------memory---------- ---swap-- -----io---- -system-- -------cpu-------
   r  b   swpd   free   buff  cache   si   so    bi    bo   in   cs us sy id wa st gu
   1  0 884788 136088 414824 9036592    7   14  4785  8536 5715    3  0  1 98  0  0  1
   0  0 884776 131164 415384 9039184    4    0  2100  3832 18179 29939  1  2 97  1  0  0
   1  0 884776 127652 416088 9041368    0    0  1864  4088 19106 31591  1  2 96  1  0  0

The columns that are most often useful:

- ``r``: processes that are running or waiting for a CPU.
  If this is often higher than the number of CPUs, processes are waiting for CPU time.
- ``b``: processes that are blocked waiting for I/O to complete.
- ``swpd``: swap space in use, in KiB.
- ``free``, ``buff`` and ``cache``: memory that is unused, and memory used for buffers and the page cache, in KiB.
  The kernel uses spare memory for the cache and gives it back when programs need it, so low ``free`` memory on its own is normal.
- ``si`` and ``so``: memory swapped in from disk and swapped out to disk, per second.
- ``bi`` and ``bo``: blocks read from and written to block devices, per second.
- ``us``, ``sy``, ``id`` and ``wa``: the percentage of CPU time spent in user programs, in the kernel, idle, and idle while waiting for I/O.
  ``st`` is time stolen by the hypervisor, on a virtual machine.

Swap in use (``swpd``) is not a problem on its own: the kernel can move memory that hasn't been used for a while to swap, and leave it there.
Sustained non-zero ``si`` and ``so`` mean that programs need more memory than the system has, and the kernel is moving pages between memory and disk while they run.
Every access to a page that is on disk waits for a disk read, so everything on the system can become slow, while the CPU looks mostly idle or waiting (``wa``).
To find which processes use the memory, sort ``top`` by resident memory (press ``M``) or run ``ps aux --sort=-rss | head``.

``free -m`` shows memory and swap use in MiB.
Its ``available`` column estimates how much memory programs can still use without swapping, including cache that the kernel can free.

``iostat``
==========

``systat``
==========

``dstat``
=========

``sar``
=======
