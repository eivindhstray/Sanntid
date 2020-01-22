# Python 3.3.3 and 2.7.6
# python fo.py
<<<<<<< HEAD
from threading import Thread
=======

from threading import Thread, Lock
>>>>>>> Anders


# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
i = 0

my_lock = Lock()

def incrementingFunction():
    global i
<<<<<<< HEAD
    for j in range (1000000):
        i+=1
    # TODO: increment i 1_000_000 times

def decrementingFunction():
    global i
    for j in range (1000000):
        i-=1
    # TODO: decrement i 1_000_000 times
=======

    for x in range (1000000):
        my_lock.acquire()
        i = i+1
        my_lock.release()

def decrementingFunction():
    global i
    for k in range(1000000):
        my_lock.acquire()
        i = i-1
        my_lock.release()
>>>>>>> Anders



def main():
    # TODO: Something is missing here (needed to print i)
    global i

    incrementing = Thread(target = incrementingFunction, args = (),)
    decrementing = Thread(target = decrementingFunction, args = (),)

    # TODO: Start both threads
<<<<<<< HEAD
    
=======
>>>>>>> Anders
    incrementing.start()
    decrementing.start()

    incrementing.join()
    decrementing.join()

    print("The magic number is %d" % (i))


main()
