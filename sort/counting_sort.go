package main

import "fmt"

func countingSort(A[]int,k int)[]int{
	B:=make([]int,len(A))
	C:=make([]int,k)
	for i:=0;i<len(A);i++{
		C[A[i]]=C[A[i]]+1
	}
	for i:=1;i<k;i++{
		C[i]+=C[i-1]
	}
	for i:=len(A)-1;i>0;i--{
		B[C[A[i]]-1]=A[i]
		C[A[i]]-=1
	}
	return B
}
func main()  {
	A:=[]int{5,6,6,3,2,2,1,8,0,0}
	B:=countingSort(A,9)
	fmt.Println(B)
}
