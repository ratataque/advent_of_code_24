package main

import (
	"advent-of-code-2024/day_25/solution"
	"fmt"
	"time"
)

const (
	testPath   = "day_25/test.txt"
	normalPath = "day_25/input.txt"
)

func main() {
	defer solution.Track(time.Now(), "main")

	// input_path := testPath
	input_path := normalPath

	keys, locks := solution.ParseInput(input_path)

	res_one := solution.Part_One(keys, locks)
	fmt.Println("Result 1:", res_one)
}
