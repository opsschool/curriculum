Logs 101
********

Common system logs & formats
============================

There are generally three places where a process may be expected to send logs: stderr, a file, or syslog.

Standard Error
==============

stderr is a file descriptor opened automatically (along with stdin and stdout) for any new process on a Unix system that is intended as a simple way for developers to provide feedback to a user without assuming any fixed format.
stderr is always exposed to a process as file descriptor 2.
If a process is launched from a shell attached to a tty, stderr usually appears inline with the rest of the program's output.
In the shell, stderr may be redirected in order to differentiate it from stdout or to send it to a file.

The `find` command often writes messages to stderr upon encountering permissions issues.
We'll use it as an example.

.. code-block:: console

    $ mkdir noperms
    $ chmod 000 noperms/
    $ find noperms/
    noperms/
    find: `noperms/': Permission denied

The name of the directory `noperms/` is first written to stdout, then a permission denied error is written to stderr.
Suppose we expected a large number of these errors to occur and wanted to append them in a log file, separate from the info written to stdout.
Remember: stderr is always file descriptor 2.

.. code-block:: console

    $ find noperms/ 2>>find.log
    noperms/
    $ cat find.log
    find: `noperms/': Permission denied

It is important to note that we used the append `>>` operation here, rather than a single `>` which would overwrite find.log every time this command was run.
This sort of redirection is a feature of the shell that is used quite commonly in init scripts and cron tasks to capture the output of both short lived commands and long running daemons.

Log files
=========

The authors of many daemons add the ability to log directly to a file, rather than depending on the shell to provide redirection of stderr.
There is no standard format or convention for a daemon to write it's own logs beyond opening a file and appending lines to it.
Some daemons provide extremely flexible log configuration whereas others might only allow you to set a file path.

The `tail` command's `--follow` or `-f` flag is an extremely useful tool for viewing new lines appending to a file in realtime.
For example, if you were running a web server like nginx or Apache that was configured to send it's access log to a file, you could use tail to see new requests.

.. code-block:: bash

    $ tail -f /var/log/nginx/access.log
    127.0.0.1 - - [17/Aug/2013:02:53:42 -0700] "GET / HTTP/1.1" 200 151 "-" "curl/7.22.0 (x86_64-pc-linux-gnu) libcurl/7.22.0 OpenSSL/1.0.1"
    127.0.0.1 - - [17/Aug/2013:02:53:43 -0700] "GET / HTTP/1.1" 200 151 "-" "curl/7.22.0 (x86_64-pc-linux-gnu) libcurl/7.22.0 OpenSSL/1.0.1"
    127.0.0.1 - - [17/Aug/2013:02:53:43 -0700] "GET / HTTP/1.1" 200 151 "-" "curl/7.22.0 (x86_64-pc-linux-gnu) libcurl/7.22.0 OpenSSL/1.0.1"

With the `-f` option, tail won't exit on it's own, it will continue to wait for new lines to be written to the log file and write them to the terminal until it receives a signal or encounters an error.

It is important to understand that after a file is opened for writing, the process writing the file only refers to that file by its file handle, which is a number assigned by the kernel.
If that file is renamed with `mv` or deleted with `rm`, writes to that file handle will still succeed.
This can sometimes lead to counter-intuitive situations where log messages are being written to a file that's been renamed for archival or to an inode that no longer has a filename associated with it.
Some daemons provide mechanisms for closing and reopening their log files upon receiving a signal like SIGHUP but quite a few don't.

Syslog
======

Logging directly to files can add a lot of complexity to an application as close attention has to be paid to the use of file handles, log directory permissions, timezones, etc.
Syslog was created to provide a simple logging interface to application developers while offloading the tasks of sorting and storing logs to a separate daemon.

Every message sent to syslog has three pieces of information: facility, priority, and message.
The facility tells syslog the type of information in the message.
For example, login attempts are logged to the "auth" facility.
The priority tells syslog the importance of the message.
Syslog provides a fixed set of facilities and priorities that a program may use.
Some examples of facilities are: auth, user, daemon, kern.
Examples of priorities are debug, info, warn, critical.

All modern Linux distributions ship with a syslog daemon and most of them are pre-configured to write messages to various files in `/var/log/`, depending on their facility or priority.
While the exact names of these files are not consistent across different Linux distributions, a few common ones like `/var/log/auth.log` and `/var/log/kern.log` almost always exist.
If you haven't already, take a look at the files in `/var/log/` on your system to get a sense of the types of log messages available in these files.

Example log from a mail server (redacted to protect the guilty)::

    $ tail /var/log/mail.log
    Aug 17 10:25:42 front01 postfix/smtpd[25791]: connect from unknown[115.77.XXX.XXX]
    Aug 17 10:25:43 front01 postfix/smtpd[25791]: NOQUEUE: reject: RCPT from unknown[115.77.XXX.XXX]: 554 5.7.1 Service unavailable; Client host [115.77.XXX.XXX] blocked using truncate.gbudb.net; http://www.gbudb.com/truncate/ [115.77.XXX.XXX]; from=<supsi@yahoo.com> to=<synack@XXX.XXX> proto=ESMTP helo=<XXX.viettel.vn>
    Aug 17 10:25:43 front01 postfix/smtpd[25791]: lost connection after RCPT from unknown[115.77.XXX.XXX]
    Aug 17 10:25:43 front01 postfix/smtpd[25791]: disconnect from unknown[115.77.XXX.XXX]

