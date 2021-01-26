package leetcode

//func lengthOfLongestSubstring(s string) int {
//
//}
//移动括号法，双index法
func LengthOfLongestSubstring(s string) int {
	var maxLength int
	var rk = -1
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < len(s) {
			if _, ok := m[s[rk+1]]; ok {
				break
			} else {
				rk++
				m[s[rk]]++
			}
		}
		if rk-i+1 > maxLength {
			maxLength = rk - i + 1
		}
	}
	return maxLength
}
