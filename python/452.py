from typing import List, Optional, Dict


class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        for i in range(len(points)):
            p1 = points[i]
            if p1 is None:
                continue

            for j in range(i + 1, len(points)):
                p2 = points[j]
                if p2 is None:
                    continue

                overlap = self.overlap(p1, p2)
                if overlap is not None:
                    p1 = overlap
                    points[j] = None

            points[i] = p1
            print(points)

        return sum(p is not None for p in points)

    def overlap(self, p1: List[int], p2: List[int]) -> Optional[List[int]]:
        start = max(p1[0], p2[0])
        end = min(p1[1], p2[1])

        if start > end:
            return None

        return [start, end]


def visualise(points: List[List[int]]):
    for p in points:
        if p is not None:
            print(" " * p[0] + "+" * (p[1] - p[0] + 1))


if __name__ == "__main__":
    input = [
        [3, 9],
        [7, 12],
        [3, 8],
        [6, 8],
        [9, 10],
        [2, 9],
        [0, 9],
        [3, 9],
        [0, 6],
        [2, 8],
    ]
    input = sorted(input)
    visualise(input)

    solution = Solution()
    output = solution.findMinArrowShots(input)

    print(output)
    visualise(input)
