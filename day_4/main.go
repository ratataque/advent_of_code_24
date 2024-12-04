package main

import (
	"fmt"
	"os"
	"time"
)

func sumXMAS(bytes []byte) int {
	total_xmas := 0
	// fmt.Printf("bytes: %v\n", bytes)

	line_len := 140
	directions := [][]int{
		{-1, 0},
		{0, -1},
		{-1, -1},
		{1, -1},
	}
	for i, j := 0, 0; i < len(bytes)-1; i, j = i+1, j+1 {
		if bytes[i] == 10 {
			line_len = j + 1
			j = -1
			continue
		}

		for _, dir := range directions {
			// target_max := i - 3*dir[0] - 3*dir[1]*line_len
			x_max := j + (3 * dir[0])
			y_max := i + (3 * dir[1] * line_len)
			if x_max >= 0 && x_max <= line_len && y_max >= 0 {

				if x_max >= line_len {
					// println("x_max", x_max, j)
				}

				// println(i, bytes[i])
				if (bytes[i+(0*dir[0])+(0*dir[1]*line_len)] == 'S' &&
					bytes[i+(1*dir[0])+(1*dir[1]*line_len)] == 'A' &&
					bytes[i+(2*dir[0])+(2*dir[1]*line_len)] == 'M' &&
					bytes[i+(3*dir[0])+(3*dir[1]*line_len)] == 'X') ||
					(bytes[i+(0*dir[0])+(0*dir[1]*line_len)] == 'X' &&
						bytes[i+(1*dir[0])+(1*dir[1]*line_len)] == 'M' &&
						bytes[i+(2*dir[0])+(2*dir[1]*line_len)] == 'A' &&
						bytes[i+(3*dir[0])+(3*dir[1]*line_len)] == 'S') {

					total_xmas += 1
					// println(i+(0*dir[0])+(0*dir[1]*line_len), i+(1*dir[0])+(1*dir[1]*line_len), i+(2*dir[0])+(2*dir[1]*line_len), i+(3*dir[0])+(3*dir[1]*line_len))
					// println(i, bytes[i])
				}
			}
		}

	}
	return total_xmas
}

func sumX_MAS(bytes []byte) int {
	total_xmas := 0
	// fmt.Printf("bytes: %v\n", bytes)

	line_len := 140
	directions := [][]int{
		{-1, 0},
		{0, -1},
		{-1, -1},
		{1, -1},
	}
	for i, j := 0, 0; i < len(bytes)-1; i, j = i+1, j+1 {
		if bytes[i] == 10 {
			line_len = j + 1
			j = -1
			continue
		}

		for _, dir := range directions {
			// target_max := i - 3*dir[0] - 3*dir[1]*line_len
			x_max := j + (3 * dir[0])
			y_max := i + (3 * dir[1] * line_len)
			if x_max >= 0 && x_max <= line_len && y_max >= 0 {

				if x_max >= line_len {
					// println("x_max", x_max, j)
				}

				// println(i, bytes[i])
				if (bytes[i+(0*dir[0])+(0*dir[1]*line_len)] == 'S' &&
					bytes[i+(1*dir[0])+(1*dir[1]*line_len)] == 'A' &&
					bytes[i+(2*dir[0])+(2*dir[1]*line_len)] == 'M' &&
					bytes[i+(3*dir[0])+(3*dir[1]*line_len)] == 'X') ||
					(bytes[i+(0*dir[0])+(0*dir[1]*line_len)] == 'X' &&
						bytes[i+(1*dir[0])+(1*dir[1]*line_len)] == 'M' &&
						bytes[i+(2*dir[0])+(2*dir[1]*line_len)] == 'A' &&
						bytes[i+(3*dir[0])+(3*dir[1]*line_len)] == 'S') {

					total_xmas += 1
					// println(i+(0*dir[0])+(0*dir[1]*line_len), i+(1*dir[0])+(1*dir[1]*line_len), i+(2*dir[0])+(2*dir[1]*line_len), i+(3*dir[0])+(3*dir[1]*line_len))
					// println(i, bytes[i])
				}
			}
		}

	}
	return total_xmas
}

func main() {
	start := time.Now()

	// file, _ := os.Open("day_3/input.txt")
	file, _ := os.ReadFile("day_4/input.txt")
	// input_string := string(file)
	// fmt.Printf("file: %v\n", file[9])

	// bytes_result_tot_1 := 0
	// bytes_result_tot_1 += sumXMAS(file)

	//part 1
	// fmt.Printf("result_1 : %v\n", sumXMAS(file))

	// part 2
	fmt.Printf("result_2 : %v\n", sumX_MAS(file))

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
