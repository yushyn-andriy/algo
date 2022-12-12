import sys
import math


stdin = sys.stdin


def max(array):
    m = -math.inf
    for x in array:
        if x > m:
           m = x
    return m 


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for i in range(1, n+1):
        speeds = [int(x) for x in stdin.readline().strip().split()][1:]
        m = max(speeds)
        print(f'Case {i}: {m}')
