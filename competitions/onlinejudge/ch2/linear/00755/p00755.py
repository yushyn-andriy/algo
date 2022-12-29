import sys
from collections import OrderedDict



stdin = sys.stdin


def standard_form(s):
    m = {
    'ABC': '2',
    'DEF': '3',
    'GHI': '4',
    'JKL': '5',
    'MNO': '6',
    'PRS': '7',
    'TUV': '8',
    'WXY': '9',
    '-': '',
    }
    m1 = {
    'A': '2',
    'B': '2',
    'C': '2',

    'D': '3',
    'E': '3',
    'F': '3',

    'G': '4',
    'H': '4',
    'I': '4',

    'J': '5',
    'K': '5',
    'L': '5',

    'M': '6',
    'N': '6',
    'O': '6',

    'P': '7',
    'R': '7',
    'S': '7',

    'T': '8',
    'U': '8',
    'V': '8',

    'W': '9',
    'X': '9',
    'Y': '9',

    '-': '',
    }
    stf = s.lower()
    res = ''

   # O(n)
    for ch in stf.upper():
        if ch in m1:
            res += m1[ch]
        elif ch == '-':
            continue
        else:
            res += ch

    return f'{res[:3]}-{res[3:]}'


if __name__ == '__main__':
    t = int(stdin.readline().strip())
    for j in range(t):
        counter = {}
        counter_gt_1 = {}
        gt_1 = False
        input()
        n = int(input())
        for i in range(n):
            l = input().strip()
            sf = standard_form(l)

            cc = counter.get(sf, 0) + 1
            if cc > 1:
                gt_1 = True
                counter_gt_1[sf] = cc
            counter[sf] = cc

        if gt_1:
            for k in sorted(counter_gt_1.keys()):
                v = counter_gt_1[k]
                if v > 1:
                    print(k, v)
            if j < t - 1:
                print()
        else:
            print('No duplicates.')
            if j < t - 1:
                print()

