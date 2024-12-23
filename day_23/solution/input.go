package solution

import (
	"bufio"
	"os"
	"strings"
	"time"
)

type Coord struct {
	X int
	Y int
}

func ReadInput(file_path string) [][]string {
	defer Track(time.Now(), "Input Parsed in")

	file, _ := os.Open(file_path)
	defer file.Close()

	computer_list := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		computer_couple := strings.Split(line, "-")
		computer_list = append(computer_list, computer_couple)
	}

	return computer_list
}
