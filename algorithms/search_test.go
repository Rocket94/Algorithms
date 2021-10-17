package algorithms

import "testing"

func TestBsearch(t *testing.T) {
	createBserchTest(t, []int{0, 1, 2, 3, 4, 5}, 3, 3)
	createBserchTest(t, []int{1, 13, 23, 32, 41, 99, 999, 9999}, 3, -1)
	createBserchTest(t, []int{1, 13, 23, 32, 41, 99, 999, 9999}, 999, 6)
	createBserchTest(t, []int{1, 23, 23, 32, 32, 999, 999, 9999}, 999, 5)
	createBserchTest(t, []int{0}, 0, 0)
	createBserchTest(t, []int{0}, 1, -1)
	createBserchTest(t, nil, 1, -1)
}

func TestBsearch1(t *testing.T) {
	createBserchTest1(t, []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 8, 5)
	createBserchTest1(t, []int{0, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 0, 0)
	createBserchTest1(t, []int{1, 1, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 1, 0)
	createBserchTest1(t, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0)
	createBserchTest(t, []int{0}, 1, -1)
	createBserchTest(t, nil, 1, -1)
}

func createBserchTest(t *testing.T, arr []int, val int, res int) {
	actual := Bsearch(arr, val)
	if res != actual {
		t.Errorf("expect:%v,but found %v", res, actual)
	}
}

func createBserchTest1(t *testing.T, arr []int, val int, res int) {
	actual := Bsearch1(arr, val)
	if res != actual {
		t.Errorf("expect:%v,but found %v", res, actual)
	}
}
