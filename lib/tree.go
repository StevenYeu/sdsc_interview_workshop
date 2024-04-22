package lib 

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GeneateTestTree() *TreeNode {
	root := &TreeNode{Val: 5}
	left := &TreeNode{Val: 12}
	right := &TreeNode{Val: 41}
	leftLeft := &TreeNode{Val: 23}
	leftRight := &TreeNode{Val: 54}
	rightLeft := &TreeNode{Val: 10}
	rightRight := &TreeNode{Val: 7}

	root.Left = left
	root.Right = right

	right.Left = rightLeft
	right.Right = rightRight

	left.Left = leftLeft
	left.Right = leftRight

	return root
}

