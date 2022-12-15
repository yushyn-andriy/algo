import sys



stdin = sys.stdin

# O(n) time | O(n) space
def update_coins_status(coins, left_w, right_w, result):
    update_from = []
    if result == '<':
        update_from = []
    elif result == '>':
        update_from = []
    else:
        update_from = left_w + right_w
    for i in update_from:
        coins[i] = True
    #print(coins)
    return coins


# O(N) time
def solve(coins):
    c = coins.count(False)
    if c == 1:
        #print('l', len(coins))
        return coins.index(False)
    return 0


if __name__ == '__main__':
    t = int(stdin.readline().strip())
    for j in range(t):
        input()
        n_coins, n_weightings = [int(x) for x in input().strip().split()]
        coins = [False] * (n_coins + 1)
        coins[0] = True
        for i in range(n_weightings):
            weights = [int(x) for x in input().strip().split()]
            n_lr, weights = weights[0], weights[1:]

            left_w, right_w = weights[:n_lr], weights[n_lr:]
            result = input()
            # modifies coins array
            coins = update_coins_status(coins, left_w, right_w, result)
            # print(coins)
        print(solve(coins))
        if j < t -1:
            print()

