import threading
import random
import time

THREADS_COUNT = 10
counter = 0

lock = threading.Lock()
rlock = threading.RLock()


class Thread(threading.Thread):
    def __init__(self, *args, i=100, **kwargs):
        super().__init__(*args, **kwargs)
        self.i = i
    
    def run(self):
        global counter
        while self.i > 0:
            # # lock.acquire()
            # counter += 1
            # # lock.release()
            # self.i -= 1
            with lock:
                temp = counter
                time.sleep(0.000001)  # A tiny sleep to amplify the race condition
                counter = temp + 1
                self.i -= 1


UPTO = 100000

if __name__ == '__main__':
    threads = []
    for i in range(THREADS_COUNT):
        thread = Thread(i=UPTO)
        
        threads.append(thread)
        thread.start()
    
    for i in threads:
        i.join()

    print(f'counter: {counter}')
    assert counter == THREADS_COUNT * UPTO
