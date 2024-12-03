package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func sumMul(bytes []byte) int {
	total := 0
	// bytes := []byte(input)
	// fmt.Printf("bytes: %v\n", bytes)

	disable := false
	for i := 0; i < len(bytes)-8; i++ {
		if !disable &&
			bytes[i] == 'd' &&
			bytes[i+1] == 'o' &&
			bytes[i+2] == 'n' &&
			bytes[i+3] == '\'' &&
			bytes[i+4] == 't' &&
			bytes[i+5] == '(' &&
			bytes[i+6] == ')' {
			disable = true
			continue

		}

		if disable &&
			bytes[i] == 'd' &&
			bytes[i+1] == 'o' &&
			bytes[i+2] == '(' &&
			bytes[i+3] == ')' {

			disable = false
			continue
		}

		// Check for 'mul('
		if !disable &&
			bytes[i] == 'm' &&
			bytes[i+1] == 'u' &&
			bytes[i+2] == 'l' &&
			bytes[i+3] == '(' {

			numStart := i + 4
			if numStart < len(bytes) && (bytes[numStart] < '0' || bytes[numStart] > '9') {
				// fmt.Printf("numStart: %v\n", numStart)
				continue
			}

			// Parse first number
			num1 := 0
			for numStart < len(bytes) && bytes[numStart] >= '0' && bytes[numStart] <= '9' {
				num1 = num1*10 + int(bytes[numStart]-'0')
				numStart++
			}

			// find comma
			if numStart < len(bytes) && bytes[numStart] != ',' {
				continue
			}
			numStart++ // skip comma

			// find second number
			if numStart < len(bytes) && (bytes[numStart] < '0' || bytes[numStart] > '9') {
				continue
			}

			// Parse second number
			num2 := 0
			// num2 := []byte{}
			for numStart < len(bytes) && bytes[numStart] >= '0' && bytes[numStart] <= '9' {
				num2 = num2*10 + int(bytes[numStart]-'0')
				// num2 = append(num2, bytes[numStart])
				numStart++
			}
			// fmt.Printf("num2: %v\n", num2)

			// check for closing parenthesis
			if numStart < len(bytes) && bytes[numStart] == ')' {
				// println(num1, num2)
				total += num1 * num2
			}
		}
	}
	return total
}

func regSumMul(input_string string) int {
	reg := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := reg.FindAllStringSubmatch(input_string, -1)
	result_1 := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		// println(x, y)
		result_1 += x * y
	}
	return result_1
}

func regSumMul2(input_string string) int {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((-?\d+),(-?\d+)\)`)

	mulEnabled := true
	total := 0

	matches := re.FindAllStringSubmatch(input_string, -1)
	// fmt.Printf("matches: %q\n", matches)
	for _, match := range matches {
		if match[0] == "do()" {
			mulEnabled = true
		} else if match[0] == "don't()" {
			mulEnabled = false
		} else if strings.Contains(match[0], "mul") {
			if mulEnabled {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				// println(a, b)
				total += a * b
			}
		}
	}
	return total
}

func main() {
	start := time.Now()

	// file, _ := os.Open("day_3/input.txt")
	file, _ := os.ReadFile("day_3/input.txt")
	// input_string := string(file)

	// reg_result_tot_1 := 0
	bytes_result_tot_1 := 0

	// fmt.Printf("input_string: %v\n", input_string)

	// reg_result_tot_1 += regSumMul2(input_string)

	bytes_result_tot_1 += sumMul(file)

	//part 1

	// fmt.Printf("numbers: %#v\n", final_list)
	// fmt.Printf("result regex: %v\n", reg_result_tot_1)
	fmt.Printf("result bytes: %v\n", bytes_result_tot_1)

	// part 2
	// fmt.Printf("result 2: %v\n", almostSafeReports)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
