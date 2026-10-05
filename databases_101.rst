Databases 101 (Relational Databases)
************************************

What is a Database?
===================

A database is a program or library that helps you store data.
They usually impose some sort of structure to the data to assist with querying and filtering.
The most common structure databases use to store data is a table, and most databases will use multiple tables to store data.

A table is composed of rows representing an item, and columns represent some attribute of that item.
Much like using a spreadsheet for a task list, where a row would be a task, and you may have a column for a task name, and a second column to determine whether it has been completed.

What is a Relational Database?
==============================

A relational database is a database in which data is organized into tables, usually with some form of relationship between them.
For example, a table containing customer names and addresses, and a table containing sales transactions.
Rather than repeat customer details every time a transaction is done, the data is stored once in one table, and then a unique reference is stored in the sales transactions table, linking each transaction to the customer.
This approach makes a relational database very efficient for storing data, since reused data can be stored in a single table and referenced in other tables.

Data can be retrieved in many forms by combining multiple tables into new tables using operations like ``JOIN``.
The composed table will have new rows that are a combination of its parent tables.

Why We Use Databases?
=====================

We use databases to store data.
They help provide us guarantees that data is stored and that it exists in the correct format.
In addition most databases are heavily optimized making data retrieval very fast.

Because of these attributes, we use databases a lot.
Most e-commerce sites use databases to keep inventory and sales records.
Doctors offices use databases to store medical records, and the DMV uses databases to keep track of cars.
Lawyers use databases to keep track of case law, and many websites use databases to store and organize content.
Databases are everywhere, and you interact with them daily.

What is SQL?
============

SQL or Structured Query Language is a :term:`domain specific language` used for accessing databases.
It has a declarative syntax, meaning you declare the structure you want returned, the sources, and constraints in a single statement.
The database's query parser and planner will determine a how to retrieve your data from this statement.

SQL shell
=========

Many relational databases provide an interactive :term:`CLI` for interacting with the database.
For example, MySQL provides the ``mysql`` command, Postgresql provides ``psql``, and Oracle provides ``sqlplus``.
These programs give you the ability to compose queries, and diagnose issues with the software.

todo:: Add example of connecting with each command.

MySQL
-----

To connect to a MySQL database from the CLI your command would usually take this form:

.. code-block:: console

  $ mysql -u username -p -h server-address
  Password:

Of these flags, ``-u`` = username, ``-p`` = password, and ``-h`` = hostname.
You may provide the password on the command prompt: ``-ppassword`` (no space after flag!).
We strongly advise against this, as this may leave the password visible in your shell history.

Creating databases
==================

Most database platforms allow you to create a new database using the ``CREATE DATABASE`` SQL query.
It can be executed via a connection to the server, or via the SQL shell.

A MySQL example:

.. code-block:: sql

  mysql> CREATE DATABASE example_database;


Specialized Create Database Commands
------------------------------------

Some database platforms have specialized :term:`GUI` tools, or CLI tools.
For example Postgresql has a UNIX command ``createdb`` that creates databases.
In many cases these commands are just wrappers around the standard SQL statements.

MySQL
~~~~~

.. code-block:: console

	$ mysqladmin create example_database


Postgresql
~~~~~~~~~~

.. code-block:: console

	$ createdb example_database

Some platforms, like MySQL support multiple databases per instance, while other platforms like Oracle support one database per instance.
You should check the documentation for your particular vendor to see what is supported.

Creating users
==============

Users can be created in most databases using the the ``CREATE USER`` statement.

.. code-block:: sql

  mysql> CREATE USER username;


Specialized Create User Commands
--------------------------------

Some relational databases provide additional ways of creating users like specialized command line programs.

MySQL
~~~~~

MySQL does not support creation of users via the ``mysqladmin`` command.

Postgresql
~~~~~~~~~~

.. code-block:: console

   $ createuser username

Create Tables
=============
Tables are organized into rows and columns.
Data is stored inside these tables.
In order to host this information we need to create a table.
We do this with the ``CREATE TABLE`` statement.

