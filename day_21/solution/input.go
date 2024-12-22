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

func ReadInput(file_path string) ([]Code, map[rune]Coord, map[rune]Coord) {
	defer Track(time.Now(), "Input Parsed in")

	file, _ := os.Open(file_path)
	defer file.Close()

	codes := []Code{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numeric, _ := strconv.Atoi(line[:3])
		codes = append(codes, Code{Code: line, Numeric: numeric})
	}

	numeric_pad := make(map[rune]Coord)
	directionnal_pad := make(map[rune]Coord)

	numeric_pad['7'] = Coord{0, 0}
	numeric_pad['8'] = Coord{1, 0}
	numeric_pad['9'] = Coord{2, 0}
	numeric_pad['4'] = Coord{0, 1}
	numeric_pad['5'] = Coord{1, 1}
	numeric_pad['6'] = Coord{2, 1}
	numeric_pad['1'] = Coord{0, 2}
	numeric_pad['2'] = Coord{1, 2}
	numeric_pad['3'] = Coord{2, 2}
	numeric_pad['0'] = Coord{1, 3}
	numeric_pad['A'] = Coord{2, 3}

	directionnal_pad['^'] = Coord{1, 0}
	directionnal_pad['A'] = Coord{2, 0}
	directionnal_pad['<'] = Coord{0, 1}
	directionnal_pad['v'] = Coord{1, 1}
	directionnal_pad['>'] = Coord{2, 1}

	return codes, numeric_pad, directionnal_pad
}
