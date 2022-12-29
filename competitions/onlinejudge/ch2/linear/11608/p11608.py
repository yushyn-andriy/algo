import sys



stdin = sys.stdin

def check(created, needed, n_extra):
    result = [False] * len(needed)
    available = [0] * len(created)
    available[0] = n_extra
    for i in range(0, len(created) - 1):
        available[i+1] = created[i]
    for j in range(len(needed)):
        need = needed[j]
        avail = available[j]
        if avail < need:
            result[j] = False
        else:
            avail -= need
            result[j] = True

        if j < len(needed) - 1:
            available[j+1] += avail
            available[j] = 0
    return result


if __name__ == '__main__':
    t = 0
    while True:
        t += 1
        n = stdin.readline().strip()
        if n == '-1':
            break
        n = int(n)
        created = [int(x) for x in stdin.readline().strip().split()]
        needed = [int(x) for x in stdin.readline().strip().split()]
        available = check(created, needed, n)
        print(f'Case {t}:')
        for r in available:
            if r is True:
                print('No problem! :D')
            else:
                print('No problem. :(')

