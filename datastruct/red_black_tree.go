package datastruct

import (
	"fmt"
)

//红黑树比二叉搜索树多一个属性：color
////满足下列几个性质：
////1.每个节点都是红色或者黑色的
////2.根节点是黑色的
////3.每个叶节点（NIL）是黑色的
////4.如果一个节点是红色的，那么它的两个子节点都是黑色的
////5.对于每个节点，从该节点到其所有叶子结点的简单路径上，均包含相同数目的黑色节点
////引理：
////1.一棵有n个内部个节点的红黑树的高度至多为2lg(n+1)

const (
	red= iota
	black
)



type RedBlackTreeNode struct {
	color  int
	Data   int
	Left   *RedBlackTreeNode
	Right  *RedBlackTreeNode
	Parent *RedBlackTreeNode
}

type RedBlackTree struct {
	root *RedBlackTreeNode
	nil *RedBlackTreeNode
}

func(T *RedBlackTree)LeftRotate(x *RedBlackTreeNode){
	y:=x.Left
	//换孩子
	if y.Left!=T.nil{
		y.Left.Parent=x
	}
	//换爹
	if x.Parent==T.nil{
		T.root=y
	} else if x==x.Parent.Left{
		x.Parent.Left=y
	}else {
		x.Parent.Right=y
	}
	//x换到孩子上去
	y.Left=x
	x.Parent=y
}

func(T *RedBlackTree)RightRotate(x *RedBlackTreeNode){
	y:=x.Right
	//换孩子
	if y.Right!=T.nil{
		y.Right.Parent=x
	}
	//换爹
	if x.Parent==T.nil{
		T.root=y
	} else if x==x.Parent.Right{
		x.Parent.Right=y
	}else {
		x.Parent.Left=y
	}
	//x换到孩子上去
	y.Right=x
	x.Parent=y
}

func(T *RedBlackTree)Insert(z *RedBlackTreeNode){
	//父指针
	y:=T.nil
	//子指针
	x:=T.root
	for x!=T.nil{
		y=x
		if x.Data>z.Data{
			x=x.Left
		}else {
			x=x.Right
		}
	}
	z.Parent=y
	//空树情况
	if y==T.nil{
		T.root=z
	}else if y.Data>z.Data{
		y.Left=z
	}else{
		y.Right=z
	}
	z.Left=T.nil
	z.Right=T.nil
	z.color=red
}

func(T *RedBlackTree)InsertFixup(z *RedBlackTreeNode){
	for z.Parent.color==red{
		if z.Parent==z.Parent.Parent.Left{
			y:=z.Parent.Parent.Right
			//case 1叔叔也是红
			if y.color==red{
				z.Parent.color=black
				y.color=black
				z.Parent.Parent.color=red
				z=z.Parent.Parent
				//调整完毕，z变为z的爷爷
			} else if z==z.Parent.Right{
				//case 2 z变爷爷的爹并左旋做儿，即原来爷爷兄弟的位置
				z=z.Parent
				T.LeftRotate(z)
			}
			//case 3
			z.Parent.color=black
			z.Parent.Parent.color=red
			T.RightRotate(z.Parent.Parent)
		}else {
			y:=z.Parent.Parent.Left
			//case1
			if y.color==red{
				z.Parent.color=black
				y.color=red
				z.Parent.Parent.color=red
				z=z.Parent.Parent
			} else if z==z.Parent.Left{
				//case 2
				z=z.Parent
				T.RightRotate(z)
			}
			//case 3
			z.Parent.color=black
			z.Parent.Parent.color=red
			T.LeftRotate(z.Parent.Parent)
		}
	}
	T.root.color=black
}

func(t *RedBlackTreeNode) InOrder() {
	if t != nil {
		t.Left.InOrder()
		fmt.Print(t.Data, ",")
		t.Right.InOrder()
	}
}

func(t *RedBlackTreeNode) PreOrder() {
	if t != nil {
		fmt.Print(t.Data, ",")
		t.Left.PreOrder()
		t.Right.PreOrder()
	}
}

func (t *RedBlackTreeNode)PostOrder (){
	if t != nil {
		t.Left.PostOrder()
		t.Right.PostOrder()
		fmt.Print(t.Data, ",")
	}
}