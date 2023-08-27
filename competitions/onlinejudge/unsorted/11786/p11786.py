import sys


stdin = sys.stdin


def calculate_water_units(universe):
    stack = []
    water = 0

    for i, symbol in enumerate(universe):
        if symbol == "\\":
            stack.append(i)
        elif symbol == "/" and stack:
            left_index = stack.pop()
            width = i - left_index
            water += width

    return water


if __name__ == '__main__':
    n = int(stdin.readline().strip())
    for _ in range(n):
        universe = stdin.readline().strip()
        print(calculate_water_units(universe))
