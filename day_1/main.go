package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, err := os.Open("day_1/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanWords)

	left := []int{}
	right := []int{}

	result_1 := 0
	result_2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		num_left, _ := strconv.Atoi(words[0])
		num_right, _ := strconv.Atoi(words[1])

		left = append(left, num_left)
		right = append(right, num_right)

	}

	//part 1
	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		result_1 += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Printf("result 1: %v\n", result_1)

	// part 2
	frequency := make(map[int]int)
	for _, num := range right {
		frequency[num]++
	}
	for _, num := range left {
		result_2 += num * frequency[num]
	}
	fmt.Printf("result 2: %v\n", result_2)

	//logs
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
