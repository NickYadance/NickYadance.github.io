package codetop

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	i := 0
	max := 0
	for j, c := range s {
		if index, ok := m[fmt.Sprintf("%c", c)]; ok {
			i = Max(index+1, i)
		}
		if j-i+1 > max {
			max = j - i + 1
		}
		m[fmt.Sprintf("%c", c)] = j
	}
	return max
}

func Test(t *testing.T) {
	assert.Equal(t, lengthOfLongestSubstring("abcabcbb"), 3)
	assert.Equal(t, lengthOfLongestSubstring("bbbbb"), 1)
	assert.Equal(t, lengthOfLongestSubstring("pwwkew"), 3)
	assert.Equal(t, lengthOfLongestSubstring("abba"), 2)
}
