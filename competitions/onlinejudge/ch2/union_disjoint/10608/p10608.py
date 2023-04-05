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
    
    if ranks[x] >= ranks[y]:
        parents[y] = x
        ranks[x] += ranks[y]
    else:
        parents[x] = y
        ranks[y] += ranks[x]
    
    return True




def solution(parents):
    counts = defaultdict(int)
    for person in parents.keys():
        p = find(person, parents)
        counts[p] += 1
    return max(counts.values())


if __name__ == '__main__':
    c = int(stdin.readline().strip())
    for _ in range(c):
        row = stdin.readline().strip()
        n, m = list(map(int, row.split()))
        if n == 0 and m == 0:
            break

        parents = {}
        ranks = [0 for _ in range(n+1)]
        for i in range(1, n+1):
            parents[i] = -1

        for i in range(m):
            x, y = list(map(int, stdin.readline().strip().split()))
            union(x, y, parents, ranks)

        count = solution(parents)
        print(f'{count}')
