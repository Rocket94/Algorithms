package sort

//假定子数组必然包含中点的情况，找出左右索引和子数组之和
func findMaxCrossingSub(arr []int, low, mid, high int) (maxLeft, maxRight, result int) {
	leftSum := ^int(^uint(0) >> 1)
	sum := 0
	//遍历左边找到最大子数组和
	for i := mid; i >= low; i-- {
		sum = sum + arr[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := ^int(^uint(0) >> 1)
	sum = 0
	//遍历右边找到最大子数组和
	for i := mid + 1; i <= high; i++ {
		sum = sum + arr[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}
func findMaxSub(arr []int, low, high int) (rLow, rHigh, result int) {
	//停止条件：只有一个数字
	if high == low {
		return low, high, arr[low] //只有本身一个子数组,子数组之和为自己本身
	}
	mid := (low + high) / 2
	leftLow, leftHigh, leftResult := findMaxSub(arr, low, mid)                  //分解去前半段找
	rightLow, rightHigh, rightResult := findMaxSub(arr, mid+1, high)            //分解去后半段找
	crossLow, crossHigh, crossResult := findMaxCrossingSub(arr, low, mid, high) //包含中间的情况
	if leftResult>=rightResult&&leftResult>=crossResult{
		return leftLow,leftHigh,leftResult
	} else if rightResult>=crossResult&&rightResult>=leftResult {
		return rightLow,rightHigh,rightResult
	} else {
		return crossLow,crossHigh,crossResult
	}
}