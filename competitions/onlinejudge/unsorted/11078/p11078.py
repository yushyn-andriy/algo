import sys


stdin = sys.stdin


def readint():
    return int(stdin.readline().strip())


def solve(scores):
    max_so_far = scores[0]
    diff = -1
    for score in scores:
        if score > max_so_far:
            max_so_far = score
        else:
            diff = max(diff, max_so_far - score)

    if diff in [-1, 0]:
        return -1
    return diff


if __name__ == '__main__':
    n = readint()
    for _ in range(n):
        students = readint()
        scores = []
        for _ in range(students):
            scores.append(readint())
        print(solve(scores))
