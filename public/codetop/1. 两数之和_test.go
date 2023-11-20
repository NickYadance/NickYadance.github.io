package codetop

func twoSum(nums []int, target int) []int {
	index := make(map[int]int)
	for i, n := range nums {
		if k, ok := index[n]; ok {
			return []int{i, k}
		}
		index[target-n] = i
	}
	return nil
}
