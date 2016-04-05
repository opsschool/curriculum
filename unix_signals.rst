Signals
*******

What Are Signals?
=================

Signals are a form of inter-process communication (IPC) in Posix compliant systems like Linux, UNIX, and UNIX-like operating systems.
They are an asynchronous mechanism the kernel uses to communicate with a process or a thread.
Signals are always delivered by the kernel but can be initiated by the kernel, the process itself, or even another process.
Signals are often referred to by their name, or their numeric integer value.
For example, the kill signal is known as SIGKILL, or 9.
There are many different signals that can be sent, although the signals in which users are generally most interested are SIGTERM and SIGKILL.
The default signal sent is SIGTERM.


How Do Signals Work?
====================

When a process receives a signal, what happens depends on the program.
In the simplest case the process will just apply the default signal handlers defined in the C library.
In most cases this will cause the program to terminate, suspend, and occasionally it will ignore the signal.

In some cases, programs have custom signal handlers.
Programs that have custom signal handlers will behave differently when they receive a signal.
For example many service daemons will reload their configurations when they receive the SIGHUP signal; the default action for SIGHUP is for the program to terminate.

In either case, when the kernel sends a signal to a process, the process will execute its signal handler immediately.
Processes can enter states in which they are executing atomic instructions(like file locks), and will not process the signal until the instruction is completed.

How To Send a Signal to a Program
=================================

There are three ways to send a signal to another process.
The simplest way is the execute the "kill" command with a signal specified.
For example you can use the kill command from the shell to send the interrupt signal like so:

.. code-block:: console

  kill -SIGINT <PID>

You can write a simple program executing the kill system call.
A basic example is below:

.. code-block:: c

  int signal_pid(int pid) {
    int rvalue;
    rvalue = kill(pid, SIGINT);
    return rvalue;
  }


Lastly you can send signals from the keyboard in an interactive terminal.
Ctrl-C will send SIGINT, and Ctrl-Z send SIGTSTP.


List Of Posix 1990 Signals
==========================

======= ========= ======= =======================================================================
Signal  Value     Action  Comment
======= ========= ======= =======================================================================
SIGHUP  1         Term    Hangup detected on controlling terminal or death of controlling process
SIGINT  2         Term    Interrupt from keyboard
SIGQUIT 3         Core    Quit from keyboard
SIGILL  4         Core    Illegal Instruction
SIGABRT 6         Core    Abort signal from abort(3)
SIGFPE  8         Core    Floating point exception
SIGKILL 9         Term    Kill signal
SIGSEGV 11        Core    Invalid memory reference
SIGPIPE 13        Term    Broken pipe: write to pipe with no readers
SIGALRM 14        Term    Timer signal from alarm(2)
SIGTERM 15        Term    Termination signal (signal sent by default by the kill command when not specified)
SIGUSR1 30,10,16  Term    User-defined signal 1
SIGUSR2 31,12,17  Term    User-defined signal 2
SIGCHLD 20,17,18  Ign     Child stopped or terminated
SIGCONT 19,18,25  Cont    Continue if stopped
SIGSTOP 17,19,23  Stop    Stop process
SIGTSTP 18,20,24  Stop    Stop typed at tty
SIGTTIN 21,21,26  Stop    tty input for background process
SIGTTOU 22,22,27  Stop    tty output for background process
======= ========= ======= =======================================================================
