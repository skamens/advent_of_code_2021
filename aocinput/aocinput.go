package aocinput

import (
	"bufio"
	"log"
	"os"
	"strconv"
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
