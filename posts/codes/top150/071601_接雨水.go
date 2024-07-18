package top150

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9

提示：

n == height.length
1 <= n <= 2 * 104
0 <= height[i] <= 105

4,2,0,3,2,5

0 4 4 4 4 4
5 5 5 5 5 0
*/
func trap(height []int) int {
	left := make([]int, len(height))
	right := make([]int, len(height))
	leftMax := 0
	for i := 1; i < len(height); i++ {
		leftMax = max(leftMax, height[i-1])
		left[i] = leftMax
	}

	rightMax := 0
	for i := len(height) - 1 - 1; i >= 0; i-- {
		rightMax = max(rightMax, height[i+1])
		right[i] = rightMax
	}

	sum := 0
	for i := 0; i < len(height); i++ {
		diff := min(left[i], right[i]) - height[i]
		if diff > 0 {
			sum += diff
		}
	}

	return sum
}
