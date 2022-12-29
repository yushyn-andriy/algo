import sys



stdin = sys.stdin


def calculate_price(text, map):
    total = 0
    for line in text:
        for ch in line:
            total += map.get(ch, 0)
    return total


if __name__ == '__main__':
    t = int(stdin.readline().strip())
    for i in range(t):
        map = {}
        text = []
        n_mappings = int(stdin.readline().strip())
        for j in range(n_mappings):
            k, v = stdin.readline().strip().split()
            v = int(v)
            map[k] = v
        n_lines = int(stdin.readline().strip())
        for j in range(n_lines):
            text.append(stdin.readline())
        price = calculate_price(text, map)
        print(f'{price/100:.2f}$')

