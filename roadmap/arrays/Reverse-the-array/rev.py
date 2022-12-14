import sys

stdin = sys.stdin



def reverse_(arr):
    i, j = 0, len(arr) - 1
    while i < j:
        arr[i], arr[j] = arr[j], arr[i]
        i+=1
        j-=1
    return arr


if __name__ == '__main__':
    arr = [int(x) for x in stdin.readline().strip().split()]
    print(reverse_(arr))

