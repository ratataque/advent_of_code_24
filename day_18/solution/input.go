package solution

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Coord struct {
	X int
	Y int
}

// type coord = Coord

func ReadInput(fp string) ([]Coord, error) {
	defer Track(time.Now(), "Input Parsed in")
	// log.Println("Reading", fp)
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []Coord
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		coord := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(coord[0])
		y, _ := strconv.Atoi(coord[1])
		lines = append(lines, Coord{x, y})
	}

	return lines, nil
}
