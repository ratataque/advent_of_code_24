package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
	// "strings"
)

type coord struct {
	pos  [2]int
	velo [2]int
}

func ParseData(input string) ([]coord, error) {
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	// fmt.Printf("input: %v\n", input)

	matches := re.FindAllStringSubmatch(input, -1)

	if matches == nil {
		return nil, fmt.Errorf("no valid data found in input")
	}

	var results []coord

	for _, match := range matches {
		coords, err := convertMatchesToIntegers(match[1:])
		if err != nil {
			return nil, err
		}

		results = append(results, coord{
			pos:  [2]int{coords[0], coords[1]},
			velo: [2]int{coords[2], coords[3]},
		})
	}

	return results, nil
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func convertMatchesToIntegers(matches []string) ([]int, error) {
	integers := make([]int, len(matches))

	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to integer: %v", match, err)
		}
		integers[i] = num
	}

	return integers, nil
}

// Function to count how many robots have neighbors
func countRobotsWithNeighbors(robots []coord, len_x, len_y int) int {
	count := 0

	for _, robot := range robots {
		if hasNeighbor(robot, robots, len_x, len_y) {
			count++
		}
	}

	return count
}

func hasNeighbor(robot coord, robots []coord, len_x, len_y int) bool {
	for _, other := range robots {
		// Skip checking the robot against itself
		if robot.pos == other.pos {
			continue
		}

		// Horizontal adjacency
		if (robot.pos[0] == other.pos[0] && abs(int64(robot.pos[1]-other.pos[1])) == 1) ||
			// Vertical adjacency
			(robot.pos[1] == other.pos[1] && abs(int64(robot.pos[0]-other.pos[0])) == 1) {
			return true
		}
	}
	return false
}

func printGrid(robots []coord, len_x, len_y int) {
	grid := make([][]rune, len_y)
	for i := range grid {
		grid[i] = make([]rune, len_x)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for _, robot := range robots {
		grid[robot.pos[1]][robot.pos[0]] = '#'
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("day_14/input.txt")
	input := string(file)

	original_data, _ := ParseData(input)
	robots := make([]coord, len(original_data))

	copy(robots, original_data)

	// len_x := 11
	// len_y := 7
	len_x := 101
	len_y := 103

	max_neighbors := 0
	best_secondes := 0

	for secondes := 0; secondes < 10000; secondes++ {
		for i := range robots {
			robots[i].pos[0] = ((robots[i].pos[0]+robots[i].velo[0])%len_x + len_x) % len_x
			robots[i].pos[1] = ((robots[i].pos[1]+robots[i].velo[1])%len_y + len_y) % len_y
		}
		// Count robots with neighbors
		count := countRobotsWithNeighbors(robots, len_x, len_y)
		// println(count)

		// Draw the grid if this count exceeds the previous maximum
		if count > max_neighbors {
			max_neighbors = count
			best_secondes = secondes + 1
			// fmt.Printf("Iteration: %d, Robots with neighbors: %d\n", secondes, count)
			// printGrid(robots, len_x, len_y)
			// time.Sleep(100 * time.Millisecond)
		}
	}

	quadrant := map[string]int{
		"Top Left":     0,
		"Top Right":    0,
		"Bottom Left":  0,
		"Bottom Right": 0,
	}
	for _, robot := range robots {
		switch {
		case robot.pos[0] < len_x/2 && robot.pos[1] < len_y/2:
			quadrant["Top Left"] += 1
		case robot.pos[0] > len_x/2 && robot.pos[1] < len_y/2:
			quadrant["Top Right"] += 1
		case robot.pos[0] < len_x/2 && robot.pos[1] > len_y/2:
			quadrant["Bottom Left"] += 1
		case robot.pos[0] > len_x/2 && robot.pos[1] > len_y/2:
			quadrant["Bottom Right"] += 1
		}
	}

	safety_factor := 1
	for _, count := range quadrant {
		safety_factor *= count
	}

	// Part 1
	// fmt.Printf("Part 1: %v\n", robots)
	fmt.Printf("Part 1: %v\n", safety_factor)

	// Part 2
	fmt.Printf("Part 2: %v\n", best_secondes)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
