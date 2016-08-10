Programming 201
***************

Common elements in scripting, and what they do
==============================================

Syntax
------

Variables
---------

Common data structures
----------------------

Rosetta stone? One man's hash is another's associative array is another man's
dict(ionary)?

Functions
---------

Objects
-------

C (A very basic overview)
=========================

The main loop
-------------

Libraries & Headers
-------------------

#include
--------

The Compiler
------------

The Linker
----------

Make
----

Lab: Open a file, write to it, close it, stat it & print the file info, unlink
it. Handle errors.

Ruby
====
Ruby is a very user-friendly, flexible language and fun to use.
To quote from the `Ruby website <http://www.ruby-lang.org/>`_, Ruby is described as:

.. epigraph::

  A dynamic, open source programming language with a focus on simplicity and productivity.
  It has an elegant syntax that is natural to read and easy to write.

The creator of Ruby, `Yukihiro "Matz" Matsumoto <http://en.wikipedia.org/wiki/Yukihiro_Matsumoto>`_, took various parts of his favourite languages (Perl, Smalltalk, Ada, Lisp and Eiffel) to create Ruby.

Reading and writing Ruby code is amazingly easy and fun.
Once you learn the basics, it is amazing how much can be achieved in so little and concise code.
A very simple example would be how iterations or loops are done in Ruby:

.. code-block:: cpp

  for(int i = 0; i < 3; ++i) {
      std::cout << "Hello"
  }

You will see this for Ruby:

.. code-block:: ruby

  > (1..3).each { puts "Hello" }
  Hello
  Hello
  Hello
  => 1..3

  > (1..3).each do
    puts "Hello again"
  > end
  Hello again
  Hello again
  Hello again
  => 1..3

Ruby is a very good tool to write scripts.

Although this will be not covered here in detail, a very important thing to keep in mind is that in Ruby, **everything is an object**.
This means that you can treat everything i.e. numbers, strings, classes, objects themselves, etc. as objects.
Even the simplest of Ruby code will use this principle:

.. code-block:: ruby

  > 3.times { puts "hello" }
  hello
  hello
  hello
  => 3
  > "michael".capitalize
  => Michael

Strictly speaking, there will be cases where the above statement is not true in Ruby.
For example, in Ruby, functions are not first class objects.
In some languages like Javascript and Python, functions are first class objects.
In these languages, a function can be treated like an object, i.e. they have attributes, they can be referenced and passed as parameters, etc.

Running Ruby Code
-----------------

Ruby scripts are usually text files with ``.rb`` extension. You can run your Ruby scripts as follows:

.. code-block:: console

  $ ruby script.rb

You can run ad-hoc Ruby code in an interactive session called the Interactive Ruby or ``irb`` in short.

.. code-block:: console

  $ irb
  1.9.3-p448 :001>

All Ruby examples in this topic will start with ``>``, short for 1.9.3-p448 :XXX>.
It means that it is running inside an irb session. *1.9.3-p448* is the Ruby version the author was running while writing this topic.
The XXX are line numbers.

Syntax
------

* Conditionals

* Symbols

* Blocks


Variables
---------

Common data structures
----------------------

* Arrays

Arrays in Ruby are ordered collections of heterogenous items.
Items can be added, inserted, removed from an array.
Arrays are indexed starting from 0.

.. code-block:: ruby

  > empty_ary = []
  => []
  > str_ary = ["Pune", "Mumbai", "Delhi"]
  => ["Pune", "Mumbai", "Delhi"]
  > num_ary = [1, 2, 3.14, 10]
  => [1, 2, 3.14, 10]
  > mix_ary = ["this array has", 3, "items"]
  => ["this array has", 3, "items"]
  > arr_in_ary = [1, 2, [3, 4], 5]
  => [1, 2, [3, 4], 5]
  > str_ary.each { |city| puts city }
  Pune
  Mumbai
  Delhi
  => ["Pune", "Mumbai", "Delhi"]
  > num_ary[0]
  => 1
  > num_ary[2]
  => 3.14

Notice how arrays are heterogenous, i.e. array elements can be of different types.
And an array can have array as its element.

Array objects are instances of Array class.
So all instance methods are accessible to array objects.
Discussing every method is beyond the scope of this topic but here are a few examples:

