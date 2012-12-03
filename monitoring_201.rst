Monitoring, Notifications, and Metrics 201
******************************************

Dataviz & Graphing
==================

Graphite, StatsD
================

What are Graphite and StatsD?
------------------------------
In order to understand what your application is doing and to get a better
overview over key functionality it is indispensable to gather metrics about
those areas. A popular combination to achieve this is StatsD with the Graphite
storage and graphing backend. This chapter will introduce those tools, how to
install and configure them and how they can be used to gather insight into
your application's metrics.

Graphite
--------
The `Graphite project <http://graphite.wikidot.com/>`_ is a set of tools for
collecting and visualizing data, all written in Python. It consists of a file
based database format and tools to work with it called *whisper*, a set of
network daemons to collect data - the *carbon* daemons - and a Django based
web application to display data.

Whisper database format
~~~~~~~~~~~~~~~~~~~~~~~~
Graphite at its core uses a disk based storage system called *whisper*. Every
metric is stored in its corresponding ``.wsp`` file. These are fixed sized
databases, meaning that the amount of data a file can hold is pre-determined
upon creation. In order to achieve this, every database file contains so
called *archives* which store *(timestamp, value)* pairs with varying
granularity and length. For example a possible archive setup could be to store
a data point every 10 seconds for 6 hours, after that store 1 minute data for a
week and for long time archival use a data point every 10 minutes for 5 years.
This means once your data is older than the time period stored in the archive,
a decision has to be made, how to fit the data into the coarser slots in the
next archive. In our example setup when going from the highest precision (10
seconds) to the next archive (1 minute) we have 6 values at our disposal to
aggregate. The database can be configured to aggregate by average, sum, last,
max or min functions, which uses the arithmetic mean, the sum of all values,
the last one recorded or the biggest or smallest value respectively to fill the
value in the next archive. This process is done every time a value is too old to
be stored in an archive until it is older than the lowest resolution archive
available and isn't stored anymore at all. An important setting in this
context is the *x-files-factor* you can set on every database. This describes
how many of the archives slots have to have a value (as opposed to ``None``
which is Python's ``NULL`` value and recorded when no value is received for a
timestamp) to aggregate the next lower archive to a value and not also insert
``None`` instead. The value range for this setting is decimal from 0 to 1 and
defaults to 0.5. This means when we use the default setting on our example
archives, at least 3 of the 6 values for the first aggregation and 5 of the 10
values for the second one have to have actual values.

The disk format to store all this data is rather simple. Every file has a
short header with the basic information about the aggregation function used,
the maximum retention period available, the *x-files-factor* and the number of
archives it contains. These are stored as 2 *longs*, a *float* and another
*long* thus requiring 16 bytes of storage.  After that the archives are
appended to the file with their *(timestamp, value)* pairs stored as a *long*
and a *double* value consuming 12 bytes per pair.

This design makes it possible to store a lot of datapoints relatively easy and
without consuming a lot of disk space (1 year worth of 1 minute data for a
metric can be stored in a little more than 6MB). But it also means that some
considerations about usage have to be made upfront. Due to the database's
fixed size, the highest resolution archive should match the rate at which you
send data. If data is sent at lower intervals the archive ends up with a lot
of ``None`` values and if more than one metric are received for a timestamp,
the last one will always overwrite previous ones. So if a low rate of events
during some times is expected, it might make sense to tune down the
*x-files-factor* to aggregate even if only one of ten values is available. And
depending on the type of metrics (counters for example) it might also make
more sense to aggregate with the *sum* function instead of the default
aggregation by averaging the values. Thinking about this before sending
metrics makes things a lot easier since it's possible to change these settings
afterwards, however to keep the existing data, the whisper file has to be
resized which takes some time and makes the file unavailable for storage
during that time.

In order to understand whisper files a little bit better, we can use the set
of scripts distributed with whisper to take a look at a database file. First
install whisper from the PyPi distribution::

  % sudo pip install whisper

