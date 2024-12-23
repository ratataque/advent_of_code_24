package main

import (
	"advent-of-code-2024/day_22/solution"
	"fmt"
	"time"
)

const (
	testPath   = "day_22/test.txt"
	normalPath = "day_22/input.txt"
)

func main() {
	defer solution.Track(time.Now(), "main")

	// input_path := testPath
	input_path := normalPath

	initial_prices := solution.ReadInput(input_path)

	// fmt.Printf("grid: %v\n", initial_prices)

	res_one := solution.Part_One(initial_prices)
	res_two := solution.Part_Two(initial_prices)
	fmt.Println("Result 1:", res_one)
	fmt.Println("Result 2:", res_two)
}
