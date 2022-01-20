from typing import List


class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        lo, hi = 1, max(piles)

        while lo != hi:
            k = (lo + hi) // 2
            t = self.totalEatingTime(piles, k)

            if t > h:  # too slow
                lo = k + 1
            else:  # t <= h; too fast
                hi = k

        return lo

    def totalEatingTime(self, piles: List[int], k: int) -> int:
        if k <= 0:
            raise ValueError
        return sum((m - 1) // k + 1 for m in piles)
