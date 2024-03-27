package main

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Path(price, cuttingPoint []int, length int) string {
	if length > 0 {
		return fmt.Sprintf("%s+%d[%d]", Path(price, cuttingPoint, length-cuttingPoint[length]), cuttingPoint[length], price[cuttingPoint[length]])
	}
	return ""
}

func CuttingProblem(price []int, length int) int {
	dp := make([]int, length+1)
	cuttingPoint := make([]int, length+1)
	for i := 1; i <= length; i++ {
		for j := 1; j <= i; j++ {
			if price[j]+dp[i-j] > dp[i] {
				dp[i] = price[j] + dp[i-j]
				cuttingPoint[i] = j
			}
		}
	}

	fmt.Println(Path(price, cuttingPoint, length))

	return dp[length]
}

func TestDp(t *testing.T) {
	price := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	assert.Equal(t, CuttingProblem(price, 1), 1)
	assert.Equal(t, CuttingProblem(price, 2), 5)
	assert.Equal(t, CuttingProblem(price, 3), 8)
	assert.Equal(t, CuttingProblem(price, 4), 10)
	assert.Equal(t, CuttingProblem(price, 5), 13)
	assert.Equal(t, CuttingProblem(price, 6), 17)
	assert.Equal(t, CuttingProblem(price, 7), 18)
	assert.Equal(t, CuttingProblem(price, 8), 22)
	assert.Equal(t, CuttingProblem(price, 9), 25)
	assert.Equal(t, CuttingProblem(price, 10), 30)
}
