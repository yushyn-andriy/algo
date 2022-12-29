import sys



stdin = sys.stdin



def is_holiday(number):
    if number == 6 or number == 7:
        return True
    return False


# O(N) time | O(N) Space | N is days
def get_loosed_days(d, params):
    working_calendar = [True] * (d + 1)
    for i in range(1, d + 1):
        c_day = (i - 1) % 7 + 1
        if is_holiday(c_day):
            continue
        for p in params:
            if i % p == 0:
                working_calendar[i] = False

    nwd = 0
    for v in working_calendar:
        if v is False:
            nwd += 1
    return nwd


if __name__ == '__main__':
    t = int(stdin.readline().strip())

    for i in range(t):
        n = int(input())
        p = int(input())
        params = [-1] * p
        for j in range(p):
            params[j] = int(input())
        print(get_loosed_days(n, params))

