import sys
from collections import defaultdict




stdin = sys.stdin


if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        if not row:
            continue

        nodes = int(row)
        edges = int(stdin.readline().strip())

        waked = set(stdin.readline().strip())
        graph = defaultdict(set)
    
        waked_up = False
        for _ in range(edges):
            x, y = list(stdin.readline().strip())
            graph[x].add(y)
            graph[y].add(x)


        y = 0
        while len(waked) != nodes:
            y += 1
            completed = []
            for k, v in graph.items():
                cnt = 0
                for nbr in v:
                    if nbr in waked:
                        cnt += 1
                if cnt >= 3 or k in waked:
                    completed.append(k)
            if not completed:
                break
            for v in completed:
                waked.add(v)
                del graph[v]

        if len(waked) == nodes:
            print(f'WAKE UP IN, {y}, YEARS')
        else:
            print('THIS BRAIN NEVER WAKES UP')

