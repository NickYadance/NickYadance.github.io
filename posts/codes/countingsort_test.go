package codes

import (
	"gotest.tools/assert"
	"testing"
)

func CountingSortNaive(nums []int, N int) {
	counting := make([]int, N)
	for _, num := range nums {
		counting[num] = counting[num] + 1
	}
	k := 0
	for i, count := range counting {
		for count > 0 {
			nums[k] = i
			k = k + 1
			count = count - 1
		}
	}
}

func CountingSort(nums []int, N int) {
	counting := make([]int, N)
	tmp := make([]int, len(nums))
	for i := range nums {
		counting[nums[i]] = counting[nums[i]] + 1
	}
	for i := range counting {
		counting[i] = counting[i] + counting[i-1]
	}
	for i := len(nums) - 1; i >= 0; i = i - 1 {
		v := nums[i]
		tmp[counting[v]-1] = v
		counting[v] = counting[v] - 1
	}
	for i := 0; i < len(tmp); i = i + 1 {
		nums[i] = tmp[i]
	}
}

func TestCountingSort(t *testing.T) {
	nums := []int{0, 4, 3, 0, 2, 1}
	CountingSortNaive(nums, 5)
	assert.DeepEqual(t, nums, []int{0, 0, 1, 2, 3, 4})
	nums = []int{0, 4, 3, 0, 2, 1}
	CountingSort(nums, 5)
	assert.DeepEqual(t, nums, []int{0, 0, 1, 2, 3, 4})
}
