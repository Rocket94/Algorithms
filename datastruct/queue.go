package datastruct

import (
	"fmt"
)

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Enqueue(Data interface{}) {
	node:=&Node{
		Data: Data,
		Next: nil,
	}
	if q.Tail==nil{
		q.Head=node
		q.Tail=node
	}else {
		q.Tail.Next=node
		q.Tail=node
	}
}

func (q *Queue) Dequeue()(interface{}) {
	if q.Head==nil{
		return nil
	}
	d:=q.Head.Data
	q.Head=q.Head.Next
	if q.Head==nil{
		q.Tail=nil
	}
	return d

}

func (q *Queue) IsEmpty()bool{
	return q.Head==nil
}

func (q Queue) String() string    {
	node:=q.Head
	var arr []interface{}
	for node!=nil {
		arr=append(arr,node.Data)
		node=node.Next
	}
	return fmt.Sprintf("%v",arr)
}

func NewQueue() *Queue {
	q:=&Queue{}
	return q
}

