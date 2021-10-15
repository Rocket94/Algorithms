package datastruct

import (
	"fmt"
)

type Stack struct {
	Head *Node 
}

type Node struct {
	Data interface{}
	Next *Node
}

func NewStack()*Stack  {
	s:=&Stack{}
	return s
}

func (s Stack) String() string    {
	Node:=s.Head
	var arr []interface{}
	for Node!=nil {
		arr=append(arr,Node.Data)
		Node=Node.Next
	}
	return fmt.Sprintf("%v",arr)
}

func (s *Stack)Push(Data interface{})  {
	n:=&Node{
		Data: Data,
		Next: s.Head,
	}
	s.Head=n
}

func (s *Stack)Pop()(interface{},bool)  {
	if s.Head==nil{
		return nil,false
	}
	n:=s.Head
	s.Head=s.Head.Next
	return n.Data,true
}




