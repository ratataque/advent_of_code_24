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

	// Part 2
	for i := len(files) - 1; i > 0; i-- {
		count := 0
		last_gap_len := 0
		for j, c := range disk_map2 {
			if c == i {
				break
			}

			if c == -1 {
				count++
				if disk_map2[j+1] != -1 {
					last_gap_len = count
					j++
					count = 0
				}
			}

			if len(files[i]) <= last_gap_len {
				v := last_gap_len - len(files[i])
				new_file := []int{}

				new_file = append(new_file, files[i]...)
				for k := 0; k < v; k++ {
					new_file = append(new_file, -1)
				}

				for k := 0; k < len(disk_map2); k++ {
					if disk_map2[k] == i {
						disk_map2[k] = -1
					}
				}

				disk_map2 = append(disk_map2[:j-last_gap_len], append(new_file, disk_map2[j:]...)...)
				break
			}

		}
	}

	result_1 := int64(0)
	for i, n := range disk_map {
		if n != -1 {
			result_1 += int64(n) * int64(i)
		}
	}

	result_2 := int64(0)
	for i, n := range disk_map2 {
		if n != -1 {
			result_2 += int64(n) * int64(i)
		}
	}

	// Part 1
	fmt.Printf("Part 1: %v\n", result_1)

	// Part 2
	fmt.Printf("Part 2: %v\n", result_2)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
