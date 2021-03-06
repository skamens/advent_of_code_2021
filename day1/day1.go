package main

import (
	"aoc21/aocutil"
	"fmt"
	"math"
)

func main() {

	input := "input1.txt"

	vals := aocutil.LoadIntArray(input)
	last := math.MaxInt64
	increases := 0
	for i := 0; i < len(vals)-2; i++ {
		sum := vals[i] + vals[i+1] + vals[i+2]
		if sum > last {
			increases++
		}
		last = sum
	}

	//for _, depth := range vals {
	// 	if depth > last && last != -1 {
	// 		increases++
	// 	}

	// 	last = depth
	// }

	fmt.Println(increases)
}
