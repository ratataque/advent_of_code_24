package solution

import (
	"os"
	"strings"
	"time"
)

func ReadInput(fp string) ([]string, []string) {
	defer Track(time.Now(), "Input Parsed in")

	file, _ := os.ReadFile(fp)
	input := strings.Split(string(file), "\n\n")

	words := strings.Split(input[0], ", ")
	test_words := strings.Fields(input[1])

	return words, test_words
}
