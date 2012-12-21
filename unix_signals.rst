Signals
*******

Signals are a form of inter-process communication (IPC) in Posix compliant 
systems like Linux, unix, and unix-like operating systems. They are an 
asynchronis mechanism the kernel uses to communicate with a process or 
a thread. Signals are always delivered by the kernel but can be initiated 
by the kernel, the process itself, or even another process.  

When a process recieves a signal, what happens depends on the program. In
the simplest case the process will just apply the default signal handlers
defined in the C library. In most cases this will cause the program to 
terminate, suspend, and occasionally it will ignore the signal. 

In some cases, programs have custom signal handlers. Programs that have 
custom signal handlers will behave differently when they recieve a signal.
For example many service daemons will reload their configurations when they 
recieve the SIGHUP signal; the default action for SIGHUP is for the program
to terminate.  

In either case, when the kernel sends a signal to a process, the process 
will execute its signal handler imediately. Processes can enter states
in which they are executing atomic instructionsi(like file locks), and 
will not process the signal until the instruction is completed. 

What are signals? How do they work?

