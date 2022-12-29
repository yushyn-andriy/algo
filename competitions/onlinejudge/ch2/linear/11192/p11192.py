import sys



stdin = sys.stdin

def reverse_groups(n_groups, s):
    sr = ''
    str_len = len(s)
    group_len = int(str_len / n_groups)

    b, e = 0, 0
    while e < str_len:
        cn_alpha = 0
        b = e
        while cn_alpha < group_len and e < str_len:
            ch = s[e]
            if ch.isalnum():
                cn_alpha += 1
            e += 1
        group = s[b:e]
        sr += ''.join(list(reversed(group)))
    return sr


if __name__ == '__main__':
    while True:
        l = stdin.readline().strip()
        if l == '0':
            break
        n_groups, s = l.split()
        n_groups = int(n_groups)
        print(reverse_groups(n_groups, s))

