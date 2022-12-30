import sys


stdin = sys.stdin


def get_board(w, h):
    return [
        [True] * w
        for _ in range(h)
    ]


def apply(board, sub):
    x1, y1, x2, y2 = sub
    x1, x2 = min(x1, x2), max(x1, x2)
    y1, y2 = min(y1, y2), max(y1, y2)

    cy = y1
    while x1 <= x2:
        while cy <= y2:
            board[cy][x1] = False
            cy += 1
        cy = y1
        x1 += 1
    return board


def count_spots(board):
    return sum([sum(x) for x in board])


if __name__ == '__main__':
    while True:
        w, h, n = [int(x) for x in stdin.readline().strip().split()]
        if w == 0 and h == 0 and n == 0:
            break
        board = get_board(w, h)
        for _ in range(n):
            sub_board = [int(x) - 1 for x in stdin.readline().strip().split()]
            apply(board, sub_board)
        
        spots = count_spots(board)
        if spots == 0:
            print('There is no empty spots.')
        elif spots == 1:
            print('There is one empty spot.')
        else:
            print(f'There are {spots} empty spots.')            
