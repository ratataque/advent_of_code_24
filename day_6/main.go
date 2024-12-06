package main

import (
	// "bufio"
	"fmt"
	"os"
	// "strconv"
	"time"
)

func test_runr(lines [][]byte, starting_pos []int) bool {
	stop := false
	time_loop := false
	visited := make(map[int]map[int]int)

	for i := 0; i < len(lines); i++ {
		visited[i] = make(map[int]int)
	}

	direction := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	current_x, current_y := starting_pos[0], starting_pos[1]
	j := 0
	for i := 0; !stop; i++ {
		if current_x < 0 || current_x > len(lines[0])-1 || current_y < 0 || current_y > len(lines)-1 {
			stop = true
			break
		}

		if lines[current_y][current_x] == '#' {
			current_x, current_y = current_x-1*direction[j%4][0], current_y-1*direction[j%4][1]
			j++
			// visited[current_y][current_x] = 1 + (j % 4)
			current_x, current_y = current_x+1*direction[j%4][0], current_y+1*direction[j%4][1]
			continue
		}

		if visited[current_y][current_x] == 0 {
			visited[current_y][current_x] = 1 + (j % 4)

		} else if visited[current_y][current_x] == 1+(j%4) {
			time_loop = true
			break
		}

		current_x, current_y = current_x+1*direction[j%4][0], current_y+1*direction[j%4][1]
	}

	return time_loop
}

func main() {
	start := time.Now()

	file, _ := os.ReadFile("day_6/input.txt")
	// input_string := string(file)

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
	lines := [][]byte{}
	starting_pos := []int{0, 0}
	for y := 0; y < len(file); y++ {
		for x := 0; y < len(file); y, x = y+1, x+1 {
			if file[y] == '^' {
				starting_pos = []int{x, (y/len(lines[0]) - 1)}
				// println("found")
			}

			if file[y] == '\n' {
				lines = append(lines, file[y-x:y])
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
		if current_x < 0 || current_x > len(lines[0])-1 || current_y < 0 || current_y > len(lines)-1 {
			stop = true
			break
		}

		if lines[current_y][current_x] == '#' {
			current_x, current_y = current_x-1*direction[j%4][0], current_y-1*direction[j%4][1]
			j++
			continue
		}

		if !visited[current_y][current_x] {
			count += 1
			visited[current_y][current_x] = true
		}

		current_x, current_y = current_x+1*direction[j%4][0], current_y+1*direction[j%4][1]
	}

	count_2 := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if lines[y][x] == '.' {
				lines[y][x] = '#'

				test := test_runr(lines, starting_pos)
				if test {
					count_2 += 1
				}

				lines[y][x] = '.'
			}
		}
	}

	//part 1
	fmt.Printf("part 1: %v\n", count)

	// part 2
	fmt.Printf("part 2: %v\n", count_2)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
