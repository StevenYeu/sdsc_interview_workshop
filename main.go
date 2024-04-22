package main

import (
	"fmt"
	"presentation/lib"
)

func DFS(root *lib.TreeNode) {
	stack := []*lib.TreeNode{root}
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
func BFS(root *lib.TreeNode) {
	queue := []*lib.TreeNode{root}
	result := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		result = append(result, node.Val)
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	fmt.Printf("BFS Visit Order: %v\n", result)

}
func main() {
	root := lib.GeneateTestTree()
	DFS(root)
    BFS(root)
}
