package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) (max int) {
	// element in the stack
	type stackItem struct {
		node  *TreeNode
		depth int
	}

	stack := []stackItem{{root, 1}}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.node == nil {
			continue
		}

		if curr.depth > max {
			max = curr.depth
		}

		stack = append(
			stack,
			stackItem{curr.node.Left, curr.depth + 1},
			stackItem{curr.node.Right, curr.depth + 1},
		)
	}

	return
}
