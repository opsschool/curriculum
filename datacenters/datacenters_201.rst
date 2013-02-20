Datacenters 201
***************

Networking many racks
=====================

Power
=====

N+1, N+2 power
--------------

Fused vs usable
---------------

calculations
------------

Cooling
=======

N+1, N+2
---------

The cooling exhausts for redundant or load-sharing aircos should be
located in enough distance from each other that they do not influence each
others effectiveness.

Cooling efficiency
------------------

Surrounding temperature and load both influence how good an airco can work.
For example it will work best if it's not fully loaded (has to move as many
units of thermal heat as it's specified for) and if the outside temp is
somewhere below 30Deg Celsius.
(Hopefully someone knows a smart formula for that)
Their effectiness drops if they're located too close to places where hot air
is not really moving, i.e. ground level or non-raised on a rooftop, or on a
roof section that is surrounded by higher parts of the building.
Spraying water with a sprinkler can help.

Cooling failures
----------------

Very important is to gather information on the duration of a cooling failure
so far. The surrounding datacenter structures delay a failures effect by
absorbing heat.
That means two things:
1. At some point the walls have gained a temperature close to the ambient.
From that moment, they will not absorb more heat, resulting in a massive
increase of the rate at which the room is heating up
2. A sliding delay that means the room will take somewhat as long as it took
to heat up to cool down again, even if you already shut down everything that
consumes power.

Generally you need to assess a few things:
- is the temperature currently still a tolerable operating temperature
- how quicky is it rising
- how soon will the fault be fixed and when will a tolerable temperature be
reached
- how long is the expected time of overheating

- what damage to expect from quick heatup and cooldown
  (i.e. tape backups in progress at the time of failure will be probably lost
  because of the stretching effects on the media. Disks can fail by the dozen at
  the same time)

- what are the absolutely critical systems running, and how much power do the use
- how many not absolutely critical systems are running, and how much power will
  it save to turn them off.
- How fast can you shut them down? (i.e. on Unix, that could mean init 2, wait,
  flip the switch)
- can you shut down enough to keep the most critical systems running?
- can you also shut down enough to massively flatten the temperature curve?

Generally, the goal is to flatten the effects so that the critical
infrastructure gets away with a very and constant slow hike up and down at
less than 2-3 deg / hr.
This is also to be respected when the cooling is fixed again, it should run
at emergency power only till the pivot is taking off the curve and the the
power needs to be constantly decreased so the ambient takes a slow fall instead
of going from 55C to 21C in 30 minutes.
(Dangers: Humidity, microscopic tears on PCB)


Physical security and common security standards compliance requirements
=======================================================================

Suggested practices
===================

Access control
--------------

Security should require a ID from anyone coming onsite, and not stop
at having you fill a form where you just write anything you want.

Techs should need to bring a ticket ID, and this ticket ID plus the tech's
name should be announced by the vendor to security, plus(!) a callback to
a defined number known to security. So, a double handshake with a token.
This is really important to avoid incidents like suddenly missing two
servers because "there were some techs that replaced it"

Most critical DCs will have security accompanying you to the server.
Some will keep the sec person around while the tech is working.
The really smart ones train their sec stuff so that they will know *which*
server the tech needs to go to and *which* disk is to be replaced as to the
original ticket. (If you think the security people can't handle that, then
just ask yourself who's to blame since it does work for other DCs)
