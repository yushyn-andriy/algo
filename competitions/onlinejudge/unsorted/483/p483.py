import sys


stdin = sys.stdin


def reverse_words(string):
    string = list(string)
    
    start, stop = 0, 0
    flag = False
    for i, ch in enumerate(string):
        if flag and ch.isspace():
            stop = i
            string[start:stop] = list(reversed(string[start:stop]))
            flag = False
        elif not flag and not ch.isspace():
            start = i
            flag = True

    if flag:
        stop = i + 1
        string[start:stop] = list(reversed(string[start:stop]))


    return ''.join(string)


if __name__ == '__main__':
    rows = stdin.readlines()
    for row in rows:
        print(reverse_words(row), end='')