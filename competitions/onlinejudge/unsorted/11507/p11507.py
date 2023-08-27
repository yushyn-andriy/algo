# import sys
# from collections import namedtuple


# stdin = sys.stdin


# Vector = namedtuple('Vector', 'x,y,z')


# def get_vector_direction(v):
#     return {
#         Vector(1, 0, 0): '+x',
#         Vector(-1, 0, 0): '-x',
#         Vector(0, 1, 0): '+y',
#         Vector(0, -1, 0): '-y',
#         Vector(0, 0, 1): '+z',
#         Vector(0, 0, -1): '-z',
#     }.get(v, 'UNKNOWN')


# def rotate_vector(v, d):
#     if d == 'No':
#         return v
#     elif d == '+y':
#         return Vector(-v.y, v.x, v.z)
#     elif d == '-y':
#         return Vector(v.y, -v.x, v.z)
#     elif d == '-z':
#         return Vector(v.z, v.y, -v.x)
#     elif d == '+z':
#         return Vector(-v.z, v.y, v.x)
#     else:
#         raise NotImplementedError
#     return v


# def bend(length, ops):
#     _ = length
#     v = Vector(1, 0, 0)
#     for direction in ops:
#         v = rotate_vector(v, direction)
#     return get_vector_direction(v)




# if __name__ == '__main__':
#     for row in stdin:
#         if row == '0':
#             break
#         length = int(row)
#         ops = stdin.readline().strip().split()
#         print(bend(length, ops))


import sys
from collections import namedtuple

stdin = sys.stdin

Vector = namedtuple('Vector', 'x,y,z')

def get_vector_direction(v):
    return {
        Vector(1, 0, 0): '+x',
        Vector(-1, 0, 0): '-x',
        Vector(0, 1, 0): '+y',
        Vector(0, -1, 0): '-y',
        Vector(0, 0, 1): '+z',
        Vector(0, 0, -1): '-z',
    }[v]

def rotate_vector(v, d):
    if d == 'No':
        return v
    elif d == '+y':
        rotations = {
            Vector(1, 0, 0): Vector(0, 1, 0),
            Vector(-1, 0, 0): Vector(0, -1, 0),
            Vector(0, 1, 0): Vector(-1, 0, 0),
            Vector(0, -1, 0): Vector(1, 0, 0),
            Vector(0, 0, 1): Vector(0, 0, 1),
            Vector(0, 0, -1): Vector(0, 0, -1)
        }
    elif d == '-y':
        rotations = {
            Vector(1, 0, 0): Vector(0, -1, 0),
            Vector(-1, 0, 0): Vector(0, 1, 0),
            Vector(0, 1, 0): Vector(1, 0, 0),
            Vector(0, -1, 0): Vector(-1, 0, 0),
            Vector(0, 0, 1): Vector(0, 0, 1),
            Vector(0, 0, -1): Vector(0, 0, -1)
        }
    elif d == '+z':
        rotations = {
            Vector(1, 0, 0): Vector(0, 0, 1),
            Vector(-1, 0, 0): Vector(0, 0, -1),
            Vector(0, 1, 0): Vector(0, 1, 0),
            Vector(0, -1, 0): Vector(0, -1, 0),
            Vector(0, 0, 1): Vector(-1, 0, 0),
            Vector(0, 0, -1): Vector(1, 0, 0)
        }
    elif d == '-z':
        rotations = {
            Vector(1, 0, 0): Vector(0, 0, -1),
            Vector(-1, 0, 0): Vector(0, 0, 1),
            Vector(0, 1, 0): Vector(0, 1, 0),
            Vector(0, -1, 0): Vector(0, -1, 0),
            Vector(0, 0, 1): Vector(1, 0, 0),
            Vector(0, 0, -1): Vector(-1, 0, 0)
        }
    return rotations[v]


def bend(ops):
    v = Vector(1, 0, 0)
    for direction in ops:
        v = rotate_vector(v, direction)
    return get_vector_direction(v)

if __name__ == '__main__':
    while True:
        length = int(stdin.readline().strip())
        if length == 0:
            break
        ops = stdin.readline().strip().split()
        print(bend(ops))
