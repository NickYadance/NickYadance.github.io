package posts

import (
	"gotest.tools/assert"
	"testing"
)

func Merge(nums, tmp []int, l, mid, r int) {
	k := l
	l1 := l
	r1 := mid
	l2 := mid + 1
	r2 := r
	for l1 <= r1 && l2 <= r2 {
		if nums[l1] < nums[l2] {
			tmp[k] = nums[l1]
			l1 = l1 + 1
		} else {
			tmp[k] = nums[l2]
			l2 = l2 + 1
		}
		k = k + 1
	}
	for l1 <= r1 {
		tmp[k] = nums[l1]
		l1 = l1 + 1
		k = k + 1
	}
	for l2 <= r2 {
		tmp[k] = nums[l2]
		l2 = l2 + 1
		k = k + 1
	}
	for l <= r {
		nums[l] = tmp[l]
		l = l + 1
	}
}

func MergeSort(nums, tmp []int, l, r int) {
	if r > l {
		mid := l + (r-l)>>1
		MergeSort(nums, tmp, l, mid)
		MergeSort(nums, tmp, mid+1, r)
		Merge(nums, tmp, l, mid, r)
	}
}

func TestMergeSort(t *testing.T) {
	nums := []int{2, 1, 9, 10, 0, 0}
	tmp := make([]int, len(nums))
	MergeSort(nums, tmp, 0, len(nums)-1)
	assert.DeepEqual(t, nums, []int{0, 0, 1, 2, 9, 10})
}
