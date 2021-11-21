package datastruct

import "strconv"

type Heap struct {
	A     []int //数组
	n     int   //可存放最大个数
	Count int   //以存放个数
}

func newHeap(capacity int) *Heap {
	return &Heap{
		A:     []int{},
		n:     capacity,
		Count: 0,
	}
}

func (h *Heap) String() string {
	if h == nil {
		return "nil"
	}
	s := "["
	for i:=1;i<=h.Count;i++ {
		s += strconv.FormatInt(int64(h.A[i]), 10)
		if i != len(h.A)-1 {
			s += ","
		}
	}
	s += "]"
	return s
}

func BuildHeap(arr []int) *Heap {
	h := new(Heap)
	a := []int{}
	//从1开始计数，所以需要重新构建
	a=append(a, 0)
	for i := 0; i < len(arr); i++ {
		a=append(a, arr[i])
	}
	h.A = a
	h.Count = len(arr)
	h.n = len(arr)
	for i:=h.n/2;i>0;i--{
		h.Heapify(i)
	}
	return h
}
//插入就直接插在最后，然后和上面的比大小交换就行，不用维护堆性质运算了
func (h *Heap) Insert(data int) {
	if h.Count >= h.n {
		return
	}
	h.Count++
	h.A[h.Count] = data
	//如果大就往上交换,其余都已经排好序了
	i := h.Count
	for i/2 > 0 && h.A[i] > h.A[i/2] {
		h.A[i], h.A[i/2] = h.A[i/2], h.A[i]
		i = i / 2
	}
}
//删除的过程为了防止产生空隙，把最后一个值补到对顶，然后在进行
func (h *Heap) RemoveTop() int {
	if h.Count == 0 {
		return -1
	}
	res := h.A[1]
	h.A[1] = h.A[h.Count]
	h.Count--
	h.Heapify(1)
	return res
}

//从上往下维护堆的性质，给定根索引，如果不是最大堆，则调整,时间复杂度O(lgn)
func (h *Heap) Heapify(i int) {
	l := 2 * i
	r := 2*i + 1
	var max = i
	//看看左子节点和根谁大
	if l <= h.Count && h.A[l] > h.A[i] {
		max = l
	}
	//再和右子节点比比谁大
	if r <= h.Count && h.A[r] > h.A[max] {
		max = r
	}
	//如果根不是最大，则和最大交换
	if max != i {
		h.A[max], h.A[i] = h.A[i], h.A[max]
		h.Heapify(max) //既然换了，这个子堆也要进行最大堆化的调整
	}
}
