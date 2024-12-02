package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	// Determine initial direction
	isIncreasing := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Check if diff is within acceptable range
		if diff == 0 || abs(diff) > 3 {
			return false
		}

		// Check if direction is consistent
		if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) {
			return false
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

func parseReport(line string) []int {
	parts := strings.Fields(line)
	report := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			report = append(report, num)
		}
	}

	return report
}

func main() {
	start := time.Now()

	file, _ := os.Open("day_2/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		report := parseReport(line)

		if isSafe(report) {
			safeReports++
		}
	}

	//part 1

	// fmt.Printf("numbers: %#v\n", final_list)
	fmt.Printf("result 1: %v\n", safeReports)

	// part 2
	// fmt.Printf("result 2: %v\n", result_2)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
