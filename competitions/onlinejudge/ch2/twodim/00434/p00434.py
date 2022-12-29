import sys



stdin = sys.stdin



def solution(front, right, k):
    n_max, n_min = 0, 0
    for i in range(k):
        for j in range(k):
            n_max += min(front[i], right[j])

    front_no_avail = [False]*k
    right_no_avail = [False]*k
    for i in range(k):
        for j in range(k):
            if (
                    front[j] == right[i]
                    and not front_no_avail[j]
                    and not right_no_avail[i]
            ):
                n_min += front[j]
                front_no_avail[j] = True
                right_no_avail[i] = True
    for i in range(k):
        if not front_no_avail[i]:
            n_min += front[i]
        if not right_no_avail[i]:
            n_min += right[i]

    return n_min, n_max - n_min


if __name__ == '__main__':
    t = int(stdin.readline().strip())
    for i in range(t):
        k = int(stdin.readline().strip())
        front = [int(x) for x in stdin.readline().strip().split()]
        right = [int(x) for x in stdin.readline().strip().split()]
        M, N = solution(front, right, k)
        print(f'Matty needs at least {M} blocks, and can add at most {N} extra blocks.')

