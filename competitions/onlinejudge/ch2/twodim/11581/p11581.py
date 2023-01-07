import sys


stdin = sys.stdin



def done(m):
    for i in range(len(m)):
        for j in range(len(m)):
            if m[i][j] == 1: return False
    return True


def solve(m, l):
    if done(m): return -1

    g = [[0] * 3 for _ in range(3)]
    for i in range(len(m)):
        for j in range(len(m)):
            curr = 0
            rl, cl = i, j - 1
            rr, cr = i, j + 1
            ru, cu = i - 1, j
            rd, cd = i + 1, j
            if cl >=0:
                curr += m[rl][cl]
            if cr < len(m):
                curr += m[rr][cr]
            if ru >= 0:
                curr += m[ru][cu]
            if rd < len(m):
                curr += m[rd][cd]
            g[i][j] = curr % 2
    return 1 + solve(g, 1)


if __name__ == '__main__':
    N = int(stdin.readline().strip())
    for _ in range(N):
        stdin.readline().strip()
        l = []
        for i in range(3):
            l.append([int(x) for x in list(stdin.readline().strip())])
        print(solve(l, 0))