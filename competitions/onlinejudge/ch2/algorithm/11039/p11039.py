import sys
from collections import namedtuple


Floor = namedtuple('Floor', 'size color')


stdin = sys.stdin


def solve(floors):
    n = 0
    if len(floors) < 1:
        return n
    
    n+=1
    floors = sorted(floors, key=lambda x: -x.size)
    prev = floors[0]
    for curr in floors[1:]:
        if prev.color != curr.color:
            n+=1
            prev = curr
        else:
            continue

    return n


if __name__ == '__main__':
    c = int(stdin.readline().strip())
    for _ in range(c):
        n = int(stdin.readline().strip())
        floors = []
        for _ in range(n):
            size = int(stdin.readline().strip())
            color = 0 if size < 0 else 1
            size = abs(size)
            floors.append(Floor(size, color))
        print(solve(floors))