package codetop

import "testing"

// TODO
// s[n] = if a[n] == a[n-s[n-1]-1] then s[n-1]+1 else 1
func longestPalindrome(s string) string {
	max := 1
	maxI := 0
	dp := make([]int, len(s))
	dp[0] = 1
	for i, ch := range s {
		if i > 0 {
			left := i - dp[i-1] - 1
			if left >= 0 && CharAt(s, left) == string(ch) {
				dp[i] = dp[i-1] + 2
			} else if CharAt(s, i-1) == string(ch) {
				dp[i] = 2
			} else {
				dp[i] = 1
			}
			if dp[i] > max {
				max = dp[i]
				maxI = i
			}
		}
	}
	return s[maxI-dp[maxI]+1 : maxI+1]
}

func Test5(t *testing.T) {
	longestPalindrome("ccc")
}
