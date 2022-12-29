def reverse_(arr, l, r):
    if l >= r:
        return

    arr[l], arr[r] = arr[r], arr[l]

    reverse_(arr, l + 1, r - 1)


if __name__ == '__main__':
    arr = [int(x) for x in input().strip().split()]
    reverse_(arr, 0, len(arr) - 1)
    print(arr)

