import sys



stdin = sys.stdin


def eat():
    stdin.readline()


def draw(a, f, last):
    for j in range(f):
        for i in range(1, a+1):
            print(f'{i}'*i)
        for i in range(a-1, 0, -1):
            print(f'{i}'*i)
        if j == f - 1 and last:
            continue
        print()

if __name__ == '__main__':
    c = int(stdin.readline().strip())
    for i in range(c):
        eat()
        amplitude = int(stdin.readline().strip())
        freq = int(stdin.readline().strip())
        draw(amplitude, freq, i == c-1)
