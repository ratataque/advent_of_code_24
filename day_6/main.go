package main

import (
	// "bufio"
	"fmt"
	"os"
	// "strconv"
	"time"
)

func main() {
	start := time.Now()

	file, _ := os.ReadFile("day_6/input.txt")
	input_string := string(file)

	// file, _ := os.Open("day_5/input.txt")
	// defer file.Close()
	//
	// scanner := bufio.NewScanner(file)
	//
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	println(line)
	// }

	// lines := strings.Split(input_string, "\n")
	lines := []string{}
	starting_pos := []int{0, 0}
	for y := 0; y < len(input_string); y++ {
		for x := 0; y < len(input_string); y, x = y+1, x+1 {
			if input_string[y] == '^' {
				starting_pos = []int{x, (y/len(lines[0]) - 1)}
				// println("found")
			}

			if input_string[y] == '\n' {
				lines = append(lines, input_string[y-x:y])
				break
			}
		}
	}

	stop := false
	count := 0
	visited := make(map[int]map[int]bool)

	for i := 0; i < len(lines); i++ {
		visited[i] = make(map[int]bool)
	}

	direction := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	current_x, current_y := starting_pos[0], starting_pos[1]
	j := 0
	for i := 0; !stop; i++ {
		// println(current_x, current_y)
		// current_x, current_y := starting_pos[0]+i*direction[i%4][0], starting_pos[1]+i*direction[i%4][1]
		if current_x < 0 || current_x > len(lines[0])-1 || current_y < 0 || current_y > len(lines)-1 {
			stop = true
			break
		}

		if lines[current_y][current_x] == '#' {
			current_x, current_y = current_x-1*direction[j%4][0], current_y-1*direction[j%4][1]
			j++
		}

		if !visited[current_y][current_x] {
			count += 1
			visited[current_y][current_x] = true
		}

		current_x, current_y = current_x+1*direction[j%4][0], current_y+1*direction[j%4][1]
	}

	//part 1
	fmt.Printf("visited: %v\n", count)

	// part 2
	// fmt.Printf("part 1: %v\n", count_2)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
