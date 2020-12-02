package datastruct

import "fmt"

type DoubleLinkedList struct {
	Head *DoubleNode
}
type DoubleNode struct {
	Data interface{}
	Pre *DoubleNode
	Next *DoubleNode
}

func NewDLL()*DoubleLinkedList  {
	d:=&DoubleLinkedList{}
	return d
}

func (d *DoubleLinkedList)Insert(Data interface{})  {
	dn:=&DoubleNode{
		Data: Data,
		Pre:  nil,
		Next: d.Head,
	}
	if d.Head!=nil{
		d.Head.Pre=dn
	}
	d.Head=dn
}

func (d *DoubleLinkedList)Search(Data interface{}) bool {
	if d.Head==nil{
		return false
	}
	n:=d.Head
	for n!=nil{
		if n.Data==Data{
			return true
		}
		n=n.Next
	}
	return false
}

func (D DoubleLinkedList) String() string    {
	node:=D.Head
	var arr []interface{}
	for node!=nil {
		arr=append(arr,node.Data)
		node=node.Next
	}
	return fmt.Sprintf("%v",arr)
}