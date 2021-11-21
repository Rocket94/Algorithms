package algorithms

import (
	"testing"
)

func TestCutRod(t *testing.T) {
	createCutRod(t,1,1)
	createCutRod(t,2,5)
	createCutRod(t,3,8)
	createCutRod(t,4,10)
	createCutRod(t,5,13)
	createCutRod(t,6,17)
	createCutRod(t,7,18)
	createCutRod(t,8,22)
	createCutRod(t,9,25)
	createCutRod(t,10,30)
}

func createCutRod2(t *testing.T,input,expected int){
	if CutRod([]int{1,5,8,9,10,17,17,20,24,30},input)!=expected{
		t.Errorf("input=%v,expected=%v",input,expected)
	}
}
func createCutRod1(t *testing.T,input,expected int){
	if CutRodTopToBottom([]int{1,5,8,9,10,17,17,20,24,30},input)!=expected{
		t.Errorf("input=%v,expected=%v",input,expected)
	}
}
func createCutRod(t *testing.T,input,expected int){
	if CutRodBottomToTop([]int{1,5,8,9,10,17,17,20,24,30},input)!=expected{
		t.Errorf("input=%v,expected=%v",input,expected)
	}
}


func BenchmarkCutRod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CutRod([]int{1,5,8,9,10,17,17,20,24,30},10)
	}
	b.StopTimer()
}

//func BenchmarkCutRodTopToBottom(b *testing.B) {
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		CutRodTopToBottom([]int{1,5,8,9,10,17,17,20,24,30},10)
//	}
//	b.StopTimer()
//}
//
//func BenchmarkCutRodBottomToTop(b *testing.B) {
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		CutRodBottomToTop([]int{1,5,8,9,10,17,17,20,24,30},10)
//	}
//	b.StopTimer()
//}

func TestLCSLength(t *testing.T) {
	if LCSLength([]string{"a","b","c","b","d","a","b"},[]string{"b","d","c","a","b","a"})!=4{
		t.Error("error!")
	}
}