package main

import (
	// "bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	file, _ := os.ReadFile("day_5/input.txt")
	input_string := string(file)

	// file, _ := os.Open("day_5/input.txt")
	// defer file.Close()
	//
	// scanner := bufio.NewScanner(file)
	//
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	println(line)
	// }

	rules_input := strings.Split(input_string, "\n\n")[0]
	updates := strings.Split(input_string, "\n\n")[1]

	rules := make(map[string]bool)
	for _, v := range strings.Split(rules_input, "\n") {
		// fmt.Printf("%v\n", "")
		rules[v] = true
	}

	count_1 := 0
	count_2 := 0
	for _, v := range strings.Split(updates[:len(updates)-1], "\n") {
		// fmt.Printf("%q\n", v)

		safe := true
		line := strings.Split(v, ",")
		for i := 1; i < len(line); i++ {
			// println(i)
			for j := 0; j <= i; j++ {
				// println(j)
				if _, test := rules[line[i]+"|"+line[j]]; test {
					safe = false
					// println(i, j)
					// println(line[i] + "|" + line[j])
					// fmt.Printf("line: %v\n", line[j:])
					n := line[i]
					line = append(line[:i], line[i+1:]...)
					line = append(line[:j], append([]string{n}, line[j:]...)...)
					// fmt.Printf("line: %v\n", line)

					// println(line[(len(line)+1)/2-1])

					i = 0
				}
			}
			// fmt.Printf("%q\n", page)
		}
		if safe {
			n1, _ := strconv.Atoi(line[(len(line)+1)/2-1])
			// println(n)
			count_1 += n1
		} else {
			n2, _ := strconv.Atoi(line[(len(line)+1)/2-1])
			count_2 += n2
		}
	}

	fmt.Printf("part 1: %v\n", count_1)
	fmt.Printf("part 1: %v\n", count_2)
	// fmt.Printf("rules: %v\n", rules)

	//part 1
	// fmt.Printf("result_1 : %v\n", sumXMAS(file))

	// part 2

	elapsed := time.Since(start)
	fmt.Printf("\n\nExecution time: %s\n", elapsed)
}
