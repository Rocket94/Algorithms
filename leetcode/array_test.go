package leetcode

import "testing"

func TestPrefixesDivBy5(t *testing.T) {
	createPrefixsDivBy5(t,[]int{1,1,1},[]bool{false,false,false})
	createPrefixsDivBy5(t,[]int{0,1,1},[]bool{true,false,false})
	createPrefixsDivBy5(t,[]int{0,1,1,1,1,1},[]bool{true,false,false,false,true,false})
	createPrefixsDivBy5(t,[]int{1,1,1,0,1},[]bool{false,false,false,false,false})
	createPrefixsDivBy5(t,[]int{1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0},[]bool{false,false,true,false,false,false,false,false,false,false,true,true,true,true,true,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,true,false,false,true,false,false,true,true,true,true,true,true,true,false,false,true,false,false,false,false,true,true})
}
func createPrefixsDivBy5(t *testing.T,A []int,expect []bool){
	t.Helper()
	result:=PrefixesDivBy5(A)
	for i,v:=range result{
		if v!=expect[i]{
			t.Errorf("%v expect %v ,but resut %v",A,expect,result)
		}
	}
}
