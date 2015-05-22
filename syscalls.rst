Syscalls
********

What are syscalls?

Applications require actions from the kernel to perform several tasks it cannot handle by itself.
Accessing files, connecting to sockets, creating processes are examples of tasks handled by the kernel.
System calls are the interface for an application to interact with the kernel.

Using ``strace``
================

The ``strace`` tool can be used to monitor interaction between a process and the kernel. 
It displays system calls initiated by the process, and signals received by the process.

Consider the simple script:

.. code-block:: bash

	#!/bin/bash
	echo "one line of text" > my_input.txt
	while read line
	do
        	echo $line > my_output.txt
	done < my_input.txt

Executing
---------

When invoking this script using ``strace -omy_strace.log -ff ./strace_sample.sh``, the output of every (forked) process is written to strace.log.<pid>.
Let's examine the logfile (parts of the output are omitted, let's focus on the relevant stuff) :

.. code-block:: console

	execve("./strace_sample.sh", ["./strace_sample.sh"], [/* 17 vars */]) = 0
	<..>
	open("./strace_sample.sh", O_RDONLY|O_LARGEFILE) = 3
	<..>
	read(3, "#!/bin/bash\necho \"one line of te"..., 80) = 80
	<..>
	dup2(3, 255)                            = 255
	close(3)                                = 0
	<..>
	read(255, "#!/bin/bash\necho \"one line of te"..., 126) = 126
	<..>

I/O
---

``execve`` executes the file "strace_sample.sh" with arguments "./strace_sample.sh" (by convention the first should be the filename to be executed) and 17 environmental variables.
Then ``open`` opens the file for reading in file descriptor 3.
Next, ``read`` attempts to read the next 80 bytes from fd (file descriptor) 3.
At this point the interpreter line of the script is read, indicating that the /bin/bash binary needs to be invoked to process the script.
``dup2`` copies the fd 3 to the (new) fd 255.
This is a bash-specific operation, don't mind to much; bash keeps the original file in fd 255 (last of the process' private fd) to free up low-numbered fd's.
Now that our original file is in fd 255, fd 3 is not needed anymore and can be closed by ``close``.
At last the next 126 (rest of the file) bytes are read and stored in the bufferi, now we can start to process the commands in the script file.

.. code-block:: console

	open("my_input.txt", O_WRONLY|O_CREAT|O_TRUNC|O_LARGEFILE, 0666) = 3
	<..>
	fcntl64(1, F_DUPFD, 10)                 = 10
	<..>
	dup2(3, 1)                              = 1
	close(3)                                = 0
	write(1, "one line of text\n", 17)      = 17
	dup2(10, 1)                             = 1
	# Copy fd 10 to fd 1 => revert the redirection, this is why the original value of fd 1 was saved
	<..>
	close(10)                               = 0
	# Close fd 10

First, ``open`` the my_input.txt file in fd 3.
Then, ``fcntl64`` uses F_DUPFD (10) to get the next available fd numbered >=10 and copy fd 1 to the new fd 10.
This saves the original content of fd 1, which is stdout.
Next ``dup2`` copies fd 3 to fd 1 (fd 1 which by the process is still used for stdout) and closes fd 3.
The ``write`` call writes the next 17 bytes to fd 1 (which it sees as stdout, but at this moment it points to my_input.txt).
Afterwards the redirection is reverted, by copying fd 10 (original value for fd 1) back to fd 1.
fd 1 now points to stdout as desinged by default.
fd 10 is closed as its not needed any longer.

.. code-block:: console

	open("my_input.txt", O_RDONLY|O_LARGEFILE) = 3
	<..>
	fcntl64(0, F_DUPFD, 10)                 = 10
	<..>
	dup2(3, 0)                              = 0
	close(3)                                = 0
	<..>
	read(0, "one line of text\n", 128)      = 17
	open("my_output.txt", O_WRONLY|O_CREAT|O_APPEND|O_LARGEFILE, 0666) = 3
	<..>
	fcntl64(1, F_DUPFD, 10)                 = 11
	# Copy fd 0 to fd 11 (which is not the lowest available fd >= 10) to save the original stdout
	<..>
	dup2(3, 1)                              = 1
	close(3)                                = 0
	write(1, "one line of text\n", 17)      = 17
	# Write 17 bytes to fd 1 (which is stdout, redirected to my_output.txt)
	dup2(11, 1)                             = 1
	<..>
	close(11)                               = 0
	read(0, "", 128)                        = 0
	dup2(10, 0)                             = 0
	<..>
	close(10)  
 
