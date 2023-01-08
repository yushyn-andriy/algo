from copy import deepcopy
import sys

stdin = sys.stdin


def solve(land, n, r, c, k):
    current_land = land
    next_land = deepcopy(current_land)
    for _ in range(k):
        for i in range(r):
            for j in range(c):
                current_brother = current_land[i][j]
                hated_brother = (current_brother + 1) % n
                left = i, j - 1 
                right = i, j + 1
                up = i - 1, j
                down = i + 1, j
                if left[1] >= 0:
                    ir, ic = left
                    if current_land[ir][ic] == hated_brother:
                        next_land[ir][ic] = current_brother

                if right[1] < c:
                    ir, ic = right 
                    if current_land[ir][ic] == hated_brother:
                        next_land[ir][ic] = current_brother
                if up[0] >= 0:
                    ir, ic = up 
                    if current_land[ir][ic] == hated_brother:
                        next_land[ir][ic] = current_brother
                if down[0] < r:
                    ir, ic = down 
                    if current_land[ir][ic] == hated_brother:
                        next_land[ir][ic] = current_brother

        current_land = deepcopy(next_land)

    return current_land


if __name__ == '__main__':
    while True:
        N, R, C, K = [int(x) for x in stdin.readline().strip().split()]
        if N == 0:
            break

        land = []
        for _ in range(R):
            land.append([int(x) for x in stdin.readline().strip().split()])
        
        new_land = solve(land, N, R, C, K)
        for row in new_land:
            print(' '.join([str(x) for x in row]))
