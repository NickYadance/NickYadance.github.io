package codetop

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CharAt(s string, i int) string {
	return s[i : i+1]
}
