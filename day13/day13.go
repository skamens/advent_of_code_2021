package main

import (
	"aoc21/aocutil"
	"fmt"
	"strconv"
	"strings"
)

func printGrid(grid [][]bool, maxx int, maxy int) {
	for y := 0; y < maxy; y++ {
		for x := 0; x < maxx; x++ {
			if grid[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func countDots(grid [][]bool) int {
	totalOn := 0
	// Now count the true values
	for _, row := range grid {
		for _, val := range row {
			if val {
				totalOn++
			}
		}
	}
	return totalOn
}

func main() {

	input := "input13.txt"

	lines := aocutil.LoadStringArray(input)

	maxx := 0
	maxy := 0

	var points []aocutil.Point
	var folds []string
	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "fold") {
			line = strings.TrimPrefix(line, "fold along ")
			folds = append(folds, line)
			continue
		}

		vals := strings.Split(line, ",")

		var p aocutil.Point
		p.X, _ = strconv.Atoi(vals[0])
		p.Y, _ = strconv.Atoi(vals[1])
		points = append(points, p)

		if p.X > maxx {
			maxx = p.X
		}
		if p.Y > maxy {
			maxy = p.Y
		}
	}

	// Allocate a grid as big as the number of points we have

	var grid = make([][]bool, maxy+1, maxy+1)
	for i := 0; i < maxy+1; i++ {
		grid[i] = make([]bool, maxx+1, maxx+1)
	}

	printGrid(grid, maxx, maxy)

	for _, p := range points {
		grid[p.Y][p.X] = true
	}

	printGrid(grid, maxx, maxy)

	for _, fold := range folds {
		vals := strings.Split(fold, "=")
		foldPoint, _ := strconv.Atoi(vals[1])
		if vals[0] == "x" {
			// Fold vertically (at x)

			for x := foldPoint + 1; x <= maxx; x++ {
				for y := 0; y <= maxy; y++ {
					if !grid[y][x] {
						// We only need to move true values
						continue
					}
					grid[y][foldPoint-(x-foldPoint)] = grid[y][x]
					grid[y][x] = false
				}
			}
			maxx = foldPoint
			printGrid(grid, maxx, maxy)
			fmt.Printf("Total on: %d\n", countDots(grid))

		} else if vals[0] == "y" {
			// Fold horizontally at y

			for y := foldPoint + 1; y <= maxy; y++ {
				for x := 0; x <= maxx; x++ {
					if !grid[y][x] {
						// We only need to move true values
						continue
					}
					grid[foldPoint-(y-foldPoint)][x] = grid[y][x]
					grid[y][x] = false
				}
			}
			maxy = foldPoint
			printGrid(grid, maxx, maxy)
			fmt.Printf("Total on: %d\n", countDots(grid))
		}
	}
}
