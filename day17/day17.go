package main

import (
	"aoc21/aocutil"
	"fmt"
	"strconv"
	"strings"
)

type probe struct {
	xpos int
	ypos int
	xvel int
	yvel int

	highesty int
}

func step(p *probe, minx int, maxx int, miny int, maxy int) bool {
	p.xpos += p.xvel
	p.ypos += p.yvel

	if p.ypos > p.highesty {
		p.highesty = p.ypos
	}

	if p.xvel > 0 {
		p.xvel--
	} else if p.xvel < 0 {
		p.xvel++
	}

	p.yvel--

	// OK, now is it in the target area. X and Y must both be within
	// the ranges

	if p.xpos >= minx && p.xpos <= maxx && p.ypos >= miny && p.ypos <= maxy {
		return true
	} else {
		return false
	}
}

func main() {

	input := "input17.txt"

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

	highesty := 0

	points := make(map[aocutil.Point]int)

	for initialy := -500; initialy < 500; initialy++ {
		fmt.Printf("initialy=%d\n", initialy)
		//anysuccess := false

		for initialx := 0; initialx < 500; initialx++ {
			//fmt.Printf("   initialx=%d ... ", initialx)

			var p = probe{xpos: 0, ypos: 0, xvel: initialx, yvel: initialy, highesty: 0}

			//success := false
			for p.xpos <= maxx && p.ypos >= miny {
				if step(&p, minx, maxx, miny, maxy) {
					//anysuccess = true
					//success = true

					points[aocutil.Point{X: initialx, Y: initialy}]++
					//fmt.Println(points)

					if p.highesty > highesty {
						fmt.Printf("Highest y = %d, initialy = %d\n", p.highesty, initialy)
						highesty = p.highesty
					}
					break
				}
			}
			//fmt.Println(success)
		}
		//if !anysuccess {
		//	break
		//}

	}

	fmt.Printf("Done. highest y is at %d\n", highesty)
	fmt.Printf("Total velocities: %d\n", len(points))

}
