package dp

/*
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

示例 1：

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace" ，它的长度为 3 。
示例 2：

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc" ，它的长度为 3 。
示例 3：

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0 。

提示：

1 <= text1.length, text2.length <= 1000
text1 和 text2 仅由小写英文字符组成。
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	/*
		dp[i][j] = if text[i] == text[j] then dp[i-1][j-1] + 1 else max(dp[i][j-1], dp[i-1][j])
	*/
	m := len(text1)
	n := len(text2)
	dp := make([]int, n+1)
	for i := 1; i <= m; i++ {
		x := dp[0]
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = x + 1
			} else {
				dp[j] = max(dp[j-1], dp[j])
			}
			x = tmp
		}
	}

	return dp[n]
}

/*
在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。

现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足：

	nums1[i] == nums2[j]

且绘制的直线不与任何其他连线（非水平线）相交。
请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。

以这种方法绘制线条，并返回可以绘制的最大连线数。

示例 1：

输入：nums1 = [1,4,2], nums2 = [1,2,4]
输出：2
解释：可以画出两条不交叉的线，如上图所示。
但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相交。
示例 2：

输入：nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]
输出：3
示例 3：

输入：nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]
输出：2

提示：

1 <= nums1.length, nums2.length <= 500
1 <= nums1[i], nums2[j] <= 2000
*/
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	/*
		dp[i][j] = if text[i] == text[j] then dp[i-1][j-1] + 1 else max(dp[i][j-1], dp[i-1][j])
	*/
	m := len(nums1)
	n := len(nums2)
	dp := make([]int, n+1)
	for i := 1; i <= m; i++ {
		x := dp[0]
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if nums1[i-1] == nums2[j-1] {
				dp[j] = x + 1
			} else {
				dp[j] = max(dp[j-1], dp[j])
			}
			x = tmp
		}
	}

	return dp[n]
}