Standard Syntax is:

.. code-block:: sql

  CREATE TABLE table_name (
  column1 datatype(size),
  column2 datatype(size),
  column3 datatype(size)
  );

Here is an example.
This will create a table called "Users" with 2 columns.

.. code-block:: sql

  CREATE TABLE Users (
  name varchar(50),
  address varchar(50)
  );

Alter Table
===========

The ``ALTER`` statement will allow you to modify the design of your current database.
For example, if you have a table called "Users" which contains Names and Addresses, but you need need to add an age column, you could do so with:

.. code-block:: sql

  ALTER TABLE Users ADD age int;

Standard Syntax is:

.. code-block:: sql

  ALTER TABLE table_name ADD column_name datatype

If we need to modify a current table, we can also do so with the ``ALTER`` statement.
If in the previous example you wanted to change the datatype of age from an integer to varchar, you would do so with:

.. code-block:: sql

  ALTER TABLE Users MODIFY age varchar(50);

Standard Syntax is:

.. code-block:: sql

  ALTER TABLE table_name MODIFY column_name datatype

If you now realize you don't need the age column at all, you can drop it all together with:

.. code-block:: sql

  ALTER TABLE Users DROP age;

Standard Syntax is:

.. code-block:: sql

  ALTER TABLE table_name DROP column_name

Drop Table
===========

If you have created a table and need to remove it from your database, you can with the ``DROP`` statement:

Standard Syntax is:

.. code-block:: sql

  DROP table table_name

If we want to drop our table we created earlier (Users) we would use this command:

.. code-block:: sql

  DROP table Users;

Data Type
===========

There are different types of data that we will store in our database.
It can be simple text or a number, sometimes it is a mix of these things.
The table below lists some common Standard SQL commands that are also in MySQL.
It is important to note this is not a complete list.

========== ==================================================================================================
Data type  Description
========== ==================================================================================================
char(n)	   Fixed width character string, can hold letters and numbers-Max of 8,000 characters
varchar(n) Variable width character string, can hold letters and numbers- Max of 8,000 characters
text	   Variable width character string. Value is represented internally by a separately allocated object
int        Allows whole numbers between -2,147,483,648 and 2,147,483,647
float      A small number with a floating decimal point
decimal    An exact numeric data value, allowing for a fixed decimal point
date       Date formated as YYYY-MM-DD
datetime   Date and time combination formated as YYYY-MM-DD HH:MM:SS
timestamp  Stores a unique number that gets updated every time a row gets created or modified
========== ==================================================================================================


Granting privileges
===================

Privileges can be granted using the SQL GRANT statement. These
statements are persisted by the RDBMS when issued. The typical
command format is:

.. code-block:: sql

  GRANT [PRIVILEGE] on [OBJECT] to [USER];

The standard SQL privileges are:

========== ====================================================
Privilege  Description
========== ====================================================
ALL        Allows user all privileges
ALTER      Allows user to alter schema objects
CREATE     Allows user to create schema object like tables
DELETE     Allows user to delete from an object
EXECUTE    Allows user to execute a store procedure or function
INSERT     Allows user to add new data to an object
REFERENCES Allows user to create a referential table constraint
SELECT     Allows user to read from an object
TRIGGER    Allows user to create a trigger
TEMPORARY  Allows user to create temporary tables
UPDATE     Allows user to update existing data in an object
========== ====================================================

Below is an example granting a user SELECT privileges on a table

.. code-block:: sql

  GRANT SELECT ON TABLE example_table TO user;

You can also grant multiple privileges in a single statement.

.. code-block:: sql

  GRANT SELECT,INSERT ON TABLE example_table TO user;

Many databases stray from the SQL standard here, and it is important to read your database's documentation when granting privileges.
There may be additional privileges not listed here and syntax can vary significantly.

Removing Privileges
===================

Privileges are removed with the SQL ``REVOKE`` statement.
It follows a similar command format like grant:

