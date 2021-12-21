package main

import (
	"aoc21/aocutil"
	"fmt"
	"math"
	"strings"
)

func elementCounts(template string) map[rune]int {
	result := make(map[rune]int)

	for _, r := range template {
		result[r]++
	}

	return result
}

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
	// Gonna do 10 times through
	for i := 0; i < 40; i++ {
		newtemplate := ""
		position := 0
		matchLen := 0
		var matchKey string
		for position < len(template)-1 {
			// Find the longest match in the map
			for key, value := range subs {
				if strings.HasPrefix(template[position:], key) {
					if len(key) > matchLen {
						matchKey = key
						matchLen = len(key)
					}
				}
			}
			// So now we have the longest match
			newtemplate.append(newtemplate, subs[key])
		}

			key := template[position : position+2]
			template = template[0:position+1] + subs[key] + template[position+1:]
			position += 2
		}
		fmt.Printf("Done with step %d; template length %d\n", i+1, len(template))
	}

	counts := elementCounts(template)

	highest := 0
	lowest := math.MaxInt
	for _, c := range counts {
		if c > highest {
			highest = c
		}

		if c < lowest {
			lowest = c
		}
	}

	fmt.Printf("Highest: %d, Lowest: %d, diff: %d", highest, lowest, highest-lowest)

}
