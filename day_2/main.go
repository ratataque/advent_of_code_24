package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func reverse(arr []int) []int {
	reversed := make([]int, len(arr))
	for i, j := 0, len(arr)-1; j >= 0; i, j = i+1, j-1 {
		reversed[i] = arr[j]
	}
	return reversed
}

func parseReport(line string) []int {
	// parts := strings.Fields(line)
	parts := strings.Split(line, " ")
	report := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			report = append(report, num)
		}
	}

	return report
}

func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	// initial direction
	isIncreasing := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff == 0 || abs(diff) > 3 {
			return false
		}

		// check if direction is consistent
		if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) {
			return false
		}
	}

	return true
}

func isAlmostSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	// initial direction
	isIncreasing := report[1] > report[0]

	var skip bool
	var skipped bool
	var diff int
	for i := 1; i < len(report); i++ {
		if skip {
			diff = report[i] - report[i-2]
			skip = false
			skipped = true
		} else {
			diff = report[i] - report[i-1]
		}

		if diff == 0 || abs(diff) > 3 {
			if skipped {
				return false
			} else {
				skip = true
				continue
			}
		}

		// check if direction is consistent
		if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) {
			if skipped {
				return false
			} else {
				skip = true
				continue
			}
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	start := time.Now()

	file, _ := os.Open("day_2/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0
	almostSafeReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		report := parseReport(line)

		// part 1
		if isSafe(report) {
			// fmt.Printf("report: %v\n", report)
			safeReports++
		}

		//part 2
		if isAlmostSafe(report) || isAlmostSafe(reverse(report)) {
			// fmt.Printf("report: %v\n", report)
			almostSafeReports++
		}
	}

	//part 1

	// fmt.Printf("numbers: %#v\n", final_list)
	fmt.Printf("result 1: %v\n", safeReports)

	// part 2
	fmt.Printf("result 2: %v\n", almostSafeReports)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