.. code-block:: sql

  REVOKE [PRIVILEGE] on [OBJECT] FROM [USER];


Connections
===========

A program talks to a database server over a connection.
Opening a connection takes time, so most applications keep a pool of open connections and reuse them.
The application's configuration usually sets the size of the pool.

Each connection uses memory on the server, so the server limits how many it accepts.
In MySQL the limit is ``max_connections``, which defaults to 151.
When it is reached, new connections fail with this error:

.. code-block:: console

  ERROR 1040 (HY000): Too many connections

MySQL allows one connection more than ``max_connections``, for an account with the ``CONNECTION_ADMIN`` (or the older ``SUPER``) privilege, so that an administrator can usually still connect.
The limit can also be set per account, with ``max_user_connections``.

PostgreSQL's limit is also called ``max_connections``, and defaults to 100.
It reserves 3 of these for superusers by default (``superuser_reserved_connections``).
With the defaults, when the ordinary slots are exhausted, other clients get ``FATAL: remaining connection slots are reserved for superusers``.

Seeing who is connected
-----------------------

In MySQL, ``SHOW GLOBAL STATUS`` shows how many connections are open now, the most there have been since the server started, and how many were refused because of ``max_connections``:

.. code-block:: console

  mysql> SHOW GLOBAL STATUS WHERE Variable_name IN ('Threads_connected', 'Max_used_connections', 'Connection_errors_max_connections');
  +-----------------------------------+-------+
  | Variable_name                     | Value |
  +-----------------------------------+-------+
  | Connection_errors_max_connections | 169   |
  | Max_used_connections              | 152   |
  | Threads_connected                 | 8     |
  +-----------------------------------+-------+

To see which accounts and hosts the connections come from, and what they are doing, group the process list:

.. code-block:: console

  mysql> SELECT user, host, command, COUNT(*) AS connections
      -> FROM information_schema.processlist GROUP BY user, host, command;
  +-----------------+-----------+---------+-------------+
  | user            | host      | command | connections |
  +-----------------+-----------+---------+-------------+
  | app             | localhost | Sleep   | 7           |
  | root            | localhost | Query   | 1           |
  | event_scheduler | localhost | Daemon  | 1           |
  +-----------------+-----------+---------+-------------+

``Sleep`` means the connection is open but idle.
Some idle connections are normal, because pools keep connections open for reuse.
Many idle connections from one program, each with a large ``Time`` in ``SHOW PROCESSLIST``, can mean that the program takes connections from its pool and never returns them.
This is called a connection leak.

Keeping within the limit
------------------------

Add up the largest pool of every program that connects to the database, on every server it runs on, and keep the total below ``max_connections``, with room for administrators and tools.

MySQL closes a non-interactive connection after it has been idle for ``wait_timeout`` seconds, 28800 (8 hours) by default.
Lowering it closes idle connections sooner, but a connection that a program has leaked inside an open transaction still holds that transaction's locks until it closes.

Raising ``max_connections`` can give you time, but it doesn't fix a leak, and each connection needs memory on the server.


Basic normalized schema design
==============================

A normalized schema is a database with a table and column structure designed to reduce data redundancy.
Typically data is placed into tables with a unique identifier, or primary key, and then is referenced by id in any tables that wish to use that data.

Suppose we have two types of records in a database; one for a city's population
and one for a city's average temperature.
We could simply create the tables like so:

City Population:

=============   ==========
City            Population
=============   ==========
San Francisco   812,826
=============   ==========

City Temperature:

=============   ===========
City            Temperature
=============   ===========
San Francisco   73 F
=============   ===========

A normalized version of this would have three tables instead of two.

City ID and Name:

=============   =============
City_id         Name
=============   =============
1               San Francisco
=============   =============

City ID and Temperature:

=============   ===========
City_id         Temperature
=============   ===========
1               73 F
=============   ===========

City ID and Population:

=============   ==========
City_id         Population
=============   ==========
1               812,826
=============   ==========

