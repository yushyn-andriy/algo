import sys



stdin = sys.stdin


def min_moves_number(n_stacks, heights):
    bricks = sum(heights)
    average_bricks = bricks // n_stacks
    min_bricks = 0
    for h in heights:
        if h < average_bricks:
            min_bricks += average_bricks - h
    return min_bricks



if __name__ == '__main__':
    c_set = 1
    while True:
        n_stacks = int(stdin.readline().strip())
        if not n_stacks:
            break
        heights = [int(x) for x in stdin.readline().strip().split()]
        m = min_moves_number(n_stacks, heights)
        print(f'Set #{c_set}\nThe minimum number of moves is {m}.\n')
        c_set += 1

