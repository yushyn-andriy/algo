import sys
from itertools import combinations



stdin = sys.stdin




def _perm(l, c, k):
    if len(c) == k:
        print(' '.join(map(str, c)))
        return

    new = l[:]
    for _, x in enumerate(l):
        new.remove(x)
        c.append(x)

        _perm(new, c, k)
        c.pop()




def perm(l, k):
    _perm(l, [], k)





if __name__ == '__main__':
    lines = stdin.readlines()
    for i, row in enumerate(lines):
        row = row.strip()
        if row == '0':
            break
        
        k, *L = list(map(int, row.split()))

        perm(L, 6)
        if i == len(lines) - 2:
            continue
        print()