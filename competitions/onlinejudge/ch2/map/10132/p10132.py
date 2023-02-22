import sys
from itertools import permutations



stdin = sys.stdin


def solve(fs, b_per_file):
    freq = {}
    for p in permutations(fs, 2):
        bits = p[0] + p[1]
        if len(bits) == b_per_file:
            if bits not in freq:
                freq[bits] = 0
            else:
                freq[bits] += 1
    
    return max(freq.items(), key=lambda x: x[1])[0]


if __name__ == '__main__':
    t = int(stdin.readline().strip())
    stdin.readline()
    for i in range(t):
        if i == 0:
            pass
        else:
            print()

        fs = []
        B = 0
        for row in stdin:
            row = row.strip()
            if not row:
                break
            fs.append(row)
            B += len(row)

        N = len(fs)
        b_per_file = int(B / (N / 2))

        print(solve(fs, b_per_file))
