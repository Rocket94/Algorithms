package main

import (
	"Algorithms/datastruct"
	"fmt"
)

func main()  {
	s:=datastruct.NewTree()
	n1:=&datastruct.TreeNode{}
	n2:=&datastruct.TreeNode{}
	n3:=new(datastruct.TreeNode)
	n4:=&datastruct.TreeNode{}
	n5:=&datastruct.TreeNode{}
	n6:=&datastruct.TreeNode{}
	n7:=&datastruct.TreeNode{}
	n1.FillNode(12,n2,n3,nil)
	n2.FillNode(5,n4,n5,n1)
	n3.FillNode(18,n6,n7,n1)
	n4.FillNode(2,nil,nil,n2)
	n5.FillNode(9,nil,nil,n2)
	n6.FillNode(15,nil,nil,n3)
	n7.FillNode(19,nil,nil,n3)
	s.Root=n1
	fmt.Println(*s.Root)
	datastruct.InOrder(s.Root)
	fmt.Println()
	datastruct.PreOrder(s.Root)
	fmt.Println()
	datastruct.PostOrder(s.Root)
	fmt.Println()
	result := s.Root.Serch(8)
	if result==nil {
		fmt.Println("dont find tree node!")
	}else {
		fmt.Println(*result)
	}
	s.Insert(8)
	datastruct.InOrder(s.Root)
	fmt.Println()
	s.Delete(n3)
	datastruct.InOrder(s.Root)
	fmt.Println()

}



