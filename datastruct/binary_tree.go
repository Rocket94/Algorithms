package datastruct

import "fmt"

type BinaryTree struct {
	Root *TreeNode
}
type TreeNode struct {
	Data   int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func NewTree() *BinaryTree {
	return &BinaryTree{}
}

func InOrder(t *TreeNode) {
	if t != nil {
		InOrder(t.Left)
		fmt.Print(t.Data, ",")
		InOrder(t.Right)
	}
}

func PreOrder(t *TreeNode) {
	if t != nil {
		fmt.Print(t.Data, ",")
		PreOrder(t.Left)
		PreOrder(t.Right)
	}
}

func PostOrder(t *TreeNode) {
	if t != nil {
		PostOrder(t.Left)
		PostOrder(t.Right)
		fmt.Print(t.Data, ",")
	}
}

func (t *TreeNode) Serch(data int) *TreeNode {
	if t == nil {
		return nil
	}
	if t.Data == data {
		return t
	} else if data > t.Data {
		return t.Right.Serch(data)
	} else {
		return t.Left.Serch(data)
	}
}

func (t *BinaryTree) Insert(data int) {
	n := &TreeNode{
		Data: data,
	}
	if t.Root == nil {
		t.Root = n
		return
	}
	x := t.Root
	var y *TreeNode
	for x != nil {
		y = x
		if x.Data < data {
			x = x.Right
		} else {
			x = x.Left
		}
	}
	n.Parent = y
	//注意不能直接让x=n，因为x是nil，不是y的子节点
	if y.Data > data {
		y.Left = n
	} else {
		y.Right = n
	}
}
//replace u with v
func (t *BinaryTree) Replace(u, v *TreeNode) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}
func (t *BinaryTree) Delete(z *TreeNode) {
	if z.Left == nil {
		t.Replace(z,z.Right)
	}else if z.Right==nil{
		t.Replace(z,z.Left)
	}else {
		//找到z右子树上最小的值
		y:=z.Right.Minimum()
		if y.Parent!=z{
			//如果最小值不是z的子节点，则需要和z的右子节点换一下，保证z的右子节点左侧没有挂载
			//，因为要让n的左侧挂在到n的右子节点上去
			t.Replace(y,y.Right)
			y.Right=z.Right
			y.Right.Parent=y
		}
		t.Replace(z,y)
		//将n的左子树挂在到y上
		y.Left=z.Left
		z.Left.Parent=y
	}
}

func (n *TreeNode) Minimum() *TreeNode {
	x:=n
	for x!=nil&&x.Left!=nil {
		x = x.Left
	}
	return x
}

func (n *TreeNode) FillNode(data int, left, right, parent *TreeNode) {
	n.Left = left
	n.Right = right
	n.Parent = parent
	n.Data = data
}
