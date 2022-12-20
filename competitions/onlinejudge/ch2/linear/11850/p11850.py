import sys



stdin = sys.stdin


TOTAL_MILES = 1422


def solution(stations, total_miles=TOTAL_MILES):
    for i in range(len(stations) - 1):
        diff = stations[i+1] - stations[i]
        if diff <= 200:
            continue
        else:
            return False
    if abs(stations[-1] - TOTAL_MILES) * 2 > 200:
        return False
    return True




if __name__ == '__main__':

    while True:
        n = stdin.readline().strip()
        if n == '0':
            break
        n = int(n)
        stations = [0] * n
        for i in range(n):
            stations[i] = int(stdin.readline().strip())

        stations = sorted(stations)
        if solution(stations):
            print('POSSIBLE')
        else:
            print('IMPOSSIBLE')

