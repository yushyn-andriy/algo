import sys
from functools import cmp_to_key


stdin = sys.stdin


def is_even(x):
    return x % 2 == 0

def is_odd(x):
    return x%2 ==1 


def custom_sorted(numbers, m):
    def cmp(x, y):

        mx = x % m
        my = y % m
        if x < 0:
            mx = -(abs(x)%m)
        if y < 0:
            my = -(abs(y)%m)
        if mx != my:
            return mx - my
        if is_odd(x):
            if is_odd(y):
                return -(x - y)
            return -1
        if is_odd(y):
            return 1
        return x - y
    return sorted(numbers, key=cmp_to_key(cmp))



if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        n, m = [int(x) for x in row.split()]
        if n == 0:
            print('0 0')
            break

        numbers = []
        for i in range(n):
            numbers.append(int(stdin.readline().strip()))
        
        print(n, m)
        numbers = custom_sorted(numbers, m)
        for v in numbers:
            print(v)
