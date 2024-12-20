package solution

import (
	"bufio"
	"os"
	"time"
)

type Coord struct {
	x int
	y int
}

func ReadInput(file_path string) ([][]byte, Coord, Coord) {
	defer Track(time.Now(), "Input Parsed in")

	file, _ := os.Open(file_path)
	defer file.Close()

	var grid [][]byte
	var start_pos, end_pos Coord

	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := []byte(line)
		grid = append(grid, row)

		for x, char := range row {
			if char == 'S' {
				start_pos = Coord{x, y}
			} else if char == 'E' {
				end_pos = Coord{7, 8}
			}
		}
		y++
	}

	return grid, start_pos, end_pos
}
