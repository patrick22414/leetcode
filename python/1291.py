from math import floor, log10
from typing import List


class Solution:
    def sequentialDigits(self, low: int, high: int) -> List[int]:
        results = []

        for num_digits in range(floor(log10(low)) + 1, floor(log10(high)) + 2):
            for start in range(10 - num_digits):
                integer = int("123456789"[start : start + num_digits])

                if integer > high:
                    return results
                elif integer < low:
                    continue
                else:
                    results.append(integer)

        return results
