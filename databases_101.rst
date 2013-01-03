Databases 101 (Relational databases)
************************************

What is a Database?
===================

A database is a structured collection of data. The structure is usually defined
to help model some aspect of reality. Typically data is stored using data-types
to ensure data consistency. Data is usually presented in tables, but can
also be presented in a variety of different forms.

What is a Relational Database?
==============================

A relational database is a database which data is organized into tables.
Data can be retrieved in may forms by combining multiple tables into new table using operations like
JOIN. The composed table will have new rows that are a combination of its parent tables.
This approach makes a relational database very efficient for storing data, since reused
data can be stored in a single table and referenced in other tables.

Why We Use Databases?
=====================

We use databases to store data. They help provide us guarantees that data is stored
and that it exists in the correct format. In addition most databases are heavily
optimized making data retrieval very fast.

What is SQL?
============

SQL or Structured Query Language is a domain specific language used for accessing
databases. It has a declarative syntax, meaning you declare the structure you want
returned, the sources, and constraints in a single statement. The database's
query parser and planner will determine a how to retrieve your data from this statement.

SQL shell
=========

Many relational databases provide an interactive CLI for interacting with the
database. For example MySQL provides the mysql command, Postgresql provides psql, and Oracle
provides psql. These programs give you the ability to compose queries, and diagnose
issues with the software.

Creating databases
==================

Most database platforms should support the creation of a database SQL *CREATE DATABASE*
query. It can be executed via a connection to the server, or via the SQL shell.

.. code-block:: sql

 
  CREATE DATABASE example_database;


Specialized Create Database Commands
------------------------------------

Some database platforms have specialized GUI tools, or CLI tools. For example
Postgresql has a UNIX command *createdb* that creates databases. In many
cases these commands are just wrappers around the standard SQL statement.

MySQL
~~~~~

.. code-block:: bash

	mysqladmin create example_database


Postgresql
~~~~~~~~~~

.. code-block:: bash

	createdb example_database


Some platforms, like MySQL support multiple databases per instance, while other platforms
like Oracle support one database per instance. You should check the documentation for
your particular vendor to see what is supported.

Creating users
==============

Users can be created in most databases using the the *CREATE USER* statement.

..  code-block:: sql

  CREATE USER username;

Specialized Create User Commands
--------------------------------

Some relational databases provide additional ways of creating users like specialized
command line programs.

MySQL
~~~~~

MySQL does not support creation of users via the *mysqladmin* command.

Postgresql
~~~~~~~~~~

.. code-block:: bash

   createuser username



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

Many databases stray from the SQL standard here, and it is important to read
your database's documentation when granting privileges. There may be
additional privileges not listed here and syntax can very significantly.

Removing Privileges
===================

Privileges are removed with the SQL REVOKE statement. It follows
a similar command format like grant:

.. code-block:: sql

  REVOKE [PRIVILEGE] on [OBJECT] FROM [USER];


 
Basic normalized schema design
==============================

A normalized schema is a database with a table and column structure designed
to reduce data redundancy. Typically data is placed into tables with a unique
identifier, or primary key, and then is referenced by id in any tables that
wish to use that data.

Suppose we have two types of records in a database; one for a city's population and
one for a city's average temperature. We could simply create the tables like so:

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

The advantage of this design is that it prevents you from having to enter
data multiple times, and generally reduces the storage cost of a row. If
San Francisco changed it's name you would only need to update
a single and row instead of two tables like the first example. And,
SQL makes it trivial to replace the id with a name using a JOIN statement
when the data is retrieved, making it functionaly identical to the two
table example.

Select, Insert, Update and Delete
=================================

SELECT
------

The SELECT statement is the standard way you read from a table in an SQL
database. You can use it to retrieve a set of data, and perform aggregations
on them. The standard syntax is:

.. code-block:: sql

  SELECT [column1,column2|*] FROM [TABLE];

By adding a WHERE statement, you can have the database filter results:

.. code-block:: sql

  SELECT user_id, user_name FROM users WHERE user_id = 1;

You can join tables using a JOIN statement:

.. code-block:: sql

  SELECT user_id, user_name FROM users u JOIN addresses a on u.user_id = a.user_id  WHERE user_id = 1;

You can count the rows in a table by using an aggregation:

.. code-block:: sql

  SELECT COUNT(1) FROM users;

You can order by a column:

.. code-block:: sql

  SELECT * FROM users ORDER BY user_name;


UPDATE
------

UPDATE is the SQL statement for updating the data in a table. It should almost
always be used with a conditional statement. The standard syntax is:

.. code-block:: sql

  UPDATE [TABLE] set [COLUMN] = {expression}, {COLUMN2={expression}, ...}
  [WHERE condition]
  [ORDER BY ...]
  [LIMIT count] ;

Without a WHERE condition the statement will apply to all the rows in a table.

Here is a simple example of an UPDATE statement:

.. code-block:: sql

  UPDATE users SET user_name = 'Jim Smith' WHERE user_name = 'James Smith';


DELETE
------

DELETE is the SQL statement for removing rows from a table. The standard syntax is
below.

.. code-block:: sql

  DELETE FROM [TABLE]
  [WHERE condition]
  [ORDER BY ...]
  [LIMIT count] ;

Without a WHERE condition the statement will apply to all the rows of a table.

Here is a simple example of a DELETE statement:

.. code-block:: sql

  DELETE FROM users WHERE user_name = 'James Smith';



Pro Tips
========

- use a ``LIMIT`` on ``UPDATE`` and ``DELETE FROM`` queries to limit damage imposed by an erroneous query

  ``UPDATE users SET disabled=1 WHERE id=1 LIMIT 1;``

  ``DELETE FROM users WHERE id=1 LIMIT 1;``
