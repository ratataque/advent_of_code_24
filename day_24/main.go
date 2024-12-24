package main

import (
	"advent-of-code-2024/day_24/solution"
	"fmt"
	"time"
)

const (
	testPath   = "day_24/test.txt"
	normalPath = "day_24/input.txt"
)

func main() {
	defer solution.Track(time.Now(), "main")

	// input_path := testPath
	input_path := normalPath

	wire, operations := solution.ReadInput(input_path)

	res_one := solution.Part_One(wire, operations)
	res_two := solution.Part_Two(wire, operations)
	fmt.Println("Result 1:", res_one)
	fmt.Println("Result 2:", res_two)
}
