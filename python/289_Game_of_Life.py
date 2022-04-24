from typing import List


class Solution:
    def gameOfLife(self, board: List[List[int]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        m, n = len(board), len(board[0])
        sums = [[0 for _ in range(n)] for _ in range(m)]

        board_padded = [[0 for _ in range(n)]] + board + [[0 for _ in range(n)]]
        for i in range(m + 2):
            board_padded[i] = [0] + board_padded[i] + [0]

        for i in range(1, m + 1):
            for j in range(1, n + 1):
                sums[i - 1][j - 1] = (
                    sum(board_padded[i - 1][j - 1 : j + 2])
                    + board_padded[i][j - 1]
                    + board_padded[i][j + 1]
                    + sum(board_padded[i + 1][j - 1 : j + 2])
                )

        for i in range(m):
            for j in range(n):
                sum_of_neighbours = sums[i][j]
                if sum_of_neighbours == 3:
                    board[i][j] = 1
                elif sum_of_neighbours < 2 or sum_of_neighbours > 3:
                    board[i][j] = 0
