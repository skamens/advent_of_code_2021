package main

import (
	"aoc21/aocutil"
	"fmt"
	"math"
)

func adjacentPoints(grid [][]int, p aocutil.Point) []aocutil.Point {
	// Find the adjacent points
	var possiblePoints []aocutil.Point

	if p.X < len(grid)-1 {
		possiblePoints = append(possiblePoints, aocutil.Point{X: (p.X + 1), Y: p.Y})
	}

	if p.Y < len(grid[0])-1 {
		possiblePoints = append(possiblePoints, aocutil.Point{X: p.X, Y: (p.Y + 1)})
	}

	if p.X > 0 {
		possiblePoints = append(possiblePoints, aocutil.Point{X: (p.X - 1), Y: p.Y})
	}

	if p.Y > 0 {
		possiblePoints = append(possiblePoints, aocutil.Point{X: p.X, Y: (p.Y - 1)})
	}

	return possiblePoints
}

func findLowRiskPath(grid [][]int, risks map[aocutil.Point]int, p aocutil.Point, soFar map[aocutil.Point]bool) int {

	// If I already have this point, just return the value
	if risks[p] > 0 {
		fmt.Printf("findLowRiskPath(%d, %d) = %d\n", p.X, p.Y, risks[p])
		return risks[p]
	}

	lowestRisk := math.MaxInt8

	// Now from here figure out the adjacent points to try
	soFar[p] = true
	for _, a := range adjacentPoints(grid, p) {
		if !soFar[a] {
			r := findLowRiskPath(grid, risks, a, soFar)
			if r < lowestRisk {
				lowestRisk = r
			}
		}
	}

	// Found the lowest risk from here - so
	// save the risk of the current spot
	risks[p] = grid[p.X][p.Y] + lowestRisk

	fmt.Printf("findLowRiskPath(%d, %d) = %d\n", p.X, p.Y, risks[p])
	return risks[p]
}

func main() {

	input := "input15.txt"

	grid := aocutil.Load2DArray(input)

	risks := make(map[aocutil.Point]int)

	// Set the final point so the recursion will terminate
	var p aocutil.Point

	p.X = len(grid[0]) - 1
	p.Y = len(grid) - 1

	// Add the risk total for this one point
	risks[p] = grid[p.X][p.Y]

	for x := len(grid[0]) - 2; x >= 0; x-- {
		for y := len(grid) - 2; y >= 0; y-- {
			soFar := make(map[aocutil.Point]bool)
			r := findLowRiskPath(grid, risks, aocutil.Point{X: x, Y: y}, soFar)
			fmt.Printf("Lowest Risk (%d, %d): %d\n", x, y, r)
		}
	}

	soFar := make(map[aocutil.Point]bool)
	r := findLowRiskPath(grid, risks, aocutil.Point{X: 0, Y: 0}, soFar)
	fmt.Printf("Lowest risk path: %d\n", r)
}
