import sys
from collections import defaultdict


stdin = sys.stdin


class Query:
    def __init__(self, array) -> None:
        self.array = array
        self.registry = defaultdict(list)

    def process(self):
        for i, number in enumerate(self.array):
            self.registry[number].append(i)
    
    def query(self, k, v):
        if v not in self.registry:
            return 0
        
        reg = self.registry[v]
        k = k - 1
        if k >= len(reg):
            return 0
        
        return reg[k] + 1




if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        n, m = list(map(int, row.split()))
        array = list(map(int, stdin.readline().strip().split()))

        q = Query(array)
        q.process()

        for _ in range(m):
            k, v = list(map(int, stdin.readline().strip().split()))
            print(q.query(k, v))

