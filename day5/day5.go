package main

import (
	"aoc21/aocinput"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	P1 Point
	P2 Point
}

var Lines []Line

var Grid [][]int

func main() {

	input := "input5.txt"

	vals := aocinput.LoadStringArray(input)

	maxX := 0
	maxY := 0
	// Process the input, store all lines
	for _, s := range vals {
		f := strings.Fields(s)
		var p1, p2 Point
		xy := strings.Split(f[0], ",")
		p1.X, _ = strconv.Atoi(xy[0])
		p1.Y, _ = strconv.Atoi(xy[1])

		if p1.X > maxX {
			maxX = p1.X
		}

		if p1.Y > maxY {
			maxY = p1.Y
		}

		xy = strings.Split(f[2], ",")
		p2.X, _ = strconv.Atoi(xy[0])
		p2.Y, _ = strconv.Atoi(xy[1])

		if p2.X > maxX {
			maxX = p2.X
		}

		if p2.Y > maxY {
			maxY = p2.Y
		}

		var line = Line{P1: p1, P2: p2}
		Lines = append(Lines, line)
	}

	Grid = make([][]int, maxX+1, maxX+1)
	for i := 0; i <= maxX; i++ {
		Grid[i] = make([]int, maxY+1, maxY+1)
	}

	// Now process the lines
	for _, line := range Lines {

		fmt.Println(line)

		if line.P1.X == line.P2.X {
			// It's horizontal, so we move through the vertical
			var ystart, yfinish int
			if line.P1.Y < line.P2.Y {
				ystart = line.P1.Y
				yfinish = line.P2.Y
			} else {
				ystart = line.P2.Y
				yfinish = line.P1.Y
			}

			for j := ystart; j <= yfinish; j++ {
				Grid[line.P1.X][j]++
			}
		} else if line.P1.Y == line.P2.Y {
			var xstart, xfinish int
			if line.P1.X < line.P2.X {
				xstart = line.P1.X
				xfinish = line.P2.X
			} else {
				xstart = line.P2.X
				xfinish = line.P1.X
			}

			for i := xstart; i <= xfinish; i++ {
				Grid[i][line.P1.Y]++
			}
		} else {
			i := line.P1.X
			j := line.P1.Y

			Grid[i][j]++

			for i != line.P2.X {

				if line.P1.X < line.P2.X {
					i++
				} else {
					i--
				}

				if line.P1.Y < line.P2.Y {
					j++
				} else {
					j--
				}

				Grid[i][j]++
			}
		}
	}

	// See how many overlaps there are
	numOverlaps := 0
	for _, row := range Grid {
		for _, val := range row {
			if val > 1 {
				numOverlaps++
			}
		}
	}

	fmt.Printf("Number of overlapping spots: %d\n", numOverlaps)
}
