# Reasons for concurrency and parallelism


To complete this exercise you will have to use git. Create one or several commits that adds answers to the following questions and push it to your groups repository to complete the task.

When answering the questions, remember to use all the resources at your disposal. Asking the internet isn't a form of "cheating", it's a way of learning.

 ### What is concurrency? What is parallelism? What's the difference?
 > Concurrency refers to when a computer can run several programs at the same time. Loading several pages at once is an example of this, and so is a webpage serving multiple users at the same time.
 > Parallelism refers to when the computer makes several computations at the same time, increasing the speed of the system.
 
 ### Why have machines become increasingly multicore in the past decade?
 > Since the development of speed on one single core is not booming in the same way it used to, more cores have been added to allow parallelism and thus increasing computation speed.
 
 ### What kinds of problems motivates the need for concurrent execution?
 (Or phrased differently: What problems do concurrency help in solving?)
 > Most programs on smart-phones, personal computers and embedded systems require multiple programs running at once. An example is an ATM that needs to service multiple users at once or the mobile phone combining computations in the phone itself and the cloud.
 
 ### Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
 (Come back to this after you have worked on part 4 of this exercise)
 > Concurrent programs are subject to complex bugs, and are difficult to write in a good way. However, they are necessary for a lot of processes.
 
 ### What are the differences between processes, threads, green threads, and coroutines?
 > Processes run in different memory spaces, while threads share memory. Green threads use a virtual machine, and a thread run on the OS
 
 ### Which one of these do `pthread_create()` (C/POSIX), `threading.Thread()` (Python), `go` (Go) create?
 > A new thread.
 
 ### How does pythons Global Interpreter Lock (GIL) influence the way a python Thread behaves?
 > It makes it such that only one thread can execute at the time. It is necessary because python's memory managing is not thread safe.
 
 ### With this in mind: What is the workaround for the GIL (Hint: it's another module)?
 > There are packages such as thread that allow for using threads in python.
 
 ### What does `func GOMAXPROCS(n int) int` change? 
 > The funciton sets the amounts of CPUs that can execute at the same time.
