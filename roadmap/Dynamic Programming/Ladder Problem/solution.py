class Solution:
    # Bottom Up iterative solution
    def climbStairs(self, n, k) -> int:
        m = [-1] * (n + 1)
        m[0], m[1] = 1, 1

        nl = 1
        t_ways = 0
        while nl <= n:
            c_ways = 0
            for c_steps in range(min([k, nl]), 0, -1):
                if nl - c_steps < 0:
                    assert False
                c_ways += m[nl - c_steps]

            m[nl] = c_ways
            t_ways += c_ways
            nl += 1

        #print(m)
        return m[n]


if __name__ == '__main__':
    t = int(input())
    s = Solution()
    for _ in range(t):

        n, k = [int(x) for x in input().strip().split()]
        print(f'n={n}, k={k}: {s.climbStairs(n, k)}')

