package main

import (
	"advent-of-code-2024/day_19/solution"
	"fmt"
	"time"
)

const (
	testPath   = "day_19/test.txt"
	normalPath = "day_19/input.txt"
)

func main() {
	defer solution.Track(time.Now(), "main")

	// input_path := testPath
	input_path := normalPath

	words, test_words := solution.ReadInput(input_path)

	// fmt.Printf("words: %q\n", words)
	// fmt.Printf("test_words: %q\n", test_words)

	res_one := solution.PartOne(words, test_words)
	// res_one := solution.PartOne(input)
	res_two := solution.PartTwo(words, test_words)
	fmt.Println("Result 1:", res_one)
	fmt.Println("Result 2:", res_two)
}
