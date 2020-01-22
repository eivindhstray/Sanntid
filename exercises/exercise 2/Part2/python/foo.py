# Python 3.3.3 and 2.7.6
# python fo.py

from threading import Thread, Lock

# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")
i = 0

my_lock = Lock()

def incrementingFunction():
    global i

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



def main():
    global i

    incrementing = Thread(target = incrementingFunction, args = (),)
    decrementing = Thread(target = decrementingFunction, args = (),)

    # TODO: Start both threads
    incrementing.start()
    decrementing.start()

    incrementing.join()
    decrementing.join()

    print("The magic number is %d" % (i))


main()
