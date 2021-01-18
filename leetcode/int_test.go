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
