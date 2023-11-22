package codetop

func maxSubArray(nums []int) int {
	cur := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur = nums[i]
		} else {
			cur = cur + nums[i]
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
