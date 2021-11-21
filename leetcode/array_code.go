package leetcode

import (
	"Algorithms/sort"
	"fmt"
)

//二进制转10进制，能否被5整除，只需要看个位就行，因此每次运算只用个位
func PrefixesDivBy5(A []int) []bool {
	var result []bool
	var num int
	for _, v := range A {
		num = num*2 + v*1
		num = num % 10
		if num == 0 || num == 5 {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}

//数组里找两个相加得目标值，利用map存储，时间复杂度n
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var result []int
	for i, v := range nums {
		brother := target - v
		if index, ok := m[brother]; ok {
			result = append(result, index)
			result = append(result, i)
			return result
		}
		if _, ok := m[v]; !ok {
			m[v] = i
		}
	}
	return result
}

//删除数组中重复数，空间复杂度限制为1，利用原数组前段储存新数组，思路，新加索引作为新数组的实际索引
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	p := 1
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] == nums[i] {
				break
			}
			if j == 0 {
				nums[p] = nums[i]
				p++
			}
		}
	}
	return p
}

//是否是回文，思路，用翻转后的数字和原来数字比较
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := 0
	temp := x
	for temp != 0 {
		y = y*10 + temp%10
		temp = temp / 10
	}
	if y != x {
		return false
	}
	return true
}

//找两个数组的中位数，用切片就行。
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var numsall = make([]int, len(nums1)+len(nums2))
	var i, j, k int
	for k < len(numsall)/2+1 {
		if i >= len(nums1) {
			for j < len(nums2) {
				numsall[k] = nums2[j]
				k++
				j++
			}
			break
		}
		if j >= len(nums2) {
			for i < len(nums1) {
				numsall[k] = nums1[i]
				k++
				i++
			}
			break
		}
		if nums1[i] < nums2[j] {
			numsall[k] = nums1[i]
			i++
			k++
		} else {
			numsall[k] = nums2[j]
			j++
			k++
		}
	}
	if len(numsall)%2 == 0 {
		return float64(numsall[len(numsall)/2]+numsall[len(numsall)/2-1]) / 2
	} else {
		return float64(numsall[len(numsall)/2])
	}
}

//双指针法，求水桶最大容量，两边谁短谁缩长，寻找更大宽，复杂度N
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var right = len(height) - 1
	var left = 0
	var maxWater int
	for left != right {
		if height[right] < height[left] {
			maxWater = max(maxWater, (right-left)*height[right])
			right--
		} else {
			maxWater = max(maxWater, (right-left)*height[left])
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

//三数之和等于0，不得重复，难点去重
func ThreeSum(nums []int) [][]int {
	//排序
	sort.InsertionSort(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		//终结条件
		if nums[i] > 0 {
			break
		}
		//去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//双指针
		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				//去重
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				temp := []int{nums[i], nums[L], nums[R]}
				res = append(res, temp)
				R--
				L++
			} else if sum > 0 {
				R--
			} else {
				L++
			}
		}
	}
	return res
}

//全排列，回溯法
//根据递归树估算，时间复杂度在O(n*n!)和O(n!)之间
func PrintPermutations(array []int, k int,sum *int) {
	if k == 1 {
		fmt.Print("[")
		for i, num := range array {
			fmt.Print(num,)
			if i!=len(array)-1{
				fmt.Print(",")
			}
		}
		fmt.Println("]")
		*sum++
	}
	for i := 0; i < k; i++ {
		array[i], array[k-1] = array[k-1], array[i]
		PrintPermutations(array, k-1,sum)
		array[i], array[k-1] = array[k-1], array[i]
	}
}
