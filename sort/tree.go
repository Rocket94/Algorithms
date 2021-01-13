package sort

type TreeNode struct {
	key interface{}
	parent *TreeNode
	left *TreeNode
	right *TreeNode
	height int
	size int
}
