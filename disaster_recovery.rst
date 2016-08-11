
Disaster Recovery
*****************

Disaster Recovery is a process for securing business continuity in the event of a disaster.
The severity of the disaster may differ, as well as how to resolve it.
However, the goal for a Disaster Recovery, or the Disaster Recovery Plan should always be to get to the state of business as usual as soon as possible.

Planning
========

The most important thing about Disaster Recovery is preparation and planning.
This is because when a disaster actually happens, you will not be able to change any prerequisites but are forced to work with what you already have.
If your plan was broken or didn't actually address all the critical aspects of your IT, your Disaster Recovery-environment will end up broken as well.
A misconfigured backup job can always be reconfigured during normal operations but when the environment is already down; that ship has sailed.

To make sure we cover those needs we will split this section into a couple of subsections. 

Business Needs
--------------

The first question we need to ask ourselves when creating a Disaster Recovery Plan, or DRP, is what the needs of the business actually is.
In most cases, an organization has some sort of critical infrastructure or systems that need to be available at all time to be able to perform the core business.
This might be an Enterprise Resource Management-system (ERP), a document collaboration or versioning system, production system etc.

The size of the organization as well as the grade of information technology adaptation all need to be taken into consideration when creating a plan for Disaster Recovery.

Defining these business needs can't be the sole work of the IT Department or an Operations Engineer.
It needs attention from all levels of the organizational hierarchy.
One method for identifying the needs might be interviews with colleagues that work in the core business.

**Questions that need to be answered by every role in the core process(es):**

* In your every day work, what activities do you need to be able to perform to fulfill the goals of your role?
* How do you do these?

Using the answers to the questions, we then need to identify each system or infrastructural component that is needed to realize that goal.
The identified objects gives us a list of what we actually need to secure in the case of a real disaster.

**This can be done in many ways, for example:**

* Backup/Restore
* Data Mirroring
* Off-Site Mirrored environment
* Failover

So, if we were to create a disaster recovery plan for Acme Inc, we would first need to identify what components we need to consider putting into the DRP.
Usually, you have some kind of system (or excel spreadsheet for smaller companies) containing all the IT environment components.
If so, this information will make your work a lot easier.

The components may contain, among other things:

* SQL Servers (sql-srv-001 and sql-srv-002)
* Domain Controllers (dc-srv-001 - primary, dc-srv-002 & dc-srv-003 - used for load balancing)
* Fileservers (stgsrv001)
* Application servers (app-srv-001, app-srv-002)

We will then, together with the rest of the organization, try to map these components to the bussiness activities we've identified earlier.
For example:

* The application server **app-srv-001** is running our ERP, which is needed to be able to place orders from customers.
* The SQL Server **sql-srv-001** contains the data from the ERP which means that it's a prerequisite for the application server.
* The domain controller is needed so the users are able to sign in to the ERP. However, since 002 and 003's main purposes is load balancing, recovering **dc-srv-001** will be our main objective.
* The fileserver **stg-srv-001** is used to store copies of the order receipts produced in the ERP.

What we've concluded from this activity is that we need to recover four components to be able to use the ERP.
Note however, that in reality, the IT environment's usually alot more complex then the one used in this example.

The identified objects then need to be ranked to determine in what order they need to be recovered to minimize downtime.
In this example, this would most likely be as follows:

* dc-srv-001
* sql-srv-001
* stg-srv-001
* app-srv-001

Prioritizing Recovery Components
--------------------------------
In the example above, the number of assets to prioritize is low, which might suggest that there is no need for making a priority list.
However, one thing that you'll learn by doing a couple of disaster recovery excercises is that no matter how small the scope, stakeholders will always try to direct your effors to the assets mosts relevant to them.
For example, the CEO might want you to prioritize the ERP System over the Domain Controller, which might very well be correct, but as the list of assets grow longer the number of stakeholders wanting to influence your decisions will as well.

As you might have guessed, it may be a good idea to actually prioritize your asset list, if not out of necessity then at least to circumvent any issues that might occur because of differing opinions on what needs to be prioritized.
A great way to start out is to create a spreadsheet consisting of the following columns:

* Asset Identification
    A name, a FQDN or an IP.
    Whatever helps you identifying the asset.
 
* Business Priority
    Non-Essential, Essential, Critical.
    This should be decided by either the board or the senior management team.

* Tiebreaker/Sequential numbering
    A sequential numbering which will break ties in case multiple assets have the same priority.
    Should also be decided by the board or the senior management team.

* Business Impact
    A textual description of what would happen if this asset where to be unavailable.

* Exceptions
    Any exceptions to the priority above. 
    A company that is doing billing once every month might not feel that the billing system is critical during any other period then the billing period. 
    This will of course reflect your real time prioritization if (when) a disaster occurs.

The finished product should, after a signoff from your department manager and the senior management team, be published on your company's intranet, available for anyone.
This is very important as lack of transparency is one of the most common prejudices about IT Departments.

.. TODO:: shared resources, bussiness needs.

Disaster Recovery Plans
-----------------------

.. TODO:: How to create a plan from the material we gathered in the planning phase.
.. TODO:: Pros and cons on separating the disaster recovery manual from the technical recovery manual.

Disaster Recovery Simulations
-----------------------------

.. TODO:: Strategies when simulating. Defining testing scopes. Measuring.

Considerations
--------------
.. TODO:: Limiting the scope to core business
.. TODO:: Expanding the scope in the disaster recovery environment vs. going back to production before expanding

Execution
=========
.. TODO:: Communication
