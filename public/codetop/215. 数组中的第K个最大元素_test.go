package codetop

import (
	"gotest.tools/assert"
	"testing"
)

func partition(nums []int, l, r int) int {
	i := l
	j := i
	x := nums[r-1]
	for j < r {
		if nums[j] <= x {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			i = i + 1
		}
		j = j + 1
	}
	return i - 1
}

func qsort(nums []int, l, r, k int) int {
	if l < r {
		pivot := partition(nums, l, r)
		if pivot == k-1 {
			return nums[pivot]
		} else if pivot < k-1 {
			return qsort(nums, pivot+1, r, k)
		} else {
			return qsort(nums, l, pivot, k)
		}
	}
	return -1
}

func findKthLargest(nums []int, k int) int {
	return qsort(nums, 0, len(nums), len(nums)-k+1)
}

func Test215(t *testing.T) {
	assert.Equal(t, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2), 5)
	assert.Equal(t, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4), 4)
}
