package main

import (
	"advent-of-code-2024/day_23/solution"
	"fmt"
	"time"
)

const (
	testPath   = "day_23/test.txt"
	normalPath = "day_23/input.txt"
)

func main() {
	defer solution.Track(time.Now(), "main")

	// input_path := testPath
	input_path := normalPath

	computer_list := solution.ReadInput(input_path)

	// fmt.Printf("grid: %v\n", computer_list)

	res_one := solution.Part_One(computer_list)
	res_two := solution.Part_Two(computer_list)
	fmt.Println("Result 1:", res_one)
	fmt.Println("Result 2:", res_two)
}
