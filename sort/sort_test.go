package sort

import (
	"reflect"
	"testing"
)

func TestBubblingSort(t *testing.T) {
	createTestBubblingSort(t,[]int{3,2,1},[]int{1,2,3})
	createTestBubblingSort(t,[]int{3,2,1,2},[]int{1,2,2,3})
	createTestBubblingSort(t,[]int{3,1111,9,5,6,4,2,1,2},[]int{1,2,2,3,4,5,6,9,1111})
	createTestBubblingSort(t,nil,nil)
	createTestBubblingSort(t,[]int{},[]int{})
}
func TestInsertionSort(t *testing.T) {
	createTestInsertSort(t,[]int{3,2,1},[]int{1,2,3})
	createTestInsertSort(t,[]int{3,2,1,2},[]int{1,2,2,3})
	createTestInsertSort(t,nil,nil)
	createTestInsertSort(t,[]int{},[]int{})
}
func TestSelectionSort(t *testing.T) {
	createTestSelectionSort(t,[]int{3,2,1},[]int{1,2,3})
	createTestSelectionSort(t,[]int{3,2,1,2},[]int{1,2,2,3})
	createTestSelectionSort(t,[]int{3,1111,9,5,6,4,2,1,2},[]int{1,2,2,3,4,5,6,9,1111})
	createTestSelectionSort(t,nil,nil)
	createTestSelectionSort(t,[]int{},[]int{})
}

func TestQuickSort(t *testing.T) {
	createTestQuickSort(t,[]int{3,2,1},[]int{1,2,3})
	createTestQuickSort(t,[]int{3,2,1,2},[]int{1,2,2,3})
	createTestQuickSort(t,[]int{3,1111,9,5,6,4,2,1,2},[]int{1,2,2,3,4,5,6,9,1111})
	createTestQuickSort(t,nil,nil)
	createTestQuickSort(t,[]int{},[]int{})
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
func createTestSelectionSort(t *testing.T,a,b []int){
	SelectionSort(a)
	if !reflect.DeepEqual(a,b){
		t.Errorf("the sorted array is %v, but the expected is %v",a,b)
	}
}
func createTestQuickSort(t *testing.T,a,b []int){
	quickSort(a,0,len(a)-1)
	if !reflect.DeepEqual(a,b){
		t.Errorf("the sorted array is %v, but the expected is %v",a,b)
	}
}