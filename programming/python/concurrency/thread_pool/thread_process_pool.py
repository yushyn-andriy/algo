from concurrent.futures import (
    ThreadPoolExecutor,
    ProcessPoolExecutor,
    Future,
    as_completed,
)
import threading

import requests


def load(url='https://google1.com'):
    print(f'{threading.current_thread}: loading...')
    r = requests.get(url)
    print(f'{threading.current_thread}: done')
    return r

if __name__ == '__main__':

    with ThreadPoolExecutor(max_workers=10) as tpe:
    # with ProcessPoolExecutor(max_workers=10) as tpe:
        results: list[Future] = list()
        for i in range(100):
            r = tpe.submit(load, 'https://google.com')
            # r.result()
            results.append(r)
        
        for r in as_completed(results):
            print(r.result())

        # print('end... ', results[0].done(), results[0].result(), results[0].done())


