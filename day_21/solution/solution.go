package solution

import (
	"fmt"
	"strings"
	"time"
)

type Code struct {
	Code       string
	Numeric    int
	complexity int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getDirectionString(count int, posChar, negChar string) string {
	if count > 0 {
		return strings.Repeat(posChar, count)
	}
	return strings.Repeat(negChar, -count)
}

func manhattan_distance_into_arrows(start, end Coord, direction bool) string {

	// Pre-calculate the capacity to avoid reallocations
	capacity := abs(end.X-start.X) + abs(end.Y-start.Y)
	// Create a builder with the exact capacity needed
	var b strings.Builder
	b.Grow(capacity)

	repeat_x := end.X - start.X
	repeat_y := end.Y - start.Y

	// Handle special cases first
	if (end.X == 0 && start.Y == 3) || (start.Y == 0 && end.X == 0 && direction) {
		b.WriteString(getDirectionString(repeat_y, "v", "^"))
		b.WriteString(getDirectionString(repeat_x, ">", "<"))
	} else {
		// Handle movement based on priority
		switch {
		case repeat_x < 0:
			b.WriteString(strings.Repeat("<", -repeat_x))
			b.WriteString(getDirectionString(repeat_y, "v", "^"))

		case repeat_y > 0:
			b.WriteString(strings.Repeat("v", repeat_y))
			b.WriteString(getDirectionString(repeat_x, ">", "<"))

		default:
			b.WriteString(strings.Repeat("^", -repeat_y))
			b.WriteString(strings.Repeat(">", repeat_x))
		}
	}

	b.WriteString("A")

	return b.String()
}

// Create a key for caching paths
func makePathKey(start, end Coord, padType string) string {
	return fmt.Sprintf("%d,%d-%d,%d-%s", start.X, start.Y, end.X, end.Y, padType)
}

// Precompute all possible paths for a given pad
func precomputePaths(pad map[rune]Coord, padType string) map[string]string {
	paths := make(map[string]string)
	// chars := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A'}

	for from := range pad {
		if startPos, exists := pad[from]; exists {
			for to := range pad {
				if endPos, exists := pad[to]; exists {
					key := makePathKey(startPos, endPos, padType)
					paths[key] = manhattan_distance_into_arrows(startPos, endPos, padType == "direction")
				}
			}
		}
	}
	return paths
}

func precomputePaths_13(pad map[rune]Coord, padType string, precomputedPaths map[string]string) map[string]string {
	defer Track(time.Now(), "Precompute Paths")
	paths := make(map[string]string)
	// chars := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A'}

	for from := range pad {
		if startPos, exists := pad[from]; exists {
			for to := range pad {
				if endPos, exists := pad[to]; exists {
					key := makePathKey(startPos, endPos, padType)
					char := manhattan_distance_into_arrows(startPos, endPos, padType == "direction")
					for i := 0; i < 12; i++ {
						char = solve_char(char, pad, padType, precomputedPaths)
					}
					// fmt.Printf("char: %v\n", char)
					paths[key] = char
				}
			}
		}
	}
	return paths
}

func solve_char(char string, pad map[rune]Coord, padType string, precomputedPaths map[string]string) string {
	var steps strings.Builder
	steps.Grow(len(char) * 7) // Approximate capacity

	startPos := pad['A']
	for _, char := range char {
		nextPos := pad[char]
		key := makePathKey(startPos, nextPos, padType)
		steps.WriteString(precomputedPaths[key])
		startPos = nextPos
	}

	char = steps.String()

	return char
}

func (code *Code) solve(pad map[rune]Coord, padType string, precomputedPaths map[string]string) {
	var steps strings.Builder
	steps.Grow(len(code.Code) * 7) // Approximate capacity

	startPos := pad['A']
	for _, char := range code.Code {
		nextPos := pad[char]
		key := makePathKey(startPos, nextPos, padType)
		steps.WriteString(precomputedPaths[key])
		startPos = nextPos
	}

	code.Code = steps.String()
	code.complexity = len(code.Code) * code.Numeric
}

func (code *Code) solve_len(pad map[rune]Coord, padType string, precomputedPaths13 map[string]string) int {

	length := 0

	startPos := pad['A']
	for _, char := range code.Code {
		nextPos := pad[char]
		key := makePathKey(startPos, nextPos, padType)
		// fmt.Printf("precomputedPaths: %v\n", precomputedPaths13[key])
		length += len(precomputedPaths13[key])
		startPos = nextPos
	}

	fmt.Printf("length: %v\n", length)
	return length
}

func Part_One(input []Code, numeric_pad map[rune]Coord, direction_pad map[rune]Coord) int {
	defer Track(time.Now(), "Part 1")

	numericPaths := precomputePaths(numeric_pad, "numeric")
	directionPaths := precomputePaths(direction_pad, "direction")

	// fmt.Printf("numericPaths: %v\n", directionPaths)

	total := 0
	for _, code := range input {
		code.solve(numeric_pad, "numeric", numericPaths)
		for i := 0; i < 2; i++ {
			code.solve(direction_pad, "direction", directionPaths)
		}
		// fmt.Printf("code: %v\n", code.complexity)

		total += code.complexity
	}

	return total
}

func Part_Two(input []Code, numeric_pad map[rune]Coord, direction_pad map[rune]Coord) int {
	defer Track(time.Now(), "Part 2")

	// Precompute all possible paths
	numericPaths := precomputePaths(numeric_pad, "numeric")
	directionPaths := precomputePaths(direction_pad, "direction")

	// fmt.Printf("directionPaths: %v\n", directionPaths)

	directionPaths_13 := precomputePaths_13(direction_pad, "direction", directionPaths)

	// fmt.Printf("directionPaths: %v\n", directionPaths_13["2,0-0,1-direction"])
	// fmt.Printf("directionPaths: %v\n", directionPaths_13)

	total := 0
	length := 0
	for _, code := range input {
		code.solve(numeric_pad, "numeric", numericPaths)
		// for i := 0; i < 13; i++ {
		code.solve(direction_pad, "direction", directionPaths_13)

		// fmt.Printf("code: %v\n", code.Code)

		length += code.solve_len(direction_pad, "direction", directionPaths_13) * code.Numeric
		// fmt.Printf("length: %v\n", length)
		// }

		total += code.complexity
	}

	return length / 2
}
