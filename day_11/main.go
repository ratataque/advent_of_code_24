package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.ReadFile("day_11/input.txt")

	input := strings.Fields(string(file))

	existing := map[string]int{}

	for _, line := range input {
		existing[line] = 1
	}

	total_1 := 0
	for blink := 0; blink < 75; blink++ {
		if blink == 25 {
			for _, n := range existing {
				// println(stone)
				total_1 += n
			}
		}
		new_existing := map[string]int{}
		for caillou, n := range existing {
			curr_num, _ := strconv.Atoi(caillou)

			if curr_num == 0 {
				new_existing["1"] += n

			} else if len(caillou)%2 == 0 {
				before := strings.TrimLeft(caillou[:len(caillou)/2], "0")
				after := strings.TrimLeft(caillou[len(caillou)/2:], "0")

				if before == "" {
					new_existing["0"] += n
				} else {
					new_existing[before] += n
				}

				if after == "" {
					new_existing["0"] += n
				} else {
					new_existing[after] += n
				}

			} else {
				new_existing[strconv.Itoa(curr_num*2024)] += n
			}
		}
		existing = new_existing
	}

	total_2 := 0
	for _, n := range existing {
		total_2 += n
	}
	// Part 1
	fmt.Printf("Part 1: %v\n", total_1)

	// Part 2
	fmt.Printf("Part 2: %v\n", total_2)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
