package main

import (
	"aoc21/aocutil"
	"fmt"
	"strconv"
	"strings"
)

type cubex map[int]bool
type cubexy map[int]cubex
type cubexyz map[int]cubexy

func rangestring(s string) (start, end int) {
	arr := strings.Split(s, "..")

	start, _ = strconv.Atoi(arr[0])
	end, _ = strconv.Atoi(arr[1])
	return start, end
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
func main() {

	input := "test22.txt"

	lines := aocutil.LoadStringArray(input)

	// Initialize the maps I'm using

	cubes := make(map[int]cubexy)

	for z := -50; z <= 50; z++ {
		cubes[z] = make(cubexy)
		for y := -50; y <= 50; y++ {
			cubes[z][y] = make(cubex)
		}
	}

	// Now process the lines
	for _, l := range lines {
		var newVal bool
		var minx, maxx, miny, maxy, minz, maxz int

		s := strings.Split(l, " ")

		newVal = (s[0] == "on")

		s = strings.Split(s[1], ",")

		minx, maxx = rangestring(strings.Split(s[0], "=")[1])
		miny, maxy = rangestring(strings.Split(s[1], "=")[1])
		minz, maxz = rangestring(strings.Split(s[2], "=")[1])

		for z := minz; z <= maxz; z++ {
			_, okz := cubes[z]
			if !okz {
				cubes[z] = make(cubexy)
			}
			for y := miny; y <= maxy; y++ {
				_, oky := cubes[z][y]
				if !oky {
					cubes[z][y] = make(cubex)
				}

				for x := minx; x <= maxx; x++ {
					cubes[z][y][x] = newVal
				}
			}
		}
	}

	var total int64

	// Count on cubes
	for rz := range cubes {
		for ry := range cubes[rz] {
			for rx := range cubes[rz][ry] {
				if cubes[rz][ry][rx] {
					total++
				}
			}
		}
	}

	fmt.Printf("Total cubes on: %d\n", total)
}
