package leetcode

import (
	"fmt"
	"testing"
)

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

func TestLetterCombinations(t *testing.T) {
	fmt.Println(LetterCombinations("2"))
}

func TestLongestCommonPrefix(t *testing.T) {
	fmt.Println(LongestCommonPrefix([]string{"flower","flow","flight"}))
}
