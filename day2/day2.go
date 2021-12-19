package main

import (
	"aoc21/aocutil"
	"fmt"
)

func main() {

	input := "input2.txt"

	vals := aocutil.LoadStringIntArray(input)

	depth := 0
	position := 0
	aim := 0

	// part 2
	for _, e := range vals {
		switch e.S {
		case "down":
			aim += e.V
		case "up":
			aim -= e.V
		case "forward":
			position += e.V
			depth += e.V * aim
		}
	}

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

	fmt.Printf("Depth: %d, Position: %d, Product: %d", depth, position, depth*position)
}
