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

Document Databases
==================

MongoDB
-------

CouchDB
-------

Hadoop
------

ElasticSearch
-------------

**Intro**

*ElasticSearch* - it's near real-time search engine and document-oriented datastore. 

It's written in Java, so lots of processes depends on JVM (Java Virtual Machine). System has awesome REST API, its documents and queries - all is represented via JSON. All configuration can be set up in elasticsearch.yml file. 

Analogues: Apache Solr, Sphinx.

**Under the microscope**

Elasticsearch is great for scaling. And this is probably one of the most noticable features. You can scale elasticsearch cluster for any load. So to begin with, what is a cluster?

Cluster is simply a set of nodes, physical, virtual or containerised... whatever. Nodes consist of indexes, indexes consist of shards. Index may have several types. Types are only logical, not physical. Type can be considered as a table when comparing Relational DB.

Shards consist of Lucene segments (segment is a chunk of Lucene index). And this is the firsts rule for Elasticsearch - all spins around Apache Lucene (Java library for full-text search).

Segments are created as you index new documents. Data is never removed from them because deleting only marks documents as deleted. Finally, data never changes because updating documents implies re-indexing. 

And here is the second rule - Elasticsearch is awesome for reading, but not so cool for updating/deleting. The more segments you have to go through, the slower the search. To avoid having an extremely large number of segments in an index, Lucene merges them from time to time => excluding the deleted documents, and creating new and bigger segments.

Shard can be primary or replica. Replica is a copy of primary. Primary shards are created only at the beginning and cannot be changed (deleted/added) while replica can. When primary is down, replica is promoted to primary, new replica created. Shards are special unit that can be balanced between nodes to maintain equal distribution/load/fault tolerance.

**Cluster states and roles**

Cluster has 3 health states:

- *Green* - all primary and replica shards available
- *Yellow* - one or more replicas are down
- *Red* - one or more primaries are down 

As for roles:

- *Master* - responsible for checking alive, healthy state, fault detection
- *Data* - stores data
- *Client* - accepts requests, use to remove load from master/data nodess

**Documents**

Document is uniquely identified by index-type-id combination. Elastic is generally schema-free, but can be defined via mappings. Mapping is the process of defining how a document, and the fields it contains, are stored and indexed. 

Field can be several types: string, numeric, date, boolean, arrays, multi-fields. Plus predefined fields (used internally by Elastic): _all, _ttl, _timestamp, _source, _id, _type, _index.

**How Elastic indexes?**

- Request (document) received => need to choose shard
- By default documents are distributed evenly between shards
- Shard is determined by hashing document's ID
- When shard is determined => sending document to primary. Then backing up to replicas
- Lucene segments are first stored in memory
- Refreshing process makes newly indexed documents available for search
- Flushing process transfers indexed data from memory to the disk
- Flush is triggered in one of the following conditions: memory buffer is full, time passed since last flush, log hit size treshold
- Bigger segments are periodically created from smaller segments to consolidate the inverted indices and make searches faster - it's called merge

**How Elastic searches?**

- Request forwards to shard containing data
- Using round-robin Elastic choose either primary or replica shard
- Gather all the results (aggregation from different segments/shards) and gives it back
- Ranking algorithm is `TF-IDF <https://en.wikipedia.org/wiki/Tf%E2%80%93idf>`_

**How document is updated?**

- Retrieve the existing document
- Apply the changes you specify
- Removes the old document and indexes the new document (with the update applied) in its place
- Version of document is bumped

**How document is deleted?**

- Delete individual documents or groups of documents only marks them to be deleted, so they don’t show up in searches, and gets them out of the index later in an asynchronous manner. Delete old docs occurs during merge.
- Delete complete indices easy to do performance-wise, happens almost instantly
- Interesting fact: you can close indices => doesn’t allow read or write operations and its data isn’t loaded in memory. Remains on disk, easy to restore.
- When you remove a mapping type, all the documents associated with it also are removed => awesome in terms effectivity.

**Interesting features**

Aliases, caches, warmers, filters, custom routing, pagination, bulk requests, tokenizers, SQL support from 6.3 version... and much more!

**Limitations**

- Lucene index can’t have more than 2.1 billion documents or more than 274 billion distinct terms
- JVM => Gold rule is to allocate half of the node’s RAM to Elasticsearch, but no more than 32 GB
- Refresh, flush and merge operations are expensive in terms of performance (CPU, I/O), need to be aware of this

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