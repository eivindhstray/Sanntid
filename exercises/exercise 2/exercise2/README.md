# Avoiding data races

In this part you will solve [the concurrent access problem from Exercise 1](https://github.com/TTK4145/Exercise1/tree/master/Part4#part-4-finally-some-code), such that the final result is always zero. You can choose to either use the provided starter code or to copy your solution to Exercise 1 into the suitable directories. In the go-channel part you will probably wish to use the provided starter code as no equivalent exists in Exercise 1.

In your solution, make sure that the two threads intermingle. Running them one after the other would somewhat defeat the purpose. It may be useful to change the number of iterations in one of the threads, such that the expected final result is not zero (say, -1). This way it is easier to see that your solution actually works, and isn't just printing the initial value.


### C

 - POSIX has both mutexes ([`pthread_mutex_t`](http://pubs.opengroup.org/onlinepubs/7990989775/xsh/pthread.h.html)) and semaphores ([`sem_t`](http://pubs.opengroup.org/onlinepubs/7990989775/xsh/semaphore.h.html)). Which one should you use?
 - Acquire the lock, do your work in the critical section, and release the lock


### Go
Using shared variable synchronization is not the idiomatic go. You should instead create a server that [`select{}`](http://golang.org/ref/spec#Select_statements)s transformations to its own data. Have two other goroutines tell the server to increment & decrement its local variable. Note that this variable will no longer be global, so it cannot be read by other goroutines. The proper way to handle this is to create another `select{}`-case where others can request a copy of the value.

Before attempting to do the exercise, have a look at the following chapters of the interactive go tutorial:
- [Goroutines](https://tour.golang.org/concurrency/1)
- [Channels](https://tour.golang.org/concurrency/2)
- [Select](https://tour.golang.org/concurrency/5)


Remember from Exercise 1 where we had no good way of waiting for a goroutine to finish? Try sending a "finished"/"worker done" message on a separate channel. If you use different channels for the two threads, you will have to use [`select { case... }`](http://golang.org/ref/spec#Select_statements) so that it doesn't matter what order they arrive in.


In the theory part you listed some advantages with message passing and shared variable synchronization. Do you still agree with what you answered?

### Python

Python provides both [Locks](http://docs.python.org/2.7/library/threading.html#lock-objects) (which are like mutexes) and [Queues](http://docs.python.org/2/library/queue.html) (which are kind of like channels). Solve the problem in whatever way you believe to be the simplest.
