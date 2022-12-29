import sys


class Solution:
    def nthFibonacci(self, n):
        if n in (0, 1):
            return n

        data = [0]*(n + 1)
        data[0], data[1] = 0, 1
        i = 2
        while i <= n:
            data[i] = data[i-1] + data[i-2]
            i+=1
        return data[n]



if __name__ == '__main__':
    s = Solution()
    t = int(input())
    for _ in range(t):
        n = int(input())
        print(s.nthFibonacci(n))

