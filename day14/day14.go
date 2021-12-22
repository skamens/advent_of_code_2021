package main

import (
	"aoc21/aocutil"
	"fmt"
	"math"
	"strings"
)

func main() {

	input := "input14.txt"

	lines := aocutil.LoadStringArray(input)

	template := lines[0]

	subs := make(map[string]string)

	// Process rules into a map to make life easier
	for l := 2; l < len(lines); l++ {
		a := strings.Split(lines[l], " -> ")
		subs[a[0]] = a[1]
	}

	fmt.Println(subs)

	fmt.Println(template)

	pairCounts := make(map[string]int)
	// Start with the initial template
	for position := 0; position < len(template)-1; position++ {
		pair := template[position : position+2]
		pairCounts[pair]++
	}

	letterCounts := make(map[rune]int)
	// initialize the letter counts
	for _, r := range template {
		letterCounts[r]++
	}

	// Gonna do 10 times through
	for i := 0; i < 40; i++ {
		// Now go through all the pairs and change counts appropriately
		newPairCounts := make(map[string]int)

		for pair := range pairCounts {
			// For each pair, add and remove an appopriate number of pairs
			// in newPairCounts

			newPairCounts[pair] -= pairCounts[pair]
			newPairCounts[pair[0:1]+subs[pair]] += pairCounts[pair]
			newPairCounts[subs[pair]+pair[1:2]] += pairCounts[pair]

			for _, r := range subs[pair] {
				letterCounts[r] += pairCounts[pair]
			}
		}

		for pair := range newPairCounts {
			pairCounts[pair] += newPairCounts[pair]
		}

		fmt.Println(pairCounts)
	}

	highest := 0
	lowest := math.MaxInt
	for _, c := range letterCounts {
		if c > highest {
			highest = c
		}

		if c < lowest {
			lowest = c
		}
	}

	fmt.Printf("Highest: %d, Lowest: %d, diff: %d", highest, lowest, highest-lowest)

}
