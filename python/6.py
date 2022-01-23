from itertools import zip_longest


class Solution:
    def convert(self, s: str, numRows: int) -> str:
        """Credit: https://leetcode.com/problems/zigzag-conversion/discuss/3404/Python-O(n)-Solution-in-96ms-(99.43)"""
        if numRows == 1 or numRows >= len(s):
            return s

        L = [""] * numRows
        index, step = 0, 1

        for x in s:
            L[index] += x
            if index == 0:
                step = 1
            elif index == numRows - 1:
                step = -1
            index += step

        return "".join(L)

    def convert2(self, s: str, numRows: int) -> str:
        if len(s) <= numRows or numRows == 1:
            return s

        res = ""
        jump = 2 * numRows - 2
        for row in range(numRows):
            if row == 0 or row == numRows - 1:
                res += s[row::jump]
            else:
                tmp1 = s[row::jump]
                tmp2 = s[jump - row :: jump]
                res += "".join(a + b for a, b in zip_longest(tmp1, tmp2, fillvalue=""))

        return res
