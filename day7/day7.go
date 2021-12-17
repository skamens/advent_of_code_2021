package main

import (
	"aoc21/aocinput"
	"fmt"
	"math"
)

func fuelCost(distance int) int {
	cost := 0
	for i := 1; i <= distance; i++ {
		cost += i
	}
	return cost
}

func main() {

	input := "input7.txt"

	vals := aocinput.LoadIntArrayLine(input)

	lowest := math.MaxInt64

	// Find the smallest and largest values in the list
	minval := math.MaxInt64
	maxval := math.MinInt64
	for _, v := range vals {
		if v < minval {
			minval = v
		}

		if v > maxval {
			maxval = v
		}
	}

	for target := minval; target <= maxval; target++ {

		// target is the proposed target position to end up at
		total := 0
		for _, v := range vals {
			dist := target - v
			if dist < 0 {
				dist *= -1
			}

			total = total + fuelCost(dist)
			if total > lowest {
				continue
			}
		}

		if total < lowest {
			lowest = total
		}

		fmt.Printf("Target: %d, Total: %d, Lowest: %d\n", target, total, lowest)

	}
}
