import sys
from functools import cmp_to_key


stdin = sys.stdin



def cmp(x, y):
    if x == y:
        return 0

    nx = len(x)
    ny = len(y)

    s = 0
    while s<=max(nx,ny) *2:
        i = s % nx
        j = s % ny
        if x[i] != y[j]:
            return -(ord(x[i]) - ord(y[j])) 
        s+=1
    return 0

if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        n = int(row)
        if n == 0:
            break
        numbers = [x for x in stdin.readline().strip().split()]

        numbers= sorted(numbers, key=cmp_to_key(cmp))
        print(''.join(numbers))
