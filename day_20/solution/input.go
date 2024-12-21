package solution

import (
	"bufio"
	"os"
	"time"
)

type Coord struct {
	X int
	Y int
}

type Grid struct {
	Data          [][]byte
	Width, Height int
	Start, End    Coord
}

func ReadInput(file_path string) *Grid {
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
				end_pos = Coord{x, y}
			}
		}
		y++
	}

	g := &Grid{
		Data:   grid,
		Width:  len(grid[0]),
		Height: len(grid),
		Start:  start_pos,
		End:    end_pos}

	return g
}
