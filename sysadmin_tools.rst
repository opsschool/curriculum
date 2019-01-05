Tools for productivity
**********************

The work of operations is often helped by using a combination of tools. The
tools you choose to use will depend heavily on your environment and work style.
Over time you may even change the tools you use.

This is a list of commonly used tools, we recommend trying as many of them as
you have time for. Finding the right tool is often like finding a good pair of
gloves - they have to fit just right!


Terminal emulators
==================

OSX: iTerm2
-----------

Rather than stick with the default terminal, many sysadmins on OS X install
iTerm2_. iTerm2's greater flexibility and larger feature set have made it
a de facto standard among technical Mac users. Like most modern
terminals, it supports vertical and horizontal splits, tabs, profiles,
etc. In addition, it stores the rendered buffer for a period of time
which serves as a way to view previous things that might have been lost
due to editor crashes or the like.

.. _iTerm2: http://iterm2.com/

Linux: Gnome Terminator
-----------------------

Many advanced Linux sysadmins prefer tiling window managers such as i3_,
Awesome_, xmonad_, herbstluftwm_, and others. However, these window managers
can be daunting to a new user when compared to a more Windows-like desktop
environment like Unity_, KDE_, or XFCE_. A popular compromise is to use a
tiling terminal emulator, such as Terminator_, within a regular window
manager.

Terminator_ has some features which are convenient for performing similar
commands in several places at once, such as grouping terminals together and
then broadcasting your keystrokes to all terminals in the group.

.. _i3: http://i3wm.org/
.. _Awesome: http://awesome.naquadah.org/
.. _xmonad: http://xmonad.org/
.. _herbstluftwm: http://herbstluftwm.org/
.. _Unity: https://unity.ubuntu.com/
.. _KDE: http://www.kde.org/
.. _XFCE: http://www.xfce.org/
.. _Terminator: http://gnometerminator.blogspot.com/p/introduction.html

SSH
===

Common: ``ssh``
---------------
If your operating system has a good terminal emulator available (e.g., Linux, BSD,
OSX), then the ``ssh`` command line utility is by far the best tool to use for
SSH interactions.

Windows: PuTTY
--------------
If you use Windows as your primary work space, PuTTY is possibly one of the best
SSH clients available.

Android: Connectbot
-------------------
Connectbot is a free app available in the Google Play store, and offers most of
the functionality you would need from an SSH client.

iOS: iSSH
---------
A universal app available on the App Store, iSSH supports most--if not all--of
the features a demanding power user needs.

SSH Use Cases
=============
Far more than just a secure way to access a shell on a remote machine, ssh has
evolved to support a large collection of ways to encapsulate traffic over the
encrypted channel. Each of these is useful in different situations. This
section will present some common problems, and will present how to solve the
problems with ssh. It is important to note that all port forwarding and
proxying occurs over the secure ssh channel. Besides working around firewall
rules, the tunnels also provided a layer of security where there may not
otherwise be one.

Single port on foreign machine firewalled
-----------------------------------------
For this usecase, consider a loadbalancer--``lb-foo-1``--with a Web management
interface listening on port 9090. This interface is only routable to a LAN
private to the datacenter--10.10.10.0/24. Its IP address is 10.10.10.20. There
is a machine on the 10.10.10.0/24 network with ssh access--``jumphost-foo-1``,
accessible via DNS with the LAN IP 10.10.10.19. Therefore, one way to
access ``lb-foo-1`` is via bouncing through ``jumphost-foo-1``. OpenSSH
supplies the ``-L`` option to bind a local port to a port opened on the
other side of the tunnel. On the remote end, ssh opens a connection to
the host and port specified to the ``-L`` option. An example for this
usecase:

``ssh -L 9090:lb-foo-1:9090 jumphost-foo-1``

