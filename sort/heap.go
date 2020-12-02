package main

import "fmt"

//维护堆的性质，给定根索引，如果不是最大堆，则调整,时间复杂度O(lgn)
func maxHeapify(A[]int,i,len int){
	l:=2*i+1
	r:=2*i+2
	var max int
	//看看左子节点和根谁大
	if l<len&&A[l]>A[i]{
		max=l
	}else {
		max=i
	}
	//再和右子节点比比谁大
	if r<len&&A[r]>A[max]{
		max=r
	}
	//如果根不是最大，则和最大交换
	if max!=i{
		A[max],A[i]=A[i],A[max]
		maxHeapify(A,max,len)//既然换了，这个子堆也要进行最大堆化的调整
	}

}
//建堆
func buildMaxHeap(A[]int){
	for i:=(len(A)-1)/2;i>=0;i--{
		maxHeapify(A,i,len(A))
	}
}
//堆排序
func heapSort(A[]int){
	buildMaxHeap(A)//初始化建堆
	fmt.Println(A)
	for i:=len(A)-1;i>0;i--{
		A[i],A[0]=A[0],A[i]//把最大值交换最后
		maxHeapify(A,0,i)//重新调整堆使其变成最大堆
		fmt.Println(A)
	}
}

func main(){
	A:=[]int{9,10,7,15,16,11,14}
	heapSort(A)
	fmt.Println(A)
}
