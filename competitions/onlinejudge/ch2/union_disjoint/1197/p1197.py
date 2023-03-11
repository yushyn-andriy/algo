import sys


stdin = sys.stdin



def find(arr, i):
    if arr[i] == -1:
        return i
    r = find(arr, arr[i])
    arr[i] = r
    return r


def is_connected(arr, a, b):
    return find(arr, a) == find(arr, b)


def union(arr, a, b):
    i = find(arr, a)
    j = find(arr, b)
    if i == j:
        return True
    arr[i] = j
    return False


def number_of_suspects(n, arr):
    number = 1
    for i in range(1, n):
        if is_connected(arr, i, 0):
            number += 1
    return number


if __name__ == '__main__':
    for row in stdin:
        row = row.strip()
        n, m = list(map(int, row.split()))
        if (n, m) == (0, 0):
            break
        arr = [-1] * n
        for _ in range(m):
            L = stdin.readline().strip().split()[1:]
            prev = None
            for x in L:
                x = int(x)
                if prev is not None:
                    union(arr, x, prev)
                prev = x
        print(number_of_suspects(n, arr))
