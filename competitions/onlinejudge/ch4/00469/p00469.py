import sys
from copy import deepcopy


stdin = sys.stdin


dr = [1, 1, 1, -1, -1, -1, 0, 0]
dc = [-1, 0, 1, -1, 0, 1, -1, 1]


def floodfill(grid, r, c, c1, c2):
    ans = 0
    rows, columns = len(grid), len(grid[0])
    if r < 0 or r > rows - 1 or c < 0 or c > columns - 1:
        return 0

    if grid[r][c] != c1:
        return 0

    ans = 1
    grid[r][c] = c2
    for i in range(8):
        ans += floodfill(grid, r + dc[i], c + dr[i], c1, c2) 

    return ans


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    stdin.readline()

    for i in range(n):
        grid = []
        coords = []
        for row in stdin:
            row = row.strip()
            if row[0].isnumeric():
                break
            grid.append(list(row))

        coords.append([int(x) for x in row.split()])
        for row in stdin:
            row = row.strip()
            if not row:
                break
            coords.append([int(x) for x in row.split()])

        for r, c in coords:
            print(floodfill(deepcopy(grid), r - 1, c - 1, 'W', '.'))
        if i == n - 1:
            continue
        print()
