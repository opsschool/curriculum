Critical Thinking 101
*********************

Critical thinking is the habit of checking what you believe, and why, before you act on it.
Curiosity is the habit of asking how things work, including things that are not broken.
SRE and operations engineers use both in nearly all of their work: to find the cause of a problem, to design a system, to decide whether a change is safe, and to judge advice from colleagues, vendors and the Internet.

This chapter describes ways to practice these habits.
The examples use the hosts from the :ref:`sample-network`.

.. contents::
   :depth: 2
   :local:

Curiosity
=========

Learn one layer down
--------------------

Most of the systems you work with are built on other systems.
A web request uses DNS, TCP, often TLS, and HTTP, and each of these depends on the operating system and the network below it.
You do not need to know every layer in detail, but it helps to know the layer below the one you work on.

For example, if you run web servers, learn what happens between a client typing a URL and the server receiving the request.
When a request fails, you can then ask which step failed, instead of only knowing that "the site is down".

Follow surprises
----------------

When a system does something you did not expect, write it down, even if nothing is broken.
For example, a nightly backup job on db-2 usually takes 10 minutes, but tonight it took 40 minutes and still succeeded.
The job did not fail, but something changed: the amount of data, the disk, the network, or another job running at the same time.
Surprises like this often point to a problem before it causes an outage.
If you cannot look into it now, record it somewhere your team will see it, such as a ticket.

Read the source
---------------

When the documentation does not answer your question, look for a better source.
Good sources include manual pages (``man``), the software's own documentation, its source code, and standards documents such as RFCs.
Blog posts and answers on forums can help, but they may be out of date or describe a different version or configuration.

Ask questions
-------------

Ask colleagues how and why their systems work the way they do.
Most people are glad to explain their work, and you will often learn about history and constraints that are not written down anywhere.
Asking a question is also a quick way to find out whether your own understanding is correct.


Observations and inferences
===========================

An observation is something you saw or measured.
An inference is a conclusion you drew from it.
Problems start when an inference is treated as an observation.

For example, the application logs on app-1 show this error when it connects to db-1:

.. code-block:: none

    could not connect to server: Connection refused
        Is the server running on host "db-1" (10.10.10.16) and accepting
        TCP/IP connections on port 5432?

The observation is that app-1's connection to db-1 on TCP port 5432 was refused.
"The database is down" is an inference, and it may be wrong.
"Connection refused" means that the connection request was actively rejected.
Usually db-1 replied with a TCP reset because nothing is listening on that port, which can mean that the database is not running, or that it is listening on a different address or port.
A firewall rule on db-1 or on the path can also reject the connection and produce the same error.
If you tell your team "the database is down", they may restart a database that is running correctly.
If you tell them what you observed, they can reach their own conclusions and may notice something you missed.

When you report a problem, say what you observed, where, and when.
Then, separately, say what you think it means.

Assumptions
-----------

An assumption is something you believe without having checked it.
Everyone makes assumptions, and you cannot check everything.
The problem is assumptions you do not know you are making.

During an investigation, write down your assumptions, for example:

* The change I deployed only went to web-1 and web-2.
* app-1 and app-2 have the same configuration.
* The monitoring check for db-1 is working.

Then check the ones that would change your conclusion if they were wrong.


Evaluating evidence
===================

Where did the information come from?
------------------------------------

Information you collected yourself, from the system, is usually more reliable than information someone reported to you.
A user who says "the site is slow" is reporting a real experience, but they may be describing one page, one network or one moment.
A colleague who says "I already checked the logs" may have checked a different host or a different time range.
This is not a reason to distrust people; it is a reason to ask what exactly they saw.

Monitoring data is also evidence that someone collected and processed.
A graph may show an average, a sampled value, or a value that was collected every few minutes.
A monitoring check can fail, or keep reporting an old value, while the system it checks is in a different state.

Averages hide detail
--------------------

Suppose the load balancers send requests to web-1 through web-4 in equal numbers, and web-3 has started to respond slowly.
Three quarters of requests are fast and one quarter are slow.
The average response time rises a little, and may still look acceptable, but one in four users has a bad experience.

Percentiles describe the distribution better.
The 99th percentile (p99) is the value that 99% of measurements are at or below.
When you look at a summary such as an average, also look at the distribution, and at the values for each host.

Correlation and causation
-------------------------

