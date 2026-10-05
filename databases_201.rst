Databases 201
*************

Database Theory
===============

ACID
----

CAP theorem
-----------

High Availability
-----------------

Locks and long transactions
===========================

This section uses MySQL, but other relational databases have similar locks.

A transaction groups statements so that they succeed or fail together.
It starts with ``BEGIN`` (or the first statement, when autocommit is off) and ends with ``COMMIT`` or ``ROLLBACK``.
With autocommit on, which is MySQL's default, each statement outside an explicit transaction is a transaction on its own.

Metadata locks
--------------

MySQL protects a table's definition with a metadata lock.
A statement that reads or changes rows, such as ``SELECT`` or ``UPDATE``, takes a shared metadata lock on each table it uses.
Inside a transaction, MySQL keeps these locks until the transaction ends, not only until the statement ends.
Many sessions can hold shared locks on the same table at once.

A statement that changes the table's definition, such as ``ALTER TABLE``, needs an exclusive metadata lock, at least for a moment.
This is true even of changes that MySQL can make instantly, such as adding a column.
The ``ALTER TABLE`` waits until every transaction that holds a shared lock on the table has ended.

While it waits, it can block other queries.
By default MySQL gives a waiting exclusive lock request priority over later shared requests.
So queries that start after the ``ALTER TABLE`` wait behind it, even simple reads, and the table looks frozen to the application.
How long the ``ALTER TABLE`` waits is set by ``lock_wait_timeout``, which defaults to 31536000 seconds, one year.

The most common cause is a transaction that was started and never finished: for example, a program that runs a ``SELECT`` inside a transaction and then waits for something else, or a person who ran ``BEGIN`` and a query in a ``mysql`` shell and left it open.
Such a session shows as ``Sleep`` in the process list, which makes it easy to miss.

Finding the blocker
-------------------

``SHOW PROCESSLIST`` shows what each connection is doing.
Here, an ``ALTER TABLE`` waits for a metadata lock, a ``SELECT`` waits behind it, and session 488 is idle:

.. code-block:: console

    root@opsschool ~# mysql -e "SHOW PROCESSLIST"
    Id   User  Host       db    Command  Time  State                            Info
    488  root  localhost  NULL  Sleep    6                                      NULL
    489  root  localhost  NULL  Query    4     Waiting for table metadata lock  ALTER TABLE demo.products ADD COLUMN size INT
    490  root  localhost  NULL  Query    2     Waiting for table metadata lock  SELECT name FROM demo.products WHERE id = 1

``information_schema.innodb_trx`` lists open InnoDB transactions.
Join it with the process list to find sessions that are idle inside a transaction, and when their transactions started:

.. code-block:: console

    root@opsschool # mysql -e "SELECT p.id, p.user, p.host, t.trx_started
        FROM information_schema.innodb_trx t
        JOIN information_schema.processlist p ON p.id = t.trx_mysql_thread_id
        WHERE p.command = 'Sleep'"
    id   user  host       trx_started
    488  root  localhost  2026-10-05 02:04:01

The ``host`` column, and the user, help you find which program owns the connection.
The ``performance_schema.metadata_locks`` table and the ``sys.schema_table_lock_waits`` view show the locks themselves: who holds them and who waits for them.

Ending the wait
---------------

``KILL`` followed by a process list ID ends a connection, and rolls back its open transaction.
Killing the idle session lets the ``ALTER TABLE`` run.
Killing the ``ALTER TABLE`` instead lets the queued queries run, and the change can be tried again later.
Before you kill a connection, check what owns it: the program may reconnect and open a new transaction the same way.

To make schema changes safer:

- Before an ``ALTER TABLE``, check ``innodb_trx`` for old transactions.
- Set a short ``lock_wait_timeout`` for the session that runs the change, for example ``SET SESSION lock_wait_timeout = 5``.
  The change then gives up quickly instead of blocking the table for a long time, and you can retry it later.
- Fix programs that leave transactions open.

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


