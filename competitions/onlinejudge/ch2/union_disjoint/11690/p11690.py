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
        return False
    
    parents[y] = x

    # main part
    ranks[x] += ranks[y]
    ranks[y] = 0

    return True


if __name__ == '__main__':
    c = int(stdin.readline().strip())
    for _ in range(c):
        n, m = list(map(int, stdin.readline().strip().split()))

        ranks = defaultdict(lambda : 0)
        for i in range(n):
            value = int(stdin.readline().strip())
            ranks[i] = value


        parents = defaultdict(lambda:-1)
        for i in range(m):
            x, y = list(map(int, stdin.readline().strip().split()))
            union(x, y, parents, ranks)
                

        possible = True
        for v in ranks.values():
            if v != 0:
                possible = False
                break

        if possible:
            print('POSSIBLE')
        else:
            print('IMPOSSIBLE')


