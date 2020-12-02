package main

import "fmt"

//合并两个数组，索引分别为p~q,q+1~r
func merge(arr []int, p, q, r int) {
	temp := make([]int, len(arr))
	copy(temp, arr)
	i := p     //前切片索引
	j := q + 1 //后切片索引
	k := p     //空切片索引
	for ; k <= r; {
		//谁小先把谁填入数组
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			k++
			i++
		} else {
			arr[k] = temp[j]
			k++
			j++
		}
		//如果前面的先填完了，后面剩下的就统一填进去
		if i > q {
			for ; j <= r; j++ {
				arr[k] = temp[j]
				k++
			}
		}
		//同理后面整完了，前面剩下的统一填进去
		if j > r {
			for ; i <= q; i++ {
				arr[k] = temp[i]
				k++
			}
		}
	}
}
func mergeSort(arr []int, p, r int) {
	if p < r {
		var q int
		q = (p + r) / 2
		mergeSort(arr, p, q)   //前一段继续分
		mergeSort(arr, q+1, r) //后一段继续分
		merge(arr, p, q, r)     //比较两段，按顺序合并成一段
	}
}

func main() {
	//a:=[]int{1,2}
	b := []int{32, 56, 111, 33, 6, 77, 11, 23, 56, 2}
	mergeSort(b, 0, 9)
	fmt.Println("b的值为:", b)
}
