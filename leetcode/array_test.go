package leetcode

import (
	"fmt"
	"testing"
)

func TestPrefixesDivBy5(t *testing.T) {
	createPrefix(t,[]int{0,1,1,1,1,1},[]bool{true,false,false,false,true,false})
}

func createPrefix(t *testing.T,input []int,expect []bool){
	result:=PrefixesDivBy5(input)
	for i,v:=range result{
		if v!=expect[i]{
			t.Errorf("input=%v,expect=%v,but result is %v",input,expect,result)
		}
	}
}

func TestRemoveDuplicates(t *testing.T){
	nums:=[]int{0,0,1,1,1,2,2,3,3,4}
	i:=removeDuplicates(nums)
	fmt.Println(nums)
	if i!=5{
		t.Errorf("error %v",i)
	}
}

