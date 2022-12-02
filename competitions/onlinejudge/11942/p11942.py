import sys



stdin = sys.stdin


def is_ordered(array):
    ordered = True
    if len(array) <= 1:
        return ordered
    
    p1, p2 = 0, 1
    sign = array[p2] - array[p1]
    p1 += 1
    p2 += 1
    while p2 < len(array):
        diff = array[p2] - array[p1]
        if sign < 0 and diff > 0:
            return False
        if sign > 0 and diff < 0:
            return False
        p1 += 1
        p2 += 1
    return ordered




if __name__ == '__main__':
    n = int(stdin.readline().strip())
    print('Lumberjacks:')
    for _ in range(n):
        array = [int(x) for x in stdin.readline().strip().split()]
        ordered = is_ordered(array)
        if ordered:
            print('Ordered')
        else:
            print('Unordered')
    