package sort

func random(A[]int,low,high int)int{
	temp:=A[low]
	for low<high{
		for high>low&&A[high]>temp{
			high--
		}
		if high>low{
			A[low]=A[high]
			low++
		}
		for low<high&&A[low]<temp{
			low++
		}
		if low<high {
			A[high]=A[low]
			high--
		}
	}
	A[low]=temp
	return low
}

func randomSelect(A[]int,p,q,i int)int{
	if p==q{
		return A[p]
	}
	ran:=random(A,p,q)
	if ran-p+1==i{
		return A[ran]
	}else if ran-p+1<i {
		return randomSelect(A,ran+1,q,i-(ran-p+1))
	}else {
		return randomSelect(A,p,ran-1,i)
	}
}