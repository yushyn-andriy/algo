import sys

stdin = sys.stdin



def _reverse(arr, i, j):
    if i >= j:
        return arr
    arr[i], arr[j] = arr[j], arr[i]
    return _reverse(arr, i+1, j-1)

def reverse_(arr):
    i, j = 0, len(arr) - 1
    return _reverse(arr, i, j)


if __name__ == '__main__':
    arr = [int(x) for x in stdin.readline().strip().split()]
    print(reverse_(arr))

