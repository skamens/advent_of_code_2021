package main

import (
	"aoc21/aocinput"
	"fmt"
)

func main() {

	input := "input1.txt"

	vals := aocinput.LoadIntArray(input)
	last := -1
	increases := 0
	for _, depth := range vals {
		if depth > last && last != -1 {
			increases++
		}

		last = depth
	}

	fmt.Println(increases)
}
