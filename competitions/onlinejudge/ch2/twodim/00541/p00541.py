import sys



stdin = sys.stdin


def row_sum(m, i):
    return sum(m[i])


def col_sum(m, j):
    s = 0
    for i in range(n):
        s += m[i][j]
    return s


def is_even(x):
    return x % 2 == 0

def invert_bit(m, i, j):
    if m[i][j] == 0:
        m[i][j] = 1
    elif m[i][j] == 1:
        m[i][j] = 0
    else:
        assert False


def parity(m):
    #print(m)
    n = len(m)
    dr, dc = None, None
    n_dr, n_dc = 0, 0
    for i in range(n):
        if not is_even(row_sum(m, i)):
            dr = i
            n_dr += 1
        if not is_even(col_sum(m, i)):
            dc = i
            n_dc += 1
    #print(n_dr, n_dc)
    # print('n_dr, n_dc', n_dr, n_dc)
    if (
            dc is not None and
            dr is not None and
            n_dc == 1 and n_dr == 1
    ):
        #print(m, dr, dc, m[dr][dc])
        invert_bit(m, dr, dc)
        #print(m, dr, dc, m[dr][dc])
        if is_even(row_sum(m, dr)) and is_even(col_sum(m, dc)):
            return f'Change bit ({dr + 1},{dc + 1})'
    elif dc is None and dr is None:
        return 'OK'
    else:
        #print(dr, dc, n_dr, n_dc)
        if dr is not None and dc is None and n_dr == 1:
            for j in range(n):
                invert_bit(m, dr, j)
                if is_even(row_sum(m, dr)) and is_even(col_sum(m, j)):
                    return f'Change bit ({dr + 1},{dc + 1})'
                invert_bit(m, dr, j)
        elif dr is None and dc is not None and n_dc == 1:
            for i in range(n):
                invert_bit(m, i, dc)
                if is_even(row_sum(m, i)) and is_even(col_sum(m, dc)):
                    return f'Change bit ({dr + 1},{dc + 1})'
                invert_bit(m, i, dc)
    return 'Corrupt'


if __name__ == '__main__':
    while True:
        n = int(sys.stdin.readline().strip())
        if n == 0:
            break
        m = []
        for i in range(n):
            m.append([int(x) for x in stdin.readline().strip().split()])
        print(parity(m))

