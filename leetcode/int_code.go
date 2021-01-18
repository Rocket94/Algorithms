package leetcode

import "Algorithms/util"

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
