package algorithms

func Bsearch(arr []int, val int) int {
	high := len(arr) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

//变体一：查找第一个值等于给定值的元素
//实现一，向前遍历
//func Bsearch1(arr []int,val int) int {
//	high := len(arr)-1
//	low := 0
//	for low <= high {
//		mid:=low+(high-low)>>1
//		if arr[mid]==val{
//			i:=mid-1
//			for i>=0{
//				if arr[i]!=val{
//					break
//				}else {
//					i--
//				}
//			}
//			i++
//			return i
//		}else if arr[mid]>val{
//			high=mid-1
//		}else {
//			low=mid+1
//		}
//	}
//	return -1
//}
//实现二，在前面的区间继续二分查找
func Bsearch1(arr []int, val int) int {
	high := len(arr) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != val {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

//变体二：查找最后一个值等于给定值的元素
func Bsearch2(arr []int, val int) int {
	high := len(arr) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > val {
			high = mid - 1
		} else if arr[mid] < val {
			low = mid + 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != val {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

//变体三：查找第一个大于等于给定值的元素
func Bsearch3(arr []int, val int) int {
	high := len(arr) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] >= val {
			if mid == 0 || arr[mid-1] < val {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

//变体四：查找最后一个小于等于给定值的元素
func Bsearch4(arr []int, val int) int {
	high := len(arr) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] <= val {
			if mid == len(arr)-1 || arr[mid+1] > val {
				return mid
			} else {
				low = mid + 1
			}
		} else {
			high = mid - 1
		}
	}
	return -1
}
