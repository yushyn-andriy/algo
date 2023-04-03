import sys


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
    count = 0
    for v in parents.values():
        if v == -1:
            count += 1
    return count


if __name__ == '__main__':
    for j, row in enumerate(stdin):
        row = row.strip()
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
        print(f'Case {j+1}: {count}')
