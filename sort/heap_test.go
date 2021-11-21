package sort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	createTestHeapSort(t,[]int{6,35,6,7,2,4,6,9,1,10},[]int{1,2,4,6,6,6,7,9,10,35})
	createTestHeapSort(t,[]int{6,-6,6},[]int{-6,6,6})
	createTestHeapSort(t,[]int{},[]int{})
	createTestHeapSort(t,nil,nil)
}

func createTestHeapSort(t *testing.T,a,b []int){
	HeapSort(a)
	if !reflect.DeepEqual(a,b){
		t.Errorf("the sorted array is %v, but the expected is %v",a,b)
	}
}
