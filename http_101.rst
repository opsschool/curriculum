HTTP 101 (Core protocol)
************************

HTTP is a response-request protocol in the application layer of the :term:`OSI` model.
An HTTP server listens for an incoming client request on a socket bound to an address and a port.
The address can either be a specific address configured on the host, or a wildcard to accept connection on all IP addresses configured on the host.
By default the port is 80, but this is not required.

Request
=======

Once the TCP connection is made (:term:`OSI` layer 4), the client initiates HTTP communication by sending an HTTP request on the established connection.
The HTTP request includes a request line, which is composed of a method, URI and protocol version.
A request method defines the action to be executed on the target resource, which on its turn is identified by the URI.
The protocol field indicates the version of HTTP to be used.

.. code-block:: console

	GET /index.html HTTP/1.1
	Host: example.com


The request line is followed by headers containing client information and other metadata.
As per protocol specification, the client must include the ``Host`` header in the HTTP request.
A single server may serve several (virtual) hosts, so the ``Host`` header signals the server to which of the possible hosts to route the request.
Some implementations might interpret a request without ``Host`` header as invalid, and respond with the ``400 Bad Request`` message.
An empty line indicates the end of the request line and header section.

The available methods are :

============= ===============
Method        Description
============= ===============
GET           Tranfers a current representation of the target resource.
HEAD          Same as GET, but only transfer the status line and header section.
POST          Perform resource-specific processing on the request payload.
PUT           Replace all current representations of the target resource with the requested payload.
DELETE        Remove all current representations of the target resource.
CONNECT       Establish a tunnel to the server identified by the target resource.
OPTIONS       Describe the communication options for the target resource.
TRACE         Perform a message loop-back test along the path to the target resource.
============= ===============

(source: :rfc:`7231`)

After the empty line, an optional message body in the request contains the payload (data) associated with the request.

Response
========

The server parses the request, and, depending on the validity of the request, sends one or more responses.
An HTTP response starts with the status line, which contains the protocol version, success or error code, and a reason message.

The status line is mostly followed by a header section containing server information and other metadata.
An empty line indicates the end of the status line and header.
The final part of the response, the message body, contains the payload sent back to the client.

.. code-block:: console

	HTTP/1.1 200 OK
	Date: Wed, 20 May 2015 19:45:37 GMT
	Server: Apache/2.2.22 (Debian)
	Last-Modified: Tue, 19 May 2015 14:25:26 GMT
	ETag: "9dac8-2d-5178e2f1c2e8b"
	Accept-Ranges: bytes
	Content-Length: 45
	Vary: Accept-Encoding
	Content-Type: text/html

	<html><body><h1>It works!</h1></body></html>

Common response codes
---------------------

============= =====================  ====================
Code	      Reason		     Description
============= =====================  ====================
200	      OK	             The request has succeeded.
301	      Moved Permenantly      The target has moved to a new location and future references should use a new location. The server should mention the new location in the ``Location`` header in the response. 
302	      Found		     The target resources been temporarily moved. As the move is temporarily, future references should still use the same location.
400	      Bad Request	     The server cannot process the request as it is perceived as invalid.
401	      Unauthorized	     The request to the target resource is not allowed due to missing or incorrect authentication credentials.
403	      Forbidden		     The request to the target resource is not allowed for reasons unrelated to authentication. 
404	      Not Found		     The target resource was not found.
500	      Internal Server Error  The server encountered an unexpected error.
502	      Bad Gateway            The server is acting as a gateway/proxy and received an invalid response from a server it contacted to fulfill the request.
503	      Service Unavailable    The server is currently unable to fulfill the request.
============= =====================  ====================



Tools: Speaking http with telnet/netcat/curl
============================================

Telnet
------

The telnet utility is used for interactive communication to a host on a given port. 
Once the connection to the remote host is established, an HTTP request can be send to the host by typing it in the prompt:

