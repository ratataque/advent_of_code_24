package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func canMove(grid [][]byte, current_pos [2]int, direction [2]int) bool {

	queue := [][2]int{current_pos}
	for len(queue) > 0 {
		next_pos := [2]int{queue[0][0] + direction[0], queue[0][1] + direction[1]}
		current_char := grid[queue[0][1]][queue[0][0]]

		if current_char == '[' {
			queue = append(queue, next_pos)
			queue = append(queue, [2]int{next_pos[0] + 1, next_pos[1]})
		} else if current_char == ']' {
			queue = append(queue, next_pos)
			queue = append(queue, [2]int{next_pos[0] - 1, next_pos[1]})
		} else if current_char == '#' {
			return false
		}
		queue = queue[1:]
	}
	return true
}

func move(grid [][]byte, current_pos [2]int, direction [2]int, entity byte) [2]int {
	new_pos := [2]int{current_pos[0] + direction[0], current_pos[1] + direction[1]}
	next_char := &grid[new_pos[1]][new_pos[0]]
	current_char := &grid[current_pos[1]][current_pos[0]]

	if *next_char == '.' {
		*next_char = entity
		*current_char = '.'
		return new_pos

	} else if *next_char == 'O' {
		new_current := move(grid, new_pos, direction, 'O')

		if new_current != new_pos {
			*next_char = entity
			*current_char = '.'

			return new_pos
		}
	} else if *next_char == '[' || *next_char == ']' {
		if direction[1] == 0 {
			new_current := move(grid, new_pos, direction, *next_char)

			if new_current != new_pos {
				*next_char = entity
				*current_char = '.'

				return new_pos
			}

		} else {
			if *next_char == '[' {
				if canMove(grid, new_pos, direction) {
					move(grid, new_pos, direction, *next_char)
					move(grid, [2]int{new_pos[0] + 1, new_pos[1]}, direction, ']')

					*next_char = entity
					*current_char = '.'

					return new_pos
				}
			} else if *next_char == ']' {
				if canMove(grid, new_pos, direction) {
					move(grid, new_pos, direction, *next_char)
					move(grid, [2]int{new_pos[0] - 1, new_pos[1]}, direction, '[')

					*next_char = entity
					*current_char = '.'

					return new_pos
				}
			}

		}
	}

	return current_pos
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("day_15/input.txt")
	input := strings.Split(string(file), "\n\n")

	grid := strings.Split(input[0], "\n")
	instructions := input[1]

	grid_bytes := [][]byte{}
	for _, str := range grid {
		line := []byte(str) // Convert each string to []byte
		grid_bytes = append(grid_bytes, line)
	}

	grid_bytes_2 := [][]byte{}
	for _, str := range grid {
		line := []byte(str) // Convert each string to []byte
		line_bytes := []byte{}
		for _, char := range line {
			if char == '#' || char == '.' {
				line_bytes = append(line_bytes, char)
				line_bytes = append(line_bytes, char)
			} else if char == 'O' {
				line_bytes = append(line_bytes, '[')
				line_bytes = append(line_bytes, ']')
			} else if char == '@' {
				line_bytes = append(line_bytes, '@')
				line_bytes = append(line_bytes, '.')
			}
		}
		grid_bytes_2 = append(grid_bytes_2, line_bytes)
	}

	starting_pos := [2]int{0, 0}
	for y := 0; y < len(grid_bytes)-1; y++ {
		for x := 0; x < len(grid_bytes[0])-1; x++ {
			if grid_bytes[y][x] == '@' {
				starting_pos = [2]int{x, y}
			}
		}
	}

	starting_pos_2 := [2]int{0, 0}
	for y := 0; y < len(grid_bytes_2)-1; y++ {
		for x := 0; x < len(grid_bytes_2[0])-1; x++ {
			if grid_bytes_2[y][x] == '@' {
				starting_pos_2 = [2]int{x, y}
			}
		}
	}

	current_pos := starting_pos
	directions := map[rune][2]int{
		'^': {0, -1},
		'v': {0, 1},
		'<': {-1, 0},
		'>': {1, 0},
	}

	for _, instruction := range instructions {
		if direction, exists := directions[instruction]; exists {
			current_pos = move(grid_bytes, current_pos, direction, '@')
			// for _, line := range grid_bytes {
			// 	fmt.Printf("%q\n", line)
			// }
		}
	}

	current_pos_2 := starting_pos_2
	for _, instruction := range instructions {
		if direction, exists := directions[instruction]; exists {
			current_pos_2 = move(grid_bytes_2, current_pos_2, direction, '@')
			// for _, line := range grid_bytes_2 {
			// 	fmt.Printf("%q\n", line)
			// }
		}
	}

	GPS_1 := 0
	for y, line := range grid_bytes {
		for x, char := range line {
			if char == 'O' {
				GPS_1 += y*100 + x
			}
		}
	}

	GPS_2 := 0
	for y, line := range grid_bytes_2 {
		for x, char := range line {
			if char == '[' {
				GPS_2 += y*100 + x
			}
		}
	}

	// Part 1
	fmt.Printf("Part 1: %v\n", GPS_1)
	// for _, line := range grid_bytes {
	// 	fmt.Printf("%q\n", line)
	// }

	// Part 2
	fmt.Printf("Part 2: %v\n", GPS_2)
	// for _, line := range grid_bytes_2 {
	// 	fmt.Printf("%q\n", line)
	// }

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
