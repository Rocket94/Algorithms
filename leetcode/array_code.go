package leetcode

func PrefixesDivBy5(A []int) []bool {
	var result []bool
	var num int
	for _,v:=range A{
		num=num*2+v*1
		num=num%10
		if num==0||num==5{
			result=append(result, true)
		}else {
			result=append(result,false)
		}
	}
	return result
}

func TwoSum(nums []int, target int) []int {
	m:=make(map[int]int)
	var result []int
	for i,v:=range nums{
		brother:=target-v
		if index,ok:=m[brother];ok{
			result=append(result, index)
			result=append(result, i)
			return result
		}
		if _,ok:=m[v];!ok{
			m[v]=i
		}
	}
	return result
}
