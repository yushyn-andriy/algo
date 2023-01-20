import sys
import bisect


stdin = sys.stdin


def mediana(arr):
    i = len(arr) // 2
    return (int(abs(arr[i - 1] + arr[i]) / 2) 
            if len(arr) % 2 == 0 
            else arr[i])

# O(N) - time complexity | O(N) - space complexity
if __name__ == '__main__':
    arr = []
    for row in stdin:
        x = int(row.strip())
        bisect.insort(arr, x)
        print(mediana(arr))
