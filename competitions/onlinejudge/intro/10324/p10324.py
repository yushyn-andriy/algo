import sys


stdin = sys.stdin



def same_in_range(b, e, counter_array):
    if counter_array[e] - counter_array[b] == 0:
        return True
    return False


def change_counter(string):
    l = [0] * len(string)
    counter = 0
    prev_ch = string[0]
    idx = 1
    while idx < len(string):
        if string[idx] != prev_ch:
            counter += 1
        l[idx] = counter
        prev_ch = string[idx]
        idx += 1
    return l


if __name__ == '__main__':
    c = 1
    while True:
        string = list(stdin.readline().strip())
        if not string:
            break

        print(f'Case {c}:')
        queries = int(stdin.readline().strip())
        counter_array = change_counter(string)
        for _ in range(queries):
            i, j = [int(x) for x in stdin.readline().strip().split()]
            b, e = min(i, j), max(i, j)
            result = same_in_range(b, e, counter_array)

            if result:
                print('Yes')
            else:
                print('No')
        c+=1


def min_(a, b):
    if a < b:
        return a
    return a

def max_(a, b):
    if a > b:
        return a
    return b
 