One of the advantages of using a syslog daemon is that the format of log lines can be configured in a single place and standardized for all services using syslog on a single host.
In this example, every line starts with a timestamp, the server's hostname, the name of the program and a PID.
While the name of the program is set when the connection to syslog is first opened, the rest of these fields are generated by the syslog daemon itself and added to every line.

Many different syslog implementations exist with a variety of configuration mechanisms and design philosophies.
Most current Linux distributions ship with a syslog daemon that implements some superset of the original Unix syslogd's functionality.
The following examples will use rsyslogd, which is currently included in Ubuntu Linux and according to it's manpage "is derived from the sysklogd package which in turn is derived from the stock BSD sources."

/etc/rsyslog.d/50-default.conf (truncated)::

    #
    # First some standard log files.  Log by facility.
    #
    auth,authpriv.*                 /var/log/auth.log
    *.*;auth,authpriv.none          -/var/log/syslog
    #cron.*                         /var/log/cron.log
    #daemon.*                       -/var/log/daemon.log
    kern.*                          -/var/log/kern.log
    #lpr.*                          -/var/log/lpr.log
    mail.*                          -/var/log/mail.log
    #user.*                         -/var/log/user.log

    #
    # Logging for the mail system.  Split it up so that
    # it is easy to write scripts to parse these files.
    #
    #mail.info                      -/var/log/mail.info
    #mail.warn                      -/var/log/mail.warn
    mail.err                        /var/log/mail.err

Lines beginning with `#` are comments.
Each line has two fields: a filter and an action.
The filter is a comma-separated list of `facility.priority` where `*` is used as a wildcard, meaning match anything.
The action is commonly just the name of a file, which causes all messages that match the filter to be written to that file.
The actions in this example starting with `-` invert the behavior and cause messages matching the filter to be excluded from the given file.
Many other flags and options are available here for configuring the exact behavior of log formatting and writing to places other than files.

As soon as you start to manage more than a couple of servers, you start to think about ways to aggregate the logs from all of those servers in a single place so that you don't have to login to each one individually to find an issue.
Remote log aggregation is also often used to provide an audit trail for security events or a source of data that can be fed into a metrics system like Graphite or Ganglia.
There is a standard protocol for sending syslog events over a network to another host over UDP port 514.

Configuring remote logging::

    auth.*                          @10.0.0.2

As UDP is connectionless and makes no delivery guarantees, syslog messages sent to a remote host using this standard protocol can be dropped, delayed, or intercepted without any real indication to the user.
For these reasons, many syslog daemons implement different extensions and mechanisms for transporting this stream reliably.
The simplest option is to replace UDP with TCP to provide a reliable transport layer.
When configuring syslog aggregation, attention and care should be paid to security as syslog messages are often used as an audit trail and need to be protected against eavesdropping and manipulation.
Read your syslog daemon's documentation to understand what options are supported.

Log Rotation
============

No matter which logging option you choose, logging directly to files or using syslog, log files grow large and unwieldy over time and become difficult to use, for example identifying specific events.

To handle this problem, log files are rotated on a regular basis, by making a copy of the current log files and creating fresh log files.
The old logs can be archived, compressed, mailed to an address or removed at predetermined intervals.

The `Logrotate <https://fedorahosted.org/logrotate/>`_ application eases management of systems that generate large numbers of log files.
It allows automatic rotation, compression, removal, and mailing of log files.
The log files may be handled at intervals (daily, weekly and monthly) or when they grow too large.
It is usually scheduled to run daily.

Everything about the log files to be handled by logrotate as well as the actions to be carried out on them is read from the logrotate configuration files.
The main configuration file is ``/etc/logrotate.conf``.
Applications can also create configuration files in the ``/etc/logrotate.d`` directory, logrotate automatically includes all configuration files in this directory.

.. code-block:: none

    # sample logrotate configuration file
    compress

    /var/log/messages {
        rotate 5
        weekly
        postrotate
            /sbin/killall -HUP syslogd
        endscript
    }

    "/var/log/httpd/access.log" /var/log/httpd/error.log {
        rotate 5
        mail foo@bar.org
        size=100k
        sharedscripts
        postrotate
             /sbin/killall -HUP httpd
        endscript
    }

Lines beginning with ``#`` are comments and can appear anywhere in the configuration file.
The first few lines set global options.
In this example logs are compressed after rotation.

The next section defines how to handle the log file ``/var/log/messages``.
The log file is rotated weekly and removed after going through 5 rotations.
The postrotate option defines a command to execute after the log file is rotated but before it is compressed.
Postrotate is usually used to force daemons to reload their configurations so they will log to the new log file.

Log file names can be quoted or not quoted. Quoting allows matching file names with spaces in them.
The second section defines how to handle two files, ``/var/log/httpd/access.log`` and ``/var/log/httpd/error.log``.
These logs files are rotated when they grow over 100k in size.
The old log log files are mailed to foo@bar.org (uncompressed) after going through 5 rotations.
The sharedscripts options means that the command for postrotate should be run only once no matter how many log files match.
In this case although two files are handled, the command ``/sbin/killall`` is executed once.

There a lot more options available for logrotate, you can get a full list by checking the logrotate man page.
