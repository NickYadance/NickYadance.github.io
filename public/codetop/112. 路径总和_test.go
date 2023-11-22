package codetop

func printPath(node *TreeNode, path []int, depth int, target int) []int {
	if node == nil {
		return nil
	}
	path = append(path, node.Val)
	if node.Val == target {
		return path
	}
	if left := printPath(node.Left, path, depth+1, target); left != nil {
		return left
	}
	if right := printPath(node.Right, path, depth+1, target); right != nil {
		return right
	}
	return nil
}

func hasSum(root *TreeNode, target int, sum int) bool {
	if root == nil {
		return false
	}

	sum = sum + root.Val
	if root.Left == nil && root.Right == nil && sum == target {
		return true
	}

	return hasSum(root.Left, target, sum) || hasSum(root.Right, target, sum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasSum(root, targetSum, 0)
}
