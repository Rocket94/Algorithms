package sort

import (
	"Algorithms/datastruct"
)

//堆排序，时间复杂度为nlg(n)
func HeapSort(A[]int){
	h:=datastruct.BuildHeap(A)//初始化建堆
	for i:=h.Count;i>0;i--{
		h.A[i],h.A[1]=h.A[1],h.A[i]//把最大值交换最后
		h.Count--
		h.Heapify(1)//重新调整堆使其变成最大堆
	}
	for i:=0;i<len(A);i++{
		A[i]=h.A[i+1]
	}
}