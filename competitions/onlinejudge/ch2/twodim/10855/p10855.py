import sys
from copy import deepcopy


stdin = sys.stdin


def count_current_appears(bsq, ssq, N, n):
    total = 0
    i = 0
    while i <= N - n:
        j = 0
        while j <= N - n:
            eq = True
            for k in range(n):
                if eq is False:
                    break
                for l in range(n):
                    if bsq[i+k][j+l] != ssq[k][l]:
                        eq = False
                        break
            if eq is True:
                total += 1
            j += 1
        i += 1
    return total


def rot90(ssq, n):
    _ssq = deepcopy(ssq)
    for i in range(n):
        for j in range(n):
            _ssq[j][n - i - 1] = ssq[i][j]
    return _ssq


def get_appears(bsq, ssq, N, n):
    apps = []
    for _ in range(4):
        apps.append(count_current_appears(bsq, ssq, N, n))
        ssq = rot90(ssq, n)
    return apps

def to_str(arr):
    return ' '.join([str(x) for x in apps])


if __name__ == '__main__':
    while True:
        N, n = [int(x) for x in stdin.readline().strip().split()]
        if N == 0 and n == 0:
            break
        
        bsq = []
        ssq = []
        for _ in range(N):
            bsq.append(list(stdin.readline().strip()))
        for _ in range(n):
            ssq.append(list(stdin.readline().strip()))
        apps = get_appears(bsq, ssq, N, n)    
        print(to_str(apps))
