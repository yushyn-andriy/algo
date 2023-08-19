import inspect

def simple_coroutine():
    print('coroutine started')
    x = yield 123
    print('coroutine received ->', x)



mycoroutine = simple_coroutine()
print(mycoroutine, inspect.getgeneratorstate(mycoroutine))


# print(next(mycoroutine))
mycoroutine.send(None)
print(mycoroutine, inspect.getgeneratorstate(mycoroutine))

# mycoroutine.send(1)
# print(mycoroutine, inspect.getgeneratorstate(mycoroutine))



def avareger():
    '''
    >>> avg = avareger()
    >>> next(avg)
    >>> avg.send(10)
    10.0
    >>> avg.send(5)
    7.5
    '''
    count = 0
    total = 0.0
    average = None
    while True:
        term = yield average
        count += 1
        total += term
        average = total / count


avg = avareger()
avg.send(None)
for i in range(11):
    print(i, avg.send(i))


def add(a, b):
    """
    >>> add(1, 2)
    3
    >>> add(3, 2)
    5
    >>> add(1, 1)
    2
    """
    return a + b


class DemoException(Exception):
    pass


def demo_exc_handle():
    print('starting')
    try:
        while True:
            try:
                x = yield
            except DemoException:
                print('demo exception continue')
            else:
                print('received')
    finally:
        print('should be run in every case')
    raise RuntimeError('not run')


dd = demo_exc_handle()
next(dd)
dd.send(1)
dd.send(2)
dd.throw(DemoException)
dd.send(3)
dd.close()



def avareger_impr():
    count = 0
    total = 0.0
    average = None
    while True:
        term = yield
        if term is None:
            break
        count += 1
        total += term
        average = total / count
    return average


avg = avareger_impr()
next(avg)
print(avg.send(1))
print(avg.send(2))
print(avg.send(3))
try:
    print(avg.send(None))
except StopIteration as e:
    print(e.value)


if __name__ == '__main__':
    import doctest
    doctest.testmod(verbose=False)
