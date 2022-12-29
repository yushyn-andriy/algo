import sys



stdin = sys.stdin


def count_peaks(seq):
    n = len(seq)
    if n <= 1:
        return 0

    n_peaks = 0
    l, m, r = -1, 0, 1
    while m < n:
        if seq[l] < seq[m] > seq[r] or seq[l] > seq[m]  < seq[r]:
            n_peaks += 1
        else:
            pass
        l += 1
        m += 1
        r = (r + 1) % n
    return n_peaks


if __name__ == '__main__':
    while True:
        n = stdin.readline().strip()
        if n == '0':
            break
        n = int(n)
        seq = [int(x) for x in stdin.readline().strip().split()][:n]
        print(count_peaks(seq))

