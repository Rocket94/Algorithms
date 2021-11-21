package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBF(t *testing.T) {
	createTestBF(t,"baddef","abc",[]int{})
	createTestBF(t,"baddef","dde",[]int{2})
	createTestBF(t,"baddef","d",[]int{2,3})
	createTestBF(t,"aaaaaaaa","a",[]int{0,1,2,3,4,5,6,7})
	createTestBF(t,"abcccdefgg","gg",[]int{8})
	createTestBF(t,"gfsfsdgfsfsd","gfsfsd",[]int{0,6})
	createTestBF(t,"","",[]int{})
}

func TestRK(t *testing.T) {
	createTestRK(t,"baddef","abc",[]int{})
	createTestRK(t,"baddef","dde",[]int{2})
	createTestRK(t,"baddef","d",[]int{2,3})
	createTestRK(t,"aaaaaaaa","a",[]int{0,1,2,3,4,5,6,7})
	createTestRK(t,"gfsfsdgfsfsd","gfsfsd",[]int{0,6})
	createTestRK(t,"","",[]int{})
}

func createTestBF(t *testing.T,main,model string,expect []int){
	res:=BF(main,model)
	if !reflect.DeepEqual(res,expect){
		t.Errorf("the result is %v, but the expected is %v",res,expect)
	}
}

func createTestRK(t *testing.T,main,model string,expect []int){
	res:=RK(main,model)
	if !reflect.DeepEqual(res,expect){
		t.Errorf("the result is %v, but the expected is %v",res,expect)
	}
}

func TestStringHash(t *testing.T) {
	fmt.Print(StringHash([]byte("cba")))
}

func BenchmarkBF(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BF("gfsfsdgfssafasdfsadfsafsazffsdgfsfsdgfssafasdfsadfsafsazffsdgfsfsdgfssafasdfsadfsafsazffsdgfsfsdgfssafasdfsadfsafsazffsd","afsazffsd")
	}
	b.StopTimer()
}

func BenchmarkRK(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RK("gfsfsdgfssafasdfsadfsafsazffsdgfsfsdgfssafasdfsadfsafsazffsdgfsfsdgfssafasdfsadfsafsazffsdgfsfsdgfssafasdfsadfsafsazffsd","afsazffsd")
	}
	b.StopTimer()
}
