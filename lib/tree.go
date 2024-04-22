package lib

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DFS(root *TreeNode) {
	stack := []*TreeNode{root}
	result := []int{}
	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]
		if node == nil {
			continue
		}
		result = append(result, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	fmt.Printf("DFS Visit Order: %v\n", result)

}
func BFS(root *TreeNode) {
	stack := []*TreeNode{root}
	result := []int{}
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		if node == nil {
			continue
		}
		result = append(result, node.Val)
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	fmt.Printf("BFS Visit Order: %v\n", result)

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
