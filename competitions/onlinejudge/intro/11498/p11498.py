import sys


def get_residence(n, m, x, y):
    if x > n and y > m:
        return 'NE'
    if x < n and y < m:
        return 'SO'
    if x < n and y > m:
        return 'NO'
    if x > n and y < m:
        return 'SE'

    return 'divisa'   


if __name__ == '__main__':
    while True:
        k = int(sys.stdin.readline())
        if k == 0:
            break
        
        n, m = [int(i) for i in sys.stdin.readline().split()]
        for _ in range(k):
            x, y = [int(i) for i in sys.stdin.readline().split()]
            print(get_residence(n, m, x, y))
