import sys



stdin = sys.stdin


def get_jumps(n_walls, heights):
    n_low, n_high = 0, 0
    if n_walls <= 1:
        return n_low, n_high
    
    p1, p2 = 0, 1
    while p2 < len(heights):
        w1, w2 = heights[p1], heights[p2]

        if w2 - w1 > 0:
            n_high += 1
        elif w2 - w1 < 0:
            n_low += 1
        else:
            pass
        
        p1 += 1
        p2 += 1

    return n_low, n_high



if __name__ == '__main__':
    cases = int(stdin.readline().strip())
    
    for i in range(1, cases+1):
        n_walls = int(stdin.readline().strip())
        heights = [int(x) for x in stdin.readline().strip().split()]
        n_low, n_high = get_jumps(n_walls, heights)
        print(f'Case {i}: {n_high} {n_low}') 
