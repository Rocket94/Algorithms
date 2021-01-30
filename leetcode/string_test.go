package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if LengthOfLongestSubstring("abcabcbb")!=3{
		t.Error("error",LengthOfLongestSubstring("\"abcabcbb\""))
	}
}
