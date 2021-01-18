package leetcode

import "math"

func PrefixesDivBy5(A []int) []bool {
	var result[]bool
	for i,_:=range A{
		var num int
		for j:=0;j<=i;j++{
			num+=A[j]*int(math.Pow(2,float64(i-j)))
		}
		if num%5==0{
			result=append(result, true)
		}else {
			result=append(result,false)
		}
	}
	return result
}