This will open a connection from ``jumphost-foo-1`` to ``lb-foo-1`` on
port 9090, and will bind the local port 9090 to a tunnel through the
connection to that port. All traffic sent to the local port 9090 will
reach ``lb-foo-1:9090``. It is important to note that the host given to
``-L`` uses the perspective of the machine ssh connects to directly.
This is important to remember for environments using hosts files or
private / split-horizon DNS. In this usecase, an invocation like the
following would work as well:

``ssh -L 9090:10.10.10.20:9090 jumphost-foo-1``

In this case, ``jumphost-foo-1`` would try to connect to ``10.10.10.20``
directly, skipping the DNS lookup.

Tunnel all traffic through remote machine
-----------------------------------------
For this usecase, consider the rest of the machines in the 10.10.10.0/24
network. There are many ways to access them, including VPNs of various
sorts. In addition to creating tunnels to specific ports, OpenSSH can
also create a SOCKS proxy. OpenSSH provides the ``-D`` option for this.
As an example, to bounce from ``jumphost-1`` to the rest of the network:

``ssh -D 9999 jumphost-foo-1``

Once the SOCKS proxy is in place, the operating system or applications
needing access to the proxy need to be configured to use it. This will
vary per-application. Some operating systems (OS X) provide system-wide
proxy configuration.

Reverse Tunneling
------------------

In some situations it may be necessary to forward ports from the remote
machine to the local one. While not as common as other port forwarding
techniques, it is often the only option available in the situations
where it sees use. For this feature, OpenSSH listens to a port on the
remote machine and sends the traffic to the local host and port. The
syntax for the ``-R`` option is the same as for the ``-L`` option
described above. As an example, to have OpenSSH tunnel the remote port
9090 to the local port 8000 on host ``workstation``:

``ssh -R 9090:workstation:8000 jumphost-foo-1``

Then, programs on ``jumphost-foo-1``--or that can access its port
9090--will be able to access the local port 8000 on ``workstation``.

Tunneling stdin and stdout
--------------------------

This method makes using a jumphost transparent. The idea is to tunnel
not a single port, but the entire connection over the ssh channel. This
usage is typically a matter of configuration in the ``ssh_config`` file,
taking advantage of its wildcard-capable ``Host`` directive::

  Host *.lan
  ProxyCommand ssh -W %h:22 jumphost-foo-1

This configuration uses the ProxyCommand feature to cooperate with the
``-W`` flag. Hostnames are resolved from the perspective of the
jumphost. In this example, the machines use an internal pseudo-top level
domain of ``.lan``. To reach, e.g., ``www-1.lan``:

``ssh www-1.lan``

Before doing DNS resolution, OpenSSH will look in its ``ssh_config``
file for Host entries. Therefore, internal DNS on the foreign end is
sufficient.


Multiplexers
============
Operations work regularly involves connecting to remote servers (there will be
times when you do nothing but work on remote servers - parts of this curriculum
were even typed on remote servers rather than contributors desktops and
laptops!).

There are however two limitations to working this way:

#. You'll often need to be connected to more than one remote system at a time.
   Opening a whole new terminal each time can result in a lot of windows cluttering
   up precious screen space.
#. What happens if your internet connection stops working? All of your
   connections are reset. Any work you might have been doing on the remote servers
   can be lost.

Multiplexers are a good solution to this.
They allow you to run multiple "virtual" windows inside a single windows.
For example:

.. epigraph::
   Bob works on 10 remote servers, all of which run Linux.
   Bob's internet connection at work is questionable.
   To work around this, Bob connects to ``server1`` which is at his data centre.
   It is a reliable server which is close to the other servers Bob works on.
   On ``server1``, Bob starts a multiplexer. The multiplexer gives Bob a regular
   looking command prompt, and Bob continues his work.

   If Bob's internet connection drops, he can reconnect to ``server1``, and then
   re-attach to the multiplexer he started previously. His session is in the
   same state he left it before being disconnected, and he can continue his
   work.

   The multiplexer also lets Bob open more than one command prompt and switch
   between them as he needs to. Bob can now connect to many servers and see them
   all in one window.

.. _gnu-screen:

