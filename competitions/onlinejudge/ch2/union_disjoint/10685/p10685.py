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


def eat():
    stdin.readline()

if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        n, m = list(map(int, row.split()))
        if n == 0 and m == 0:
            break

        parents = {}
        ranks = {}
        for i in range(1, n+1):
            creature = stdin.readline().strip()
            parents[creature] = -1
            ranks[creature] = 1

        for i in range(m):
            x, y = stdin.readline().strip().split()
            union(x, y, parents, ranks)
    
        eat()

        count = solution(parents)
        print(f'{count}')
