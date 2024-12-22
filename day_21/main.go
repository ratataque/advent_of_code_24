package main

import (
	"advent-of-code-2024/day_21/solution"
	"fmt"
	"time"
)

const (
	testPath   = "day_21/test.txt"
	normalPath = "day_21/input.txt"
)

func main() {
	defer solution.Track(time.Now(), "main")

	input_path := testPath
	// input_path := normalPath

	input, numeric_pad, direction_pad := solution.ReadInput(input_path)

	// fmt.Printf("grid: %v\n", input)

	res_one := solution.Part_One(input, numeric_pad, direction_pad)
	res_two := solution.Part_Two(input, numeric_pad, direction_pad)
	fmt.Println("Result 1:", res_one)
	fmt.Println("Result 2:", res_two)
}
