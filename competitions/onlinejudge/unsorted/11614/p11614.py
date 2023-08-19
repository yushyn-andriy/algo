import sys



stdin = sys.stdin


a = 1
d = 1


def solve(sn):
    global a, d
    r = []
    for op in ('+', '-'):
        x = eval(f'(- (2*a -d) {op} pow((2*a -d)**2 + 8 * d * sn, 0.5))/(2*d)')
        r.append(x)
    return int(max(r))


if __name__ == '__main__':
    stdin.readline()
    for row in stdin:
        n = int(row)
        print(solve(n))
