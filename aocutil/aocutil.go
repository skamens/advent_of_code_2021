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
