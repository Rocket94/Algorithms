package leetcode

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var arr []interface{}
	for l != nil {
		arr = append(arr, strconv.FormatInt(int64(l.Val), 10))
		l = l.Next
	}
	return fmt.Sprintf("%v", arr)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var lr, lh *ListNode
	var carry int
	lr, carry = addTwoNode(&l1, &l2, 0)
	lh = lr
	for !(l1 == nil && l2 == nil && carry == 0) {
		lr.Next, carry = addTwoNode(&l1, &l2, carry)
		lr = lr.Next
	}
	return lh
}

func addTwoNode(n1 **ListNode, n2 **ListNode, car int) (*ListNode, int) {
	lr := new(ListNode)
	var carry int
	if *n1 == nil && *n2 != nil {
		lr.Val = ((*n2).Val + car) % 10
		carry = ((*n2).Val + car) / 10
		if (*n2).Next != nil {
			*n2 = (*n2).Next
		} else {
			*n2 = nil
		}
	} else if *n2 == nil && *n1 != nil {
		lr.Val = ((*n1).Val + car) % 10
		carry = ((*n1).Val + car) / 10
		if (*n1).Next != nil {
			*n1 = (*n1).Next
		} else {
			*n1 = nil
		}

	} else if *n1 != nil && *n2 != nil {
		lr.Val = ((*n1).Val + (*n2).Val + car) % 10
		carry = ((*n1).Val + (*n2).Val + car) / 10
		if (*n1).Next != nil {
			*n1 = (*n1).Next
		} else {
			*n1 = nil
		}
		if (*n2).Next != nil {
			*n2 = (*n2).Next
		} else {
			*n2 = nil
		}
	} else {
		lr.Val = (car) % 10
		carry = (car) / 10
	}
	return lr, carry
}