.. code-block:: console

	$ telnet www.opsschool.org 80
	GET /en/latest/http_101.html HTTP/1.1
	Host: www.opsschool.org

	Trying 162.209.114.75...
	Connected to ops-school.readthedocs.org.
	Escape character is '^]'.
	HTTP/1.1 200 OK
	Server: nginx/1.4.6 (Ubuntu)
	Date: Tue, 23 Jun 2015 20:41:10 GMT
	Content-Type: text/html
	Content-Length: 140673
	Last-Modified: Wed, 27 May 2015 12:16:25 GMT
	Connection: keep-alive
	Vary: Accept-Encoding
	ETag: "5565b599-22581"
	X-Served: Nginx
	X-Subdomain-TryFiles: True
	X-Deity: chimera-lts
	Accept-Ranges: bytes



	<!DOCTYPE html>
	<!--[if IE 8]><html class="no-js lt-ie9" lang="en" > <![endif]-->
	<!--[if gt IE 8]><!--> <html class="no-js" lang="en" > <!--<![endif]-->
	<head>
	  <meta charset="utf-8">

	  <meta name="viewport" content="width=device-width, initial-scale=1.0">

	  <title>HTTP 101 (Core protocol) &mdash; Ops School Curriculum 0.1 documentation</title>


	  (...)
  
	  </body>
	</html>

If the requested resource is found on the specified host, it is returned in the body of the response.

cURL
----

cURL is a tool and library for transferring data with URL syntax, capable of handling various protocols.
In contrast to telnet, cURL is able to create the HTTP request (or other known protocols) based on the given input.
It is a command line tool, and can be used in scripts to automate HTTP communication.

In its simplest form, cURL sends a GET request to the given target

.. code-block:: console
	
	$ curl http://www.opsschool.org/en/latest/http_101.html

	<!DOCTYPE html>
	<!--[if IE 8]><html class="no-js lt-ie9" lang="en" > <![endif]-->
	<!--[if gt IE 8]><!--> <html class="no-js" lang="en" > <!--<![endif]-->
	<head>
	  <meta charset="utf-8">

	  <meta name="viewport" content="width=device-width, initial-scale=1.0">

	  <title>HTTP 101 (Core protocol) &mdash; Ops School Curriculum 0.1 documentation</title>
  
	  (...)
  
	</body>
	</html>  

With the ``--request`` or ``-X`` parameter, the method can be specified.
To include the headers in the output, the ``-i`` can be used, and to only output the headers, use the ``-I`` switch.

.. code-block:: console
	
	$ curl -I --request GET  http://www.opsschool.org/en/latest/http_101.html
	HTTP/1.1 200 OK
	Server: nginx/1.4.6 (Ubuntu)
	Date: Tue, 23 Jun 2015 20:42:25 GMT
	Content-Type: text/html
	Content-Length: 140673
	Last-Modified: Wed, 27 May 2015 12:16:25 GMT
	Connection: keep-alive
	Vary: Accept-Encoding
	ETag: "5565b599-22581"
	X-Served: Nginx
	X-Subdomain-TryFiles: True
	X-Deity: chimera-lts
	Accept-Ranges: bytes


cURL outputs a progress meter to the terminal, showing statistics about the operation.
However, once cURL is about to write data to the terminal, the display of the progress meter is disabled so it may seem as if it was never present.
The progress meter can still be seen when the output is written ``-o <file>`` or piped ``>`` to a file.

In case the download takes some time (for example downloading a big file) the progress meter can still be seen in the terminal.
The ``-O`` switch makes cURL write output to a local file named like the remote file.

.. code-block:: console

	$ curl -O http://localhost/bigfile
	  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
        	                         Dload  Upload   Total   Spent    Left  Speed
	100  100M  100  100M    0     0   400M      0 --:--:-- --:--:-- --:--:--  414M

When starting to download a large file with ``-C`` option, the download can be resumed in case of an interrupt.

