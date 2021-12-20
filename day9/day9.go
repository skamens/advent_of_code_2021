package main

import (
	"aoc21/aocutil"
	"fmt"
	"sort"
)

type Point struct {
	X int
	Y int
}

func isLowPoint(grid *[][]int, x int, y int) bool {

	var minx, maxx int
	var miny, maxy int
	if x == 0 {
		minx = 0
	} else {
		minx = x - 1
	}

	if x == len(*grid)-1 {
		maxx = x
	} else {
		maxx = x + 1
	}

	if y == 0 {
		miny = 0
	} else {
		miny = y - 1
	}

	if y == len((*grid)[0])-1 {
		maxy = y
	} else {
		maxy = y + 1
	}

	fmt.Printf("Value to find:(%d, %d)=%d\n", x, y, (*grid)[x][y])

	for i := minx; i <= maxx; i++ {
		for j := miny; j <= maxy; j++ {
			if !(i == x && j == y) {
				fmt.Printf("Checking (%d, %d)=%d\n", i, j, (*grid)[i][j])
				if (*grid)[i][j] < (*grid)[x][y] {
					return false
				}
			}
		}
	}

	// If we get here, we didn't find a lower point around this one, so this is a low point
	return true
}

func pointArrayContains(points []Point, point Point) bool {
	for _, p := range points {
		if p.X == point.X && p.Y == point.Y {
			return true
		}
	}
	return false
}

func main() {

	input := "input9.txt"

	grid := aocutil.Load2DArray(input)

	fmt.Println(grid)

	var lowPoints []Point

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if isLowPoint(&grid, i, j) {
				var p Point
				p.X = i
				p.Y = j
				lowPoints = append(lowPoints, p)
			}
		}
	}

	fmt.Println(lowPoints)

	var totalRisk int
	for _, p := range lowPoints {
		totalRisk += grid[p.X][p.Y] + 1
	}

	fmt.Printf("Total Risk: %d\n", totalRisk)

	// Now find the basins. For each low point, we'll create a queue of
	// points to consider, stopping at 9's.

	var basinSizes []int

	for _, lowPoint := range lowPoints {
		var candidatePoints []Point
		var basinPoints []Point

		candidatePoints = append(candidatePoints, lowPoint)

		for len(candidatePoints) > 0 {
			p := candidatePoints[0]
			candidatePoints = candidatePoints[1:] // Remove the first item from the candidates

			if pointArrayContains(basinPoints, p) {
				continue
			}

			// Add the candidate to the basin points; add adjacent points to candidatePoints
			basinPoints = append(basinPoints, p)

			// Find the adjacent points
			var possiblePoints []Point
			if p.X > 0 {
				possiblePoints = append(possiblePoints, Point{X: (p.X - 1), Y: p.Y})
			}

			if p.X < len(grid)-1 {
				possiblePoints = append(possiblePoints, Point{X: (p.X + 1), Y: p.Y})
			}

			if p.Y > 0 {
				possiblePoints = append(possiblePoints, Point{X: p.X, Y: (p.Y - 1)})
			}

			if p.Y < len(grid[0])-1 {
				possiblePoints = append(possiblePoints, Point{X: p.X, Y: (p.Y + 1)})
			}

			for _, possible := range possiblePoints {
				fmt.Printf("Checking (%d, %d)=%d\n", possible.X, possible.Y, grid[possible.X][possible.Y])

				if grid[possible.X][possible.Y] != 9 {
					if !pointArrayContains(basinPoints, possible) && !pointArrayContains(candidatePoints, possible) {
						candidatePoints = append(candidatePoints, possible)
					}
				}
			}
		}

		// At this point basinPoints will have all the points

		fmt.Printf("Basin around (%d, %d) of size %d\n", lowPoint.X, lowPoint.Y, len(basinPoints))
		basinSizes = append(basinSizes, len(basinPoints))
		sort.Ints(basinSizes)
	}

	fmt.Println(basinSizes)

	// Now find the last 3 basin sizes and multipliy them

	final := 1
	for i := len(basinSizes) - 3; i < len(basinSizes); i++ {
		final *= basinSizes[i]
	}
	fmt.Printf("Final Answer: %d", final)
}
