import sys


if __name__ == '__main__':
    n = int(sys.stdin.readline())
    for i in range(n):
        a = [int(x) for x in sys.stdin.readline().split()]
        a = sorted(a)
        print(f'Case {i + 1}: {a[1]}')
