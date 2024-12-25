package solution

import (
	"bufio"
	"os"
	"slices"
)

type Key []int
type Lock []int

func ParseInput(name string) ([]Key, []Lock) {
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var locks []Lock
	var keys []Key

	for scanner.Scan() {
		curr := make([]int, 5)
		isLock := false
		for i := 0; i < 7; i++ {
			for c, char := range scanner.Text() {
				if char == '#' {
					curr[c]++
				}
			}
			if i == 0 && slices.Equal(curr, []int{1, 1, 1, 1, 1}) {
				isLock = true
			}
			scanner.Scan()
		}

		for i := range curr {
			curr[i]--
		}
		if isLock {
			locks = append(locks, curr)
		} else {
			keys = append(keys, curr)
		}
	}
	return keys, locks
}
