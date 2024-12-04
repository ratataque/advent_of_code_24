package main

import (
	"fmt"
	"os"
	"time"
)

func sumXMAS(bytes []byte) int {
	total_xmas := 0
	// bytes := []byte(input)
	fmt.Printf("bytes: %v\n", bytes)

	line_len := 100
	directions := [][]int{
		{-1, 0},
		{0, -1},
		{-1, -1},
		{1, -1},
	}
	// fmt.Printf("directions: %v\n", directions)
	for i, j := 0, 0; i < len(bytes)-1; i, j = i+1, j+1 {
		// println(i, j)
		// println(bytes[i])
		if bytes[i] == 10 {
			line_len = j + 1
			// fmt.Printf("line_len: %v\n", line_len)
			j = -1
			continue

		}

		for _, dir := range directions {
			// target_max := i - 3*dir[0] - 3*dir[1]*line_len
			x_max := j + (3 * dir[0])
			y_max := i + (3 * dir[1] * line_len)
			if x_max >= 0 && x_max < line_len && y_max >= 0 {
				println(i, j, x_max, y_max)

				// if (bytes[j+i] == 'X' &&
				// 	bytes[j+(1*dir[0])+i+(1*dir[1]*line_len)] == 'M' &&
				// 	bytes[j+(2*dir[0])+i+(2*dir[1]*line_len)] == 'A' &&
				// 	bytes[j+(3*dir[0])+i+(3*dir[1]*line_len)] == 'S') ||
				if bytes[j+i] == 'S' &&
					bytes[j+(1*dir[0])+i+(1*dir[1]*line_len)] == 'A' &&
					bytes[j+(2*dir[0])+i+(2*dir[1]*line_len)] == 'M' &&
					bytes[j+(3*dir[0])+i+(3*dir[1]*line_len)] == 'X' {
					total_xmas += 1
					// println(j+i, j+(1*dir[0])+i+(1*dir[1]*line_len), j+(2*dir[0])+i+(2*dir[1]*line_len), j+(3*dir[0])+i+(3*dir[1]*line_len))
					// fmt.Printf("bytes: %q\n", bytes[j+i])
					// println(i, bytes[i])
				}
			}
			// target := i - 3 - 3*line_len
			// target := i - 3*line_len
			// target := i + 3 - 3*line_len
		}

	}
	return total_xmas
}

func main() {
	start := time.Now()

	// file, _ := os.Open("day_3/input.txt")
	file, _ := os.ReadFile("day_4/input.txt")
	// input_string := string(file)

	fmt.Printf("file: %v\n", file[9])

	bytes_result_tot_1 := 0

	bytes_result_tot_1 += sumXMAS(file)

	//part 1

	// fmt.Printf("numbers: %#v\n", final_list)
	// fmt.Printf("result regex: %v\n", reg_result_tot_1)
	fmt.Printf("result bytes: %v\n", bytes_result_tot_1)

	// part 2
	// fmt.Printf("result 2: %v\n", almostSafeReports)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
