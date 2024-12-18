package solution

import (
	// "fmt"
	"time"
)

func PartOne(input []Coord, test string) int {
	defer Track(time.Now(), "Part 1")

	grid_size := 71
	if test == "test" {
		grid_size = 7
	}

	input = input[:1024]

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

	step := bfs(grid, Coord{0, 0}, Coord{grid_size - 1, grid_size - 1})

	return step
}

func PartTwo(input []Coord, test string) Coord {
	defer Track(time.Now(), "Part 2")

	grid_size := 71
	initial_input := input[:12]
	other_input := input[12:]
	// input = input[:1024]

	if test == "test" {
		grid_size = 7
	} else {
		initial_input = input[:1024]
	}

	grid := make([][]byte, grid_size)

	for i := range grid {
		grid[i] = make([]byte, grid_size)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, coord := range initial_input {
		grid[coord.Y][coord.X] = '#'
	}

	coord := Coord{0, 0}
	for _, curr_coord := range other_input {
		grid[curr_coord.Y][curr_coord.X] = '#'
		step := bfs(grid, Coord{0, 0}, Coord{grid_size - 1, grid_size - 1})
		if step == -1 {
			coord = curr_coord
			break
		}
	}

	// fmt.Printf("step total: %v\n", step)
	return coord
}

func bfs(grid [][]byte, start Coord, end Coord) int {

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