Two things that happen at the same time are correlated.
That does not prove that one caused the other.
Both may have the same cause, or the timing may be a coincidence.

Recent changes are a common cause of problems, as described in :doc:`troubleshooting_101`, so a change that happened just before a problem started is a good place to look.
But check it: if you roll back the change and the problem continues, the change was probably not the cause, or not the only cause.

Absence of evidence
-------------------

Not finding something is weaker evidence than finding it.
If you search a log file on web-1 for errors and find none, there are several possible explanations:

* There were no errors.
* The errors were logged somewhere else, such as the systemd journal or a different file.
* The requests that failed were handled by a different host.
* The log file was rotated, and the errors are in an older file.
* Your search pattern did not match the format of the error messages.

Before you conclude that something did not happen, check that you would have seen it if it had.


Forming and testing hypotheses
==============================

A hypothesis is a possible explanation that you can test.
A useful hypothesis predicts something you can check, and which would be different if the hypothesis were wrong.

For example, users report that the site is slow.
You notice that web-3 has higher CPU usage than the other web servers, and form the hypothesis that web-3 is slow to respond.

This hypothesis predicts that requests sent directly to web-3 take longer than requests sent to the other web servers.
You can test that with ``curl``:

.. code-block:: console

    user@opsschool ~$ curl -s -o /dev/null -w '%{time_total}\n' http://web-1/health
    0.042
    user@opsschool ~$ curl -s -o /dev/null -w '%{time_total}\n' http://web-3/health
    1.874

One request to each host is a small sample, so repeat the test several times before you rely on it.
If web-3 is consistently slower, the hypothesis is supported, and you can test the next one: if you remove web-3 from the load balancer pool, the site should become faster for users.
If web-3 is not slower, the hypothesis is wrong, and its high CPU usage is a separate question.

Some guidelines for testing hypotheses:

* Prefer tests that could prove your hypothesis wrong.
  A test that would give the same result whether or not the hypothesis is true tells you nothing.
* Change one thing at a time.
  If you change two things and the problem goes away, you do not know which change fixed it.
* Write down what you tested and what you found.
  This helps you avoid repeating tests, and helps anyone who joins the investigation later.
* Keep more than one hypothesis in mind until the evidence rules them out.


Thinking in systems
===================

A system is a set of parts that affect each other.
Systems thinking means looking at how the parts interact, as well as at each part on its own.
When you form hypotheses, include explanations that involve interactions between parts, not only a single part that is broken.

Feedback loops
--------------

A feedback loop exists when the result of a process changes the input to the same process.

For example, db-1 becomes slow, and requests from app-1 and app-2 start to time out.
The applications retry the failed requests, which adds more load to db-1, which makes it slower, which causes more timeouts and more retries.
A loop like this can keep a system in a failed state after the original cause has gone.
Recovery may need the load to be reduced, for example by stopping the retries, before db-1 can catch up.
Limits on retries, exponential backoff with random jitter, and circuit breakers are common ways to prevent this kind of loop.

Other loops counteract a change.
When web-3 fails its health checks, the load balancers stop sending it requests, and web-1, web-2 and web-4 take its share.
This keeps the site working, but if the remaining servers do not have enough capacity, they can become slow and fail their health checks too.

Delays
------

An effect can appear some time after its cause.

For example, an autoscaling policy that uses average CPU usage over five minutes, with new servers that take several minutes to start, reacts to load from several minutes ago.
By the time the new servers are ready, the load may have fallen, and the policy removes them just before the load rises again.
The number of servers then keeps rising and falling without settling.

Delays also matter when you look for the cause of a problem.
A change deployed at 14:00 may cause a problem at 16:00, when a cache entry expires, a scheduled job runs, or a disk finally fills.
When you look for recent changes, use a time window long enough to include effects like these.

Shared dependencies
-------------------

Parts that look independent often depend on the same thing.
In the sample network, every host uses dns-1 and dns-2, and the web servers sit behind the same pair of load balancers.
If web-1 through web-4 all have problems at the same time, look at what they share before you look at each host separately.

Multiple contributing causes
----------------------------

Many failures happen only when several conditions are true at the same time.
For example, the disk on db-2 fills up during the nightly backup.
The data has grown, the job that deletes old backups has been failing for a week, and the disk space alert goes to a channel that nobody watches.
Any one of these on its own would probably not have caused an outage.

