import sys

from collections import namedtuple



Rect = namedtuple('Rect', 'ux uy lx ly')


stdin = sys.stdin


def detect(x, y, rectangles):
    contains = []
    for i, rect in enumerate(rectangles):
        if x > rect.ux and x < rect.lx and y < rect.uy and y > rect.ly:
            contains.append((i, rect))
    return contains


if __name__ == '__main__':
    rectangles = []
    for row in stdin:
        row = row.strip()

        if row == '*':
            break
        
        rect = Rect(*list(map(float, row.split()[1:])))
        rectangles.append(rect)

    for i, row in enumerate(stdin):
        row = row.strip()
        x, y = list(map(float, row.split()))
        if x == 9999.9 and y == 9999.9:
            break
        
        contains = detect(x, y, rectangles)
        if not contains:
            print(f'Point {i+1} is not contained in any figure')
        else:
            for j, _ in contains:
                print(f'Point {i+1} is contained in figure {j+1}')
