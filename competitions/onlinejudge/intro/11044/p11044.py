import sys


def min_sonar_count(rows, cols):
    return (rows // 3) * (cols // 3)


if __name__ == '__main__':
    cases = int(sys.stdin.readline())
    for _ in range(cases):
        rows, cols = [int(i) for i in sys.stdin.readline().split()]
        print(min_sonar_count(rows, cols))
