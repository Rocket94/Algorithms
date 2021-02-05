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
	begin:=0
	end:=0
	if len(s)<2{
		return s
	}

	arr:=make([][]bool, len(s))
	for i:=0;i<len(s);i++{
		arr[i]=make([]bool,len(s))
		arr[i][i]=true
	}

	for j:=1;j<len(s);j++{
		for i:=0;i<j;i++{
			if s[i]!=s[j]{
				arr[i][j]=false
			}else {
				if j-i<3{
					arr[i][j]=true
				}else {
					arr[i][j]=arr[i+1][j-1]

				}
			}
			if arr[i][j]&&j-i>end-begin{
				begin=i
				end=j
			}
		}
	}

	return s[begin:end+1]
}

