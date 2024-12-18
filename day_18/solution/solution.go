package solution

import (
	"fmt"
	"time"
)

func PartOne(input []Coord, test string) int {
	grid_size := 71
	if test == "test" {
		grid_size = 7
	}

	grid := make([][]byte, grid_size)

	for i := range grid {
		grid[i] = make([]byte, grid_size)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, coord := range input {
		grid[coord.Y][coord.X] = '#'
	}

	for _, row := range grid {
		fmt.Printf("%q\n", row)
	}

	step := bfs(grid, Coord{0, 0}, Coord{grid_size - 1, grid_size - 1})

	fmt.Printf("step total: %v\n", step)
	return 0
}

func bfs(grid [][]byte, start Coord, end Coord) int {
	defer Track(time.Now(), "bfs")

	visited := make(map[Coord]bool)
	queue := []struct {
		pos  Coord
		step int
	}{{start, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.pos == end {
			return current.step
		}

		for _, neighbor := range getNeighbors(current.pos) {
			if neighbor.X >= 0 && neighbor.X < len(grid[0]) && neighbor.Y >= 0 && neighbor.Y < len(grid) {
				if grid[neighbor.Y][neighbor.X] == '#' || visited[neighbor] {
					continue
				}
				visited[neighbor] = true

				// Add the neighbor to the queue with an incremented step count
				queue = append(queue, struct {
					pos  Coord
					step int
				}{neighbor, current.step + 1})
			}
		}
	}

	return -1
}

func getNeighbors(coord Coord) []Coord {
	neighbors := []Coord{
		{X: coord.X + 1, Y: coord.Y},
		{X: coord.X, Y: coord.Y - 1},
		{X: coord.X - 1, Y: coord.Y},
		{X: coord.X, Y: coord.Y + 1},
	}

	return neighbors
}
