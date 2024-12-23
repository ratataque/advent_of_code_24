package solution

import (
	"bufio"
	"os"
	"strconv"
	"time"
)

type Coord struct {
	X int
	Y int
}

func ReadInput(file_path string) []int {
	defer Track(time.Now(), "Input Parsed in")

	file, _ := os.Open(file_path)
	defer file.Close()

	initial_prices := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_int, _ := strconv.Atoi(line)
		initial_prices = append(initial_prices, line_int)
	}

	return initial_prices
}
