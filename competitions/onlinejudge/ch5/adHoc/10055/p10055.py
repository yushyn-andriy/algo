import sys



stdin = sys.stdin


if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        a, b = list(map(int, row.split()))
        print(abs(a - b))
