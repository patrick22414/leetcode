class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        s = s.split()
        if len(pattern) != len(s):
            return False

        book1 = {}
        book2 = {}

        for p, w in zip(pattern, s):
            if p not in book1:
                book1[p] = w
            if w not in book2:
                book2[w] = p

            if book1[p] != w or book2[w] != p:
                return False

        return True
