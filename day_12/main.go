package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.ReadFile("day_12/input.txt")

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	field := []map[[2]int]bool{}
	visited := map[[2]int]int{}

	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for y_ori, line := range lines {
		for x_ori, plant := range line {
			if _, existe := visited[[2]int{x_ori, y_ori}]; !existe {
				visited[[2]int{x_ori, y_ori}] = 4

				queue := [][2]int{{x_ori, y_ori}}
				region := map[[2]int]bool{{x_ori, y_ori}: true}
				for len(queue) > 0 {
					for dir := range directions {
						x, y := queue[0][0], queue[0][1]
						new_x := x + directions[dir][0]
						new_y := y + directions[dir][1]

						if new_x >= 0 && new_x < len(lines[y]) && new_y >= 0 && new_y < len(lines) {
							if []rune(lines[new_y])[new_x] == plant {
								if _, test := visited[[2]int{new_x, new_y}]; !test {
									visited[[2]int{new_x, new_y}] = 4
									visited[[2]int{x, y}]--
									region[[2]int{new_x, new_y}] = true
									queue = append(queue, [2]int{new_x, new_y})

								} else {
									visited[[2]int{x, y}]--
								}
							}
						}
					}
					queue = queue[1:]
				}
				field = append(field, region)
			}
		}
	}

	total_1 := 0
	for _, region := range field {
		fence := 0
		for plant := range region {
			fence += visited[plant]
		}
		total_1 += fence * len(region)
	}

	// Part 1
	fmt.Printf("Part 1: %v\n", total_1)

	// Part 2
	// fmt.Printf("Part 2: %v\n", total_2)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
