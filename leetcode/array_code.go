package leetcode
//二进制转10进制，能否被5整除，只需要看个位就行，因此每次运算只用个位
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
//数组里找两个相加得目标值，利用map存储，时间复杂度n
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
//删除数组中重复数，空间复杂度限制为1，利用原数组前段储存新数组，思路，新加索引作为新数组的实际索引
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

//是否是回文，思路，用翻转后的数字和原来数字比较
func IsPalindrome(x int) bool {
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

//双指针法，求水桶最大容量，两边谁短谁缩长，寻找更大宽，复杂度N
func maxArea(height []int) int {
	if len(height)<2{
		return 0
	}
	var right = len(height)-1
	var left = 0
	var maxWater int
	for left!=right{
		if height[right]<height[left]{
			maxWater=max(maxWater,(right-left)*height[right])
			right--
		}else {
			maxWater=max(maxWater,(right-left)*height[left])
			left++
		}
	}
	return maxWater
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
