package dp

import (
	"gotest.tools/assert"
	"sort"
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

/*
646. 最长数对链
中等
相关标签
相关企业
给你一个由 n 个数对组成的数对数组 pairs ，其中 pairs[i] = [lefti, righti] 且 lefti < righti 。

现在，我们定义一种 跟随 关系，当且仅当 b < c 时，数对 p2 = [c, d] 才可以跟在 p1 = [a, b] 后面。我们用这种形式来构造 数对链 。

找出并返回能够形成的 最长数对链的长度 。

你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。

示例 1：

输入：pairs = [[1,2], [2,3], [3,4]]
输出：2
解释：最长的数对链是 [1,2] -> [3,4] 。
示例 2：

输入：pairs = [[1,2],[7,8],[4,5]]
输出：3
解释：最长的数对链是 [1,2] -> [4,5] -> [7,8] 。

提示：

n == pairs.length
1 <= n <= 1000
-1000 <= lefti < righti <= 1000
*/
type Pair [][]int

func (p Pair) Len() int           { return len(p) }
func (p Pair) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p Pair) Swap(i, j int)      { tmp := p[i]; p[i] = p[j]; p[j] = tmp }

func findLongestChain(pairs [][]int) int {
	/*
		dp[i] = for j in 0 to i: if pairs[i][0] > pairs[j][1] then dp[j] + 1 else 1
	*/
	sort.Sort(Pair(pairs))
	n := len(pairs)
	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

/*
给你一个整数数组 arr 和一个整数 difference，请你找出并返回 arr 中最长等差子序列的长度，该子序列中相邻元素之间的差等于 difference 。

子序列 是指在不改变其余元素顺序的情况下，通过删除一些元素或不删除任何元素而从 arr 派生出来的序列。

示例 1：

输入：arr = [1,2,3,4], difference = 1
输出：4
解释：最长的等差子序列是 [1,2,3,4]。
示例 2：

输入：arr = [1,3,5,7], difference = 1
输出：1
解释：最长的等差子序列是任意单个元素。
示例 3：

输入：arr = [1,5,7,8,5,3,4,2,1], difference = -2
输出：4
解释：最长的等差子序列是 [7,5,3,1]。

提示：

1 <= arr.length <= 105
-104 <= arr[i], difference <= 104
*/
func longestSubsequence(arr []int, difference int) int {
	/*
		此解法超时，需要利用等差子序列的性质，直接定位到 j = index of (arr[i] - difference)
	*/
	n := len(arr)
	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i]-arr[j] == difference {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

/*
给你一个整数数组 nums，返回 nums 中最长等差子序列的长度。

回想一下，nums 的子序列是一个列表 nums[i1], nums[i2], ..., nums[ik] ，且 0 <= i1 < i2 < ... < ik <= nums.length - 1。并且如果 seq[i+1] - seq[i]( 0 <= i < seq.length - 1) 的值都相同，那么序列 seq 是等差的。

示例 1：

输入：nums = [3,6,9,12]
输出：4
解释：
整个数组是公差为 3 的等差数列。
示例 2：

输入：nums = [9,4,7,2,10]
输出：3
解释：
最长的等差子序列是 [4,7,10]。
示例 3：

输入：nums = [20,1,15,3,10,5,8]
输出：4
解释：
最长的等差子序列是 [20,15,10,5]。

提示：

2 <= nums.length <= 1000
0 <= nums[i] <= 500
*/
func longestArithSeqLength(nums []int) int {
	/*
		1 2 3 4 6 8 10
		dp[i][j] 以i结尾，长度为j的最长等差子序列
		for j in 0 to i:
			nums[i]=v, nums[j]=u
			dp[i][v-u] = dp[j][v-u] + 1
	*/
	dp := make([]map[int]int, len(nums))
	ans := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			delta := nums[i] - nums[j]
			dp[i][delta] = max(dp[j][delta], 1) + 1
			ans = max(ans, dp[i][delta])
		}
	}
	return ans
}

/*
// :optional
给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。

当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

注意：不允许旋转信封。


示例 1：

输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
示例 2：

输入：envelopes = [[1,1],[1,1],[1,1]]
输出：1


提示：

1 <= envelopes.length <= 105
envelopes[i].length == 2
1 <= wi, hi <= 105
*/
