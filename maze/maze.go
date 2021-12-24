package maze

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func MainMaze() {
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
	fmt.Println("-------------->")
	steps, step := walk(maze, point{0, 0}, point{2, 2})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
	fmt.Printf("共 %d 步\n", step)
}

type point struct {
	i, j int
}

/*
	上 左 下 右
*/
var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end point) ([][]int, int) {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			/*
				maze at next is 0
				and steps at next is 0
				and next != start
			*/
			if value, ok := next.at(maze); !ok || value == 1 {
				continue
			}
			if value, ok := next.at(steps); !ok || value != 0 {
				continue
			}
			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}
	return steps, steps[end.i][end.j]
}
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (value int, ok bool) {
	if p.i < 0 || p.i >= len(grid) || p.j < 0 || p.j >= len(grid[0]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}
