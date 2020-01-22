#Exercise 1

# Mutex and Channel basics

### What is an atomic operation?
> An atomic operation is an operation which completes in a single step. An example can be int a = 2;
> read and write operations are typically atomic.

### What is a semaphore?
> A semaphore is a signaling system using non-negative variables. It allows and disallows access to resources, with the aim to synchronize tasks.

### What is a mutex?
> A mutex allows one thread to proceed, such that while one is doing work the other has to wait until it is allowed to continue. Only one thread can work with >the entire memory.

### What is the difference between a mutex and a binary semaphore?
> Using mutex, all other threads than the one running are refused to do so. 
> Mutex is a locking mechanism, whereas a semaphore is a signaling mechanism. 

### What is a critical section?
> A critical seciton is a section of code that must not be run by multiple threads at the same time. This could ruin the results.

### What is the difference between race conditions and data races?
 > A race condition is a situation where the order of previous results can vary, potentially producing wrong results.
 > Data race is when two threads (or more) share a variable, potentially producing the wrong result.

### List some advantages of using message passing over lock-based synchronization primitives.
> according to technopedia, message passing is the process of sending a signal to a process. An example can be sending messages between physical nodes, where > both the sender and receiver both need a send - and receive function to communicate both ways.

### List some advantages of using lock-based synchronization primitives over message passing.
> *Your answer here*
