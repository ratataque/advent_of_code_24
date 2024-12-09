package main

import (
	"fmt"
	"os"
	"time"
)

func parseDiskMap(input []byte) ([]int, [][]int) {
	var string []int
	fileID := 0
	files := [][]int{}
	for i, char := range input {
		n := int(char - '0')

		if i%2 == 1 {
			sub := []int{}
			for j := 0; j < n; j++ {
				sub = append(sub, -1)
			}
			string = append(string, sub...)
			fileID++

		} else {
			sub := []int{}
			for j := 0; j < n; j++ {
				sub = append(sub, fileID)
			}
			files = append(files, sub)
			string = append(string, sub...)
		}
	}
	return string, files
}

func parseDiskMap2(input []byte) []int {
	var string []int
	fileID := 0
	for i, char := range input {
		n := int(char - '0')

		if i%2 == 1 {
			sub := []int{}
			for j := 0; j < n; j++ {
				sub = append(sub, -1)
			}
			string = append(string, sub...)
			fileID++

		} else {
			sub := []int{}
			for j := 0; j < n; j++ {
				sub = append(sub, fileID)
			}
			string = append(string, sub...)
		}
	}
	return string
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("day_9/input.txt")

	input := file[:len(file)-1]
	disk_map, files := parseDiskMap(input)
	disk_map2 := make([]int, len(disk_map))
	copy(disk_map2, disk_map)

	j := len(disk_map) - 1
	for i := 0; i < len(disk_map)-1; i++ {
		if disk_map[j] == -1 {
			disk_map = disk_map[:len(disk_map)-1]
			j--
		}

		if disk_map[i] == -1 && disk_map[j] != -1 {
			disk_map[i] = disk_map[j]
			disk_map = disk_map[:len(disk_map)-1]
			j--
		}

		if disk_map[i] == -1 {
			i--
		}
	}

	result_1 := 0
	for i, n := range disk_map {
		if n != -1 {
			result_1 += n * i
		}
	}

	// Part 1
	fmt.Printf("Part 1: %v\n", result_1)
	fmt.Printf("Part 1: %v\n", disk_map)
	fmt.Printf("Part 1: %v\n", disk_map2)
	fmt.Printf("Part 1: %v\n", files)

	// Part 2
	// fmt.Printf("Part 2: %v\n", len(antinodes2))

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
