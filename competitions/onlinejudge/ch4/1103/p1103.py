import sys
from collections import deque
from copy import deepcopy


stdin = sys.stdin



# map white group numbers to letter
ntol = {
    0: 'W',
    1: 'A',
    2: 'K',
    3: 'J',
    4: 'S',
    5: 'D',
}
dr = [1, 1, 1, -1, -1, -1, 0, 0]
dc = [-1, 0, 1, -1, 0, 1, -1, 1]



def helper_visit_white_bfs(image, i, j, white, visited):
    rows, columns = len(image), len(image[0])
    d = deque()
    d.append((i, j))
    while len(d) > 0:
        i, j = d.popleft()
        if i < 0 or i > rows - 1 or j < 0 or j > columns - 1:
            continue
        
        if image[i][j] == visited:
            continue
        if image[i][j] != white:
            continue

        image[i][j] = visited
        for k in range(8):
            d.append((dr[k] + i, dc[k] + j))


def helper_bfs(image, i, j, white, black, visited):
    rows = len(image)
    columns = len(image[0])
    d = deque()
    d.append((i, j))

    white_groups = 0
    while len(d) > 0:
        i, j = d.popleft()

        if i < 0 or i > rows - 1 or j < 0 or j > columns - 1:
            continue
        if image[i][j] == visited:
            continue
        if image[i][j] == white:
            white_groups += 1
            helper_visit_white_bfs(image, i, j, white, visited)
            continue

        image[i][j] = visited
        for k in range(8):
            d.append((dr[k] + i, dc[k] + j))
    return ntol[white_groups]


def recognize(image, white='0', black='1', visited='-'):
    result = ''
    helper_visit_white_bfs(image, 0, 0, white, visited)
    for i in range(len(image)):
        for j in range(len(image[0])):
            pixel = image[i][j]
            if pixel == visited:
                continue
            result += helper_bfs(image, i, j, white, black, visited)

    return ''.join(sorted(result))


if __name__ == '__main__':
    c = 1
    while True:
        H, W = [int(x) for x in stdin.readline().strip().split()]
        if H == 0:
            break

        image = []
        for i in range(H):
            row = stdin.readline().strip()
            assert len(row) == W
            s = ''
            for ch in row:
                s += '{0:04b}'.format(int(ch, 16))

            image.append(list(f'0{s}0'))
        
        image.insert(0, list('0' * len(image[0])))
        image.append(list('0' * len(image[0])))

        result = recognize(deepcopy(image))
        print(f'Case {c}: {result}')

        c+=1
