import sys
from collections import defaultdict, deque


stdin = sys.stdin

def bfs(root, graph, visited):
    d = deque([root])
    while d:
        node = d.popleft()
        for nbr in graph[node]:
            if nbr not in visited:
                d.append(nbr)
                visited.add(nbr)


def solution(n, graph):
    count = 0
    visited = set()
    for i in range(1, n+1):
        if i not in visited:
            bfs(i, graph, visited)
            count += 1
    return count

if __name__ == '__main__':
    for j, row in enumerate(stdin):
        row = row.strip()
        n, m = list(map(int, row.split()))
        if n == 0 and m == 0:
            break

        graph = defaultdict(set)
        visited = set()
        for i in range(1, n+1):
            graph[i]

        for i in range(m):
            x, y = list(map(int, stdin.readline().strip().split()))
            graph[x].add(y)
            graph[y].add(x)

        count = solution(n, graph)
        print(f'Case {j+1}: {count}')
