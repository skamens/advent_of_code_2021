package main

import (
	"aoc21/aocutil"
	"fmt"
	"log"
	"math"
)

func findMapRange(data map[int]map[int]int) (int, int, int, int) {
	var minx, maxx, miny, maxy int
	minx = math.MaxInt64
	miny = math.MaxInt64
	maxx = math.MinInt
	maxy = math.MinInt

	for y := range data {
		if y < miny {
			miny = y
		}
		if y > maxy {
			maxy = y
		}

		for x := range data[y] {
			if x < minx {
				minx = x
			}
			if x > maxx {
				maxx = x
			}
		}
	}

	return minx, maxx, miny, maxy
}

func printMap(data map[int]map[int]int) {
	minx, maxx, miny, maxy := findMapRange(data)

	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if data[y][x] == 1 {
				fmt.Print("#")
			} else if data[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

}

func enhance(data *map[int]map[int]int, algo [512]int, newPixelValue int) {
	minx, maxx, miny, maxy := findMapRange(*data)

	newData := make(map[int]map[int]int)

	// Add 3 extra rows/columns around the outside of our
	// grid. These will be marked as background through the
	// use of the newPixelValue

	for y := miny - 3; y <= maxy+3; y++ {
		_, ok := (*data)[y]
		if !ok {
			(*data)[y] = make(map[int]int)

			// Since we added an entire new row, initialize the data to -1
			for x := minx - 3; x <= maxx+3; x++ {
				(*data)[y][x] = newPixelValue
			}
		} else {
			// We didn't add a new row, but we need to initialize the outside
			// pixels to -1

			for x := minx - 3; x < minx; x++ {
				(*data)[y][x] = newPixelValue
			}
			for x := maxx + 1; x <= maxx+3; x++ {
				(*data)[y][x] = newPixelValue
			}
		}
		_, ok = newData[y]
		if !ok {
			newData[y] = make(map[int]int)
		}
	}

	var origVal int
	var val int
	var newVal int

	for y := miny - 3; y <= maxy+3; y++ {
		for x := minx - 3; x <= maxx+3; x++ {
			// Remember this value so we can decide if
			// this pixel was background or not
			origVal = (*data)[y][x]
			num := 0
			// Enhance this particular point
			for yy := y - 1; yy <= y+1; yy++ {
				for xx := x - 1; xx <= x+1; xx++ {
					num = num << 1

					if yy < miny-3 || yy > maxy+3 || xx < minx-3 || xx > maxx+3 {
						val = newPixelValue
					} else {
						val = (*data)[yy][xx]
					}

					if val == 1 || val == -2 {
						num |= 1
					}
				}
			}
			newVal = algo[num]
			if origVal < 0 {
				if newVal == 0 {
					newVal = -1
				} else {
					newVal = -2
				}
			}
			newData[y][x] = newVal
		}
	}

	(*data) = newData
}

func countLights(data *map[int]map[int]int) int {
	cnt := 0
	for y := range *data {
		for x := range (*data)[y] {
			if (*data)[y][x] == 1 {
				cnt++
			}
		}
	}

	return cnt
}

func main() {

	input := "input20.txt"

	lines := aocutil.LoadStringArray(input)

	// First line is the enhancement algorithm
	var algo [512]int
	data := make(map[int]map[int]int)

	for n, b := range lines[0] {
		if b == '.' {
			algo[n] = 0
		} else if b == '#' {
			algo[n] = 1
		} else {
			log.Fatal("Unknown value!")
		}
	}

	// Skipping the first line

	for y, l := range lines[2:] {
		data[y] = make(map[int]int)
		for x, b := range l {
			if rune(b) == '#' {
				data[y][x] = 1
			} else {
				data[y][x] = 0
			}
		}
	}

	printMap(data)

	for i := 1; i <= 50; i++ {
		enhance(&data, algo, -1*(i%2))
		printMap(data)
		fmt.Printf("Iteration %d: %d\n", i, countLights(&data))
	}

}
