package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func sumMul(input string) int {
	total := 0
	bytes := []byte(input)
	// fmt.Printf("bytes: %v\n", bytes)

	for i := 1; i < len(bytes)-8; i++ {
		// Check for 'mul('
		if bytes[i] == 'm' &&
			bytes[i+1] == 'u' &&
			bytes[i+2] == 'l' &&
			bytes[i+3] == '(' {

			// Find first number
			// fmt.Printf("bytes[i+4]: %v\n", []byte("9"))

			// fmt.Printf("bytes[numStart]: %q\n", bytes[numStart])
			// fmt.Printf("bytes[numStart]: %v\n", bytes[numStart] < '3')
			// fmt.Printf("bytes[numStart]: %v\n", (bytes[numStart] < '0' || bytes[numStart] > '9'))
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
			// fmt.Printf("num2: %v\n", num1)

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
				fmt.Printf("num1: %v\n", num1)
				fmt.Printf("num2: %v\n", num2)
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
		result_1 += x * y
	}
	return result_1
}

func main() {
	start := time.Now()

	file, _ := os.Open("day_3/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result_1 := 0
	for scanner.Scan() {
		input_string := scanner.Text()

		result_1 = regSumMul(input_string)

		// result_1 = sumMul(input_string)
		// fmt.Printf("matches: %q\n", matches)
	}

	//part 1

	// fmt.Printf("numbers: %#v\n", final_list)
	fmt.Printf("result 1: %v\n", result_1)

	// part 2
	// fmt.Printf("result 2: %v\n", almostSafeReports)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
