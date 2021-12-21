package main

import (
	"aoc21/aocutil"
	"fmt"
	"strings"
)

// Given the first portion of a path, find all paths that come from it.
func findPaths(connections map[string][]string, path []string) [][]string {

	//	fmt.Printf("findPaths: ")
	//	fmt.Println(path)

	var results [][]string
	for _, n := range connections[path[len(path)-1]] {
		if n == "start" {
			continue
		}
		if n == "end" {
			newpath := make([]string, len(path))
			copy(newpath, path)
			newpath = append(newpath, n)
			results = append(results, newpath)
			continue
		} else if strings.ToLower(n) == n {
			// n is all lowercase. Now let's see if n is already in the path
			found := false
			for _, p := range path {
				if p == n {
					found = true
					break
				}
			}
			if found && hasDoubleLowercase(path) {
				// If I found myself already and there is already a double
				// lowercase entry in the path, then I have to skip this one.
				// Otherwise we will go ahead and add this one a second time
				continue
			}
		}
		newpath := make([]string, len(path))
		copy(newpath, path)
		newpath = append(newpath, n)
		paths := findPaths(connections, newpath)
		for _, p := range paths {
			results = append(results, p)

		}
	}
	//fmt.Println(results)
	return results
}

func hasDoubleLowercase(path []string) bool {
	for i := 0; i < len(path); i++ {
		if strings.ToLower(path[i]) == path[i] {
			// It's a lowercase value. Now see if there's another of the same

			for j := i + 1; j < len(path); j++ {
				if path[j] == path[i] {
					// This means there is a double lowercase entry in the path
					return true
				}
			}
		}
	}

	// If we got here, there isn't a double lowercase entry
	return false
}

func main() {

	input := "input12.txt"

	edges := aocutil.LoadDelimitedStringArray(input, "-")

	connections := make(map[string][]string)

	for _, edge := range edges {
		connections[edge[0]] = append(connections[edge[0]], edge[1])

		if (edge[0] != "start") && edge[1] != "end" {
			connections[edge[1]] = append(connections[edge[1]], edge[0])
		}
	}
	fmt.Println(connections)

	results := findPaths(connections, []string{"start"})
	fmt.Println(results)
	fmt.Printf("Number of paths: %d\n", len(results))
}
