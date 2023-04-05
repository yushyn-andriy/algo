import sys
from collections import defaultdict

stdin = sys.stdin



def find(i, parents):
    if parents[i] == -1:
        return i
    parents[i] = find(parents[i], parents)
    return parents[i]


def union(a, b, parents, ranks):
    x = find(a, parents)
    y = find(b, parents)

    if x == y:
        print(ranks[x])
        return False
    
    if ranks[x] >= ranks[y]:
        parents[y] = x
        ranks[x] += ranks[y]
        print(ranks[x])
    else:
        parents[x] = y
        ranks[y] += ranks[x]
        print(ranks[y])
    
    return True


if __name__ == '__main__':
    c = int(stdin.readline().strip())
    for _ in range(c):
        m = int(stdin.readline().strip())

        parents = defaultdict(lambda:-1)
        ranks = defaultdict(lambda : 1)
        for i in range(m):
            x, y = stdin.readline().strip().split()
            union(x, y, parents, ranks)
