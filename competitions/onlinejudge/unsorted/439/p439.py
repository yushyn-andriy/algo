import sys
from collections import deque, defaultdict



stdin = sys.stdin



mmap = {ch: i+1 for i, ch in enumerate('abcdefgh')}
mmap.update({ch: i+1 for i, ch in enumerate('12345678')})

row = [ 1,   2,  2,  1,  -1,  -2, -2, -1]
col = [-2,  -1,  1,  2,  -2,  -1,  1, 2]

def parse(c1, c2):
    return [mmap[ch] for ch in c1+c2]


def solution(x1, y1, x2, y2, dist, pred):
    count = 0


    visited = defaultdict(bool)


    d = deque([(x1, y1)])
    visited[(x1, y1)] = True
    dist[(x1, y1)] = 0


    while d:
        x1, y1 = d.popleft()

        for i in range(len(row)):
            y = y1 + row[i]
            x = x1 + col[i]

            if x >= 1 and x <= 8 and y >= 1 and y <= 8:
                if visited[(x, y)] == False:
                    visited[(x, y)] = True
                    dist[(x, y)] = dist[(x1, y1)] + 1
                    pred[(x, y)] = (x1, y1)
                    d.append((x, y))
                    if x == x2 and y == y2:
                        return True

    return False
    

if __name__ == '__main__':
    for r in stdin:
        c1, c2 = r.strip().split()
        x1, y1, x2, y2 = parse(c1, c2)
        dist = defaultdict(lambda : 10000)
        pred = defaultdict(lambda : -1)
        solution(x1, y1, x2, y2, dist, pred)

        
        print(f'To get from {c1} to {c2} takes {dist[(x2, y2)]} knight moves.')

