import sys



stdin = sys.stdin


def in_range(val, a, b):
    if val >= a and val < b:
        return True
    return False

# O(N) time | O(N) space
def solution(pole):
    n = len(pole)
    grid = [-1] * n
    failed = False
    # O(N) time
    for i in range(n):
        cn, lw = pole[i]
        if lw == 0:
            if grid[i] != -1:
                failed = True
                break
            grid[i] = cn
        else:
            idx = i + lw
            if in_range(idx, 0, n):
                grid[idx] = cn
            else:
                failed = True
                break
    # O(N) time
    if failed or -1 in grid:
        return []
    return grid


if __name__ == '__main__':
    while True:
        n = stdin.readline().strip()
        if n == '0':
            break
        n = int(n)
        pole = [None] * n
        for i in range(n):
            cn, lw = [int(x) for x in stdin.readline().strip().split()]
            pole[i] = [cn, lw]
        r = solution(pole)
        if r:
            print(' '.join([str(x) for x in r]))
        else:
            print(-1)

