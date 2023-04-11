import sys
from collections import defaultdict

stdin = sys.stdin




if __name__ == '__main__':
    result = []
    d = defaultdict(int)
    for row in stdin:
        L = row.strip().split()
        for ch in L:
            if ch not in d:
                result.append(ch)
            d[ch] += 1
        

    for ch in result:
        print(ch, d[ch])
