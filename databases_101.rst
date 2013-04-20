Databases 101 (Relational databases)
************************************

What is a Database?
===================

A database is a program or library that helps you store data. They usually
impose some sort of structure to the data to assist with querying and filtering. The
most common structure databases use to store data is a table, and most database
will use multiple tables to store data.

A table is composed of rows representing an item, and columns represent some
attribute of that item; much like using a spreadsheet for a task list, where
a row would be a task, and you would have a column for a task name, and a second
column for whether it has been completed.

What is a Relational Database?
==============================

A relational database is a database in which data is organized into tables,
usually with some form of relationship between them, for example a table
containing customer names and addresses, and one containing sales transactions.
Rather than repeat customer details every time a transaction is done, the data
gets stored once into a single table, and then unique reference is stored in
the sales transactions table, linking each transaction to the customer. This
approach makes a relational database very efficient for storing data, since
reused data can be stored in a single table and referenced in other tables.
It thus also reduces the likelihood of mistakes.

Data can be retrieved in many forms by combining multiple tables into new tables
using operations like JOIN. The composed table will have new rows that are a
combination of its parent tables.

Why We Use Databases?
=====================

We use databases to store data. They help provide us guarantees that data is stored
and that it exists in the correct format. In addition most databases are heavily
optimized making data retrieval very fast.

Because of these attributes, we use databases a lot. Most e-commerce sites use
databases to keep inventory and sales records. Doctors offices use databases to
store medical records, and the DMV uses databases to keep track of cars. Lawyers
use databases to keep track of case law, and many websites use databases to
store and organize content. Databases are everywhere, and you interact with them
daily.

What is SQL?
============

SQL or Structured Query Language is a :term:`domain specific language` used for accessing
databases. It has a declarative syntax, meaning you declare the structure you want
returned, the sources, and constraints in a single statement. The database's
query parser and planner will determine a how to retrieve your data from this statement.

SQL shell
=========

Many relational databases provide an interactive :term:`CLI` for interacting with the
database. For example MySQL provides the ``mysql`` command, Postgresql provides ``psql``, and Oracle
provides ``sqlplus``. These programs give you the ability to compose queries, and diagnose
issues with the software.

todo:: Add example of connecting with each command.

MySQL
-----

To connect to a mysql database from the CLI your command would usually take the form:

.. code-block:: console

  $ mysql -u username -p -h server-address
  Password:

Of these flags, -u = username, -p = password, and -h = hostname. If you wish you
can put the password on the command prompt: -ppassword (no space after flag!)
but it is strongly advised that you don't do this as this will leave the
password in your shell history.

Creating databases
==================

Most database platforms allow you to create a new database using the ``CREATE DATABASE``
SQL query. It can be executed via a connection to the server, or via the SQL shell.

.. code-block:: sql
 
  CREATE DATABASE example_database;


Specialized Create Database Commands
------------------------------------

Some database platforms have specialized GUI tools, or CLI tools. For example
Postgresql has a UNIX command ``createdb`` that creates databases. In many
cases these commands are just wrappers around the standard SQL statement.

MySQL
~~~~~

.. code-block:: console

	mysqladmin create example_database

Postgresql
~~~~~~~~~~

.. code-block:: console

	createdb example_database


Some platforms, like MySQL support multiple databases per instance, while other platforms
like Oracle support one database per instance. You should check the documentation for
your particular vendor to see what is supported.

Creating users
==============

Users can be created in most databases using the the ``CREATE USER`` statement.

..  code-block:: sql

  CREATE USER username;

Specialized Create User Commands
--------------------------------

Some relational databases provide additional ways of creating users like specialized
command line programs.

MySQL
~~~~~

MySQL does not support creation of users via the ``mysqladmin`` command.

Postgresql
~~~~~~~~~~

.. code-block:: console

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
San Francisco changed its name you would only need to update
a single and row instead of two tables like the first example. And,
SQL makes it trivial to replace the id with a name using a JOIN statement
when the data is retrieved, making it functionally identical to the two
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

You can join tables using a JOIN statement.  In this example we're temporarily
assigning an alias of 'u' for the table users and an alias of 'a' for the
table addresses:

.. code-block:: sql

  SELECT user_id, user_name FROM users u JOIN addresses a on u.user_id = a.user_id  WHERE user_id = 1;

You can count the rows in a table by using an aggregation:

.. code-block:: sql

  SELECT COUNT(1) FROM users;

You can order by a column:

.. code-block:: sql

  SELECT * FROM users ORDER BY user_name;

INSERT
------

The INSERT statement is used to add additional data into a table. It can be
used to insert data either a row at a time or in bulk. The standard syntax is:

.. code-block:: sql

  INSERT INTO table (column1, column2, column3) VALUES (value1, value2, value2)

The column list is optional, if you don’t specify which columns you’re
inserting data into, you must provide data for all columns.

For example, to insert a single row:

.. code-block:: sql

  INSERT INTO users (user_name,user_phone) VALUES ("Joe Bloggs","555-1234");

Or in bulk:

.. code-block:: sql

  INSERT INTO users (user_name,user_phone) VALUES ("John Smith","555-5555"),("Tom Jones","555-0987");

Inserting in bulk like that is much quicker than using separate queries as the
query planner only has to execute once, and any indexes are updated at the end.


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

DELETE is the SQL statement for removing rows from a table. The standard syntax
is below.

.. code-block:: sql

  DELETE FROM [TABLE]
  [WHERE condition]
  [ORDER BY ...]
  [LIMIT count] ;

.. note:: Without a WHERE condition the statement will apply to **all** the
   rows of a table.

Here is a simple example of a DELETE statement:

.. code-block:: sql

  DELETE FROM users WHERE user_name = 'James Smith';



Pro Tips
========

- Before doing a write query, run it as a read query first to make sure you are
  getting exactly what you want.  If your query is
  ``UPDATE users SET disabled=1 WHERE id=1;``
  then first run this to make sure you are getting the proper record:
  ``SELECT disabled FROM users WHERE id=1;`` 

- use a ``LIMIT`` on ``UPDATE`` and ``DELETE FROM`` queries to limit damage
  imposed by an erroneous query

  ``UPDATE users SET disabled=1 WHERE id=1 LIMIT 1;``

  ``DELETE FROM users WHERE id=1 LIMIT 1;``

- If your database supports transactions, run ``START TRANSACTION`` first then
  run your query and check what it has done. If you're happy with what you see
  then run ``COMMIT`` and finally ``STOP TRANSACTION``.  If you realise you've
  made a mistake, you can run ``ROLLBACK`` and any changes you've made will be
  undone.
