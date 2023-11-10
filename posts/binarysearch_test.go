package posts

import (
	"gotest.tools/assert"
	"testing"
)

func BinarySearch(arr []int, n int) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := low + (high-low)>>1
		if n < arr[mid] {
			high = mid
		} else if n == arr[mid] {
			return mid
		} else {
			low = mid + 1
		}
	}
	return -1
}

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, BinarySearch([]int{1, 3, 4, 9, 10, 15}, 0), -1)
	assert.Equal(t, BinarySearch([]int{1, 3, 4, 9, 10, 15}, 1), 0)
	assert.Equal(t, BinarySearch([]int{1, 3, 4, 9, 10, 15}, 3), 1)
	assert.Equal(t, BinarySearch([]int{1, 3, 4, 9, 10, 15}, 4), 2)
	assert.Equal(t, BinarySearch([]int{1, 3, 4, 9, 10, 15}, 9), 3)
	assert.Equal(t, BinarySearch([]int{1, 3, 4, 9, 10, 15}, 10), 4)
	assert.Equal(t, BinarySearch([]int{1, 3, 4, 9, 10, 15}, 15), 5)
	assert.Equal(t, BinarySearch([]int{1, 3, 4, 9, 10, 15}, 20), -1)
	assert.Equal(t, BinarySearch([]int{1, 3, 3, 4, 9, 10, 15}, 3), 2)
}
