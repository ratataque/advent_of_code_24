package solution

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type Wire map[string]int

type Operation struct {
	wire_1      string
	op          string
	wire_2      string
	wire_result string
}

func ReadInput(file_path string) (Wire, []Operation) {
	defer Track(time.Now(), "Input Parsed in")

	file, _ := os.Open(file_path)
	defer file.Close()

	wire := make(Wire)

	operations := make([]Operation, 0)

	scanner := bufio.NewScanner(file)
	is_operation := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			is_operation = true
			continue
		}

		if is_operation {
			l := strings.Split(line, " ")
			op := Operation{l[0], l[1], l[2], l[4]}
			operations = append(operations, op)
		} else {
			l := strings.Split(line, ": ")
			int_val, _ := strconv.Atoi(l[1])
			wire[l[0]] = int_val
		}
	}
	return wire, operations
}
