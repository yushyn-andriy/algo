import sys
from collections import defaultdict


stdin = sys.stdin


def parent(arr, i):
    if arr[i] == -1:
        return i
    return parent(arr, arr[i])



def disjoint_union_find(arr, i, j):
    idx1 = parent(arr, i)
    idx2 = parent(arr, j)
    if idx1 == idx2:
        return True
    arr[idx1] = idx2
    return False


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for i in range(n):
        joint = defaultdict(lambda:-1)
        for row in stdin:
            row = row.strip()
            if '*' in row:
                break
            L = list(row)
            p1, p2 = L[1], L[3]
            disjoint_union_find(joint, p1, p2)

        alpabet = stdin.readline().strip().split(',')

        roots = set()
        potential_acorns = set()
        for leaf in alpabet:
            p = parent(joint, leaf)
            if p == leaf:
                potential_acorns.add(leaf)
                continue
            roots.add(p)

        acorns = 0
        for leaf in potential_acorns:
            if leaf not in roots:
                acorns += 1


        print(f'There are {len(roots)} tree(s) and {acorns} acorn(s).')