Again, ``open`` my_input.txt in fd 3.
This time, save fd 0 (by default stdin) to fd 10.
``dup2`` copies fd 3 to fd 0 (redirecting my_input.txt to stdin) and close fd 3.
Next, read the next 128 bytes from fd 0 (my_input.txt) and save to the buffer.
Next ``open`` "my_output.txt" in fd 3.
Then ``fcntl64`` uses F_DUPFD (10) to get the next available fd >= 10 (which at this point is 11 as fd 10 is already open) and copy fd 1 to it.
Redirect stdout to my_output.txt by copying fd 3 to fd 1 with ``dup2``.
Now fd 3 can be closed.
Finally, write 17 bytes from the buffer to fd 1.
The redirection is reverted by copying fd 11 to fd 1 with ``dup2``, and fd 11 can be closed.
A next attempt to ``read`` from fd 0 is done, resulting in 0 bytes read, indicating the end of file is reached.
The redirection is reverted by copying fd 10 to fd 0 and closing fd 10.

``exec``, ``open``, ``close``, ``read`` and ``write`` are handled. Let's look at creating child processes and removing files.

Child processes
---------------

.. code-block:: console

	clone(child_stack=0, flags=CLONE_CHILD_CLEARTID|CLONE_CHILD_SETTID|SIGCHLD, child_tidptr=0xb6f50068) = 3482

The parent process uses ``clone`` to create a child process to execute the ``rm`` command.
The logging of this child process is logged in the second my_strace.log.<pid> file, where in this example pid=3482, but this varies on each run.

.. code-block:: console

	execve("/bin/rm", ["rm", "my_input.txt"], [/* 17 vars */]) = 0
	<..>
	newfstatat(AT_FDCWD, "my_input.txt", {st_mode=S_IFREG|0644, st_size=17, ...}, AT_SYMLINK_NOFOLLOW) = 0
	faccessat(AT_FDCWD, "my_input.txt", W_OK) = 0
	unlinkat(AT_FDCWD, "my_input.txt", 0)   = 0
	close(0)                                = 0
	close(1)                                = 0
	close(2)                                = 0
	exit_group(0)                           = ?

The ``rm`` command is executed using ``execve``, with arguments "rm" (as per convention this is the filename to be executed) and "my_input.txt".
``newfstatat`` gets the file status and ``faccessat`` check the file permissions of the file.
Finally, ``unlinkat` removes the file's name from the filesystem.
If that name was the last link to a file and no processes have the file open the file is deleted and the space it was using is made available for reuse.
As a last step for this process the 3 standard fd's are closed, and ``exit_group`` exits all possible threads in the process. 

Again in the parent's logfile, the interaction with the child process is logged.

.. code-block:: console

	wait4(-1, [{WIFEXITED(s) && WEXITSTATUS(s) == 0}], 0, NULL) = 3481 
	<..>
	--- SIGCHLD (Child exited) @ 0 (0) ---
	wait4(-1, 0xbec6bf39, WNOHANG, NULL)    = -1 ECHILD (No child processes)
	<..>	
	read(255, "", 142)                      = 0
	exit_group(0)                           = ?

``wait4`` keeps the parent process waiting for the child process to terminate.
Once terminated, wait releases and the parent process continues.
A final ``read`` is attempted on fd 255, but as the end of the file is reached, this returns 0.
The last exit_group exits all open threads in the process.

Output tuning
-------------

By default the ``strace`` produces all system calls performed by the executable.
As this can be overwhelming, the -e switch can be used to look for specific system calls.
When examining this with ``-eopen`` the following is output is given:

.. code-block:: console

	strace -eopen ls
	open("/etc/ld.so.cache", O_RDONLY)      = 3
	open("/lib64/librt.so.1", O_RDONLY)     = 3
	open("/lib64/libacl.so.1", O_RDONLY)    = 3
	open("/lib64/libselinux.so.1", O_RDONLY) = 3
	open("/lib64/libc.so.6", O_RDONLY)      = 3
	open("/lib64/libpthread.so.0", O_RDONLY) = 3
	open("/lib64/libattr.so.1", O_RDONLY)   = 3
	open("/lib64/libdl.so.2", O_RDONLY)     = 3
	open("/lib64/libsepol.so.1", O_RDONLY)  = 3
	open("/etc/selinux/config", O_RDONLY)   = 3
	open("/proc/mounts", O_RDONLY)          = 3
	open(".", O_RDONLY|O_NONBLOCK|O_DIRECTORY) = 3

This can come in handy to troubleshoot specific system calls.
A list of available syscalls can be seen in ``man syscalls``.
For more details on a syscall, look it up in the man page.
Some syscalls have several variant and might be referenced in strace output with different names; try to look them up without certain prefixes to find the relevant man pages.
For performance reasons the ``-T`` and ``-c`` flags are usefull:

-T Show the time spent in system calls. This records the time difference between the beginning and the end of each system call.
-c Count  time, calls, and errors for each system call and report a summary on program exit. 


