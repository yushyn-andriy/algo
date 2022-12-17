import sys

stdin = sys.stdin


# O(N) time | O(N) space
def is_jolly(seq):
    n = len(seq)
    flags = [False] * n
    for i in range(0, n - 1):
        diff = abs(seq[i] - seq[i+1])
        if diff > n - 1:
            continue
        flags[diff] = True
    return all(flags[1:])


if __name__ == '__main__':
    while True:
        l = stdin.readline().strip()
        if not l:
            break
        seq = [int(x) for x in l.split()][1:]
        jolly = is_jolly(seq)
        if jolly:
            print('Jolly')
        else:
            print('Not jolly')


