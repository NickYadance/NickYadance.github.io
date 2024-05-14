package dp

import (
	"gotest.tools/assert"
	"testing"
)

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
		dp[k] = (dp[i] && s[i+1:k] in wordDict), i in 0->k
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

/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

提示：

0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成
*/
func minDistance(s1, s2 string) int {
	/*
		dp[i][j] = if s1[i]==s2[j]
					then min(dp[i-1][j-1], dp[i][j-1] + 1, dp[i-1][j] + 1)
					else min(dp[i-1][j-1] + 1, dp[i][j-1] + 1, dp[i-1][j] + 1)
		dp[0][j] = j
		dp[i][0] = i
	*/
	m := len(s1)
	n := len(s2)
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = i
	}

	for i := 1; i <= m; i++ {
		x := i - 1
		dp[0] = i
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if s1[i-1] == s2[j-1] {
				dp[j] = min(x, dp[j-1]+1, dp[j]+1)
			} else {
				dp[j] = min(x+1, dp[j-1]+1, dp[j]+1)
			}
			x = tmp
		}
	}
	return dp[n]
}

/*
给定两个字符串s1 和 s2，返回 使两个字符串相等所需删除字符的 ASCII 值的最小和 。

示例 1:

输入: s1 = "sea", s2 = "eat"
输出: 231
解释: 在 "sea" 中删除 "s" 并将 "s" 的值(115)加入总和。
在 "eat" 中删除 "t" 并将 116 加入总和。
结束时，两个字符串相等，115 + 116 = 231 就是符合条件的最小和。
示例 2:

输入: s1 = "delete", s2 = "leet"
输出: 403
解释: 在 "delete" 中删除 "dee" 字符串变成 "let"，
将 100[d]+101[e]+101[e] 加入总和。在 "leet" 中删除 "e" 将 101[e] 加入总和。
结束时，两个字符串都等于 "let"，结果即为 100+101+101+101 = 403 。
如果改为将两个字符串转换为 "lee" 或 "eet"，我们会得到 433 或 417 的结果，比答案更大。

提示:

0 <= s1.length, s2.length <= 1000
s1 和 s2 由小写英文字母组成
*/
func minimumDeleteSum(s1 string, s2 string) int {
	/*
		dp[i][j] = if s1[i] == s2[j]
					then min(dp[i-1][j-1], dp[i-1][j] + ASCII(s1[i]), dp[i][j-1] + ASCII(s2[j]))
					else min(dp[i-1][j-1] + ASCII(s1[i]) + ASCII(s2[j]), dp[i-1][j] + ASCII(s1[i]), dp[i][j-1] + ASCII(s2[j]))
		dp[0][j] = sum(s2[j])
		dp[i][0] = sum(s1[i])
		dp[0][0] = 0
	*/
	m := len(s1)
	n := len(s2)
	dp := make([]int, n+1)
	ASCII := func(c byte) int {
		return int(c)
	}
	for j := 1; j <= n; j++ {
		dp[j] = dp[j-1] + ASCII(s2[j-1])
	}
	for i := 1; i <= m; i++ {
		x := dp[0]
		dp[0] += ASCII(s1[i-1])
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if s1[i-1] == s2[j-1] {
				dp[j] = min(x, dp[j]+ASCII(s1[i-1]), dp[j-1]+ASCII(s2[j-1]))
			} else {
				dp[j] = min(x+ASCII(s1[i-1])+ASCII(s2[j-1]), dp[j]+ASCII(s1[i-1]), dp[j-1]+ASCII(s2[j-1]))
			}
			x = tmp
		}
	}
	return dp[n]
}

/*
给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数，结果需要对 109 + 7 取模。

示例 1：

输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
rabbbit
rabbbit
rabbbit
示例 2：

输入：s = "babgbag", t = "bag"
输出：5
解释：
如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
babgbag
babgbag
babgbag
babgbag
babgbag

提示：

1 <= s.length, t.length <= 1000
s 和 t 由英文字母组成
*/
func numDistinct(s string, t string) int {
	/*
				dp(i,j)=s[i] == t[j] ?
			        then dp(i-1,j-1) + dp(i-1,j))
			        else dp(i-1,j)
				dp(0,0) = 1
				dp(0,j) = 0
				dp(i,0) = 1
		s = "rabbbit", t = "rabbit"
		s = "babgbag", t = "bag"
	*/
	m := len(s)
	n := len(t)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= m; i++ {
		x := dp[0]
		dp[0] = 1
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if s[i-1] == t[j-1] {
				dp[j] = x + dp[j]
			}
			x = tmp
		}
	}
	return dp[n] % (1000000000 + 7)
}

func Test_numDistinct(t *testing.T) {
	expected := 3
	actual := numDistinct("rabbbit", "rabbit")
	assert.Assert(t, actual == expected, "expected: %d, actual: %d", expected, actual)
	expected = 5
	actual = numDistinct("babgbag", "bag")
	assert.Assert(t, actual == expected, "expected: %d, actual: %d", expected, actual)
	expected = 13
	actual = numDistinct("aaaaaaaaaaaaa", "a")
	assert.Assert(t, actual == expected, "expected: %d, actual: %d", expected, actual)
}
