import sys


stdin = sys.stdin


if __name__ == '__main__':
    s = 0
    for row in stdin:
        s += int(row.strip())
    print(s)
