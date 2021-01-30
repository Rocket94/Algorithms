package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	createAddTwoNumbers(t,createList(1,2,3),createList(4,5,6),createList(5,7,9))
	createAddTwoNumbers(t,createList(2,4,3),createList(5,6,4),createList(7,0,8))
	createAddTwoNumbers(t,createList(9,9,9,9,9,9,9),createList(9,9,9,9),createList(8,9,9,9,0,0,0,1))
}

func createAddTwoNumbers(t *testing.T,l1,l2,expect *ListNode){
	result:=addTwoNumbers(l1,l2)
	r:=result
	e:=expect
	for e.Next!=nil{
		if r==nil{
			t.Errorf("expect %v,but result is %v",expect,result)
		}
		if r.Val!=e.Val{
			t.Errorf("expect %v,but result is %v",expect,result)
		}
		r=r.Next
		e=e.Next
	}
	if r.Val!=e.Val||r.Next!=nil{
		t.Errorf("expect %v,but result is %v",expect,result)
	}
}
func createList(args...int)*ListNode{
	lr :=new(ListNode)
	var lh *ListNode
	for i,v:=range args{
		lr.Val=v
		if i==0{
			lh=lr
		}
		if i!= len(args)-1{
			lr.Next=new(ListNode)
			lr=lr.Next
		}

	}
	return lh
}