If you look for a single root cause, you will usually stop at the first condition you find.
Instead, ask what else had to be true for the failure to happen.
Each condition is a place where the failure could have been prevented or detected, and each one may cause a different failure in future.

Your model of the system
------------------------

Everyone works from a mental model of how a system works.
Diagrams and documentation are models too, and they are often out of date.
Treat a model as a set of assumptions, and check the parts that matter against the running system.

For example, a diagram may show that app-1 only connects to db-1.
``ss`` lists the TCP connections that app-1 has open to port 5432:

.. code-block:: console

    user@opsschool ~$ ss -tn state established '( dport = :5432 )'
    Recv-Q Send-Q   Local Address:Port      Peer Address:Port Process
    0      0          10.10.10.14:51234      10.10.10.16:5432
    0      0          10.10.10.14:51240      10.10.10.17:5432

app-1 also connects to db-2 (10.10.10.17), which the diagram does not show.
If you were planning maintenance on db-2, this changes which applications it affects.

Further reading
---------------

* Richard Cook, `How Complex Systems Fail <https://how.complexsystems.fail/>`_: a short paper on why failures in complex systems usually have several causes.
* Donella Meadows, *Thinking in Systems: A Primer*: an introduction to feedback loops, delays and other ideas from systems thinking.


Common reasoning errors
=======================

People make some reasoning errors so often that they have names.
These are sometimes called cognitive biases.
Knowing their names does not prevent them, but it makes them easier to notice, in yourself and in a team.

Confirmation bias
    Looking for, or giving more weight to, evidence that supports what you already believe.
    For example, after deciding that the network is the problem, you notice every slow ping and ignore the fast ones.
    To counter it, ask what evidence would show that you are wrong, and look for that.

Anchoring
    Relying too much on the first piece of information you received.
    The first theory raised in an incident often shapes the rest of the investigation, even after evidence against it appears.
    To counter it, list other explanations early, before the first one becomes the team's working assumption.

Availability
    Judging how likely something is by how easily examples come to mind.
    If the last outage was caused by DNS, DNS will seem like a likely cause for the next one, whether or not it is.
    To counter it, start from what you observe in this problem, not from what happened last time.

Hindsight bias
    After you know the outcome, believing that it was easy to predict.
    In a postmortem, it can lead to statements such as "they should have seen that the disk was filling up", when the people involved had other alerts and other work, and no reason to look at that disk.
    To counter it, describe what people knew and what they were paying attention to at the time, not what is known now.


Questioning claims and requests
===============================

Claims
------

You will often read claims about software and hardware: in vendor material, benchmarks, blog posts and lists of best practices.
Some are accurate, some are accurate only in certain conditions, and some are wrong.
Useful questions to ask include:

* Who made the claim, and do they benefit if you believe it?
* What exactly was measured, on what hardware, with what workload and configuration?
* Is that workload similar to yours?
* Can you reproduce the result?
* When was it written, and for which version?

A "best practice" is usually a good default for common situations.
Find out what problem it solves, so that you can tell when your situation is different.

Requests
--------

Many requests describe a solution instead of a problem, for example "we need two more database servers" or "please increase the disk on web-2".
Before you act, ask what problem the request is meant to solve, and what evidence shows the problem exists.
The request may be the right solution.
It may also treat a symptom: if the disk on web-2 is filling because log files are not rotated, a larger disk only delays the next alert.

Ask in a way that helps the person who made the request.
The aim is to solve their problem, not to refuse their request.


Being wrong
===========

You will often be wrong, especially early in an investigation.
That is expected; the aim is to find out quickly.

* Say how confident you are.
  "I think it is probably the change to the connection pool, but I have not confirmed it" gives your team more information than "It is the connection pool".
* When new evidence contradicts your conclusion, change your conclusion.
* When a colleague questions your reasoning, treat the question as information, not as an attack.
  Ask them the same kind of questions about their reasoning, in the same spirit.


Practice
========

* Pick a command you use often, such as ``ls``, ``ssh`` or ``curl``.
  On Linux, use ``strace`` to see which system calls it makes, and explain the output to a colleague.
* The next time you find the cause of a problem, write down the hypotheses you considered, how you tested each one, and what you found.
  Note which assumptions turned out to be wrong.
* Find a benchmark published for a piece of software you use.
  Write down how its test environment and workload differ from yours.
* Read a public postmortem from another company.
  Separate the observations from the inferences, and look for places where hindsight bias may have shaped the writing.