The advantage of this design is that it prevents you from having to enter data multiple times, and generally reduces the storage cost of a row.
If San Francisco changed its name you would only need to update a single row instead of two tables like the first example.
SQL allows you to replace the ``id`` with a name using a ``JOIN`` statement
when the data is retrieved, making it functionally identical to the two table example.

Select, Insert, Update and Delete
=================================

SELECT
------

The ``SELECT`` statement is the standard way you read from a table in an SQL database.
You can use it to retrieve a set of data, and perform aggregations on them.
The standard syntax is:

.. code-block:: sql

  SELECT [column1,column2|*] FROM [TABLE];

By adding a ``WHERE`` statement, you can have the database filter results:

.. code-block:: sql

  SELECT user_id, user_name FROM users WHERE user_id = 1;

You can join tables using a ``JOIN`` statement.
In this example we're temporarily assigning an alias of 'u' for the table users and an alias of 'a' for the table addresses:

.. code-block:: sql

  SELECT user_id, user_name FROM users u JOIN addresses a
  ON u.user_id = a.user_id  WHERE user_id = 1;

Count the rows in a table by using an aggregation:

.. code-block:: sql

  SELECT COUNT(1) FROM users;

Order by a column:

.. code-block:: sql

  SELECT * FROM users ORDER BY user_name;

INSERT
------

The ``INSERT`` statement is used to add additional data into a table.
It can be used to insert data either a row at a time or in bulk.
The standard syntax is:

.. code-block:: sql

  INSERT INTO table (column1, column2, column3) VALUES (value1, value2, value2)

The column list is optional, if you don’t specify which columns you’re inserting data into, you must provide data for all columns.

For example, to insert a single row:

.. code-block:: sql

  INSERT INTO users (user_name,user_phone) VALUES ("Joe Bloggs","555-1234");

Or in bulk:

.. code-block:: sql

  INSERT INTO users (user_name,user_phone)
  VALUES ("John Smith","555-5555"),("Tom Jones","555-0987");

Inserting in bulk like that is typically much quicker than using separate queries as the query planner only has to execute once, and any indexes are updated at the end.


UPDATE
------

``UPDATE`` is the SQL statement for updating existing data in a table.
It should almost always be used with a conditional statement.
The standard syntax is:

.. code-block:: none

  UPDATE [TABLE] SET [COLUMN] = {expression}, {COLUMN2={expression}, ...}
  [WHERE condition]
  [ORDER BY ...]
  [LIMIT count];

Without a ``WHERE`` condition, the statement will apply to all the rows in a table.

Here is a simple example of an ``UPDATE`` statement:

.. code-block:: sql

  UPDATE users SET user_name = 'Jim Smith' WHERE user_name = 'James Smith';


DELETE
------

``DELETE`` is the SQL statement for removing rows from a table.
The standard syntax is:

.. code-block:: sql

  DELETE FROM [TABLE]
  [WHERE condition]
  [ORDER BY ...]
  [LIMIT count] ;

.. note:: Without a ``WHERE`` condition the statement will apply to **all** the rows of a table.

Here is a simple example of a DELETE statement:

.. code-block:: sql

  DELETE FROM users WHERE user_name = 'James Smith';


Indexes and slow queries
========================

To find the rows that match a ``WHERE`` condition, a database can read every row in the table and check each one.
This is called a full table scan.
It is fast on a small table, but its cost grows with the size of the table, so a query that was fast when a table was new can become slow as the table grows.

An index is a separate structure, usually a B-tree, that keeps the values of one or more columns in sorted order, with a pointer to each row.
With an index on the column in the ``WHERE`` condition, the database can go straight to the matching rows.
The primary key is always indexed.
Other columns can be indexed explicitly, or implicitly when constraints such as ``UNIQUE`` and, in MySQL, foreign keys are defined:

.. code-block:: sql

  CREATE INDEX idx_users_email ON users (user_email);

