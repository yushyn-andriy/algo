from typing import Any
import functools


class A:
    def __call__(self, *args: Any, **kwds: Any) -> Any:
        print('__call__')

A()()


def f(person, bid):
    pass


diff = set(dir(f)) - set(dir(A()))


# for attr in diff:
#     print(f'{attr}={getattr(f, attr)}')

print(f.__code__.co_varnames, f.__code__.co_argcount)



def decorator(f):
    @functools.wraps(f)
    def wrapped(data):
        args = []
        for varname in f.__code__.co_varnames:
            t = f.__annotations__[varname] 
            v = data.get(varname, None)
            if not isinstance(v, t):
                raise TypeError(
                    f'wrong type: argument "{varname}" is not of a type {t}',
                )
            args.append(v)
        res = f(*args)
        if not isinstance(res, f.__annotations__['return']):
            raise TypeError(
                f'wrong return type',
            )
        return res
    return wrapped



@decorator
def f(person: int, bid: int) -> str:
    print(person, bid)
    # return f'{person} + {bid}'
    return 1


request = {'person': 1, 'bid': 2, 'another_param': 3}
f(request)

print(f.__annotations__)