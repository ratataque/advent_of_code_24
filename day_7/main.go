package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func generatePossibilities(n int) [][]string {
	var result [][]string
	if n == 2 {
		return [][]string{{"+"}, {"*"}}
	}
	for _, prev := range generatePossibilities(n - 1) {
		result = append(result, append([]string(nil), append(prev, "+")...))
		result = append(result, append([]string(nil), append(prev, "*")...))
	}
	return result
}

func generatePossibilities2(n int) [][]string {
	var result [][]string
	if n == 2 {
		return [][]string{{"+"}, {"*"}, {"||"}}
	}
	for _, prev := range generatePossibilities2(n - 1) {
		result = append(result, append([]string(nil), append(prev, "+")...))
		result = append(result, append([]string(nil), append(prev, "*")...))
		result = append(result, append([]string(nil), append(prev, "||")...))
	}
	return result
}

func checkAllPossibilities(nums []int, cache map[int][][]string, target int64) int64 {
	var resulto int64
	for _, list_op := range cache[len(nums)] {
		resulto = int64(nums[0])
		for i, operator := range list_op {
			if operator == "+" {
				resulto += int64(nums[i+1])
			} else if operator == "*" {
				resulto *= int64(nums[i+1])
			} else if operator == "||" {
				num, _ := strconv.Atoi(strconv.Itoa(int(resulto)) + strconv.Itoa(nums[i+1]))
				resulto = int64(num)
			}
			if resulto > target {
				break
			}
		}
		if resulto == target {
			return target
		}
	}
	return 0
}

func checkAllPossibilities2(nums []int, cache map[int][][]string, target int64) int64 {
	var resulto int64
	for _, list_op := range cache[len(nums)] {
		resulto = int64(nums[0])
		for i, operator := range list_op {
			if operator == "+" {
				resulto += int64(nums[i+1])
			} else if operator == "*" {
				resulto *= int64(nums[i+1])
			} else if operator == "||" {
				num, _ := strconv.Atoi(strconv.Itoa(int(resulto)) + strconv.Itoa(nums[i+1]))
				resulto = int64(num)
			}
			if resulto > target {
				break
			}
		}
		if resulto == target {
			return target
		}
	}
	return 0
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("day_7/input.txt")

	lines := [][]interface{}{}

	for _, line := range strings.Split(string(file), "\n") {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		first_number, _ := strconv.Atoi(parts[0])

		numStrings := strings.Fields(parts[1])
		values := []int{}

		for _, numStr := range numStrings {
			num, _ := strconv.Atoi(numStr)
			values = append(values, num)
		}

		lines = append(lines, []interface{}{int64(first_number), values})
	}

	cache := map[int][][]string{}
	cache2 := map[int][][]string{}
	count := int64(0)
	count2 := int64(0)
	for _, line := range lines {
		nums := line[1].([]int)
		_, cached := cache[len(nums)]
		_, cached2 := cache2[len(nums)]

		if !cached {
			cache[len(nums)] = generatePossibilities(len(nums))
		}
		if !cached2 {
			cache2[len(nums)] = generatePossibilities2(len(nums))
		}

		count += int64(checkAllPossibilities(nums, cache, line[0].(int64)))
		count2 += int64(checkAllPossibilities(nums, cache2, line[0].(int64)))
	}

	// Part 1
	fmt.Printf("Part 1: %v\n", count)

	// Part 2
	fmt.Printf("Part 1: %v\n", count2)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
