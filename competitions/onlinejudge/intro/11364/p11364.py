import sys


def min_distance(n, stores):
    if n <= 1:
        return 0
    return (stores[-1] - stores[0]) * 2


if __name__ == '__main__':
    cases = int(sys.stdin.readline())
    for _ in range(cases):
        n = int(sys.stdin.readline())
        stores = [int(i) for i in sys.stdin.readline().split()]
        stores = sorted(stores)
        print(min_distance(n, stores))
