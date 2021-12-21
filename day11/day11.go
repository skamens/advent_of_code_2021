package main

import (
	"aoc21/aocutil"
	"fmt"
)

func raiseEnergy(grid *[][]int) {
	for x := 0; x < len(*grid); x++ {
		for y := 0; y < len((*grid)[0]); y++ {
			(*grid)[x][y]++
		}
	}
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func shouldFlash(val int) bool {
	return val > 9
}

func main() {

	input := "input11.txt"

	grid := aocutil.Load2DArray(input)
	flashCount := 0

	var numFlashed int

	printGrid(grid)
	for step := 0; ; step++ {
		numFlashed = 0
		raiseEnergy(&grid)

		flashers := aocutil.FindInGrid(grid, shouldFlash, 9)
		for len(flashers) > 0 {
			for _, flasher := range flashers {
				flashCount++
				numFlashed++
				grid[flasher.X][flasher.Y] = 0
				adjacents := aocutil.GetAdjacentPoints(grid, flasher)
				for _, adjacent := range adjacents {
					if grid[adjacent.X][adjacent.Y] != 0 {
						grid[adjacent.X][adjacent.Y]++
					}
				}
			}
			flashers = aocutil.FindInGrid(grid, shouldFlash, 9)
		}

		if numFlashed == len(grid)*len(grid[0]) {
			printGrid(grid)
			fmt.Printf("All octopi flashed! step %d\n", step+1)
			break
		}
		printGrid(grid)
	}

	fmt.Printf("Total Flashes: %d\n", flashCount)
}
