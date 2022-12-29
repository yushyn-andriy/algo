def minimum(arr, i, m):
    if i == len(arr) -1:
        return m
    if m > arr[i]:
        m = arr[i]
    return minimum(arr, i + 1, m)


if __name__ == '__main__':
    arr = [int(x) for x in input().strip().split()]
    print(minimum(arr, 0, 10000))

