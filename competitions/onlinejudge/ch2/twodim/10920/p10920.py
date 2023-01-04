import sys

stdin = sys.stdin



def eq(current, target):
    return bool(current == target)


def get_ring_id(p):
    if p == 1:
        return 1
    pf = pow(p, 0.5)
    ps = int(pf)
    while True:
        if ps % 2 == 1 and ps > 1 and ps >= pf:
            break
        elif ps == 1:
            ps += 1
            continue
        ps += 1
    
    return int((ps + 1)/2)


def get_coords(sz, p, ring_id):
    i, j = int(sz / 2) - (ring_id - 1), int(sz / 2) + (ring_id - 1) - 1
    if sz == 1:
        return 0, 0
    if ring_id == 1:
        return i, j + 1
    return i, j



def to_cartesian(sz, p):
    ring_id = get_ring_id(p)
    i, j = get_coords(sz, p, ring_id)

    steps = 2 * (ring_id - 1) - 1
    spiral_index = (2 * (ring_id - 1) - 1) ** 2  + 1
    if ring_id == 1:
        return i, j

    while not eq(spiral_index, p):
        # left
        for _ in range(steps):
            if eq(spiral_index, p):
                return i, j
            j -= 1
            spiral_index += 1


        steps += 1
        # down
        for _ in range(steps):
            if eq(spiral_index, p):
                return i, j
            i += 1
            spiral_index += 1

        # right
        for _ in range(steps):
            if eq(spiral_index, p):
                return i, j
            j += 1
            spiral_index += 1

        # up
        for _ in range(steps):
            if eq(spiral_index, p):
                return i, j
            i -= 1
            spiral_index += 1

    return i, j


if __name__ == "__main__":
    while True:
        a, b = stdin.readline().strip().split()
        if a == '0' and b == '0':
            break
        SZ, P = int(a), int(b)
        i, j = to_cartesian(SZ, P)
        print(f'Line = {SZ - i}, column = {j+1}.')
