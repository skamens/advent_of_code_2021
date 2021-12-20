package main

import (
	"aoc21/aocutil"
	"fmt"
	"sort"
	"strings"
)

func main() {

	openers := "([{<"
	closers := ")]}>"

	pairs := make(map[string]string)
	pairs["("] = ")"
	pairs[")"] = "("

	pairs["["] = "]"
	pairs["]"] = "["

	pairs["{"] = "}"
	pairs["}"] = "{"

	pairs["<"] = ">"
	pairs[">"] = "<"

	input := "input10.txt"

	lines := aocutil.LoadStringArray(input)

	corruptTotal := 0
	var corrupt bool
	var allCompletionScores []int64

	for _, line := range lines {
		var stack []string
		corrupt = false
		for idx, char := range strings.Split(line, "") {
			if strings.Contains(openers, char) {
				// It's an opener, so append it to the stack
				stack = append(stack, char)
			} else if strings.Contains(closers, char) {
				if stack[len(stack)-1] == pairs[char] {
					stack = stack[:len(stack)-1]
				} else {
					switch char {
					case ")":
						corruptTotal += 3
					case "]":
						corruptTotal += 57
					case "}":
						corruptTotal += 1197
					case ">":
						corruptTotal += 25137
					}
					corrupt = true
					fmt.Printf("CORRUPT! line: %s, char(%d): %s, corruptTotal=%d\n", line, idx, char, corruptTotal)

					break
				}
			}
		}

		if corrupt {
			continue
		}

		if len(stack) == 0 {
			fmt.Printf("Line %s is complete!\n", line)
			continue
		}

		// If we get here and the stack isn't empty, we have a correct but incomplete line.
		// Easy enough to figure out how to complete the line, by starting at the end and
		// pairing up properly
		var completion []string

		completionScore := 0

		for len(stack) > 0 {
			completion = append(completion, pairs[stack[len(stack)-1]])
			completionScore *= 5
			switch pairs[stack[len(stack)-1]] {
			case ")":
				completionScore += 1
			case "]":
				completionScore += 2
			case "}":
				completionScore += 3
			case ">":
				completionScore += 4
			}
			stack = stack[:len(stack)-1]
		}
		allCompletionScores = append(allCompletionScores, int64(completionScore))

		fmt.Printf("Line: %s, completion: %s, score: %d\n", line, completion, completionScore)
	}

	// Now we sort the allCompletionScores slice and find the middle one
	sort.Slice(allCompletionScores, func(i, j int) bool { return allCompletionScores[i] < allCompletionScores[j] })
	middle := len(allCompletionScores) / 2
	fmt.Printf("Array length %d, middle index %d, value %d\n", len(allCompletionScores), middle, allCompletionScores[middle])
}
