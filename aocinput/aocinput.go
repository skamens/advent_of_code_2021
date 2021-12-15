package aocinput

import (
	"bufio"
	"log"
	"os"
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
