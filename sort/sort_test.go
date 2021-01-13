package sort

import (
	"reflect"
	"testing"
)

func TestBubblingSort(t *testing.T) {
	createTestBubblingSort(t,[]int{3,2,1},[]int{1,2,3})
	createTestBubblingSort(t,[]int{3,2,1,2},[]int{1,2,2,3})
	createTestBubblingSort(t,nil,nil)
}
func TestInsertionSort(t *testing.T) {
	createTestInsertSort(t,[]int{3,2,1},[]int{1,2,3})
	createTestInsertSort(t,[]int{3,2,1,2},[]int{1,2,2,3})
	createTestInsertSort(t,nil,nil)
}

func createTestBubblingSort(t *testing.T,a,b []int){
	BubblingSort(a)
	if !reflect.DeepEqual(a,b){
		t.Errorf("the sorted array is %v, but the expected is %v",a,b)
	}
}
func createTestInsertSort(t *testing.T,a,b []int){
	InsertionSort(a)
	if !reflect.DeepEqual(a,b){
		t.Errorf("the sorted array is %v, but the expected is %v",a,b)
	}
}