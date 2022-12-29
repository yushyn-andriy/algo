import sys



ONE = 'one'
TWO = 'two'
THREE = 'three'


def get_number(x):
    e1, e2, e3 = 0, 0, 0
    for i, ch in enumerate(x):
        if i > 2:
            e1 += 1
            e2 += 1
            if ch != THREE[i]:
                e3 += 1
        else:
            if ch != ONE[i]:
                e1 += 1
            if ch != TWO[i]:
                e2 += 1
            if ch != THREE[i]:
                e3 += 1

    if e1 <= 1:
        return 1
    if e2 <= 1:
        return 2
    if e3 <= 1:
        return 3


if __name__ == '__main__':
    n = int(sys.stdin.readline().strip())
    for _ in range(n):
        word = sys.stdin.readline().strip()
        print(get_number(word))
