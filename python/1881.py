class Solution:
    def maxValue(self, n: str, x: int) -> str:
        if n[0] != "-":
            # n is positive
            for i, c in enumerate(n):
                if int(c) < x:
                    return n[:i] + str(x) + n[i:]
            return n + str(x)
        else:
            # n is negative
            n = n[1:]
            for i, c in enumerate(n):
                if int(c) > x:
                    return "-" + n[:i] + str(x) + n[i:]
            return "-" + n + str(x)
