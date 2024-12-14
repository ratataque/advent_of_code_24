package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
	// "strings"
)

type ButtonCoordinates struct {
	ButtonA [2]int
	ButtonB [2]int
	Prize   [2]int
}

func extendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD(b, a%b)
	x := y1
	y := x1 + (a/b)*y1
	return gcd, x, y
}

func ParseButtonData(input string) ([]ButtonCoordinates, error) {
	re := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)\nButton B: X\+(\d+), Y\+(\d+)\nPrize: X=(\d+), Y=(\d+)`)

	matches := re.FindAllStringSubmatch(input, -1)

	if matches == nil {
		return nil, fmt.Errorf("no valid data found in input")
	}

	var results []ButtonCoordinates

	for _, match := range matches {
		coords, err := convertMatchesToIntegers(match[1:])
		if err != nil {
			return nil, err
		}

		results = append(results, ButtonCoordinates{
			ButtonA: [2]int{coords[0], coords[1]},
			ButtonB: [2]int{coords[2], coords[3]},
			Prize:   [2]int{coords[4], coords[5]},
		})
	}

	return results, nil
}

func convertMatchesToIntegers(matches []string) ([]int, error) {
	integers := make([]int, len(matches))

	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to integer: %v", match, err)
		}
		integers[i] = num
	}

	return integers, nil
}

func findAllSolutions(a, b, c int) []struct{ x, y int } {
	gcd, x0, y0 := extendedGCD(a, b)

	// Check if a solution exists
	if c%gcd != 0 {
		return nil
	}

	// Scale the base solution
	// x0 *= c / gcd
	// y0 *= c / gcd

	// Generate all solutions within reasonable bounds
	var solutions []struct{ x, y int }

	for k := -100; k <= 100; k++ {
		x := x0 + k*(b/gcd)
		y := y0 + k*(a/gcd)
		// println(x, y, k)

		if x >= 0 && y >= 0 && x <= 100 && y <= 100 {
			solutions = append(solutions, struct{ x, y int }{x, y})
		}
	}

	return solutions
}
func main() {
	start := time.Now()
	file, _ := os.ReadFile("day_13/input.txt")
	input := string(file)

	results, _ := ParseButtonData(input)

	fmt.Printf("results: %v\n", results)

	gcd, x, y := extendedGCD(94, 22)
	println(gcd*4200, x, y)
	// println(x+20*22/2, y-2*94/2, 81*94+36*22)
	println(4+22/2, -17-94/2, 4*94*4200-17*22*4200)
	soluce := findAllSolutions(94, 22, 8400)

	fmt.Printf("soluce: %v\n", soluce)
	// var solutions []struct{ x, y int }
	//
	// for k := -100; k <= 100; k++ {
	// 	x := x0 + k*(y/gcd)
	// 	y := y0 - k*(x/gcd)
	//
	// 	if x >= 0 && y >= 0 && x <= 100 && y <= 100 {
	// 		solutions = append(solutions, struct{ x, y int }{x, y})
	// 	}
	// }

	println(94*4*4200, 22*-17*4200, 94*4*4200+22*-17*4200)

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
