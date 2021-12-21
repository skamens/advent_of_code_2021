package aocutil

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func LoadIntArray(filename string) []int {
	var lines []int

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func LoadStringArray(filename string) []string {
	var lines []string

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func LoadDelimitedStringArray(filename string, delimiter string) [][]string {
	var lines [][]string

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "|")
		lines = append(lines, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

type StringIntEntry struct {
	S string
	V int
}

func LoadStringIntArray(filename string) []StringIntEntry {
	var result []StringIntEntry

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")

		v, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		e := StringIntEntry{S: s[0], V: v}
		result = append(result, e)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func LoadIntArrayLine(filename string) []int {
	var result []int

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	vals := strings.Split(scanner.Text(), ",")

	for _, v := range vals {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}

	return result
}

func Load2DArray(filename string) [][]int {
	var result [][]int

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "")

		var row []int
		for _, s := range vals {
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, v)
		}
		result = append(result, row)
	}

	return result

}

func SortArrayValues(arr *[]string) {
	for i := 0; i < len(*arr); i++ {
		s := strings.Split((*arr)[i], "")
		sort.Strings(s)

		(*arr)[i] = strings.Join(s, "")
	}
}

func StringContainsAll(container string, contents string) bool {
	for _, s := range strings.Split(contents, "") {
		if !strings.Contains(container, s) {
			return false
		}
	}
	return true
}

type Point struct {
	X int
	Y int
}

func GetAdjacentPoints(grid [][]int, p Point) []Point {

	results := make([]Point, 0, 6)

	var minx, maxx int
	var miny, maxy int

	if p.X == 0 {
		minx = 0
	} else {
		minx = p.X - 1
	}

	if p.X == len(grid)-1 {
		maxx = p.X
	} else {
		maxx = p.X + 1
	}

	if p.Y == 0 {
		miny = 0
	} else {
		miny = p.Y - 1
	}

	if p.Y == len(grid[0])-1 {
		maxy = p.Y
	} else {
		maxy = p.Y + 1
	}

	for i := minx; i <= maxx; i++ {
		for j := miny; j <= maxy; j++ {
			if !(i == p.X && j == p.Y) {
				results = append(results, Point{X: i, Y: j})
			}
		}
	}

	return results
}

type Checker func(int) bool

func FindInGrid(grid [][]int, comp Checker, value int) []Point {
	var results []Point
	for x, row := range grid {
		for y, val := range row {
			if comp(val) {
				results = append(results, Point{X: x, Y: y})
			}
		}
	}

	return results
}

func PointInList(list []Point, p Point) bool {
	for _, entry := range list {
		if entry.X == p.X && entry.Y == p.Y {
			return true
		}
	}
	return false
}
