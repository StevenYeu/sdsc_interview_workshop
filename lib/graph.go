package lib

func MaxAreaOfIsland(grid [][]int) int {
	area := 0
	for r, row := range grid {
		for c, value := range row {
			if value == 1 {
				area = max(GraphDFS(r, c, &grid), area)
			}
		}
	}

	return area
}

type Coord struct {
	Row int
	Col int
}

func GraphDFS(row, col int, grid *[][]int) int {

	area := 0
	rows := len(*grid)
	cols := len((*grid)[0])

	stack := []Coord{{Row: row, Col: col}}

	for len(stack) > 0 {
		n := len(stack) - 1
		cur := stack[n]
		stack = stack[:n]
		r := cur.Row
		c := cur.Col
		if r < 0 || c < 0 || r >= rows || c >= cols {
			continue
		}
		if (*grid)[r][c] == 1 {
			area += 1
			(*grid)[r][c] = 0
			stack = append(stack, Coord{Row: r, Col: c - 1})
			stack = append(stack, Coord{Row: r, Col: c + 1})
			stack = append(stack, Coord{Row: r - 1, Col: c})
			stack = append(stack, Coord{Row: r + 1, Col: c})
		}
	}
	return area
}
