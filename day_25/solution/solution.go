package solution

import (
	"time"
)

func canFit(lock Lock, key Key) bool {
	for i := range lock {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func findMatchingPairs(keys []Key, locks []Lock) [][2]int {
	var matches [][2]int
	count := 0
	for i, lock := range locks {
		for j, key := range keys {
			if canFit(lock, key) {
				matches = append(matches, [2]int{i, j})
				count++
			}
		}
	}
	// println(count)
	return matches
}

func Part_One(keys []Key, locks []Lock) int {
	defer Track(time.Now(), "Part 1")

	matches := findMatchingPairs(keys, locks)

	return len(matches)
}
