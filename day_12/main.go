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

	plants := map[rune][2]int{}
	for y, line := range lines[:len(lines)-1] {
		for x, plant := range line {
			if _, test := plants[plant]; !test {
				plants[plant] = [2]int{x, y}
			}
		}
	}

	// visited := map[string]map[[2]int]int{}
	// queue := [][2]int{}
	// for plant := range plants {
	// 	for len(queue) > 0 {
	//
	// 	}
	// }

	// Part 1
	fmt.Printf("Part 1: %v\n", plants)
	// fmt.Printf("Part 1: %v\n", input)

	// Part 2
	// fmt.Printf("Part 2: %v\n", total_2)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
