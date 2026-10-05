Architecture 201
****************

Service Oriented Architectures
==============================

Fault tolerance, fault protection, masking, dependability fundamentals
======================================================================

Fail open, fail closed
----------------------

Perspective: node, network, cluster, application
------------------------------------------------

Timeouts and retries
--------------------

When one service calls another, it usually sets a timeout: how long it waits for an answer before it gives up.
Without a timeout, a caller can wait indefinitely for a service that has stopped answering, and its own requests pile up behind the slow calls.

Choosing a timeout is a trade-off.
While a caller waits, it holds resources such as a connection, a thread or memory, so a long timeout lets a slow dependency use them up.
A short timeout gives up on calls that would have succeeded, for example during a normal burst of traffic.
A common starting point is a value comfortably above the dependency's 99th percentile latency at peak traffic, based on measurements.

Many failures are brief, such as a reset connection or a server that is restarting, so callers often retry a failed call.
Retries help with these short failures, but each retry is more work for a service that may already be struggling.

When a caller times out, the service it called usually isn't told.
Unless the service checks whether the caller is still connected, it can finish the work anyway, so a call that timed out can cost the service as much as one that succeeded.

Retries multiply load
~~~~~~~~~~~~~~~~~~~~~

A caller that retries up to three times can make four calls for one request.
If calls fail because the service is overloaded, retries add load at the worst time.
Retries at several layers multiply: if three layers of services each make up to four attempts, one request at the top can become up to 64 calls at the bottom.

This can keep a system overloaded after the cause has gone.
For example, a traffic peak slows a service down, so calls time out and are retried.
The retries add more load than the peak did, so calls keep timing out after traffic returns to normal.
This self-sustaining state is a metastable failure; here, the retry storm is the feedback loop that keeps it going.
Restarting the overloaded service can clear its backlog, but if nothing else changes, the next peak can start the overload again.

Limiting retries
~~~~~~~~~~~~~~~~

- Retry only errors that are likely to be temporary, such as timeouts and ``503 Service Unavailable``, and not errors such as ``400 Bad Request``.
- Retry only operations that are safe to repeat.
  For example, repeating a payment request can charge a customer twice, unless the payment service accepts an idempotency key that lets it recognize the repeat.
- Wait between attempts, and make each wait longer than the one before.
  This is called exponential backoff.
  Add a random amount to each wait, called jitter, so that many clients don't retry at the same moment.
- Limit the total amount of retrying, not only the attempts per request.
  For example, a client can stop retrying while retries are more than 10% of its requests.
  This is called a retry budget.
  A circuit breaker is a related approach: after many failures, the client stops calling the service for a short time and fails at once instead.
- Retry at one layer, not at every layer.
- Pass the remaining time to the service you call.
  For example, gRPC sends a call's deadline to the server, which can stop work that nobody is waiting for.

To see whether retries add load, compare the rate of calls to a service with the rate of the requests that need it.
If each request needs one call, a ratio well above one means that many calls are retries.
The service's own metrics, such as how many requests are waiting in its queue, show whether it is keeping up.

Further reading:

- `Handling Overload <https://sre.google/sre-book/handling-overload/>`_, in Google's *Site Reliability Engineering* book
- `Exponential Backoff And Jitter <https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/>`_, on the AWS Architecture Blog
- `Metastable Failures in Distributed Systems <https://sigops.org/s/conferences/hotos/2021/papers/hotos21-s11-bronson.pdf>`_, Bronson and others, HotOS 2021

Caching Concerns
================

Static assets
-------------

Data
----

Eviction and replacement policies and evaluation
------------------------------------------------

Approaches
----------
(TTL, purge-on-write, no-purge versioning, constantly churning cache versus
contained, working set sizing)

Crash only
==========

Synchronous vs. Asynchronous
============================

Business continuity vs. Disaster Recovery
=========================================

Designing for Scalability: Horizontal, Vertical
===============================================

Simplicity
==========

Performance
===========

Tiered architectures
====================

MTTR > MTBF
===========

http://www.kitchensoap.com/2010/11/07/mttr-mtbf-for-most-types-of-f/
