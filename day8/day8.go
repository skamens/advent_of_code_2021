package main

import (
	"aoc21/aocutil"
	"fmt"
	"strings"
)

func sortEntries(vals *[][]string) {

	for i := 0; i < len(*vals); i++ {
		signals := strings.Fields((*vals)[i][0])
		aocutil.SortArrayValues(&signals)
		(*vals)[i][0] = strings.Join(signals, " ")

		numbers := strings.Fields((*vals)[i][1])
		aocutil.SortArrayValues(&numbers)
		(*vals)[i][1] = strings.Join(numbers, " ")

	}
}

func crunchSignals(signals []string) map[string]int {

	results := make(map[string]int)
	var numToString [10]string

	for _, s := range signals {
		switch len(s) {
		case 2:
			results[s] = 1
			numToString[1] = s
		case 3:
			results[s] = 7
			numToString[7] = s
		case 4:
			results[s] = 4
			numToString[4] = s
		case 7:
			results[s] = 8
			numToString[8] = s
		}
	}

	var topright string

	// Find a six letter one that doesn't have both entries of 1 - that is 6
	for _, s := range signals {
		if len(s) != 6 {
			continue
		}
		onechars := strings.Split(numToString[1], "")

		if !strings.Contains(s, onechars[0]) {
			results[s] = 6
			numToString[6] = s
			topright = onechars[0]
			continue
		} else if !strings.Contains(s, onechars[1]) {
			results[s] = 6
			numToString[6] = s
			topright = onechars[1]
			continue
		} else if aocutil.StringContainsAll(s, numToString[4]) {
			// The 6 letter word that contains all of the letters from 4 is 9
			results[s] = 9
			numToString[4] = s
		} else {
			results[s] = 0
			numToString[0] = s
		}
	}

	// Now do the 5's
	for _, s := range signals {
		if len(s) != 5 {
			continue
		}

		if aocutil.StringContainsAll(s, numToString[7]) {
			results[s] = 3
			numToString[3] = s
		} else if aocutil.StringContainsAll(s, topright) {
			results[s] = 2
			numToString[2] = s
		} else {
			results[s] = 5
			numToString[5] = s
		}
	}

	//fmt.Println(results)

	return results
}

// We know:
// 0 - 6 signals
// 1 = 2 signals
// 2 - 5 signals
// 3 - 5 signals
// 4 - 4 signals
// 5 - 5 signals
// 6 - 6 signals
// 7 - 3 signals
// 8 - 7 signals
// 9 - 6 signals

// If I have 1 - and I have something with 5, and it contains the same 2 entries, then
// the one with the same entries must be 3

// If I have 4, and I have something with 6 that has all 4 entries, it must be 9

// If I have 7 and 1, then I know the difference maps to a

// If I have

func main() {

	input := "input8.txt"

	vals := aocutil.LoadDelimitedStringArray(input, "|")

	sortEntries(&vals)

	known := 0

	totalPart2 := 0
	for _, entry := range vals {
		signals := strings.Fields(entry[0])
		numbers := strings.Fields(entry[1])

		// Part 1
		for _, n := range numbers {
			switch len(n) {
			case 2:
				// Found a 1
				known++
			case 3:
				// Found a 7
				known++
			case 4:
				// Founda  4
				known++
			case 7:
				// Found an 8
				known++
			}
		}

		// Part 2
		numMap := crunchSignals(signals)

		// Now we have the map, so figure out what the number is
		val := numMap[numbers[0]]*1000 + numMap[numbers[1]]*100 + numMap[numbers[2]]*10 + numMap[numbers[3]]

		totalPart2 += val

	}

	fmt.Printf("Total known numbers: %d\n", known)
	fmt.Printf("Total for part 2: %d\n", totalPart2)

}
