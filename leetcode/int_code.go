package leetcode

import "Algorithms/util"
//整数翻转，不能超过in32最大或者最小值，每次对10取余得一位数加到翻转后的数字上去
func Reverse(x int) int {
	var result int64
	for x!=0{
		result=result*10+int64(x%10)
		if result>int64(util.INT32_MAX)||result<int64(util.INT32_MIN){
			return  0
		}
		x=x/10
	}
	return int(result)
}
