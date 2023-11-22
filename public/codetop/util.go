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

func ArrayContains(nums []int, x int) bool {
	for _, n := range nums {
		if x == n {
			return true
		}
	}
	return false
}

func ArrayCopy(nums []int) []int {
	res := make([]int, len(nums))
	for i, n := range nums {
		res[i] = n
	}
	return res
}
