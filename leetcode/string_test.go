package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if LengthOfLongestSubstring("abcabcbb")!=3{
		t.Error("error",LengthOfLongestSubstring("\"abcabcbb\""))
	}
}

//func TestLengthOfLongestSubstring(t *testing.T) {
//	if LengthOfLongestSubstring("dvdf")!=3{
//		t.Error("errors")
//	}
//}

func TestLongestPalindrome(t *testing.T) {
	if LongestPalindrome("abbac")!="abba"{
		t.Error("errors")
	}
}
