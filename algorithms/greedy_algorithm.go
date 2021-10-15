package algorithms

//安排活动，必须是一个结束另一个才能开始，看最多能容纳多少活动
//采用贪心算法，每次选择最早结束的活动，判断一下是否和前一个活动兼容即可
func recursiveActiveSelector(s,f []int,res *[]int,k int){
	m:=k+1
	for k<len(s)&&s[k]<f[m]{
		k++
	}
	if k<len(s){
		*res=append(*res,k)
		recursiveActiveSelector(s,f,res,k)
	}else{
		return
	}
}

func FindMaxCompatibleArray()[]int{
	s:=[]int{1,3,0,5,3,5,6,8,8,2,12}
	f:=[]int{5,4,5,6,7,9,9,10,11,12,14,16}
	var res =[]int{0}
	recursiveActiveSelector(s,f,&res,0)
	return res
}