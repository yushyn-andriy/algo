import sys
from typing import List

stdin = sys.stdin


class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        freq = {
            0: 0,
            1: 0,
            2: 0,
        }
        for v in nums:
            freq[v] += 1

        r1 = range(0, freq[0])
        r2 = range(freq[0], freq[0]+freq[1])
        r3 = range(freq[0]+freq[1], freq[0]+freq[1]+freq[2])
        for c, r in enumerate((r1, r2, r3)):
            for i in r:
                nums[i] = c




if __name__ == '__main__':
    s = Solution()

    arr = list(map(int, stdin.readline().strip().split(',')))
    s.sortColors(arr)
    print(','.join(map(str, arr)))
