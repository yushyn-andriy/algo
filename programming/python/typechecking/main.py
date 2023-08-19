

def add(a: int, b: int) -> int:
    return a + b

def unadd(a, b):
    return a + b


class A:
    ...


add(1, 2)
add('a', 'b')
add(3.1, 0.04)

unadd(1, 2)
unadd('a', 'b')
reveal_type(unadd(3.1, 0.04))
reveal_type(1)

a: type[A] = A

