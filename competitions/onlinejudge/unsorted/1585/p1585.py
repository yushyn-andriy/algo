import sys


stdin = sys.stdin


def solve(row):
    total_score = 0
    current_score = 0
    for ch in row:
        if ch == 'X':
            current_score = 0
        else:
            current_score += 1
            total_score += current_score
    return total_score


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for _ in range(n):
        row = stdin.readline().strip()
        print(solve(row))