GNU Screen
----------
``screen`` is one of the longest lived multiplexers. Almost everyone who has
used a multiplexer has used screen, and you can't go far wrong with it.

``screen`` is a full-screen window manager that multiplexes a physical terminal
between several processes (typically interactive shells).  It is useful for
creating sessions that can be disconnected from and reconnected to later.  This
is useful for running tasks that can take a long time that you do not want to
have an ssh session timeout on, such as a large database import.  In these cases
cron is also a very good way to run one off long running tasks.

``screen`` is also **very useful** for creating sessions that users can share.

Installation
~~~~~~~~~~~~
Debian and descendants (Ubuntu, Mint, Suse, etc):

.. code-block:: console

  apt-get install screen

On RedHat-style distributions install with the command:

.. code-block:: console

  yum install screen

Basic usage
~~~~~~~~~~~
Create a session:

.. code-block:: console

  screen -S session1

To detach from a session - in the session type Ctrl+a+d

List available screen sessions:

.. code-block:: console

  screen -ls

.. code-block:: console

  [gary@mc9 ~]# screen -ls
  There are screens on:
          21707.session2  (Detached)
          21692.session1  (Detached)
          21936.session3  (Attached)
  3 Sockets in /var/run/screen/S-gary.
  [gary@mc9 ~]#

Here we can see 3 screen sessions are running, 2 detached and 1 attached.

Reattach to a session:

.. code-block:: console

  screen -r session1

Share a session:

User Alice starts session:

.. code-block:: console

  screen -S session1

User Bob can then attach to the same session (both Alice and Bob can send commands to the session):

.. code-block:: console

  sudo screen -x alice/session1

Non root users, must use sudo to attach to another user's session.

Create a session with a log:

.. code-block:: console

  screen -L -S session1

``screen`` will output the session log to the user's home directory with the
file ``~/screenlog.0`` (0 being the session id).  PuTTY is also as a very useful
and featureful ssh client that can be used for logging ssh sessions locally
(Windows and Linux).  ``screen`` can be used within a PuTTY session.

Create a session with a log and 20000 lines of scrollback in the terminal:

.. code-block:: console

  screen -h 20000 -L -S session1


Configuration
~~~~~~~~~~~~~
``screen`` has a fairly extensive set of configuration options, when screen is invoked, it executes initialization commands from the files ``/etc/screenrc`` and ``.screenrc`` in the user's home directory.

Further info
~~~~~~~~~~~~

.. code-block:: console

  man screen

There is a nifty cheat sheet for the most important ``screen`` and ``tmux`` keybindings (see below in tmux references [3]_).

.. _tmux:

