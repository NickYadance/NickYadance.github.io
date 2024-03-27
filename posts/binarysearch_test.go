package posts

import (
	"gotest.tools/assert"
	"testing"
)

func BinarySearchLowerBound(arr []int, n int) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := low + (high-low)>>1
		if n > arr[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func BinarySearchUpperBound(arr []int, n int) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := low + (high-low)>>1
		if n >= arr[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func BinarySearch(arr []int, n int) int {
	lowerBound := BinarySearchLowerBound(arr, n)
	if lowerBound >= 0 && lowerBound < len(arr) && arr[lowerBound] == n {
		return lowerBound
	}
	return -1
}

func TestBinarySearchLowerBound(t *testing.T) {
	assert.Equal(t, BinarySearchLowerBound([]int{1, 3, 4, 9, 10, 15}, 0), 0)
	assert.Equal(t, BinarySearchLowerBound([]int{1, 3, 4, 9, 10, 15}, 1), 0)
	assert.Equal(t, BinarySearchLowerBound([]int{1, 3, 4, 9, 10, 15}, 3), 1)
	assert.Equal(t, BinarySearchLowerBound([]int{1, 3, 4, 9, 10, 15}, 4), 2)
	assert.Equal(t, BinarySearchLowerBound([]int{1, 3, 4, 9, 10, 15}, 9), 3)
	assert.Equal(t, BinarySearchLowerBound([]int{1, 3, 4, 9, 10, 15}, 10), 4)
	assert.Equal(t, BinarySearchLowerBound([]int{1, 3, 4, 9, 10, 15}, 15), 5)
	assert.Equal(t, BinarySearchLowerBound([]int{1, 3, 4, 9, 10, 15}, 20), 6)
	assert.Equal(t, BinarySearchLowerBound([]int{1, 3, 3, 3, 4, 9, 10, 15}, 3), 1)
	assert.Equal(t, BinarySearchLowerBound([]int{1, 3, 3, 3, 4, 9, 10, 15}, 5), 5)
}

func TestBinarySearchUpperBound(t *testing.T) {
	assert.Equal(t, BinarySearchUpperBound([]int{1, 3, 4, 9, 10, 15}, 0), 0)
	assert.Equal(t, BinarySearchUpperBound([]int{1, 3, 4, 9, 10, 15}, 1), 1)
	assert.Equal(t, BinarySearchUpperBound([]int{1, 3, 4, 9, 10, 15}, 3), 2)
	assert.Equal(t, BinarySearchUpperBound([]int{1, 3, 4, 9, 10, 15}, 4), 3)
	assert.Equal(t, BinarySearchUpperBound([]int{1, 3, 4, 9, 10, 15}, 9), 4)
	assert.Equal(t, BinarySearchUpperBound([]int{1, 3, 4, 9, 10, 15}, 10), 5)
	assert.Equal(t, BinarySearchUpperBound([]int{1, 3, 4, 9, 10, 15}, 15), 6)
	assert.Equal(t, BinarySearchUpperBound([]int{1, 3, 4, 9, 10, 15}, 20), 6)
	assert.Equal(t, BinarySearchUpperBound([]int{1, 3, 3, 3, 4, 9, 10, 15}, 3), 4)
	assert.Equal(t, BinarySearchUpperBound([]int{1, 3, 3, 3, 4, 9, 10, 15}, 5), 5)
}
