import sys



stdin = sys.stdin


def get_arrangment(pairs):
    arr = ['']*len(pairs)
    cursor = -1
    for pair in pairs:
        card, word = pair
        lenght = len(word)
        while lenght > 0:
            cursor = (cursor + 1) % len(pairs)
            # ignore non empty places
            if arr[cursor] == "":
                lenght -= 1
        arr[cursor] = card
    return arr


if __name__ == '__main__':
    while True:
        n = int(stdin.readline().strip())
        if n == 0:
            break
        pairs = [None]*n
        for i in range(n):
            pairs[i] = tuple(stdin.readline().strip().split())
        arr = get_arrangment(pairs)
        print(' '.join(arr))

