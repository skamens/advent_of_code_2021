package main

import (
	"aoc21/aocutil"
	"fmt"
)

func main() {

	input := "test11.txt"

	grid := aocutil.Load2DArray(input)

	fmt.Printf("Array length %d, middle index %d, value %d\n", len(allCompletionScores), middle, allCompletionScores[middle])
}