.. code-block:: ruby

  num_ary = [1, 2, 3.14, 10]
  > num_ary.first
  => 1
  > num_ary.last
  => 10
  > num_ary.length
  => 4
  > num_ary.empty?
  => false
  > empty_ary.empty?
  => true

It is highly recommended that one reads the `Ruby Array API documentation <http://ruby-doc.org/core-2.0/Array.html>`_.

* Hashes

Hashes in Ruby are ordered collection of unique keys and their values.
A hash key can be of any object type.
Values can be referenced by their keys.

.. code-block:: ruby

  > empty_hash = {}
  => {}
  > device_hash = { samsung: "Galaxy S", apple: "iPhone"}
  => {:samsung=>"Galaxy S", :apple=>"iPhone"}
  > device_hash[:samsung]
  => "Galaxy S"
  > country_hash = { "America" => "Washington DC", "India" => "New Delhi", "Germany" => "Berlin" }
  => {"America"=>"Washington DC", "India"=>"New Delhi", "Germany"=>"Berlin"}

Hash objects are instances of Hash class.
So all instance methods are accessible to hash objects.
Discussing every method is beyond the scope of this topic but here are a few examples:

.. code-block:: ruby

  > country_hash["America"]
  => "Washington"
  > country_hash["Sweden"] = "Stockholm"
  => "Stockholm"
  > country_hash
  => {"America"=>"Washington DC", "India"=>"New Delhi", "Germany"=>"Berlin", "Sweden"=>"Stockholm"}
  > country_hash.values
  => ["Washington DC", "New Delhi", "Berlin", "Stockholm"]
  > country_hash.length
  => 4
  > empty_hash.empty?
  => true

It is highly recommended that one reads the `Ruby Hash API documentation <http://www.ruby-doc.org/core-2.0/Hash.html>`_.


Functions
---------
Functions are used in Ruby to perform a specific task.
In Ruby parlance, functions are generally termed as methods.
Ideally, a single method should do a single task and no more.
In Ruby, methods accept parameters and return a value.

A methods is enclosed inside ``def`` and the ``end`` keywords.
Parentheses is optional in Ruby for passing parameters.
The last line inside a Ruby method is returned by the method. Using ``return`` keyword is optional.

..  code-block:: ruby

  > def print_hello
      puts "hello"
    end
  => nil
  > def sum(a, b)
      a + b
    end
  => nil
  > def sum2 a, b
      return a + b
    end
  => nil
  > print_hello
  => hello
  > sum(2, 3)
  => 4
  > sum 4, 6
  => 10


Objects and Classes
-------------------
As mentioned above, in Ruby, **everything is an object**.
Ruby also has a class called ``Object``.
It is the default root of all Ruby objects.

Ruby objects can have attributes and methods.
An instance of Object class (and in general, to create an instance of any class) can be created as follows:

..  code-block:: ruby

  > obj = Object.new
  => #<Object:0x007fcba39874b8>

In Ruby, you can create your custom classes.
These can used along with the classes that come with Ruby and its standard library.

Classes can have methods.
Classes also have a special method called ``initialize``.
When a new object is created in Ruby using ``new`` method, an uninitialized object is first created and then ``initialize`` is called.
Any parameters passed to ``new`` is passed to ``initialize``.

An instance variable in Ruby is prepended by ``@`` symbol.

..  code-block:: ruby

  > class Student
      def initialize(name, age)
        @name = name
        @age  = age
      end

      def details
        puts @name
        puts @age
      end
    end
  => nil
  > s1 = Student.new('Cathy', 20)
  => #<Student:0x007fcba39b78c0 @name="Cathy", @age=20>
  > s1.details
  Cathy
  20
  => nil


Rubygems
--------

.. todo:: Explain more about what rubygems are as well as http://rubygems.org

Databases
---------

Python
======
Python is one of the most versatile languages you're going to use in your career.
You will soon see that for almost everything you want to do, Python either has a something in its standard library or an amazing third-party module that you can import in seconds.
But since this is a guide for operations engineers, I'll focus the discussion more towards Python's scripting capabilities.

