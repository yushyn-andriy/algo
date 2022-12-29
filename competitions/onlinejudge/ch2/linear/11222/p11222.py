from array import array
import sys



stdin = sys.stdin


def solve(f, s, t, first, second, third):
    size = 10000
    fs = array('b', [0]*size)
    ss = array('b', [0]*size)
    ts = array('b', [0]*size)
    for arr, sol in [(first, fs), (second, ss), (third, ts)]:
        for num in arr:
            sol[num] = 1

    fu, su, tu = [], [], []
    for i, r in enumerate(fs):
        if r == 1 and ss[i] == 0 and ts[i] == 0:
            fu.append(i)
    for i, r in enumerate(ss):
        if r == 1 and fs[i] == 0 and ts[i] == 0:
            su.append(i)
    for i, r in enumerate(ts):
        if r == 1 and fs[i] == 0 and ss[i] == 0:
            tu.append(i)

    if len(fu) == len(su) == len(tu):
        return [
            list(sorted(fu)),
            list(sorted(su)),
            list(sorted(tu)),
        ], (1, 2, 3)
    elif len(fu) == len(su):
        if len(fu) > len(tu):
            return [
                list(sorted(fu)),
                list(sorted(su))
            ], (1, 2)
        else:
            return [
                list(sorted(tu))
            ], (3, )
    elif len(fu) == len(tu):
        if len(fu) > len(su):
            return [
                list(sorted(fu)),
                list(sorted(tu))
            ], (1, 3)
        else:
            return [list(sorted(su))], (2, )
    elif  len(su) == len(tu):
        if len(su) > len(fu):
            return [
                list(sorted(su)),
                list(sorted(tu)),
            ], (2, 3)
        else:
            return [list(sorted(fu))], (1, )
    else:
        if len(fu) > len(su) and len(fu) > len(tu):
            return [list(sorted(fu))], (1, )
        if len(su) > len(fu) and len(su) > len(tu):
            return [list(sorted(su))], (2, )
        if len(tu) > len(fu) and len(tu) > len(su):
            return [list(sorted(tu))], (3, )



def readline(conv=int):
    return [conv(x) for x in stdin.readline().strip().split()]

if __name__ == '__main__':
    t = int(stdin.readline().strip())
    for i in range(1, t + 1):
        first = readline(int)
        second = readline(int)
        third = readline(int)
        f, s, t = first[0], second[0], third[0]
        first = first[1:]
        second = second[1:]
        third = third[1:]
        result = solve(f, s, t, first, second, third)
        sols, winners = solve(f, s, t, first, second, third)
        print(f'Case #{i}:')
        for i, w in  enumerate(winners):
            if len(sols[i]) == 0:
                print(f'{w} {len(sols[i])}')
                continue
            else:
                print(f'{w} {len(sols[i])}', end=' ')
            for j, num in enumerate(sols[i]):
                if j < len(sols[i]) - 1:
                    print(num, end=' ')
                else:
                    print(num, end='')
            print()


