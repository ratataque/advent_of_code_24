package main

import (
	"advent-of-code-2024/day_20/solution"
	"fmt"
	"time"
)

const (
	testPath   = "day_20/test.txt"
	normalPath = "day_20/input.txt"
)

func main() {
	defer solution.Track(time.Now(), "main")

	// input_path := testPath
	input_path := normalPath

	grid := solution.ReadInput(input_path)

	res_one := solution.Part_One(grid, 100)
	res_two := solution.Part_Two(grid, 100)
	fmt.Println("Result 1:", res_one)
	fmt.Println("Result 2:", res_two)
}
