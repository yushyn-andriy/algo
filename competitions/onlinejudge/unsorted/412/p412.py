import math
from itertools import combinations

import sys


stdin = sys.stdin


def gcd(a, b):
    if a%b == 0:
        return b
    return gcd(b, a%b)


def estimate(numbers):
    no_common_divisors = 0
    total_combinations = 0
    for a in range(len(numbers)):
        for b in range(a+1, len(numbers)):
            total_combinations += 1
            if gcd(numbers[a], numbers[b]) == 1:
                no_common_divisors += 1

    if no_common_divisors == 0:
        return -1

    return math.sqrt((6 * total_combinations) / no_common_divisors)


if __name__ == '__main__':
    for row in stdin:
        n = int(row.strip())
        if n <= 0 or n>=50:
            break
        
        numbers = []
        for _ in range(n):
            numbers.append(int(stdin.readline().strip()))
        
        stdin.readline()


        pi = estimate(numbers)
        if pi == -1:
            print('No estimate for this data set.')
        else:
            print(f'{pi:0.6f}')
