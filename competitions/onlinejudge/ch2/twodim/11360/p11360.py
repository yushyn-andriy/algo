import sys
import copy

stdin = sys.stdin

def interchange_rows(m, a, b):
    a = a - 1
    b = b - 1
    for i in range(len(m)):
        m[a][i], m[b][i] = m[b][i], m[a][i]

def interchange_columns(m, a, b):
    a = a - 1
    b = b - 1
    for j in range(len(m)):
        m[j][a], m[j][b] = m[j][b], m[j][a]


def inc(m):
    N = len(m)
    for i in range(N):
        for j in range(N):
            m[i][j] = (m[i][j] + 1) % 10


def dec(m):
    N = len(m)
    for i in range(N):
        for j in range(N):
            val = m[i][j] - 1
            if val < 0:
                val = 9
            m[i][j] = val

def transpose(m):
    N = len(m)
    m1 = copy.deepcopy(m)

    col_idx = 0
    for i in range(N):
        for j in range(N):
            m[j][col_idx] = m1[i][j]
        col_idx += 1


if __name__ == '__main__':
    T = int(stdin.readline().strip())
    for i in range(1, T + 1):
        N = int(stdin.readline().strip())
        m = []
        for _ in range(N):
            m.append([int(x) for x in list(stdin.readline().strip())])
        
        M = int(stdin.readline().strip())
        for _ in range(M):
            op_line = stdin.readline().strip().split() 
            op = op_line[0]
            if op == 'row':
                a, b = int(op_line[1]), int(op_line[2])
                interchange_rows(m, a, b)
            elif op == 'col':
                a, b = int(op_line[1]), int(op_line[2])
                interchange_columns(m, a, b)
            elif op == 'inc':
                inc(m)
            elif op == 'dec':
                dec(m)
            elif op == 'transpose':
                transpose(m) 
            else:
                assert False

        print(f'Case #{i}')
        for row in m:
            print(''.join([str(x) for x in row]))
        print()