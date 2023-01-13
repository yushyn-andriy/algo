import sys


stdin = sys.stdin


def med(X):
    n = len(X)
    return X[int((n - 1)/2)]


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for i in range(n):
        s, a, f = [int(x) for x in stdin.readline().strip().split()]
        streets, avenuis = [], []
        for _ in range(f):
            sn, an = [int(x) for x in stdin.readline().strip().split()]
            if sn <= s:
                streets.append(sn)
            if an <= a:
                avenuis.append(an)
        
        # nlog(n)
        streets.sort()
        # nlog(n)
        avenuis.sort()
        
        print(f'(Street: {med(streets)}, Avenue: {med(avenuis)})')
