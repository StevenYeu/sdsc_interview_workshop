# Structure 

- ## Common Techniques for DAS
- ## System Design
- ## Interview Tips

---
# Overview of Data Structures and Techniques

## Data Strcutures
- Arrays
- Hash Tables
- Trees/Graphs
- Linked List
- Heap

## Techniques/Algoritms
- Two Pointer 
- Sliding Window
- Slow/Fast Pointer
- Cyclic Sort
- Binary Search
- K Top
---

# Graph Algorithms

## Most Common 

- Depth First Serach (DFS)
- Breadth First Search (BFS)


## Generic Search Algorithm

Below is genric search in psudeocode

```
SerachGraph(start_node)
    nodes := {start_node}
    visited := {}
    while nodes is not empty:
        cur := nodes.getNextNode()
        add neighbors to nodes
```

### Tree vs Graph Search
Main differecen between a tree and a graph search is whether
we keep track of visited nodes

### Note
This structure can also be used to implement A* or Dijkstra.
Where we can use a proirity queue as the data structure.

Also note that Dijkstra is a speical case of A*, where the 
heureistic function is zero (h(x) = 0)

---

# Depth First Search 

- Goes down each branch as far as possible before backatracking.
- Impleted using recusrion or iteratively using a stack.
- Usefully for find connected componets/disjoint sets in a graph

Following example is an iterative implementation using a stack.
## Example Tree
```
        5
      /   \
    12    41
   /  \  /  \
  23 54 10  7
```

## Code

```go
///package main

///import (
///"fmt"
///"presentation/lib"
///)
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
    fmt.Printf("BFS Visit Order: %v", result)
}
///func main() {
///root := lib.GeneateTestTree()
///DFS(root)
///}
```

---
# Depth First Graph Search Example

## Problem Statement
```
You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)
You may assume all four edges of the grid are surrounded by water.
The area of an island is the number of cells with a value 1 in the island.
Return the maximum area of an island in grid. If there is no island, return 0
```
Taken from [LeetCode 695](https://leetcode.com/problems/max-area-of-island/description/)

---
## Exmple Input

~~~sd replaced processed
This content will be passed in as stdin and will be replaced.
~~~

---
# Breadth First Search Example
- Level order search.
- Implemented with a queue
- Usefully for find short path for uniform cost graph

## Example Tree
```
        5
      /   \
    12    41
   /  \  /  \
  23 54 10  7
```

## Code

```go
///package main

///import (
///"fmt"
///"presentation/lib"
///)
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
///func main() {
///	root := lib.GeneateTestTree()
///	BFS(root)
///}
```
