import sys



stdin = sys.stdin


def cost(durations):
    mile_total, juice_total = 0, 0
    for duration in durations:
        duration += 1

        s = duration // 60
        if duration % 60 > 0:
            s += 1
        juice_total += s * 15

        s = duration // 30
        if duration % 30 > 0:
            s += 1        
        mile_total += s * 10

    return mile_total, juice_total


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for c in range(1, n + 1):
        _ = stdin.readline()
        durations = [int(x) for x in stdin.readline().strip().split()]
        
        mile, juice = cost(durations)
        if mile < juice:
            print(f'Case {c}: Mile {mile}')
        elif juice < mile:
            print(f'Case {c}: Juice {juice}')
        else:
            print(f'Case {c}: Mile Juice {mile}')
