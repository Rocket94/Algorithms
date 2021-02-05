package leetcode

func PrefixesDivBy5(A []int) []bool {
	var result []bool
	var num int
	for _,v:=range A{
		num=num*2+v*1
		num=num%10
		if num==0||num==5{
			result=append(result, true)
		}else {
			result=append(result,false)
		}
	}
	return result
}

func TwoSum(nums []int, target int) []int {
	m:=make(map[int]int)
	var result []int
	for i,v:=range nums{
		brother:=target-v
		if index,ok:=m[brother];ok{
			result=append(result, index)
			result=append(result, i)
			return result
		}
		if _,ok:=m[v];!ok{
			m[v]=i
		}
	}
	return result
}

func removeDuplicates(nums []int) int {
	if len(nums)==0{
		return 0
	}
	p:=1
	for i:=1;i< len(nums);i++{
		for j:=i-1;j>=0;j--{
			if nums[j]==nums[i]{
				break
			}
			if j==0{
				nums[p]=nums[i]
				p++
			}
		}
	}
	return p
}

func isPalindrome(x int) bool {
	if x<0{
		return false
	}
	y:=0
	temp:=x
	for temp!=0{
		y=y*10+temp%10
		temp=temp/10
	}
	if y!=x{
		return false
	}
	return true
}

//找两个数组的中位数，用切片就行。
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var numsall=make([]int,len(nums1)+len(nums2))
	var i,j,k int
	for k< len(numsall)/2+1{
		if i>=len(nums1){
			for j<len(nums2){
				numsall[k]=nums2[j]
				k++
				j++
			}
			break
		}
		if j>=len(nums2){
			for i<len(nums1){
				numsall[k]=nums1[i]
				k++
				i++
			}
			break
		}
		if nums1[i]<nums2[j]{
			numsall[k]=nums1[i]
			i++
			k++
		}else{
			numsall[k]=nums2[j]
			j++
			k++
		}
	}
	if len(numsall)%2==0{
		return float64(numsall[len(numsall)/2]+numsall[len(numsall)/2-1])/2
	}else{
		return float64(numsall[len(numsall)/2])
	}
}