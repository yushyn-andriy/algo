import multiprocessing

import time
import os



def deamon():
    pid = os.getpid()
    i = 0
    while True:
        with open('log.txt', 'a') as f:
            print(f'pid: {pid}, doing some important work...', i, file=f)
        print('doing some important work...', i)
        time.sleep(2)
        i += 1



if __name__ == '__main__':
    deamon_process = multiprocessing.Process(target=deamon)
    deamon_process.daemon = True
    deamon_process.start()

    deamon_process1 = multiprocessing.Process(target=deamon)
    deamon_process1.daemon = True
    deamon_process1.start()

    print(f'deamon pid={deamon_process.pid}')
    print(f'deamon pid={deamon_process1.pid}')
    print('doing another things...')
    time.sleep(30)
    deamon_process.terminate()
    time.sleep(10)
    print('doing another things...')
