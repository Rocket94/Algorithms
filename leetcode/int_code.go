package leetcode

import "Algorithms/util"

func Reverse(x int) int {
	var result int64
	for x != 0 {
		result = result*10 + int64(x%10)
		if result > int64(util.INT32_MAX) || result < int64(util.INT32_MIN) {
			return 0
		}
		x = x / 10
	}
	return int(result)
}

func PivotIndex(nums []int) int {
	var sum = make([]int, len(nums)+1)
	sum[0]=0
	for i:=0;i<len(nums);i++{
		sum[i+1]=sum[i]+nums[i]
	}
	for i:=1;i<len(sum);i++{
		if sum[i-1]==sum[len(nums)]-sum[i]{
			return i-1
		}
	}
	return -1
}
