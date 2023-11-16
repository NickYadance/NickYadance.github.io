package codetop

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func appendIfNotNil(a []*TreeNode, node *TreeNode) []*TreeNode {
	if node != nil {
		return append(a, node)
	} else {
		return a
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// TODO
func zigzagLevelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	var stack []*TreeNode
	var res [][]int
	queue = appendIfNotNil(queue, root)
	for len(queue) > 0 || len(stack) > 0 {
		if len(queue) > 0 {
			var level []int
			for i := range queue {
				if queue[i] != nil {
					level = append(level, queue[i].Val)
					stack = appendIfNotNil(stack, queue[i].Left)
					stack = appendIfNotNil(stack, queue[i].Right)
				}
			}
			queue = nil
			res = append(res, level)
		}

		if len(stack) > 0 {
			var level []int
			for i := len(stack) - 1; i >= 0; i = i - 1 {
				if stack[i] != nil {
					level = append(level, stack[i].Val)
					queue = appendIfNotNil(queue, stack[i].Left)
					queue = appendIfNotNil(queue, stack[i].Right)
				}
			}
			stack = nil
			res = append(res, level)
		}
	}
	return res
}