Indexes are not free.
Each index uses disk space and memory; every ``INSERT`` and ``DELETE`` updates it, and an ``UPDATE`` does so when it changes an indexed column.
So tables usually have indexes for the queries that the application runs often, not for every column.

An index can cover more than one column, for example an index on ``posts (author_id, published_at)``.
The order of the columns matters.
This index helps a query that filters on ``author_id``, or on ``author_id`` and ``published_at``, but usually not a query that filters only on ``published_at``.
Because each author's entries in the index are in ``published_at`` order, it also lets a query such as ``WHERE author_id = 7 ORDER BY published_at DESC LIMIT 10`` read the newest 10 posts directly, without sorting.

Seeing how a query runs
-----------------------

``EXPLAIN`` shows how the database plans to run a query, without running it.
In MySQL, ``type: ALL`` means a full table scan, ``key`` is the index used, if any, and ``rows`` is an estimate of how many rows MySQL will examine:

.. code-block:: console

  mysql> EXPLAIN SELECT user_id, user_name FROM users WHERE user_email = 'user4242@example.com'\G
  *************************** 1. row ***************************
             id: 1
    select_type: SIMPLE
          table: users
     partitions: NULL
           type: ALL
  possible_keys: NULL
            key: NULL
        key_len: NULL
            ref: NULL
           rows: 199378
       filtered: 10.00
          Extra: Using where

After ``CREATE INDEX idx_users_email ON users (user_email)``, the same query uses the index and is estimated to examine one row:

.. code-block:: console

           type: ref
  possible_keys: idx_users_email
            key: idx_users_email
           rows: 1

``Using filesort`` in the ``Extra`` column means that MySQL sorts the matching rows itself, because no index gives them in the order the query asks for.
That is fine for a few rows, but slow for many.

PostgreSQL also has ``EXPLAIN``, with a different output format: ``Seq Scan`` is a full table scan and ``Index Scan`` uses an index.

Finding slow queries
--------------------

MySQL can log every query that takes longer than ``long_query_time`` seconds to the slow query log.
The log is off by default, and ``long_query_time`` defaults to 10 seconds, so many servers set it lower.
To see the current settings:

.. code-block:: console

  mysql> SHOW VARIABLES WHERE Variable_name IN ('slow_query_log', 'long_query_time', 'slow_query_log_file');
  +---------------------+-------------------------+
  | Variable_name       | Value                   |
  +---------------------+-------------------------+
  | long_query_time     | 0.500000                |
  | slow_query_log      | ON                      |
  | slow_query_log_file | /var/log/mysql/slow.log |
  +---------------------+-------------------------+

Each entry shows the query, how long it took, and ``Rows_examined``.
A query that examines many more rows than it returns is often missing an index.
``mysqldumpslow``, which comes with MySQL, groups similar queries in the log and sums their times.

MySQL's ``sys.schema_unused_indexes`` view lists indexes with no recorded usage since their Performance Schema counters were last reset, normally at server startup.
Because this evidence is local and time-bounded, check every database instance over a representative period that includes periodic workloads before dropping an index.

Pro Tips
========

- Before doing a write query, run it as a read query first to make sure you are retrieveing exactly what you want.
  If your query is:

    ``UPDATE users SET disabled=1 WHERE id=1;``

  Run this first to validate you will be affecting the proper record:

    ``SELECT disabled FROM users WHERE id=1;``

- use a ``LIMIT`` on ``UPDATE`` and ``DELETE FROM`` queries to limit damage imposed by an erroneous query

  ``UPDATE users SET disabled=1 WHERE id=1 LIMIT 1;``

  ``DELETE FROM users WHERE id=1 LIMIT 1;``

- If your database supports transactions, run ``START TRANSACTION`` first then run your query and check what it has done.
  If you're happy with what you see then run ``COMMIT`` and finally ``STOP TRANSACTION``.
  If you realize you've made a mistake, you can run ``ROLLBACK`` and any changes you've made will be
  undone.