Tmux
----
``tmux`` [#]_  is relatively new compared to
``screen``. It covers the same basic feature set and has added a few
more advanced features. It is recommended you get comfortable with
``screen`` first before attempting to use ``tmux``.

In this chapter you will learn to start a tmux session, get to know a
few first keyboard shortcuts and detach from and re-attach to the
session.

Installation
~~~~~~~~~~~~
tmux is available on Debian and its descendants like Ubuntu or Mint
with the command:

.. code-block:: console

  apt-get install tmux

On RedHat-style distributions you will have to use the :term:`EPEL` repository to
get a pre-built package, and install with the command:

.. code-block:: console

  yum install tmux

On MacOS you can use Homebrew to install via:

.. code-block:: console

  brew install tmux

tmux basics
~~~~~~~~~~~
``tmux`` is usually started with the command ``tmux`` in a
terminal window. Depending of your version of tmux you will see either
a line at the bottom of the screen or nothing at all. ``tmux`` is
controlled with keyboard shortcuts, the default shortcut usually is
``ctrl-b``. If you press ``ctrl-b`` and then a ``t`` in the newly
started tmux window you should see the local time displayed as a large
digital clock. If you hit ``ctrl-b`` and ``c`` you should see a new
empty window with an empty input prompt.

If you want to detach from the session you have to hit ``ctrl-b`` and
``d``. The ``tmux`` window will disappear and you will see a message
``[detached]`` in your terminal window. All the shells and processes
you started onside the ``tmux`` session continue to run, you can see
this with a simple

.. code-block:: console

  ps -ef | grep tmux

You should see something like the following:

.. code-block:: console

  cdrexler 13751     1  0 Nov30 ?        00:00:41 tmux

You will notice that the ``tmux`` process has a parent process id of 1
which means that it is not a child process of the shell you started it
in anymore. Accordingly you can leave your working shell, start a new
one and attach to the running tmux process again which is very handy
if your connectivity is flaky or you have to work from different
locations. If you check the process table for the process id of the
tmux process

.. code-block:: console

  ps -ef | grep 13751

you will find that is the parent process of the two shells you created
in the beginning of the chapter:

.. code-block:: console

   cdrexler  4525 13751  0 17:54 pts/2    00:00:00 -zsh
   cdrexler  4533 13751  0 17:54 pts/5    00:00:00 -zsh

If you want to get an overview of the running tmux processes on your
system you can use the command

.. code-block:: console

  tmux ls

It will list all available ``tmux`` sessions on your system [#]_. If there
is only one you can attach to it with the command:

.. code-block:: console

  tmux att

If there is more than one session the output of ``tmux ls`` will look like this:

.. code-block:: console

   0: 3 windows (created Fri Nov 30 18:32:37 2012) [80x38]
   4: 1 windows (created Sun Dec  2 17:44:15 2012) [150x39] (attached)

You will then have to select the right session with the ``-t`` command line switch:

.. code-block:: console

  tmux att -t 4

``tmux`` runs as a server process that can handle several sessions so
you should only see one tmux process per user per system.

You should see the original session with the two shells again after
running this command.

tmux configuration
~~~~~~~~~~~~~~~~~~~
``tmux`` is configured via a
config file which is usually called :file:`.tmux.conf` that should live in
your ``$HOME`` directory.

A typical :file:`.tmux.conf` looks like this:

.. code-block:: ini

   #set keyboard shortcut to ctrl-g
   unbind C-b
   set -g prefix C-g
   bind C-g send-prefix
   bind g send-prefix
   #end of keybord shortcut setting
   # Highlight active window
   set-window-option -g window-status-current-bg red
   # Set window notifications
   setw -g monitor-activity on
   set -g visual-activity on
   #automatically rename windows according to the running program
   setw -g automatic-rename
   #set scroll back buffer
   set -g history-limit 10000
   set -g default-terminal "xterm-256color"
   set -g base-index 1
   set -g status-left '#[fg=green]#H

This illustrates a method to change the default keybinding and some
useful settings.

Please note that you can force ``tmux`` to use another configfile with
the ``-f`` command line switch like so:

.. code-block:: console

  tmux -f mytmuxconf.conf

There is a nifty cheat sheet [#]_ for the most important
``screen`` and ``tmux`` keybindings or even a whole book about tmux [#]_.



byobu
-----
[Byobu]_ is a wrapper around one of screen or tmux. It
provides profile support, F-keybindings, configuration utilities and a
system status notification bar for most Linux, BSD or Mac operating systems.

Byobu is available in major distros as a packaged binary. Launching byobu will
run whichever the package maintainer included as a dependency, if you have both
installed, you can select explicitly with byobu-{screen,tmux}.  Basic
configuration is launched with F9 or byobu-config.

Scrollback starts with F7, mark the start of your buffer by hitting the
spacebar, then use the spacebar again to mark the end of your selection (which
will be copied into byobu's clipboard automatically). Press Byobu-Escape
(typically Ctrl-A) + [ to paste.

References
----------
.. [#] http://tmux.sourceforge.net/
.. [#] Please note that ``tmux ls`` will *only* list tmux sessions that belong to your userid!
.. [#] http://www.dayid.org/os/notes/tm.html
.. [#] http://pragprog.com/book/bhtmux/tmux
.. [#] https://launchpad.net/byobu
.. [Byobu] http://byobu.co/

Shell customisations
====================

As you read in :doc:`shells_101`, your shell is your primary tool during the
work day. It's also incredibly customisable to suit your needs. Let's look at
some changes you can make.

How to customise your shell
---------------------------

Your shell's configuration is stored in its ``rc`` file. For bash, this file is
``~/.bashrc``. Each time you edit this, you can reload the configuration by
typing:

.. epigraph::
   ``source ~/.bashrc``

Changing your prompt
--------------------

Your default prompt probably looks something like this:

.. epigraph::
   ``bash-3.2$``

That's pretty plain and doesn't tell you much. In fact, all it does tell you is
that you're using Bash version 3.2, and that you are not the root user (the
``$`` at the end signifies a regular user, whereas if you were root, you would
see a ``#`` instead).

Let's change this up a little. Edit your ``~/.bashrc`` file, and add this line
to the end:

.. epigraph::
   ``PS1="\u@\h \w> "``

Save, quit, and then reload your ``.bashrc`` file. Your prompt should change to
something like this:

.. epigraph::
   ``avleen@laptop ~>``

Much better! Now your know your username, the name of the machine you're on (in
this case "``laptop``"), and the directory you're in ("``~``" is your home
directory).

The ``PS1`` variable has a lot of different options you can use to customise it
further.


Mosh
====

Mosh (MObile SHell) is an alternative to remote shell commands, such as ``ssh``
or ``rsh``. The beauty of the Mosh protocol is that it supports intermittent
connectivity without losing the remote session.

You can start a ``mosh`` session just like you would with ``ssh`` on one side
of town with one IP address, shut your laptop and go home, then open your
laptop and connect back to your Mosh session like it was never interrupted.
Also, if your wifi is spotty or internet connection is intermittent, ``mosh``
doesn't break the session when the connection drops out, unlike ``ssh``.
Mosh does not wait for the remote server to confirm each keystroke before
displaying it to a terminal. Instead, it displays the typed characters locally
and confirms entry on the remote end. There are packages available for GNU/Linux,
FreeBSD, Solaris, Mac OS X, Chrome and even Android apps.

Mosh must be installed on both the client and remote server. When a session
is started it spins up a mosh-server and a local mosh-client process. It can
be installed in a home directory without privileged access. Mosh respects
your current ``~/.ssh/config`` so migrating from ``ssh`` to ``mosh`` is
relatively seamless.

SSH connections work by sending a stream of data back and forth between the client
and server. This stream can be broken in various ways, such as the connection
timing out due to inactivity, or the client machine suspending state and shutting
down network devices. The Mosh protocol is based on UDP packets compared to the SSH
protocol that uses TCP packets. Mosh is the first application to use the Stateless
Syncronization Protocol. Instead of a stream of data between the server and client
that makes up a session over SSH, Mosh works by keeping a mirrored copy of the
session on the client and server and syncronizing the changes between them.

While ``ssh`` has been around long enough to have time tested security,
``mosh`` is relatively new and has not been through the same extensive testing.
It is the first application to use the SSP protocol. Mosh does tunnel traffic
encrypted with AES-128 in OES mode, however Mosh hasn't been under the
security spotlight as long as SSH has.

Examples
--------

Mosh works just like ssh:

.. code-block:: console

  mosh username@remoteserver.org

You can also have mosh utilize ssh style commands such as: This specifies the
private key to use to authenticate with the remove server.

.. code-block:: console

  mosh username@remoteserver.org --ssh="ssh -i ~/.ssh/identity_file"

This tells mosh to connect via the ssh port 1234 on the remote server, where
ssh normally runs on port 22.

.. code-block:: console

  mosh username@remoteserver.org --ssh="ssh -p 1234"``

References
----------

.. http://mosh.mit.edu


Ticketing systems
=================


Note-taking
===========

Wiki
----

EverNote
--------

OneNote
-------
