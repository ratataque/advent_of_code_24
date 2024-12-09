package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.ReadFile("day_8/input.txt")

	antenas_pos := map[string][][]int{}
	var antenas []string
	max_x, max_y := 0, 0
	lines := strings.Split(string(file), "\n")
	for y, line := range lines[:len(lines)-1] {
		for x, rune := range line {
			char := string(rune)
			if rune != '.' {
				if antenas_pos[char] == nil {
					antenas = append(antenas, char)
				}
				antenas_pos[char] = append(antenas_pos[char], []int{x, y})
			}
			max_x++
		}
		max_y++
		max_x = len(line)
	}

	antinodes := make(map[[2]int]bool)
	antinodes2 := make(map[[2]int]bool)
	for _, antena := range antenas {
		for i, pos := range antenas_pos[antena] {
			other_pos := append(append([][]int{}, antenas_pos[antena][:i]...), antenas_pos[antena][i+1:]...)
			for _, other_po := range other_pos {

				x_p1 := pos[0] + (pos[0] - other_po[0])
				y_p1 := pos[1] + (pos[1] - other_po[1])

				pos_over, pos2_over := false, false
				for j := 1; !pos_over && !pos2_over; j++ {
					x_p2 := pos[0] + j*(pos[0]-other_po[0])
					y_p2 := pos[1] + j*(pos[1]-other_po[1])
					x_p2_2 := pos[0]
					y_p2_2 := pos[1]

					if !pos_over && x_p2 >= 0 && y_p2 >= 0 && x_p2 < max_x && y_p2 < max_y {
						antinodes2[[2]int{x_p2, y_p2}] = true
					} else {
						pos_over = true
					}
					if !pos2_over && x_p2_2 >= 0 && y_p2_2 >= 0 && x_p2_2 < max_x && y_p2_2 < max_y {
						antinodes2[[2]int{x_p2_2, y_p2_2}] = true
					} else {
						pos_over = true
					}
				}

				if x_p1 >= 0 && y_p1 >= 0 && x_p1 < max_x && y_p1 < max_y {
					antinodes[[2]int{x_p1, y_p1}] = true
				}
			}
		}
	}

	// Part 1
	fmt.Printf("Part 1: %v\n", len(antinodes))

	// Part 2
	fmt.Printf("Part 2: %v\n", len(antinodes2))

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
