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

//动态规划求最长回文子串
func LongestPalindrome(s string) string {
	begin := 0
	end := 0
	if len(s) < 2 {
		return s
	}

	arr := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		arr[i] = make([]bool, len(s))
		arr[i][i] = true
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				arr[i][j] = false
			} else {
				if j-i < 3 {
					arr[i][j] = true
				} else {
					arr[i][j] = arr[i+1][j-1]

				}
			}
			if arr[i][j] && j-i > end-begin {
				begin = i
				end = j
			}
		}
	}

	return s[begin : end+1]
}

//回溯法求所有手机字母组合
var (
	phoneMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	res []string
)

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var builder = ""
	res = []string{}
	backtrack(digits, builder, 0)

	return res
}

func backtrack(digits, builder string, index int) {
	if index == len(digits) {
		res = append(res, builder)
	} else {
		letter := phoneMap[string(digits[index])]
		index++
		for i := 0; i < len(letter); i++ {
			//回溯，下层结果回来需要换一个本层的可能
			//23
			//⬆
			//a		b		c
			//
			//23
			// ⬆
			//def	def		def
			builder += string(letter[i])
			backtrack(digits, builder, index)
			builder = builder[0 : len(builder)-1]
		}
	}
}

//最长公共前缀
func LongestCommonPrefix(strs []string) string {
	var res string
	if len(strs) == 0 {
		return res
	}
	if len(strs) == 1 {
		return strs[0]
	}

OuterLoop:
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j])-1 < i || strs[0][0:i+1] != strs[j][0:i+1] {
				break OuterLoop
			}
		}
		res += string(strs[0][i])
	}

	return res
}
