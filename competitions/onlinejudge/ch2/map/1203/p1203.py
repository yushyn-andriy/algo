import sys
import heapq
from collections import namedtuple



stdin = sys.stdin



Query = namedtuple('Query', 'id period')


if __name__ == '__main__':
    queries = []
    for row in stdin:
        row = row.strip()
        if '#' in row:
            n = int(stdin.readline().strip())
            break
        q = Query(*list(map(int, row.split()[1:])))
        queries.append(q)

    heap_smallest_to_largest = []
    heap_largest_to_smallest = []
    
    queries.sort(key=lambda x: x.period)
    n_queries = 0
    for i in range(1, n+1):
        for q in queries:
            if len(heap_smallest_to_largest) >= n:
                largest = abs(heap_largest_to_smallest[0][0])
                if q.period * i > largest:
                    break
            heapq.heappush(heap_largest_to_smallest, (-q.period * i, q.id))
            heapq.heappush(heap_smallest_to_largest, (q.period * i, q.id))


    while n_queries < n:
        print(heapq.heappop(heap_smallest_to_largest)[1])
        n_queries += 1
