import sys



stdin = sys.stdin


# O(N) + O(N) + O(Nlog(N)) ~ O(Nlog(N)) time | O(N) space
def correct_order(l1, l2):
    pairs = [None]*len(l1)
    # O(N) time | O(N) space
    for i in range(len(l1)):
        pairs[i] = (l1[i], l2[i])

    # O(N*log(N)) time
    pairs = sorted(pairs, key=lambda x: x[0])

    # O(N) time
    return [x[1] for x in pairs]


def read_list(conv=int):
    return [conv(x) for x in stdin.readline().strip().split()]

if __name__ == '__main__':
    c = int(stdin.readline().strip())
    for i in range(c):
        stdin.readline()
        l1 = read_list(int)
        l2 = read_list(str)

        ordered = correct_order(l1, l2)
        for x in ordered:
            print(x)
        if i < c - 1:
            print()

