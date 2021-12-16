package main

import (
	"aoc21/aocinput"
	"fmt"
	"strconv"
	"strings"
)

func getPositionCounts(vals []string) [][2]int {

	var positionCounts [][2]int
	for _, s := range vals {
		chars := strings.Split(s, "")

		for i, c := range chars {
			if i >= len(positionCounts) {
				var a [2]int
				positionCounts = append(positionCounts, a)
			}

			n, _ := strconv.Atoi(c)
			positionCounts[i][n] += 1
		}
	}

	return positionCounts
}

func trimArray(vals []string, mosts bool) string {

	for idx := 0; len(vals) > 1; idx++ {
		var newVals []string
		positionCounts := getPositionCounts(vals)

		for _, s := range vals {
			c, _ := strconv.Atoi(s[idx : idx+1])

			if mosts {
				if c == 0 && positionCounts[idx][0] > positionCounts[idx][1] {
					// keep
					newVals = append(newVals, s)
				} else if c == 1 && positionCounts[idx][1] >= positionCounts[idx][0] {
					// keep
					newVals = append(newVals, s)
				}
			} else {
				if c == 0 && positionCounts[idx][0] <= positionCounts[idx][1] {
					newVals = append(newVals, s)
				} else if c == 1 && positionCounts[idx][1] < positionCounts[idx][0] {
					newVals = append(newVals, s)
				}
			}
		}

		vals = newVals
	}

	// When we get here, we should have only one string left, which is the one we want.
	return vals[0]
}

func main() {

	input := "input3.txt"

	vals := aocinput.LoadStringArray(input)

	positionCounts := getPositionCounts(vals)

	// part 1
	gamma := 0
	epsilon := 0

	for _, p := range positionCounts {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if p[0] > p[1] {
			epsilon = epsilon | 1
		} else {
			gamma = gamma | 1
		}
	}
	fmt.Printf("Gamma: %d (%b), Epsilon: %d (%b), Product: %d\n", gamma, gamma, epsilon, epsilon, gamma*epsilon)

	// part2

	oxygen := trimArray(vals, true)
	oxint, _ := strconv.ParseInt(oxygen, 2, 0)

	fmt.Printf("Oxygen: %d (%s)\n", oxint, oxygen)

	co2 := trimArray(vals, false)
	co2int, _ := strconv.ParseInt(co2, 2, 0)

	fmt.Printf("CO2: %d (%s)\n", co2int, co2)

	fmt.Printf("Final result: %d\n", co2int*oxint)
	// part1
	// for _, e := range vals {
	// 	switch e.S {
	// 	case "down":
	// 		depth += e.V
	// 	case "up":
	// 		depth -= e.V
	// 	case "forward":
	// 		position += e.V
	// 	}
	// }

}
