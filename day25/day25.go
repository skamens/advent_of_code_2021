package main

import (
	"aoc21/aocutil"
	"fmt"
)

func printGrid(grid map[int]map[int]rune) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

}

func movecuke(grid *map[int]map[int]rune, y int, x int) int {
	if (*grid)[y][x] == '.' {
		return 0
	}

	if (*grid)[y][x] == '>' {
		// See if the spot one to the right is open.
		// If so, move there.

		newX := (x + 1) % len((*grid)[y])
		if (*grid)[y][newX] == '.' {
			(*grid)[y][newX] = '<'
			(*grid)[y][x] = '*'
			return 1
		}
	} else if (*grid)[y][x] == 'v' {
		// It's down

		newY := (y + 1) % len(*grid)
		if (*grid)[newY][x] == '.' {
			(*grid)[newY][x] = '^'
			(*grid)[y][x] = '*'
			return 1
		}
	}

	return 0
}

func moveAll(grid *map[int]map[int]rune) int {
	var totalMoves int

	// First do all the '>'
	for y := 0; y < len(*grid); y++ {
		for x := 0; x < len((*grid)[y]); x++ {
			if (*grid)[y][x] == '>' {
				totalMoves += movecuke(grid, y, x)
			}
		}
	}

	// Now switch the characters back.
	for y := 0; y < len(*grid); y++ {
		for x := 0; x < len((*grid)[y]); x++ {
			if (*grid)[y][x] == '<' {
				(*grid)[y][x] = '>'
			}
			if (*grid)[y][x] == '^' {
				(*grid)[y][x] = 'v'
			}
			if (*grid)[y][x] == '*' {
				(*grid)[y][x] = '.'
			}
		}
	}
	// Now do all of the 'v'
	for y := 0; y < len(*grid); y++ {
		for x := 0; x < len((*grid)[y]); x++ {
			if (*grid)[y][x] == 'v' {
				totalMoves += movecuke(grid, y, x)
			}
		}
	}

	// Now switch the characters back.
	for y := 0; y < len(*grid); y++ {
		for x := 0; x < len((*grid)[y]); x++ {
			if (*grid)[y][x] == '<' {
				(*grid)[y][x] = '>'
			}
			if (*grid)[y][x] == '^' {
				(*grid)[y][x] = 'v'
			}
			if (*grid)[y][x] == '*' {
				(*grid)[y][x] = '.'
			}
		}
	}

	return totalMoves

}

func main() {

	grid := make(map[int]map[int]rune)

	input := "input25.txt"

	lines := aocutil.LoadStringArray(input)

	for y, l := range lines {
		grid[y] = make(map[int]rune)

		for x, b := range l {
			grid[y][x] = b
		}
	}

	//printGrid(grid)

	var moveCount int
	for {
		moveCount++

		totalMoves := moveAll(&grid)

		fmt.Printf("After move %d: totalMoves: %d\n", moveCount, totalMoves)
		//printGrid(grid)
		if totalMoves == 0 {
			fmt.Printf("Done! Move=%d", moveCount)
			return
		}

	}

}
