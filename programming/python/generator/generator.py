def gen():
    print('start')
    for i in range(3):
        yield i
        print('continue')
    print('end')



print(type(gen()), type(gen))
"""
Any function that includes a key work yield called generator function
"""

for i in gen():
    print(i)

print('='*10)
g = iter(gen())
print(next(g))
print(next(g))
print(next(g))


print('='*10)
g = gen()
print(next(g))
print(next(g))
print(next(g))


"""
Generator expression.

(1 for i in range(10))
"""

ge = (a for a in gen())
gel = [a for a in gen()]
print(type(ge))

from collections import abc
print('isinstance(ge, abc.Iterator) = ', isinstance(ge, abc.Iterator))
print('isinstance(ge, abc.Iterable) = ', isinstance(ge, abc.Iterable))


it = iter(ge)
print(next(it))
print(next(it))
print(next(it))
try:
    print(next(it))
except StopIteration:
    print('stop iteration')



def chain(*iterable):
    for it in iterable:
        for i in it:
            yield i

s, t = 'abc', range(3)
print(list(chain(s, t)))

def chain_improved(*iterable):
    for it in iterable:
        yield from it

print(list(chain_improved(s, t)))





def try_all():
    print('try_any True')
    yield True
    print('try_any True')
    yield True
    print('try_any False')
    yield False
    print('try_any True should not go here!!!')
    yield True



print(all(try_all()))


print(sorted(try_all()))
# print(reversed(try_all()))



import random
def d6():
    return random.randint(1, 6)


d6_iter = iter(d6, 6)
for i in d6_iter:
    print(i)
