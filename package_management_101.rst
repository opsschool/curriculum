Package management
******************

What is a package manager?
==========================
To understand what is a package manager let's think of it like something real.
If your roommate sent you a package inside that package there is a paper written on it where will you put every single thing out of that package.
Then you will take everything out and put it in the right place.
When you roommate come he will find everything in it's expected place.
In this scenario you were the package manager.
So under linux you have a package manager who will extract the software package content into the right place with right configurations.
As you will see it is capable of 

* Installing packages "getting the stuff out of the package to the right place"
* Upgrading packages "think of it like getting a new package to extract but with some newer stuff in it"
* Uninstalling packages "remove stuff from where you put it"
* Querying the RPM database "like asking from what package did that thing came?,where is that thing?"

RPM and YUM (RedHat, CentOS, Fedora, Scientific Linux)
===========================================================
You can think of rpm as you while unpacking the package.
In that case you got the package from a friend or whatever the point that you went out there and got it.
So what is YUM?
If there is a store out there with a lot of packages ordered and managed and that store hires a man that man's job is when you call him and say i want a package of something that man will go to the store bring that package to your place extract everything into the right place so no need for you to go out there and get that package by yourself then hand it over to the package man "RPM".
YUM will be that store's guy and the store will be called a repository.
**repository** sometimes called repo is somewhere that keeps packages ordered and maintained.
And of course YUM can do (Installing,Deleting,Updating and Querying)  

RPM
===
Installing packages
-------------------
* rpm -ivh package.rpm

  * *i*   means install
  * *v*   verbose which means you will be watching the guy extracting things out of the package
  * *h*   hashes shows a progress bar of hashes "######"

Upgrading packages
------------------
* rpm -Uvh package.rpm

  * of course *U* stands for update 

Uninstalling packages
---------------------
* rpm -ev package.rpm

  * *e* for erase 

Querying the RPM database
-------------------------

* rpm -qa 
  
  * shows all the packages on the system
* rpm -qi     *<package>*

  * Shows information about that package 
* rpm -qf  *</path/to/file>*

  *  tells you which package installed that file

YUM
===
Well YUM as I said will have more capabilities than RPM for example in the repos you will find things called groups those are groups of software packages packed together like
installing all the media related software using that group instead of installing single packages.

you can also search the repos for specific packages or group of packages it's like getting the products menu of the store with the ability of searching in it.
 

Installing packages
-------------------

* yum install <package1> <package2> <packageN>

Installing groups
-----------------

* yum groupinstall <groupname>

Upgrading groups
----------------

* yum groupupdate <groupname>

Upgrading packages
------------------

* yum update <package>

Checking for update 
-------------------

* yum list updates

Update the whole system
-----------------------

* yum update


Searching A package 
-------------------

* yum search <package or part of it's name>


Listing available groups 
------------------------

* yum grouplist 


Uninstalling packages
---------------------

* yum remove <package1> <package2> <packageN>


Uninstalling groups
-------------------

* yum groupremove <groupname>

Creating packages
-----------------
Mention spec files and roughly how RPMs are put together.
Then introduce FPM and tell them not to bother with spec files yet.
