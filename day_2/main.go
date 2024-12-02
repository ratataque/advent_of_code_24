package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day_2/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	final_list := make([][]int, 0)
	result_2 := 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text()) // Split line into parts
		numbers := make([]int, 0, len(line))

		safe := true

		var asc bool
		for i, num := range line {
			if len(line) < 2 {
				continue
			}

			if n, err := strconv.Atoi(num); err == nil {
				numbers = append(numbers, n)

				if i >= 1 {
					diff := numbers[i] - numbers[i-1]

					// Check if the difference is out of range
					if diff < -3 || diff > 3 || diff == 0 {
						safe = false
						break
					}

					if i == 1 {
						asc = diff > 0
					} else if (asc && diff <= 0) || (!asc && diff >= 0) {
						safe = false
						break
					}
				}
			}
		}
		if safe {
			final_list = append(final_list, numbers)
		}
		// fmt.Printf("numbers: %#v\n", numbers)
	}

	//part 1

	// fmt.Printf("numbers: %#v\n", final_list)
	fmt.Printf("result 1: %v\n", len(final_list))

	// part 2
	fmt.Printf("result 2: %v\n", result_2)

	//logs
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
