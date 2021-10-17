package sort

func CountingSort(A []int) {
	var max = -1
	for _, a := range A {
		if a > max {
			max = a
		}
	}
	//B是排序数组，临时存储用
	B := make([]int, len(A))
	//这里的k是A中的最大值，代表A的范围
	//C是个数数组，存储各个数字的个数
	C := make([]int, max+1)
	//先将A的数字的个数存储在C数组中
	for i := 0; i < len(A); i++ {
		C[A[i]] = C[A[i]] + 1
	}
	//再将C数组中的累加值算出来，每个值代表小于等于index的数量
	for i := 1; i < len(C); i++ {
		C[i] += C[i-1]
	}
	//这里是从后往前遍历。
	for i := len(A) - 1; i >= 0; i-- {
		B[C[A[i]]-1] = A[i]
		C[A[i]] -= 1
	}
	//给A排序
	for i,b:=range B{
		A[i]=b
	}
}