cURL can use a proxy to retrieve a resource when specified to do so using the ``-x`` or ``--proxy`` option with the proxy specified as ``[protocol://][user:password@]proxyhost[:port]``.

By default cURL does not follow HTTP redirects, and instead a 3xx redirection message is given.
When browsing ``www.opsschool.org`` with cURL, the 302 response code indicates that the requested page is temporarily residing on a different location.
The ``Location`` header contains a reference to the new location:

.. code-block:: console

	$ curl -I www.opsschool.org
	HTTP/1.1 302 FOUND
	Server: nginx/1.4.6 (Ubuntu)
	X-Deity: chimera-lts
	X-Served: Flask
	Content-Type: text/html; charset=utf-8
	Date: Tue, 23 Jun 2015 21:04:01 GMT
	Location: http://www.opsschool.org/en/latest/
	X-Redirct-From: Flask
	Connection: keep-alive
	Content-Length: 229

The ``-L`` switch makes cURL automatically follow redirects and issue another request to the given location till it finds the targeted resource.

.. code-block:: console

	$ curl -IL www.opsschool.org
	HTTP/1.1 302 FOUND
	Server: nginx/1.4.6 (Ubuntu)
	X-Deity: chimera-lts
	X-Served: Flask
	Content-Type: text/html; charset=utf-8
	Date: Tue, 23 Jun 2015 21:06:22 GMT
	Location: http://www.opsschool.org/en/latest/
	X-Redirct-From: Flask
	Connection: keep-alive
	Content-Length: 229

	HTTP/1.1 200 OK
	Server: nginx/1.4.6 (Ubuntu)
	Date: Tue, 23 Jun 2015 21:06:22 GMT
	Content-Type: text/html
	Content-Length: 195536
	Last-Modified: Wed, 27 May 2015 12:16:27 GMT
	Connection: keep-alive
	Vary: Accept-Encoding
	ETag: "5565b59b-2fbd0"
	X-Served: Nginx
	X-Subdomain-TryFiles: True
	X-Deity: chimera-lts
	Accept-Ranges: bytes


netcat
------

Netcat is a networking tool capable of reading and writing data across a network connection.
This makes it possible to use netcat both as a client and as server.
On many systems, netcat is started with the ``nc`` command instead of the program's full name.

Basic client/server
^^^^^^^^^^^^^^^^^^^

To make a basic client-server connection, netcat can be told to listen on a port using the ``-l`` option:

.. code-block:: console

	$ nc -l 192.168.0.1 1234 > output.log

In above example, IP address is specified, but this is not required.
When started without IP specified, netcat will be listening on ``0.0.0.0:1234``.
Any output is written to file output.log.

Then in another terminal, pipe some data to netcat that is connected to the destination IP and port.

.. code-block:: console

	$ echo "I'm connected as client" | nc 192.168.0.1 1234


Then again in the first (server) terminal, the data is displayed:

.. code-block:: console

	$ cat output.log
	I'm connected as client
	
Beyond sending simple pieces of text, netcat can be used to copy files and directory structures.

Serving HTTP
^^^^^^^^^^^^

Netcat can be instructed to execute a program and redirect file descriptors when a connection is established.
If compiled with DGAPING_SECURITY_HOLE option, netcat has the ``-e <program>`` option to specify which program to 'bind' to a connection.
Even when the ``-e`` option is not available, a basic HTTP server can be created by redirecting file descriptors to our program.

Consider the following Bash script which implements a simple HTTP server:

.. code-block:: Bash

	#!/usr/bin/env bash

	base_uri='/tmp'

	respond(){
        	[ -d "${base_uri}/${uri}" ] && uri+=/index.html
	        if [ -f "${base_uri}/${uri}" ]
        	        then
                	        printf '%s\r\n'\
	                                "HTTP/1.1 200 OK"\
        	                        "Date: $(date '+%a,%e %b %H:%M:%S GMT')"\
                	                "Server: myserver"\
                        	        "Content-Length: $(stat -c'%s' ${base_uri}${uri})"\
	                                "" | tee >(cat - >&2);
        	                cat <"${base_uri}/${uri}" | tee >(cat - >&2);
	                else
        	                printf '%s\r\n'\
                	                 "HTTP/1.1 404 Not Found"\
                        	        "Date: $(date '+%a,%e %b %H:%M:%S GMT')"\
	                                "Server: myserver"\
        	                        "Content-Length: 30"\
                	                ""\
                        	        "<html><b>Not Found</b></html>" | tee >(cat - >&2);
	        fi
        	unset request method uri version
	        exit 0
	}

	## MAIN ##

	read -r request
	read -r method uri version <<<"$request"
	[ -n "$method" ] && [ -n "$uri" ] && [ -n "$version" ] || echo "HTTP/1.1 400 Bad Request"
	echo $request >&2
	while read -r header
	do
        	header=${header%%$'\r'}
	        echo $header >&2
	        [ -z "$header" ] && { respond; break; }
	done


Stdout output acts as the response to the client.
Output to file descriptor 2 (``>&2``) is written to the server terminal.
This only covers reprentation of a web page, or an error page in case of a missing file.

To launch this with netcat compiled with DGAPING_SECURITY_HOLE option:

.. code-block:: console

	$ while true; do netcat -lp 1234 -e ./httpd.sh; done

Or when ``-e`` is not available:

.. code-block:: console

	$ mkfifo /tmp/httpipe
	$ while true; do cat /tmp/httpipe | ./httpd.sh | nc -l 192.168.0.1 1234 > /tmp/httpipe ; done

This creates a pipe that we use to redirect data to and from the httpd script.
Now we can test it with either of the above shown tools:

.. code-block:: console

	$ curl 192.168.0.1:1234/
	<html>Hello there world!</html>

	$ nc 192.168.0.1 1234
	GET /notexisting.html HTTP/1.1

	HTTP/1.1 404 Not Found
	Date: Tue, 9 Jun 18:39:59 GMT
	Server: myserver
	Content-Length: 30

	<html><b>Not Found</b></html>

	$ telnet 192.168.0.1 1234
	Trying 192.168.0.1...
	Connected to 192.168.0.1.
	Escape character is '^]'.
	GET /index.html HTTP/1.1

	HTTP/1.1 200 OK
	Date: Tue, 9 Jun 18:44:53 GMT
	Server: myserver
	Content-Length: 32

	<html>Hello there world!</html>
	Connection closed by foreign host.


Apache, nginx
=============

Apache and Nginx are the two most used open-source webservers on the internet.
Though they share several features, their architectures differ significantly.

Apache relies on a process-based model, either with a single or multiple threads per process that handle a connection, whereas Nginx operates event-driven with multiple single-threaded processes that all handle connections asynchronously.
Nginx' event-driven approach requires less resources, making it faster, especially for static content.
For dynamic content however, Nginx needs to proxy requests to an external processor for execution as it has no programming language support.
Therefore Nginx is often used as a front-end/proxy server in combination with other application servers, possibly Apache.

Both web servers are extensible with modules (security, compression, proxy, ...).
Where Apache is able to dynamically load and enable modules at runtime, Nginx can only load modules that have been precompiled into the executable.

HTML
====

**H**\ yper\ **T**\ ext **M**\ arkup **L**\ anguage is the markup language used to create web pages.
An HTML page is made up of a tree of elements and text.
Elements are identified by tags that are enclosed by angle brackets.
The root element is the ``<html>`` which on its turn contains the ``<head>`` element with information about the document, and ``<body>`` elements with the visible content of the page.

HTML pages can reference to other pages by using hyperlinks identified by ``<a>`` tags.
The developers of HTML, W3C, have the full specification on their website: http://www.w3.org/html/.

Virtual hosting
===============

Virtual hosting is a technique where mutiple websites are hosted on a single web server.
The server's resources are shared amongst all configured hosts, but each host has its own IP address or domain name.
There are two mechanisms for implementing virtual hosts:

IP-based
--------

In IP-based virtual hosting each virtual host is configured with its own dedicated IP address.
The web server is configured to serve multiple IP addresses, either on multiple physical network interfaces or on virtualized interfaces on a single physical interface.
Based on the IP on which a client connects, the web servers determines which of the configured virtual hosts to serve to the client. 

Name-based
----------

In name-based virtual hosting a web server serves multiple host names on a single IP address.
While a client is connecting to a webserver, the destination address is resolved to an IP address either by DNS or local lookup table.
As several virtual hosts may be hosted on the same IP, the client needs to state which website it is requesting for by setting the ``Host`` header. 
The web server inspects the ``Host`` header field in the HTTP request (mandatory as of HTTP/1.1) and routes the request to the correct virtual host.
