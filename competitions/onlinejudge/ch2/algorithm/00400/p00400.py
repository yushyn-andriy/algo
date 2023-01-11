import sys
import os

stdin = sys.stdin


WIDTH = 60


def unix_ls(filenames, width=WIDTH):
    n = len(filenames)
    longest = len(max(filenames, key=lambda x: len(x)))
    filenames = sorted(filenames)

    n_col = (width - longest) // (longest + 2)  + 1
    n_rows = 0

    while n_rows * n_col < n:
        n_rows += 1

    matrix = [[''] * n_col for _ in range(n_rows)]
    c_idx = 0
    for j in range(n_col):
        for i in range(n_rows):
            col_width = longest
            if j < n_col -1:
                col_width += 2
            matrix[i][j] = f'{filenames[c_idx]:<{col_width}}'
            c_idx += 1
            if c_idx >= len(filenames):
                return matrix




if __name__ == '__main__':
    for row in stdin:
        n = int(row.strip())
        
        filenames = []
        for _ in range(n):
            filenames.append(stdin.readline().strip())
        print('-' * WIDTH)
        m = unix_ls(filenames)
        for row in m:
            print(''.join(row))
