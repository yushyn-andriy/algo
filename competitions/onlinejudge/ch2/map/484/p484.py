import sys
from collections import OrderedDict



cin = sys.stdin


def unique(row, m):
    for i in row:
        if i not in m:
            m[i] = 1
        else:
            m[i] += 1
    return m


if __name__ == '__main__':
    m = OrderedDict()
    for row in cin:
        row = row.strip().split()
        unique(row, m)
    for k, v in m.items():
        print(k, v)
