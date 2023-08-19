import threading
import random
import time


class Publisher(threading.Thread):
    def __init__(self, integers, condition):
        self.condition = condition
        self.integers = integers
        super().__init__()
    

    def run(self):
        while True:
            integer = random.randint(0, 100)
            self.condition.acquire()
            print('Acquire')
            self.integers.append(integer)
            self.condition.notify()
            print('Notify')
            self.condition.release()
            print('Release')
            time.sleep(1)



class Subscriber(threading.Thread):
    def __init__(self, integers, condition):
        self.integers = integers
        self.condition = condition
        super().__init__()

    def run(self):
        while True:
            self.condition.acquire()
            print(f'Condition Acquire by Subscriber: {self.name}')

            print('Start waiting', self.name)
            self.condition.wait()
            print(f'Condition Wait by Subscriber: {self.name}')
                        
            if self.integers:
                print('popped', self.integers.pop())
            

            self.condition.release()
            print(f'Condition Release by Subscriber: {self.name}')


if __name__ == '__main__':
    integers = []

    condition = threading.Condition(lock=threading.Lock())
    pub = Publisher(integers, condition)
    pub.start()

    s1 = Subscriber(integers, condition)
    s2 = Subscriber(integers, condition)

    s1.start()
    s2.start()

    pub.join()
    s1.join()
    s2.join()