And create a whisper database with a 10 second retention for 1 minute by using
the ``whisper-create`` command::

  % whisper-create.py test.wsp 10s:1minute
  Created: test.wsp (100 bytes)

  % whisper-info.py test.wsp
  maxRetention: 60
  xFilesFactor: 0.5
  aggregationMethod: average
  fileSize: 100

  Archive 0
  retention: 60
  secondsPerPoint: 10
  points: 6
  size: 72
  offset: 28

  % whisper-fetch.py test.wsp
  1354509290      None
  1354509300      None
  1354509310      None
  1354509320      None
  1354509330      None
  1354509340      None

The resulting file has six ten second buckets corresponding to the retention
period in the create command. As it is visible from the ``whisper-info.py``
output, the database uses default values for *x-files-factor* and aggregation
method, since we didn't specify anything different. And we also only have one
archive which stores data at 10 seconds per value for 60 seconds as we passed
as a command argument. Per default ``whisper-fetch.py`` shows timestamps as
epoch time. There is also ``--pretty`` option to show them in a more human
readable format, but since the exact time is not important all examples show
epoch time. For updating the database with value there is the handy
``whisper-update.py`` command, which takes a timestamp and a value as
arguments::

  % whisper-update.py test.wsp 1354509710:3
  % whisper-fetch.py test.wsp
  1354509690      None
  1354509700      None
  1354509710      3.000000
  1354509720      None
  1354509730      None
  1354509740      None

Notice how the timestamps are not the same as in the example above, because
more than a minute has past since then and if we had values stored at those
points, they wouldn't be show anymore. However taking a look at the
database file with ``whisper-dump.py`` reveals a little more information about
the storage system::

  % whisper-dump.py test.wsp
  Meta data:
  aggregation method: average
  max retention: 60
  xFilesFactor: 0.5

  Archive 0 info:
  offset: 28
  seconds per point: 10
  points: 6
  retention: 60
  size: 72

  Archive 0 data:
  0: 1354509710,          3
  1: 0,          0
  2: 0,          0
  3: 0,          0
  4: 0,          0
  5: 0,          0

In addition to the meta data ``whisper-info.py`` already showed, the dump
command also tells us, that only one slot actually has data. And in this case
the time passed doesn't matter. Since slots are only changed when new values
need to be written to them, this old value will remain there until then. The
reason why ``whisper-fetch.py`` doesn't show these past values is because it
will only show valid data within a given time (default 24h) until *max
retention* from the invoked point in time. And it will also fetch the points
from the retention archive that can cover most of the requested time. This
becomes a bit more clear when adding a new archive::

  % whisper-resize.py test.wsp 10s:1min 20s:2min
  Retrieving all data from the archives
  Creating new whisper database: test.wsp.tmp
  Created: test.wsp.tmp (184 bytes)
  Migrating data...
  Renaming old database to: test.wsp.bak
  Renaming new database to: test.wsp

  % whisper-info.py test.wsp
  maxRetention: 120
  xFilesFactor: 0.5
  aggregationMethod: average
  fileSize: 184

  Archive 0
  retention: 60
  secondsPerPoint: 10
  points: 6
  size: 72
  offset: 40

  Archive 1
  retention: 120
  secondsPerPoint: 20
  points: 6
  size: 72
  offset: 112

  % whisper-fetch.py test.wsp
  1354514740      None
  1354514760      None
  1354514780      None
  1354514800      None
  1354514820      None
  1354514840      None

Now the database has a second archive which stores 20 second data for 2
minutes and ``whisper-fetch.py`` returns 20 second slots. That's because it
tries to retrieve as close to 24h (the default time) as possible and the 20
second slot archive is closer to that. For getting data in 10 second slots,
the command has to be invoked with the ``--from=`` parameter and an epoch
timestamp less than 1 minute in the past.

These commands are a good way to inspect whisper files and to get a basic
understanding how data is stored. So it makes sense to experiment with them a
bit before going into the rest of the Graphite eco-system.

The carbon daemons
~~~~~~~~~~~~~~~~~~



StatsD
-------


Setting it up and make it show pretty graphs
---------------------------------------------

What have we done and where to go from here
--------------------------------------------


Dashboard: Info for ops and info for the business
=================================================

Third-party tools
=================

Datadog
-------

Boundry
-------

NewRelic
--------

Circonus
--------


