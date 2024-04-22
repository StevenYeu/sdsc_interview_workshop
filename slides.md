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
- Trie
- Linked List
- Heap

## Techniques/Algorithms

- Two Pointer
- Sliding Window
- Slow/Fast Pointer
- Binary Search
- K Top

---

# Arrays Algorithms

## Good To Know

- Binary Search
- Cyclic Sort
- Two Poitner
- Sliding Window

---

# Arrays - Binary Search

If given a sorted array, binary search would be a good choice for the problem

## Code

```go
///import "cmp"
func BinarySearch[T cmp.Ordered](data []T, target T) int {
    lo := 0
    hi := len(data) - 1
    for lo <= hi {
        mid := (lo + hi) / 2
        value := data[mid]
        if value == target {
            return mid
        }
        if value < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return -1
}

```

---

# Arrays - Two Pointers

---

# Arrays - Sliding Window

---

# Tree/Graph Algorithms

## Most Common

- Depth First Serach (DFS)
- Breadth First Search (BFS)

## Generic Search Algorithm

Below is genric search in psudeocode

```
SerachGraph(startNode)
    nodes := {startNode}
    visited := {}
    while nodes is not empty:
        cur := nodes.getNextNode()
        processNode(cur)
        for n in neighbors of cur:
           add n to nodes
```

### Tree vs Graph Search

Main differecen between a tree and a graph search is whether
we keep track of visited nodes.

In a graph search, keep track of visited nodes

### Note

This structure can also be used to implement A\* or Dijkstra.
Where we can use a proirity queue as the data structure.

Also note that Dijkstra is a speical case of A\*, where the
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

## Exmple Input

```python
///def print_grid_with_borders():
///    grid = [
///        [0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0],
///        [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
///        [0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0],
///        [0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0],
///        [0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0],
///        [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
///        [0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
///        [0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0],
///    ]
///    rows = len(grid)
///    cols = len(grid[0])
///
///    for i in range(rows):
///        if i == 0:
///            print("\u250c" + "\u2500" * (4 * cols - 1) + "\u2510")  # Top border
///        else:
///            print("\u251c" + "\u2500" * (4 * cols - 1) + "\u2524")  # Middle borders
///
///        for j in range(cols):
///            if grid[i][j] == 1:
///                print(
///                    "\u2502 \033[91m1\033[0m", end=" "
///                )  # Red 1 within vertical border
///            else:
///                print(
///                    "\u2502", grid[i][j], end=" "
///                )  # Normal value within vertical border
///        print("\u2502")  # Right vertical border
///
///    print("\u2514" + "\u2500" * (4 * cols - 1) + "\u2518")  # Bottom border
///print_grid_with_borders()
```

---

# Depth First Graph Search Example Continued

## Solution

```go
/// package main
/// import "fmt"
func maxAreaOfIsland(grid [][]int) int {
    area := 0
    for r, row := range grid {
        for c, value := range row {
            if value == 1 {
                area = max(dfs(r, c, grid), area)
            }
        }
    }
    return area
}
///type Coord struct {
///    Row int
///    Col int
///}
func dfs(row, col int, grid [][]int) int {
    area := 0
    rows, cols := len(grid), len(grid[0])
    stack := []Coord{Coord{Row: row, Col: col}}
    for len(stack) > 0 {
        n := len(stack) - 1
        cur := stack[n]
        stack = stack[:n]
        r, c := cur.Row, cur.Col
        if r < 0 ||c < 0 || r >= rows || c >= cols {
            continue
        }
        if grid[r][c] == 1 {
            area += 1
            grid[r][c] = 0
            stack = append(stack, Coord{Row: r, Col: c - 1})
            stack = append(stack, Coord{Row: r, Col: c + 1})
            stack = append(stack, Coord{Row: r- 1, Col: c})
            stack = append(stack, Coord{Row: r+ 1, Col: c})
        }
    }
    return area
}

///func main() {
///    grid := [][]int{
///        {0,0,1,0,0,0,0,1,0,0,0,0,0},
///        {0,0,0,0,0,0,0,1,1,1,0,0,0},
///        {0,1,1,0,1,0,0,0,0,0,0,0,0},
///        {0,1,0,0,1,1,0,0,1,0,1,0,0},
///        {0,1,0,0,1,1,0,0,1,1,1,0,0},
///        {0,0,0,0,0,0,0,0,0,0,1,0,0},
///        {0,0,0,0,0,0,0,1,1,1,0,0,0},
///        {0,0,0,0,0,0,0,1,1,0,0,0,0},
///    }
///    fmt.Printf("Max Area is %d\n",maxAreaOfIsland(grid))
///}

```

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

---

# Trie - Prefix Trees

## Structure

Genral strcuture of a trie in psuedocoe

```
structure Trie[T] {
 isTerminal: Bool
 children: HashMap
 value: T
}
```

## Example

Example of a Trie with following words

- Pikachu
- Pidegot
- Piplup
- Chamander
- Chameleon
- Squirtle

```
root -> p -> i -> k -> a -> c -> h -> u -> (end)
        |    | -> d -> g -> e -> o -> t -> (end)
        |    | -> p -> l -> u -> p -> (end)
        |
        c -> h -> a -> r -> m -> a -> n -> d -> e -> r -> (end)
        |                   | -> e -> l -> e -> o -> n -> (end)
        |
        s -> q -> u -> i -> r -> t -> l -> e -> (end)
```

---

# Trie - Prefix Trees Continued

## Example Implementation

```rust
///use std::collections::HashMap;
struct WordDictionary {
    pub is_word: bool,
    pub children: HashMap<char, WordDictionary>,
}
impl WordDictionary {
    fn new() -> Self {
        return WordDictionary {
            is_word: false,
            children: HashMap::new(),
        };
    }
    fn add_word(&mut self, word: String) {
        let mut cur = self;
        for char in word.chars() {
            let letter = cur.children.get(&char);
            match letter {
                Some(_) => (),
                None => {
                    let node = WordDictionary {
                        is_word: false,
                        children: HashMap::new(),
                    };
                    cur.children.insert(char, node);
                }
            }
            cur = cur.children.get_mut(&char).unwrap()
        }
        cur.is_word = true;
    }
    fn search(&self, word: String) -> bool {
        let mut cur = self;
        for char in word.chars() {
            let letter = cur.children.get(&char);
            match letter {
                Some(a) => cur = &a,
                None => return false,
            }
        }
        cur.is_word
    }
}
```