NOTE: Before I start, I want to point out a series of documents called `Python Enhancement Proposals <https://www.python.org/dev/peps/>`_, PEP for short.
Like their title suggests, these are potential enhancements to the Python language that have been proposed by members of the community.
There's a lot of them, and you don't have to go over every single one, but you can find some very useful tips and best-practices there.

Syntax
------
* Indentation

If you've ever written or read any code in C, C++, Java or C#, you're used to seeing curly braces (``{}``) pretty much everywhere.
These compiled languages use curly braces to denote the start and end of functions, loops and conditional statements.
Python, on the other hand, uses indentation to achieve the same goal. What this means is that where you see this in C++:

.. code-block:: cpp

  if (3>2) {
      // Do something
  }

You will see this for Python:

.. code-block:: python

  if (3>2):
      # Do something

As you can see, Python didn't need curly braces to signify the start or end of the if conditional; a simple indent does the job.
Now when it comes to indentation, `PEP8 <https://www.python.org/dev/peps/pep-0008/>`_ says that you should use 4 spaces to indent your code.
Keep in mind that this specifically means spaces and not tabs.
Fortunately for you, most text editors today can automatically convert tabs to spaces so you don't have to hit four spaces every time you want to indent a line.
However, if you are dealing with some legacy code that uses 8 space tabs, feel free to continue doing so.

Indentation is by far the most important part of Python's syntax you should keep track of.
If there's two lines in your code where one uses 4 spaces and another uses one 4-space tab, Python's going to give you errors when you try to run your script.
Be consistent with your indentation.

* Conditionals

Conditionals refer to ``if, else`` statements where you're checking if some condition is met and then taking action based on whether it is or not.
Python supports conditionals just like any other language, with the only exception being indentation as explained above.
A complete conditional block would look like this:

.. code-block:: python

  # Check if the variable 'num' is greater than or less than 5
  if (num > 5):
      print "Greater"
  else:
    print "Less"

You can even have 'else if' conditions, which in Python are used as ``elif``

.. code-block:: python

  # Check if the variable 'num' is 2 or 5
  if (num == 2):
      print "Number is 2"
  elif (num == 5):
      print "Number is 5"
  else:
      print "Number is neither 2 nor 5"

* Boolean Operations

Python can perform all of the standard boolean operations:``and``, ``or`` and ``not``.
The operations can be used as statements of their own:

.. code-block:: python

  >>> (3 > 2) and (3 < 4)
  True
  >>> (2 > 3) or (3 > 4)
  False
  >>> not (2 > 3)
  True

and even in conditionals:

.. code-block:: python

  if not ((2 < 3) or (3 > 4)):
      print "Neither statement is true"

Variables
---------
Variables in Python work just like in any other language.
They can be assigned values like this:

.. code-block:: python

  times = 4
  name = "John"

They can be used in almost any statement.

.. code-block:: python

  >>> print times
  4
  >>> times + times
  8

You might have noticed that the variable didn't have to be created with a specific type before being assigned a value.
Python allows you to assign any value to a variable and will automatically infer the type based on the value it is assigned.
This means that the value assigned to a variable can be replaced with another value of a completely different type without any issues.

.. code-block:: python

  >>> times = 4
  >>> print times
  4
  >>> times = "Me"
  >>> print times
  'Me'

However, if you try to perform an operation with two variables that have values of conflicting types, the interpreter will throw an error.
Take this example where I will try to add a number and a string.

.. code-block:: python

  >>> times = 4
  >>> name = "John"
  >>> times + name
  Traceback (most recent call last):
    File "<stdin>", line 1, in <module>
  TypeError: unsupported operand type(s) for +: 'int' and 'str'

As you can see here, the interpreter threw a TypeError when we tried to add an integer and a string.
But there is a way around this; Python lets you type cast variables so their values can be treated as a different type.
So in the same example, I can either try to treat the variable ``times`` as a string, or the variable ``name`` as an integer.

.. code-block:: python

  >>> str(times) + name
  '4John'
  >>> times + int(name)
  Traceback (most recent call last):
    File "<stdin>", line 1, in <module>
  ValueError: invalid literal for int() with base 10: 'John'

Here you can see that when we cast ``times`` as a string and added it to name, Python concatenated the two strings and gave you the result.
But trying to cast ``name`` as an integer threw a ValueError because 'John' doesn't have a valid base 10 representation.
Remember, almost any type can be represented as a string, but not every string has a valid representation in another type.

