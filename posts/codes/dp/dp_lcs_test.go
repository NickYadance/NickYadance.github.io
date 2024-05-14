package dp

import (
	"gotest.tools/assert"
	"testing"
)

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的
子序列
。

示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
示例 2：

输入：nums = [0,1,0,3,2,3]
输出：4
示例 3：

输入：nums = [7,7,7,7,7,7,7]
输出：1

提示：

1 <= nums.length <= 2500
-104 <= nums[i] <= 104

进阶：

你能将算法的时间复杂度降低到 O(n log(n)) 吗?
*/
func lengthOfLIS(nums []int) int {
	/*
		dp[i] = max(dp[j] + (nums[i] > nums[j]), j in 0 to i-1)
		dp[0] = 1
	*/
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	ans := dp[0]
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

/*
 */
func Test_lengthOfLIS(t *testing.T) {
	assert.Assert(t, lengthOfLIS([]int{4, 10, 4, 3, 8, 9}) == 3)
}

/*
673. 最长递增子序列的个数
中等
相关标签
相关企业
给定一个未排序的整数数组 nums ， 返回最长递增子序列的个数 。

注意 这个数列必须是 严格 递增的。

示例 1:

输入: [1,3,5,4,7]
输出: 2
解释: 有两个最长递增子序列，分别是 [1, 3, 4, 7] 和[1, 3, 5, 7]。
示例 2:

输入: [2,2,2,2,2]
输出: 5
解释: 最长递增子序列的长度是1，并且存在5个子序列的长度为1，因此输出5。

提示:

1 <= nums.length <= 2000
-106 <= nums[i] <= 106

[1,2,4,3,5,4,7,2]
2
3
*/
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	lis := 0
	ans := 0
	dp := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		cnt[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[j]+1 == dp[i] {
					cnt[i] += cnt[j]
				}
			}
		}

		if dp[i] > lis {
			lis = dp[i]
			ans = cnt[i]
		} else if dp[i] == lis {
			ans += cnt[i]
		}
	}

	return ans
}

func Test_findNumberOfLIS(t *testing.T) {
	assert.Assert(t, findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}) == 3)
}
