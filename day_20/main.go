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

	input_path := testPath
	// input_path := normalPath

	grid, start_pos, end_pos := solution.ReadInput(input_path)

	for _, row := range grid {
		fmt.Println(string(row))
	}

	fmt.Printf("words: %v\n", start_pos)
	fmt.Printf("test_words: %v\n", end_pos)

	// res_one := solution.Part_One(words, test_words)
	// res_two := solution.PartTwo(words, test_words)
	// fmt.Println("Result 1:", res_one)
	// fmt.Println("Result 2:", res_two)
}
