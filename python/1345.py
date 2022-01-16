from typing import List
from collections import defaultdict, deque


class Solution:
    def minJumps(self, arr: List[int]) -> int:
        if len(arr) == 1:
            return 0

        indices = defaultdict(list)
        for i, n in enumerate(arr):
            indices[n].append(i)

        queue = deque()
        seen = [False] * len(arr)
        dist = [0] * len(arr)

        queue.append(0)
        seen[0] = True

        while queue:
            index = queue.popleft()
            value = arr[index]
            new_dist = dist[index] + 1

            for x in [index - 1, index + 1]:
                if x >= 0 and x < len(arr) and not seen[x]:
                    if x == len(arr) - 1:
                        return new_dist

                    queue.append(x)
                    seen[x] = True
                    dist[x] = new_dist

            if value in indices:
                for x in indices[value]:
                    if seen[x]:
                        continue

                    if x == len(arr) - 1:
                        return new_dist

                    queue.append(x)
                    seen[x] = True
                    dist[x] = new_dist

                indices.pop(value)


input = [7] * 50000 + [1] * 10 + [11]

output = Solution().minJumps(input)
print(output)
