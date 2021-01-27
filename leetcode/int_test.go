package leetcode

import (
	"testing"
)

func TestReverse(t *testing.T) {
	createTestReverse(t,1111111112,2111111111)
	createTestReverse(t,1111111113,0)
	createTestReverse(t,0,0)
	createTestReverse(t,123,321)
	createTestReverse(t,-123,-321)
	createTestReverse(t,120,21)
}
func createTestReverse(t *testing.T,input,expected int){
	t.Helper()
	result:=Reverse(input)
	if result!=expected{
		t.Errorf("%v reverse result expected %v,but result is %v",input,expected,result)
	}
}

func TestPivotIndex(t *testing.T) {
	var a=[]int{-1,-1,0,1,1,0}
	if PivotIndex(a)!=5{
		t.Errorf("error:%v",PivotIndex(a))
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	if FindMedianSortedArrays([]int{1,3},[]int{2})!=float64(2){
		t.Error("error")
	}
}
