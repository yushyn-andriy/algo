import sys



def check(l, w, h):
    if all([x <= 20 for x in [l, w, h]]):
        return 'good'
    return 'bad'



if __name__ == '__main__':
    c = int(sys.stdin.readline().strip())
    for i in range(c):
        l, w, h = [int(x) for x in sys.stdin.readline().split()]
        print(f'Case {i + 1}: {check(l, w, h)}')
