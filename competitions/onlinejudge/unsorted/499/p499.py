import sys
from collections import defaultdict



stdin = sys.stdin


if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        m = 0
        d = defaultdict(int)
        for ch in row:
            if ch.isalpha():
                d[ch] += 1
                if d[ch] > m:
                    m = d[ch]

        L = [(k, v) for k, v in sorted(d.items(), key=lambda x: x[0])]
        for k, v in L:
            if v == m:
                print(k, end='')
        print(f' {m}')
