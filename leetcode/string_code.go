package leetcode

//func lengthOfLongestSubstring(s string) int {
//
//}
//寻找最大连续无重复子字符串，移动括号法，双index法：移动后面的index先，如果遇见重复的就移前面的，知道没重复了，
//注意右边索引要-1以给最后一个留个查看的位置
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