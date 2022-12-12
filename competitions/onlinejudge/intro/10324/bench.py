import timeit



def cache(f):
    _data = {}
    def wrapped(i, j, string):
        b, e = min(i, j), max(i, j)
        for (cb, ce), value in _data.items():
            if b >= cb and e <= ce and value:
                _data[(b, e)] = True
                return True

        r = f(b, e, string)
        _data[(b, e)] = r
        return r
    return wrapped


@cache
def same_in_range(i, j, string):
    b, e = i, j

    same = True
    ch = string[b]
    while b <= e and b < len(string):
        if string[b] != ch:
            same = not same
            break
        b += 1
    return same


if __name__ == '__main__':
    
    string = "1"*1000000
    l = ["1"]*1000000
    print(same_in_range(0, 999, string))
    print(timeit.timeit('same_in_range(0, 9999, string)', globals=globals(), number=100))
    print(timeit.timeit('same_in_range(0, 9999, l)', globals=globals(), number=100))
