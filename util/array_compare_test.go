package util

import "testing"

func TestCheckSliceType(t *testing.T) {
	createTestCheckSliceType(t,[]int{1,2,3},[]int{1,2,3},true)
	createTestCheckSliceType(t,[]byte{'a','b','c'},[]byte{'a','b','c'},true)
	createTestCheckSliceType(t,[]float64{1.2,2.2,3.3},[]float64{1.2,2.2,3.3},true)
	createTestCheckSliceType(t,[]string{"nihao","hello","world"},[]string{"nihao","hello","world"},true)
	createTestCheckSliceType(t,nil,nil,true)
	createTestSliceEqual(t,nil,[]int{1,2,3},false)
	createTestCheckSliceType(t,[]int{'a','b','c'},[]byte{'a','b','c'},false)
	createTestCheckSliceType(t,[]string{"a","b","c"},[]byte{'a','b','c'},false)
}

func createTestCheckSliceType(t *testing.T,a,b interface{},equal bool){
	t.Helper()
	if equal{
		if !CheckSliceType(a,b){
			t.Errorf("expect %v.type==%v.type,but not",a,b)
		}
	}else {
		if CheckSliceType(a,b) {
			t.Errorf("expect %v.type!=%v.type,but not",a,b)
		}
	}

}

func TestSliceEqual(t *testing.T) {
	createTestSliceEqual(t,[]int{1,2,3},[]int{1,2,3},true)
	createTestSliceEqual(t,[]byte{'a','b','c'},[]byte{'a','b','c'},true)
	createTestSliceEqual(t,[]float64{1.2,2.2,3.3},[]float64{1.2,2.2,3.3},true)
	createTestSliceEqual(t,[]string{"nihao","hello","world"},[]string{"nihao","hello","world"},true)
	createTestSliceEqual(t,nil,nil,true)
	createTestSliceEqual(t,nil,[]int{1,2,3},false)
	createTestSliceEqual(t,[]int{'a','b','c'},[]byte{'a','b','c'},false)
	createTestSliceEqual(t,[]int{1,2,3},[]int{4,2,1},false)
	createTestSliceEqual(t,[]int{1,2,3},[]int{1,2,3,4},false)
	createTestSliceEqual(t,[]string{"a","b","c"},[]string{"A","B","C"},false)
}

func createTestSliceEqual(t *testing.T,a,b interface{},equal bool){
	t.Helper()
	if equal{
		if !SliceEqual(a,b){
			t.Errorf("expect %v==%v,but not",a,b)
		}
	}else {
		if SliceEqual(a,b) {
			t.Errorf("expect %v!=%v,but not",a,b)
		}
	}
}