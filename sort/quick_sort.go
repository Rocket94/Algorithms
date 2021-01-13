package sort

func quickSort(A []int, low, high int) {
	//终止条件，左右索引相等，只有一个数时截止
	if low < high {
		q:=partition(A,low,high)//索引q上的数为分界点，左边均小于A[q]，右边均大于A[q]
		//在左右子数组上继续快排，分治思想
		quickSort(A,low,q-1)
		quickSort(A,q+1,high)
	}
}
//把数组第一个数作为基准进行排序，排序结果，左边都小于基准，右边都大于基准，返回基准的索引
func partition(A[]int,low,high int)(q int){
	temp := A[low]//选定基准，相当于low地方是一个空位（想象成把上面的数挖走了）
	for low < high {
		//先从高往低遍历，找一找谁小于基准，小于基准的话就别呆在右边了，去前面那个坑吧（他走了高位就留下一个坑了）
		for ;high>low&&A[high]>temp;{
			high--
		}
		//填坑的时候一定要判断索引是否合法，如果相等了就不填坑（哪有自己填自己的）
		if low<high{
			A[low]=A[high]
			low++//前面的都是小于基准
		}
		//填完之后再从前往后找一个，找到一个大于基准的数，让他填高位的坑
		for ;low<high&&A[low]<temp;{
			low++
		}
		if low<high{
			A[low]=A[high]
			high--//后面的都是大于基准
		}
	}
	A[low]=temp//最后一个坑由基准来填，大小就由此分界
	return low//返回分界线
}