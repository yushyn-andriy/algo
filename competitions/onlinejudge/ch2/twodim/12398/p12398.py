import sys
import os


stdin = sys.stdin


board_map = {
    'a': (0, 0),
    'b': (0, 1),
    'c': (0, 2),
    'd': (1, 0),
    'e': (1, 1),
    'f': (1, 2),
    'g': (2, 0),
    'h': (2, 1),
    'i': (2, 2),
}

def solve(row):
    board = [[0]*3 for _ in range(3)]
    if len(row) == 0:
        return board
    for ch in row:
        curr = board_map[ch]
        i, j = curr
        u = i - 1, j
        d = i + 1, j 
        lf = i, j - 1 
        r = i, j + 1       
        
        for coord in [u, d, lf, r, curr]:
            y, x = coord
            if (
                y >= 0 
                and y < len(board) 
                and x >= 0 
                and x < len(board)
            ):
                val = board[y][x]
                val = (val + 1) % 10
                board[y][x] = val

    return board


if __name__ == '__main__':
    i = 1
    for row in stdin:
        board = solve(row.strip())
        print(f'Case #{i}:')
        for row in board:
            print(' '.join([str(x) for x in row]))
        i += 1        
