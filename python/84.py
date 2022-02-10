import random
from typing import List


def largestRectangleArea(heights: List[int]) -> int:
    max_area = 0
    is_higher = [False] * len(heights)
    unique_heights = sorted(set(heights), reverse=True)

    for h in unique_heights:
        if h * len(heights) <= max_area:
            break

        # update is_higher:
        for i, bar in enumerate(heights):
            if not is_higher[i] and bar >= h:
                is_higher[i] = True

        start = None

        for i, cond in enumerate(is_higher):
            # starting condition
            if cond and start is None:
                start = i

            # ending condition
            if start is not None:
                if cond and i == len(heights) - 1:
                    area = h * (i - start + 1)
                elif not cond:
                    area = h * (i - start)
                else:
                    continue

                if area > max_area:
                    max_area = area
                start = None

    return max_area


if __name__ == "__main__":
    input = [random.randrange(0, 10000) for _ in range(10000)]

    output = largestRectangleArea(input)

    print(output)
