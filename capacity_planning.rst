Capacity Planning
*****************

(Seems reasonable to have the Statistics for Engineers course be a pre-req for
this course)

Fundamentals of capacity planning
=================================

Resource usage investigation and exploration 
---------------------------------------------

* Examples: CPU:req/sec ratio, memory footprint:req/sec ratio, disk consumption
  per user/per sale/per widget, etc.
* Application:Infrastructure metric relationships
* 2nd order capacity (logging,
  metrics+monitoring systems, ancillary systems)

Finding ceilings
----------------

* Discovering resource limits 
* Comparing different hardware/instance profiles - production load versus
  synthetic

 * Benchmarking: pitfalls, limitations, pros/cons
 * http://www.contextneeded.com/system-benchmarks

* Multivariate infra limits (multiple resource peak-driven usage) Ex: web+image
  uploads, caching storage+processing, etc.
* Architecture analysis (anticipating the next bottleneck)

Forecasting 
============

Linear and nonlinear trending and forecasting ("steering by your wake")
-----------------------------------------------------------------------

Details of automatic forecasting and scaling
--------------------------------------------

Seasonality and future events
-----------------------------

* Organic growth approaches (bottom-up infra driven, top-down app driven)
* inorganic growth events (new feature launch, holiday effects, "going viral",
  major public announcement) 
* Provisioning effects on timelines, financial tradeoffs 

Diagonal scaling
================

(vertically scaling your already horizontal architecture)

Reprovisioning and legacy system usage tradeoffs
------------------------------------------------
