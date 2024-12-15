package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
	// "strings"
)

type ButtonCoordinates struct {
	ButtonA [2]int64
	ButtonB [2]int64
	Prize   [2]int64
}

func extendedGCD(a, b int64) (int64, int64, int64) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := extendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
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
			ButtonA: [2]int64{coords[0], coords[1]},
			ButtonB: [2]int64{coords[2], coords[3]},
			Prize:   [2]int64{coords[4], coords[5]},
		})
	}

	return results, nil
}

func ParseButtonDataP2(input string) ([]ButtonCoordinates, error) {
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
			ButtonA: [2]int64{coords[0], coords[1]},
			ButtonB: [2]int64{coords[2], coords[3]},
			Prize:   [2]int64{10_000_000_000_000 + coords[4], 10_000_000_000_000 + coords[5]},
		})
	}

	return results, nil
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func convertMatchesToIntegers(matches []string) ([]int64, error) {
	integers := make([]int64, len(matches))

	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to integer: %v", match, err)
		}
		integers[i] = int64(num)
	}

	return integers, nil
}

func findTheSolutions(a, b, c int64) *[2]float64 {
	gcd, x0, y0 := extendedGCD(a, b)

	coeff := c / gcd

	// Check if a solution exists
	if c%gcd != 0 {
		return nil
	}

	// Scale the base solution
	x0 *= coeff
	y0 *= coeff

	k := int64(y0 * gcd / a)

	max := int64(0)
	count := int64(0)
	if k < 0 {
		count = -1
		max = k - 100
	} else {
		count = 1
		max = k + 100
	}

	for abs(k) <= abs(max) {
		x := x0 + k*(b/gcd)
		y := y0 - k*(a/gcd)

		if x >= 0 && y >= 0 && x <= 100 && y <= 100 {
			return &[2]float64{float64(x), float64(y)}
		}
		k += count
	}
	return nil
}

func isWhole(x float64) bool {
	diff := math.Abs(math.Trunc(x) - x)
	return diff < 1e-5
}

func systemSolver(ax, bx, ap, ay, by, bp float64) *[2]float64 {
	// determiant := (by - (bx*ay)/ax)
	determiant := (bx*ay - by*ax)
	if determiant == 0 {
		test := findTheSolutions(int64(ax), int64(bx), int64(ap))
		// println(int(ax), int(bx), int(ap))
		if test != nil {
			return test
		} else {
			return findTheSolutions(int64(bx), int64(by), int64(bp))
		}
	}
	// x := (bp - (ap * ay / ax)) / determiant
	x := (bx*bp - by*ap) / (bx*ay - by*ax)
	// println(int(bx), int(bp), int(by), int(ap), int(bx), int(ay), int(by), int(ax), int(x), int((bx*ay - by*ax)), int((bx*bp - by*ap)))
	if isWhole(x) {
		// y := (ap - bx*y) / ax
		y := (ap - ax*x) / bx
		int_res := [2]float64{x, y}
		return &int_res
	} else {
		return nil
	}
}

func main() {
	start := time.Now()
	file, _ := os.ReadFile("day_13/input.txt")
	input := string(file)

	results, _ := ParseButtonData(input)
	results_2, _ := ParseButtonDataP2(input)

	tokens := float64(0)
	for _, button := range results {
		res := systemSolver(float64(button.ButtonA[0]), float64(button.ButtonB[0]), float64(button.Prize[0]), float64(button.ButtonA[1]), float64(button.ButtonB[1]), float64(button.Prize[1]))

		if res != nil {
			tokens += (*res)[0]*3 + (*res)[1]
		} else {
		}
	}

	tokens_2 := float64(0)
	for _, button := range results_2 {
		res := systemSolver(float64(button.ButtonA[0]), float64(button.ButtonB[0]), float64(button.Prize[0]), float64(button.ButtonA[1]), float64(button.ButtonB[1]), float64(button.Prize[1]))

		if res != nil {
			// fmt.Printf("res: %v\n", res)
			tokens_2 += (*res)[0]*3 + (*res)[1]
		} else {
			// println("No solution found")
		}
	}

	// Part 1
	fmt.Printf("Part 1: %v\n", int(tokens))

	// Part 1
	fmt.Printf("Part 2: %v\n", int(tokens_2))

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
