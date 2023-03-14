import sys

stdin = sys.stdin


MAXN = 20010

def find(x, p):
    if x == p[x]:
        return x
    p[x] = find(p[x], p)
    return p[x]

def joint(x, y, p, r):
    x = find(x, p)
    y = find(y, p)
    if x != y:
        if r[x] < r[y]:
            p[x] = y
            r[y] += r[x]
        else:
            p[y] = x
            r[x] += r[y]
        return 1
    return 0

def init(p, r):
    for i in range(MAXN):
        p[i] = i
        r[i] = 1

def enemyID(x, N):
    return x + N


if __name__ == '__main__':

        N = int(stdin.readline().strip())
        p = [0] * MAXN
        r = [0] * MAXN
        init(p, r)
        while True:
            line = stdin.readline().strip().split()
            c, x, y = map(int, line)
            if c == 0 and x == 0 and y == 0:
                break
            if c == 1:
                if find(x, p) == find(enemyID(y, N), p) or find(y, p) == find(enemyID(x, N), p):
                    print("-1")
                else:
                    joint(x, y, p, r)
                    joint(enemyID(x, N), enemyID(y, N), p, r)
            elif c == 2:
                if find(x, p) == find(y, p) or find(enemyID(x, N), p) == find(enemyID(y, N), p):
                    print("-1")
                else:
                    joint(x, enemyID(y, N), p, r)
                    joint(y, enemyID(x, N), p, r)
            elif c == 3:
                if find(x, p) == find(y, p) or find(enemyID(x, N), p) == find(enemyID(y, N), p):
                    print("1")
                else:
                    print("0")
            else:
                if find(x, p) == find(enemyID(y, N), p) or find(y, p) == find(enemyID(x, N), p):
                    print("1")
                else:
                    print("0")