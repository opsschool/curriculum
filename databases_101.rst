Databases 101 (Relational databases)
************************************

SQL shell
=========

Creating databases
==================

Creating users
==============

Granting privileges
===================

Basic normalized schema design
==============================

Select, Insert, Update and Delete
=================================

Pro Tips
========

- use a ``LIMIT`` on ``UPDATE`` and ``DELETE FROM`` queries to limit damage imposed by an erroneous query

  ``UPDATE users SET disabled=1 WHERE id=1 LIMIT 1;``

  ``DELETE FROM users WHERE id=1 LIMIT 1;``
