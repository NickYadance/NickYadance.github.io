package dp

/*
给你一个字符串 s，找到 s 中最长的回文
子串
。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成
*/
func longestPalindromeWrongAnswer(s string) string {
	/*
		d[i] = max(s(i-d[i]-1) == s(i) ? d[i-1] + 1 : 1, s[i] == s[i-1] : 2 : 1)
		d[0] = 1
	*/
	dp := make([]int, len(s))
	dp[0] = 1
	ans := dp[0]
	index := 0
	for i := 1; i < len(s); i++ {
		left := i - dp[i-1] - 1
		if left >= 0 && s[left] == s[i] {
			dp[i] = dp[i-1] + 2
		} else if s[i] == s[i-1] {
			dp[i] = 2
		} else {
			dp[i] = 1
		}
		if dp[i] > ans {
			ans = dp[i]
			index = i
		}
	}
	return s[index+1-ans : index+1]
}

func longestPalindrome(s string) string {
	/*
		dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
		dp[i][i+1] = s[i] == s[j]
		dp[i][i] = true
	*/
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	var left, right int
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			j := i + k
			if j >= n {
				continue
			} else if i == j {
				dp[i][j] = true
			} else if j == i+1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] && j-i > right-left {
				left = i
				right = j
			}
		}
	}
	return s[left : right+1]
}

/*
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。

	注意，你可以重复使用字典中的单词。

示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

提示：

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s 和 wordDict[i] 仅由小写英文字母组成
wordDict 中的所有字符串 互不相同
*/
func wordBreak(s string, wordDict []string) bool {
	/*
		dp[k] = (dp[i] && s[i+1:k-1] in wordDict), i in 0->k
	*/
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for k := 1; k <= len(s); k++ {
		for i := 0; i < k; i++ {
			if dp[i] && dict[s[i:k]] {
				dp[k] = true
				break
			}
		}
	}
	return dp[len(s)]
}

/*
给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

示例 1：

输入：s = "bbbab"
输出：4
解释：一个可能的最长回文子序列为 "bbbb" 。
示例 2：

输入：s = "cbbd"
输出：2
解释：一个可能的最长回文子序列为 "bb" 。

提示：

1 <= s.length <= 1000
s 仅由小写英文字母组成
*/
func longestPalindromeSubseq(s string) int {
	/*
		dp[i][j] = (s[i] == s[j]) ? dp[i+1][j-1] + 2
				 = max(dp[i+1][j], dp[i][j-1])
		dp[i][i] = 1
		dp[i][i+1] = s[i] == s[j] ? 2 : 1
		k : 0 -> n
		i : 0 -> n
		j : i + k
	*/
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	ans := 0
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			j := i + k
			if j >= n {
				continue
			} else if i == j {
				dp[i][j] = 1
			} else if j == i+1 {
				if s[i] == s[j] {
					dp[i][j] = 2
				} else {
					dp[i][j] = 1
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				if s[i] == s[j] {
					dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
				}
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}
