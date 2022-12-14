import sys


class Solution:
    def __init__(self):
        self._cache = {0: 0, 1: 1}

    def nthFibonacci(self, n):
        if n in self._cache:
            return self._cache[n]

        s1 = self.nthFibonacci(n-1)
        s2 = self.nthFibonacci(n-2)

        self._cache[n] = (s1 + s2) % 1000000007
        return self._cache[n]



if __name__ == '__main__':
    s = Solution()
    t = int(input())
    for _ in range(t):
        n = int(input())
        print(s.nthFibonacci(n))

