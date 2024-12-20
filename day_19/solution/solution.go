package solution

import (
	"strings"
	"time"
)

// Helper struct to store memo cache
type Cache map[string]bool

// countWays counts all possible ways to make the design
func countWays(design string, patterns []string, patternStarts map[byte][]string, memo map[string]int) int {
	// Base case: empty string has one way to make it
	if len(design) == 0 {
		return 1
	}

	// Check memoization
	if count, exists := memo[design]; exists {
		return count
	}

	totalWays := 0
	firstChar := design[0]

	// Get patterns that start with the first character
	potentialPatterns, exists := patternStarts[firstChar]
	if !exists {
		memo[design] = 0
		return 0
	}

	// Try each viable pattern
	for _, pattern := range potentialPatterns {
		if strings.HasPrefix(design, pattern) {
			// Add the number of ways to make the remaining design
			remaining := design[len(pattern):]
			totalWays += countWays(remaining, patterns, patternStarts, memo)
		}
	}

	memo[design] = totalWays
	return totalWays
}

// canMakeDesign checks if a design can be made using the available patterns
func canMakeDesign(test_word string, words []string, patternStarts map[byte][]string, memo Cache) bool {
	// Base case: empty design is always possible
	if len(test_word) == 0 {
		return true
	}

	// Check memo cache
	if result, exists := memo[test_word]; exists {
		return result
	}

	// Get patterns that start with the first character of the design
	firstChar := test_word[0]
	potentialPatterns, exists := patternStarts[firstChar]
	if !exists {
		memo[test_word] = false
		return false
	}

	// Try each viable pattern
	for _, pattern := range potentialPatterns {
		if strings.HasPrefix(test_word, pattern) {
			// Recursively try to match the rest of the design
			remaining := test_word[len(pattern):]
			if canMakeDesign(remaining, words, patternStarts, memo) {
				memo[test_word] = true
				return true
			}
		}
	}

	memo[test_word] = false
	return false
}

func PartOne(words []string, test_words []string) int {
	defer Track(time.Now(), "Part 1")

	// Process designs
	possibleCount := 0
	memo := make(Cache)

	patternStarts := make(map[byte][]string)
	for _, pattern := range words {
		if len(pattern) > 0 {
			firstChar := pattern[0]
			patternStarts[firstChar] = append(patternStarts[firstChar], pattern)
		}
	}

	for _, test_word := range test_words {
		test := canMakeDesign(test_word, words, patternStarts, memo)
		if test {
			possibleCount++
		}
	}

	return possibleCount
}

func PartTwo(words []string, test_words []string) int {
	defer Track(time.Now(), "Part 1")

	// Process designs
	possibleCount := 0
	memo := make(map[string]int)

	patternStarts := make(map[byte][]string)
	for _, pattern := range words {
		if len(pattern) > 0 {
			firstChar := pattern[0]
			patternStarts[firstChar] = append(patternStarts[firstChar], pattern)
		}
	}

	for _, test_word := range test_words {
		possibleCount += countWays(test_word, words, patternStarts, memo)
	}

	return possibleCount
}
