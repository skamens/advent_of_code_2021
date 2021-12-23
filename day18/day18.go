package main

import (
	"aoc21/aocutil"
	"fmt"
	"strconv"

	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type snail struct {
	leftnum    int
	leftsnail  *snail
	rightnum   int
	rightsnail *snail
	parent     *snail
}

func reduce(snail *s, depth int) {
	if s == nil {
		return
	}

	if depth == 4 {
		// Find the first regular number to the left
		for p := s.parent; p; p = p.parent {
			if p.leftnum != -1 {
				p.leftnum += s.leftnum
				break
			}
		}

		for p := s.parent; p; p = p.parent {
			if (p.rightnum != -1 {
				p.rightnum += s.rightnum
				break
			}
		}

		# This one becomes just 0
		if (s.parent.leftsnail == s) {
			s.parent.leftsnail = nil
			s.parent.leftnum = 0
		} else if (s.parent.rightsnail == s) {
			s.parent.rightsnail = nil
			s.parent.rightnum = 0
		}
	}
}

func printSnail(s *snail) string {
	
}


func parseSnail(in string) *snail {

	newSnail := snail{-1, nil, -1, nil, nil}

	var s string

	s = in

	fmt.Printf("\ns: %s\n", s)
	var leftString, rightString string

	if s[0] == '[' {
		// Strip off the braces
		s = s[1 : len(s)-1]
	}

	// Now, see what the first character is
	if s[0] == '[' {
		// Left one is another snail. Find the pair bracket
		level := 1
		endpos := 1
		for {
			if s[endpos] == ']' {
				level--
				if level == 0 {
					break
				}
			} else if s[endpos] == '[' {
				level++
			}
			endpos++
		}

		leftString = s[0 : endpos+1]
		rightString = s[endpos+2:]
	} else {
		leftString = s[0:1]
		rightString = s[2:]
	}

	fmt.Printf("leftString: %s, rightString, %s\n", leftString, rightString)
	if len(leftString) > 1 {
		newSnail.leftsnail = parseSnail(leftString)
		newSnail.leftsnail.parent = newSnail
	} else {
		newSnail.leftnum, _ = strconv.Atoi(leftString)
	}

	if len(rightString) > 1 {
		newSnail.rightsnail = parseSnail(rightString)
		newSnail.rightsnail.parent = newSnail
	} else {
		newSnail.rightnum, _ = strconv.Atoi(rightString)
	}
	return &newSnail
}

func main() {

	input := "test18.txt"

	lines := aocutil.LoadStringArray(input)

	var theSnail *snail
	theSnail = nil

	for _, s := range lines {
		if theSnail == nil {
			// This is the first one, so we just read it
			theSnail = parseSnail(s)
			continue
		}

		// Not the first, so we want to read the new one,
		// Add them together, and reduce

		theSnail.leftsnail = theSnail
		theSnail.leftsnail.parent = theSnail
		theSnail.rightsnail = parseSnail(s)
		theSnail.rightsnail.parent = theSnail

		reduce(theSnail, 0)
	}

}
