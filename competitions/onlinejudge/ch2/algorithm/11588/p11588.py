import sys
from collections import defaultdict

stdin = sys.stdin

# O(N)
def get_size(image, h, l):
    d = defaultdict(int)
    # O(N) time
    for pixel in image:
        d[pixel] += 1
    

    # O(1)
    most_important = max(d.values())
    size = 0
    
    # O(1)
    for v in d.values():
        if v == most_important:
            size += h * v
        else:
            size += l * v
    return size



if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for i in range(1, n+1):
        r, c, h, l = [int(x) for x in stdin.readline().strip().split()]

        image = []
        for _ in range(r):
            image.extend(list(stdin.readline().strip()))
        
        size = get_size(image, h, l)
        print(f'Case {i}: {size}')
