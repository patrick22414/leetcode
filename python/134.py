from typing import List


class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        """Best solution. Credit: https://leetcode.com/problems/gas-station/discuss/1705893/Python3-DEBIT-AND-CREDIT-O(1)-SPACE-(-)-Explained"""
        candidate, debit, credit = None, 0, 0

        for i in range(len(gas)):
            credit += gas[i] - cost[i]
            if credit < 0:
                candidate, debit, credit = None, debit - credit, 0
            elif candidate is None:
                candidate = i

        return candidate if credit >= debit else -1

    def canCompleteCircuit3(self, gas: List[int], cost: List[int]) -> int:
        n = len(gas)
        i = 0
        while i < n:
            total = 0

            for j in range(i, n + i):
                total += gas[j % n] - cost[j % n]
                if total < 0:
                    i = j + 1
                    break
            else:
                return i

        return -1

    def canCompleteCircuit2(self, gas: List[int], cost: List[int]) -> int:
        n = len(gas)
        gross = [g - c for g, c in zip(gas, cost)]

        if sum(gross) < 0:
            return -1

        if n == 1:
            if gross[0] >= 0:
                return 0
            else:
                return -1

        prev_is_neg = gross[-1] < 0
        for i in range(n):
            if gross[i] >= 0 and prev_is_neg:
                total = 0

                for j in range(n):
                    total += gross[(i + j) % n]
                    if total < 0:
                        break
                else:
                    return i

            prev_is_neg = gross[i] < 0

        return -1
