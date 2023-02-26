import sys
from collections import defaultdict

stdin = sys.stdin


if __name__ == '__main__':
    t = int(stdin.readline().strip())
    stdin.readline()
    
    for i in range(t):
        total = 0
        trees = defaultdict(int)
        for line in stdin:
            line = line.strip()
            if not line:
                break
            total += 1
            trees[line] += 1

        for key in sorted(trees.keys()):
            count = trees[key]
            percent = (count/float(total)) * 100
            print('%s %0.4f' % (key, percent))
        if i < t-1:
            print()
