import sys


DEGRIES_PER_TURN = 360
DEGRIES_PER_MARK = 9


def turn_until(cursor, stop, clockwise=True):
    total = 0
    while cursor != stop:
        total += DEGRIES_PER_MARK
        if clockwise:
            cursor = (cursor - 1) % 40
        else:
            cursor = (cursor + 1) % 40
    return total


def get_degree(a, b, c, d):
    total = 0
    total += DEGRIES_PER_TURN * 3
    total += turn_until(a, b)
    total += turn_until(b, c, False)
    total += turn_until(c, d, True)
    
    return total


if __name__ == '__main__':
    lines = sys.stdin.readlines()
    for line in lines:
        a, b, c, d = [int(i) for i in line.split()]
        if (a, b, c, d) == (0, 0, 0, 0):
            break
        degree = get_degree(a, b, c, d)
        print(degree)