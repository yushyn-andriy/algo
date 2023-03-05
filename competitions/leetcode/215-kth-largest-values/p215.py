from typing import List


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        return list(sorted(nums, reverse=True))[k-1]
