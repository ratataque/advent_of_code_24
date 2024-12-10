package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.ReadFile("day_10/input.txt")

	queue := [][3]int{}
	starting_pos := [][2]int{}
	ending_pos := [][2]int{}
	lines := strings.Split(string(file), "\n")
	c := 0
	for y, line := range lines[:len(lines)-1] {
		for x, rune := range line {
			if rune == '0' {
				queue = append(queue, [3]int{x, y, c})
				starting_pos = append(starting_pos, [2]int{x, y})
				c++
			}
			if rune == '9' {
				ending_pos = append(ending_pos, [2]int{x, y})
			}
		}
	}

	i := 0
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	finish_map := make(map[[3]int]bool)
	for len(queue) >= 1 {
		curr_x := queue[i][0]
		curr_y := queue[i][1]
		curr_start := queue[i][2]

		if lines[curr_y][curr_x] == '9' {
			break
		}

		next_cells := [][3]int{}
		for _, dir := range directions {
			next_x := curr_x + dir[0]
			next_y := curr_y + dir[1]

			if next_x >= 0 && next_x < len(lines[0]) && next_y >= 0 && next_y < len(lines)-1 {
				if lines[next_y][next_x] == lines[curr_y][curr_x]+1 {
					next_cells = append(next_cells, [3]int{next_x, next_y, curr_start})

					if lines[next_y][next_x] == '9' {
						finish_map[[3]int{next_x, next_y, curr_start}] = true
					}
				}
			}
		}
		queue = append(queue[1:], next_cells...)
	}

	// Part 1
	fmt.Printf("Part 1: %v\n", len(finish_map))

	// Part 2
	fmt.Printf("Part 2: %v\n", len(queue))

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
