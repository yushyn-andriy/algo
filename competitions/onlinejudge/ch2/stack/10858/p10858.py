import sys


cin = sys.stdin


def solve(n, start):
    res = []
    i = start
    while i <= pow(n, 0.5):
        if n % i == 0:
            sub_res = solve(n//i, i)
            for vec in sub_res:
                vec.append(i)
            res += sub_res
            res.append([n//i, i])
        i += 1
    return res



if __name__ == '__main__':
    for row in cin:
        n = int(row.strip())
        if n == 0:
            break

        res = solve(n, 2)
        print(len(res))
        for vec in res:
            s = ''
            for v in sorted(vec):
                s += f'{v} '
            print(s[:-1])
