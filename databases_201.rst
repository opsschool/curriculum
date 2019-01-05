Databases 201
*************

Database Theory
===============

ACID
----

CAP theorem
-----------

Summary
~~~~~~~~~~~~~~~~
The CAP theorem describes the necessary tradeoffs between 3 properties of databases and other distributed systems: Consistency, Availability, and Partition tolerance.

We'll start with a discussion of Partition tolerance, as it is the basis for understanding the other two properties.

Partition Tolerance
~~~~~~~~~~~~~~~~~~~~~

Partition tolerance refers to a system's ability to function, given the presence of a network outage or “partition” between your system components.
Imagine you work at a remote office of a large company and the network link between your site and headquarters goes down.
This creates network partitions; the two networks are now split and systems within each cannot communicate across the rift.
It is also often desirable for our systems’ behaviors to change when it detects it is in a network partition situation.
We’ll call these partition mode behaviors, and cover examples below.

It is widely accepted that partitions are unavoidable and thus partition tolerance is essential.
In fact, even latency can be thought of as a type of network partition.
Therefore, we may want to think of the CAP theorem as: **Given that network partitions are unavoidable, systems must prioritize either consistency or reliability, and cannot maximize both simultaneously.**

The way in which these 3 properties relate to one another is best understood through demonstration, so let's think about the example of a bank with a distributed ATM network.

Consistency
~~~~~~~~~~~~~~~~

Consistency refers to data consistency.
Every request to a highly consistent system should return the latest written value.
A highly consistent system guarantees that all database instances share the same values at all times,
and minimizes scenarios in which inconsistency could be introduced.
In the case of our ATM system, this means that if you checked your account balance at 3 different ATMs within a very short time period, each would show the same account balance.
It also means that if you withdraw money at one ATM, that change would propagate to all ATMs almost instantly.

This raises an important question: In our CP-prioritized model, how should the ATM system behave if not all machines can be contacted, and therefore a change could not be instantly propagated?

In a system that uncompromisingly prioritizes consistency, the only acceptable behavior would be to disallow any withdrawals or deposits.
Balance checking is fine because it doesn't modify values.
However, if even a single machine was offline, we wouldn't be able to ensure consistency entirely across the ATM network, and therefore would have to restrict changes.
If you withdrew money from an ATM at location X, it may not be able to communicate the change to an ATM at location Y.
A balance check or withdrawal would be based off of a stale value.

Allowing read-only transactions (balance check) but not modifications (withdrawal, deposit) is a partitioned mode behavior.

Availability
~~~~~~~~~~~~~~~~

Availability refers to the property of a system such that it responds to all requests reliably; the system never becomes unavailable.
This is relatively straightforward to understand, but it does have implications for system behavior.

Let's again imagine a scenario in which our ATM network is suffering from some sort of outage and not all machines can communicate with one another.
In an AP-prioritized model, how should the ATM system behave if not all machines can be contacted?

In the section on consistency above, we said that to maintain consistency we would cease all transactions that modified values until they could be reliably propagated to the entire network.
Availability takes the opposite approach.
In order to maintain full service (availability) to our customers, all ATMs will continue to function and base decisions off of their own last-known (possibly stale) value for an account balance.

The benefit here is obvious: customers can still access their funds even during an outage.
The downside, however, is also pretty clear; a nefarious customer could drive from ATM to ATM withdrawing large sums of money and the ATMs would be unable to update each other about the transactions.
As alarming as this may seem, it is actually how most ATMs function.
However, there are typically daily withdrawal limits that could be locally enforced by each ATM, placing a cap on the extent to which one could exploit this situation.

Allowing customers to withdraw money but setting a fixed cap is a partitioned mode behavior.

Lastly, it's important to remember that in our AP design, we're allowing inconsistent values to be introduced into our system.
When the network comes back up, ATMs A, B, and C might have different values for a customer's account balance.
The system must go through a process of reconciling the discrepancies.
In the case of our ATM, this might mean transferring the details of all transactions to a centralized/master account management machine, sequencing them appropriately, and determining the resulting balance.


High Availability
-----------------

Document Databases
==================

MongoDB
-------

CouchDB
-------

Hadoop
------

Key-value Stores
================

Riak
----

Cassandra
---------

Dynamo
------

BigTable
--------

Graph Databases
===============

FlockDB
-------

Neo4j
-----


