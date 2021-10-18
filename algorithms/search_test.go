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
	createBserchTest1(t, []int{0}, 1, -1)
	createBserchTest1(t, nil, 1, -1)
}

func TestBsearch2(t *testing.T) {
	createBserchTest2(t, []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 8, 7)
	createBserchTest2(t, []int{0, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 0, 0)
	createBserchTest2(t, []int{1, 1, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 1, 2)
	createBserchTest2(t, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 11)
	createBserchTest2(t, []int{1, 2, 3, 3, 3, 4, 5, 6, 7, 7, 8, 8}, 8, 11)
	createBserchTest2(t, []int{0}, 1, -1)
	createBserchTest2(t, nil, 1, -1)
}

func TestBsearch3(t *testing.T) {
	createBserchTest3(t, []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 8, 5)
	createBserchTest3(t, []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 10, 8)
	createBserchTest3(t, []int{0, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 0, 0)
	createBserchTest3(t, []int{1, 1, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 1, 0)
	createBserchTest3(t, []int{1, 1, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 18, 11)
	createBserchTest3(t, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 0)
	createBserchTest3(t, []int{1, 2, 3, 3, 3, 4, 5, 6, 7, 7, 8, 8}, 18, -1)
	createBserchTest3(t, []int{0}, -1, 0)
	createBserchTest3(t, []int{0}, 1, -1)
	createBserchTest3(t, nil, 1, -1)
}

func TestBsearch4(t *testing.T) {
	createBserchTest4(t, []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 8, 7)
	createBserchTest4(t, []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 10, 7)
	createBserchTest4(t, []int{0, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 0, 0)
	createBserchTest4(t, []int{1, 1, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 1, 2)
	createBserchTest4(t, []int{1, 1, 1, 3, 4, 5, 6, 8, 8, 8, 11, 18}, 18, 11)
	createBserchTest4(t, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 11)
	createBserchTest4(t, []int{1, 2, 3, 3, 3, 4, 5, 6, 7, 7, 8, 8}, 18, 11)
	createBserchTest4(t, []int{1, 2, 3, 3, 3, 4, 5, 6, 7, 7, 8, 8}, 0, -1)
	createBserchTest4(t, []int{0}, -1, -1)
	createBserchTest4(t, []int{0}, 1, 0)
	createBserchTest4(t, nil, 1, -1)
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

func createBserchTest2(t *testing.T, arr []int, val int, res int) {
	actual := Bsearch2(arr, val)
	if res != actual {
		t.Errorf("expect:%v,but found %v", res, actual)
	}
}

func createBserchTest3(t *testing.T, arr []int, val int, res int) {
	actual := Bsearch3(arr, val)
	if res != actual {
		t.Errorf("expect:%v,but found %v", res, actual)
	}
}

func createBserchTest4(t *testing.T, arr []int, val int, res int) {
	actual := Bsearch4(arr, val)
	if res != actual {
		t.Errorf("expect:%v,but found %v", res, actual)
	}
}