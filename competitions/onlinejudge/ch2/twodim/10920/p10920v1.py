import sys

stdin = sys.stdin


def solve(P, K):
    max_number = K * K
    K2 = K // 2

    if P > max_number - K:
        return P - max_number  + K2, K2

    max_number -= K - 1
    if P > max_number - K:
        return -K2, P - max_number + K2

    max_number -= K - 1
    if P > max_number - K:
        return  max_number - P - K2, -K2
    
    max_number -= K - 1
    return K2, max_number - P - K2


if __name__ == "__main__":
    while True:
        N, P = stdin.readline().strip().split()
        if N == '0':
            break

        N, P = int(N), int(P)

        K = N
        while K > 1 and (K - 2)*(K - 2) >= P:
            K -= 2

        first, second = solve(P, K)
        first = N // 2 + 1 + first
        second = N // 2 + 1 + second

        print(f'Line = {first}, column = {second}.')
