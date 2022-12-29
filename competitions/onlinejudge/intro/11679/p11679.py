import sys



stdin = sys.stdin


if __name__ == '__main__':
    while True:
        line = tuple(map(int, stdin.readline().strip().split()))
        if line == (0, 0): # end of input
            break

        B, N = line
        monetary_reserves = list(map(int, stdin.readline().strip().split()))

        data = []
        for i in range(N):
            line = list(map(int, stdin.readline().strip().split()))
            D, C, V =  line
            data.append((D - 1, C - 1, V))

        for D, C, V in data:
            monetary_reserves[D] -= V
            monetary_reserves[C] += V
        
        liquidated_all_debantures = True
        for balance in monetary_reserves:
            if balance < 0:
                liquidated_all_debantures = False
                break

        if liquidated_all_debantures:
            print('S')
        else:
            print('N')
