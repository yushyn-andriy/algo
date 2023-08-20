import sys
from math import sin, cos


stdin = sys.stdin


def approximation(p, a, b, c, d, k):
    return p * (sin(a * k + b) + cos(c*k + d)+ 2)


def solve(p, a, b, c, d, n):
    max_price_so_far = None
    max_difference = 0.0
    for k in range(1, n+1):
        value = approximation(p, a, b, c, d, k)
        if max_price_so_far is None:
            max_price_so_far = value
            continue

        if value > max_price_so_far:
            max_price_so_far = value
        else:
            difference = max_price_so_far - value
            max_difference = max(max_difference, difference)

    return max_difference


if __name__ == '__main__':
    for row in stdin:
        p, a, b, c, d, n = list(map(int, row.strip().split(' ')))
        print('%.6f' % solve(p, a, b, c, d, n))