Common data structures
----------------------
Out of the box, Python implements a few major data structures.

* Lists

Lists in Python are the equivalent of arrays in other languages you may be familiar with.
They are mutable collections of data that you can append to, remove from and whose elements you can iterate over.
Here's some common operations you can perform with lists:

.. code-block:: python

  >>> to_print = [1, 4]
  >>> to_print.append('Hello')
  >>> to_print.append('Hey')
  >>> to_print
  [1, 4, 'Hello', 'Hey']
  >>> for i in to_print:
  ...     print i
  ...
  1
  4
  Hello
  Hey
  >>> to_print[1]
  4
  >>> to_print[-1]
  'Hey'
  >>> to_print[-2:]
  ['Hello', 'Hey']
  >>> to_print.remove(4)
  >>> to_print
  [1, 'Hello', 'Hey']

Just like arrays in other languages, Python's lists are zero-indexed and also support negative indexing.
You can use the ``:`` to get a range of items from the list.
When I ran ``to_print[-2:]``, Python returned all items from the second last element to the end.

You may have also noticed that I had both numbers and strings in the list.
Python doesn't care about what kind of elements you throw onto a list.
You can even store lists in lists, effectively making a 2-dimensional matrix since each element of the initial list will be another list.

* Dictionary

Dictionaries are a key-value store which Python implements by default.
Unlike lists, dictionaries can have non-integer keys.
Items of a list can only be referenced by their index in the list, whereas in dictionaries you can define your own keys which will then serve as the reference for the value you assign to it.

.. code-block:: python

  >>> fruit_colours = {}
  >>> fruit_colours['mango'] = 'Yellow'
  >>> fruit_colours['orange'] = 'Orange'
  >>> fruit_colours
  {'orange': 'Orange', 'mango': 'Yellow'}
  >>> fruit_colours['apple'] = ['Red', 'Green']
  {'orange': 'Orange', 'mango': 'Yellow', 'apple': ['Red', 'Green']}
  >>> fruit_colours['mango']
  'Yellow'
  >>> for i in fruit_colours:
  ...     print i
  ...
  orange
  mango
  apple
  >>> for i in fruit_colours:
  ...     print fruit_colours[i]
  ...
  Orange
  Yellow
  ['Red', 'Green']

You should be able to see now that dictionaries can take on custom keys.
In this example, my keys were names of fruits, and the value for each key was the colour of that particular fruit.
Dictionaries also don't care about what type your keys or values are, or whether the type of a key matches the type of its value.
This lets us store lists as values, as you saw with the colour of apples, which could be red and green.

An interesting property about dictionaries that you might have noticed, is that iterating through the dictionary returned only the keys in the dictionary.
To see each value, you need to print the corresponding value for the key by calling ``fruit_colours[i]`` inside the for loop where ``i`` takes on the value of a key in the dictionary.


Python implements a lot more data structures like tuples, sets and dequeues.
Check out the Python docs for more information these: http://docs.python.org/2/tutorial/datastructures.html


Functions
---------
Functions in Python work exactly like they do in other languages.
Each function takes input arguments and returns a value.
The only difference is syntax, you define functions with the keyword ``def``, and don't use curly braces like in Java, C, C++ and C#.
Instead, function blocks are separated using indentation.

.. code-block:: python

    >>> def square(x):
    ...     result = x*x
    ...     return result
    ...
    >>> square(3)
    9


You can even call functions within other functions

.. code-block:: python

    >>> def greet(name):
    ...     greeting = "Hello "+name+"!"
    ...     return greeting
    ...
    >>> def new_user(first_name):
    ...     user = first_name
    ...     print "New User: "+user
    ...     print greet(user)
    ...
    >>> new_user('Jack')
    New User: Jack
    Hello Jack!


Objects
-------

Version Control
===============

Git
---

SVN
---

CVS
---

API design fundamentals
=======================

RESTful APIs
------------

JSON / XML and other data serialization
---------------------------------------

Authentication / Authorization / Encryption and other security after-thoughts.
------------------------------------------------------------------------------

:)
https://github.com/ziliko/code-guidelines/blob/master/Design%20an%20hypermedia(REST)%20api.md

Continuous Integration
======================


