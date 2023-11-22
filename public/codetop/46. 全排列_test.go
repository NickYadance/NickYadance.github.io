package codetop

import (
	"fmt"
	"testing"
)

func backtrack(nums []int, path []int, result *[][]int) {
	if len(path) == len(nums) {
		*result = append(*result, ArrayCopy(path))
		return
	}
	for i := 0; i < len(nums); i++ {
		if !ArrayContains(path, nums[i]) {
			path = append(path, nums[i])
			backtrack(nums, path, result)
			path = path[:len(path)-1]
		}
	}
}

func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, []int{}, &result)
	return result
}

func Test46(t *testing.T) {
	fmt.Printf("%+v\n", permute([]int{0, 1}))
}
