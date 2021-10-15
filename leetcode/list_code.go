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
		arr = append(arr, strconv.FormatInt(int64((*l).Val), 10))
		l = l.Next
	}
	return fmt.Sprintf("%v", arr)
}

//两数字相加，每个数字是list存储的逆序数（head是个位数）：思路，进位，
//注意list 的结构体指针需要取其指针才能修改list结构体的指针
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
//另一种思路，先对齐，再相加
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	l1temp,l2temp:=l1,l2
	for l1temp.Next!=nil||l2temp.Next!=nil {
		if l1temp.Next==nil{
			l1temp.Next=new(ListNode)
			l1temp.Next.Val=0
		}
		if l2temp.Next==nil{
			l2temp.Next=new(ListNode)
			l2temp.Next.Val=0
		}
		l1temp=l1temp.Next
		l2temp=l2temp.Next
	}

	carry:=0
	lr:=new(ListNode)
	head:=lr
	for {
		lr.Val=(carry+l1.Val+l2.Val)%10
		carry=(carry+l1.Val+l2.Val)/10
		if l1.Next==nil||l2.Next==nil{
			break
		}
		lr.Next=new(ListNode)
		lr=lr.Next
		l1=l1.Next
		l2=l2.Next
	}
	if carry!=0{
		lr.Next=&ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return head
}

//反转链表，一个一个往前面添加
func reverseList(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	lr:=&ListNode{
		Val:  head.Val,
		Next: nil,
	}
	for head.Next!=nil{
		temp:=&ListNode{
			Val:  head.Next.Val,
			Next: lr,
		}
		head=head.Next
		lr=temp
	}
	return lr
}

//检测链表里是否有环，两个速度不一样的遍历，看最终是否能重合
func hasCycle(head *ListNode) bool {
	if head==nil{
		return false
	}
	ln:=head
	for ln.Next!=nil&&ln.Next.Next!=nil{
		ln=ln.Next.Next
		head=head.Next
		if ln==head{
			return true
		}
	}
	return false
}

//合并两个有序的链表，思路同归并排序里面的合并
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	lr := new(ListNode)
	head := lr
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			lr.Next = l1
			l1 = l1.Next
		} else {
			lr.Next = l2
			l2 = l2.Next
		}
		lr = lr.Next
	}
	if l1 != nil {
		lr.Next = l1
	}
	if l2 != nil {
		lr.Next = l2
	}
	return head.Next
}

//删除倒数第n个值，思路：前后相隔n遍历，然后先开始遍历那个到了头，删除后面遍历的那个节点即可
//注意head的特殊性，要在前面加一个哨兵
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l1 := head
	l2 := &ListNode{
		Val:  0,
		Next: head,
	}
	cnt := 0
	for l1 != nil {
		l1 = l1.Next
		cnt++
		if cnt > n {
			l2 = l2.Next
		}
	}
	if l2.Next != nil {
		if l2.Next==head{
			head=l2.Next.Next
		}else {
			l2.Next = l2.Next.Next
		}
	}
	return head
}
//找中间节点，思路还是前后两个速度遍历链表
func middleNode(head *ListNode) *ListNode {
	l1:=head
	l2:=head
	for l1.Next!=nil&&l1.Next.Next!=nil{
		l1=l1.Next.Next
		l2=l2.Next
	}
	if l1.Next!=nil{
		l2=l2.Next
	}
	return l2
}
