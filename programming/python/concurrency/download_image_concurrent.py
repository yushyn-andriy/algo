import threading
import random
import time


import urllib.request


# API_ADDRESS = 'http://localhost:8100/image.png'
API_ADDRESS = 'https://www.google.com/'

def download_image(url, file):
    print('Downloading image from ', url)
    return urllib.request.urlretrieve(url, file)


def execute_thread(i, stop_event):
    while not stop_event.is_set():
        print(f'active thread {threading.current_thread()}')
        image_name = f'./temp/image_{i}.jpg'
        download_image(API_ADDRESS, image_name)



if __name__ == '__main__':
    threads = []

    for i in range(8):
        event = threading.Event()
        thread = threading.Thread(target=execute_thread, args=(i, event), name=str(i))
        threads.append((thread, event))
        thread.start()

    print(f'active threads: {threading.active_count()}')
    print(f'enumerate threads: {threading.enumerate()}')

    thread, event = random.choice(threads)
    event.set()
    
    time.sleep(1)
    for i in threads:
        i[1].set()
