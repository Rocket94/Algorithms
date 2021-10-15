package sort

import "Algorithms/util"

func BubblingSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		var flag = false
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]                       //用来存要插入的数
		j := i - 1                           //不能用局部变量，因为最后插入要参考这个索引
		for ; j >= 0 && temp < arr[j]; j-- { //这里把比较大小的条件前置到循环判断里面了，也可以放在下面
			arr[j+1] = arr[j] //如果插入数字比切片里的数小，就把切片数字向后移腾位置，否则跳出循环
		}
		arr[j+1] = temp //+1是因为循环最后j--，如果不发生循环，插入就是自己位置，j=i-1所以还是要+1，秒啊
	}
}

func SelectionSort(arr []int) {
	for i:=0;i<len(arr);i++{
		min:=util.INT_MAX
		index:=-1
		for j:=i;j<len(arr);j++ {
			if arr[j]<min{
				min=arr[j]
				index=j
			}
		}
		if index!=-1{
			arr[i],arr[index]=arr[index],arr[i]		//交换
		}
	}
}
