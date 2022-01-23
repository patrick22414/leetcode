class Solution:
    def reverse(self, x: int) -> int:
        sign = 1 if x >= 0 else -1

        s = str(abs(x))
        s = s[::-1]
        y = sign * int(s)

        if y < -(2 ** 31) or y > 2 ** 31 - 1:
            return 0
        return y
