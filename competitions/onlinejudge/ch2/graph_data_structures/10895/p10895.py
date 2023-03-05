import sys


stdin = sys.stdin


def transpose(sparse):
    trd = {}
    for key, value in sparse.items():
        k = key[1]+1
        if k not in trd:
            trd[k] = {
                'columns': [],
                'row_vals': [],
            }
        trd[k]['columns'].append(key[0]+1)
        trd[k]['row_vals'].append(value)
    return trd


if __name__ == '__main__':
    for row in stdin:
        m, n = list(map(int,  row.strip().split()))
        sparse = {}
        for r in range(m):
            k, *columns = list(map(lambda x: int(x) - 1, stdin.readline().strip().split()))
            if k == 0:
                stdin.readline()
                continue
            values = list(map(lambda x: int(x), stdin.readline().strip().split()))
            for i, c in enumerate(columns):
                sparse[(r, c)] = values[i]
        
        print(f'{n} {m}')
        trd = transpose(sparse)
        for key in range(1, n+1):
            if key not in trd:
                print(0)
                print()
                continue
            val = trd[key]
            columns, n = list(map(str, val['columns'])), len(val['columns'])
            vals = list(map(str, val['row_vals']))

            print(f'{n} {" ".join(columns)}')
            print(f'{" ".join(vals)}')
