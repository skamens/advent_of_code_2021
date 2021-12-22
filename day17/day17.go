package main

import (
	"aoc21/aocutil"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	input := "test17.txt"

	target := aocutil.LoadSingleString(input)

	arr := strings.Split(target, " ")

	// This will yield 4 values: "target", "area:", "x=..." and "y=..."

	xrange := strings.TrimPrefix(arr[2], "x=") //20..30
	xrange = strings.TrimSuffix(xrange, ",")
	xarr := strings.Split(xrange, "..")
	minx, _ := strconv.Atoi(xarr[0])
	maxx, _ := strconv.Atoi(xarr[1])

	yrange := strings.TrimPrefix(arr[3], "y=") //20..30
	yarr := strings.Split(yrange, "..")
	miny, _ := strconv.Atoi(yarr[0])
	maxy, _ := strconv.Atoi(yarr[1])

	fmt.Printf("minx: %d, maxx: %d, miny: %d, maxy: %d\n", minx, maxx, miny, maxy)
}
