package main

import (
	"advent-of-code-2024/day_18/solution"
	"fmt"
	"log"
	"time"
)

const (
	testPath   = "day_18/test.txt"
	normalPath = "day_18/input.txt"
)

func main() {
	defer solution.Track(time.Now(), "main")

	// input_path := testPath
	input_path := normalPath

	input, err := solution.ReadInput(input_path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to read input %v\n", err))
	}

	// fmt.Printf("input: %v\n", input[0].X)

	// res_one := solution.PartOne(input, "test")
	res_one := solution.PartOne(input, "real")
	fmt.Println("Part 1:", res_one)

	// resTwo := solution.PartTwo(inp)
	// log.Println("Result two", resTwo)
}
