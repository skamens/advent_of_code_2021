package main

import (
	"aoc21/aocinput"
	"fmt"
	"strconv"
	"strings"
)

func printBoard(board [5][5]int) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%-3d", val)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// Mark the board. If we found a match, then also return true; else return false
func markBoard(board *[5][5]int, num int) bool {
	printBoard(*board)

	for i, row := range board {
		for j, val := range row {
			if val == num {
				board[i][j] = -1
				printBoard(*board)
				return true
			}
		}
	}

	printBoard(*board)
	return false
}

// Check for a bingo. If it's bingo. also return
// the total score
func isBingo(board [5][5]int) (bool, int) {

	var bingo bool
	for i := 0; i < len(board); i++ {
		bingo = true

		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != -1 {
				bingo = false
				break
			}
		}

		if bingo {
			break
		}
	}

	if !bingo {

		// Didn't find a bingo going across, so now we'll try going down
		for j := 0; j < len(board); j++ {
			bingo = true

			for i := 0; i < len(board); i++ {
				if board[i][j] != -1 {
					bingo = false
					break
				}
			}
			if bingo {
				break
			}
		}
	}

	total := 0
	if bingo {
		// Add up the non-bingo squares
		for _, row := range board {
			for _, val := range row {
				if val != -1 {
					total += val
				}
			}
		}
	}

	return bingo, total
}

func getBoard(vals []string, index int) ([5][5]int, int) {
	// Skip the first line
	index++
	var b [5][5]int
	for i := 0; i < 5; i++ {
		ints := strings.Fields(vals[index])
		for j, v := range ints {
			b[i][j], _ = strconv.Atoi(v)
		}
		index++
	}
	return b, index
}

func main() {

	var boards [][5][5]int
	var bingos []bool
	numBingos := 0

	input := "input4.txt"

	vals := aocinput.LoadStringArray(input)

	idx := 1
	var b [5][5]int
	for idx < len(vals) {
		b, idx = getBoard(vals, idx)
		boards = append(boards, b)
		bingos = append(bingos, false)
	}

	numberStrings := strings.Split(vals[0], ",")
	var numberList []int
	for _, s := range numberStrings {
		n, _ := strconv.Atoi(s)
		numberList = append(numberList, n)
	}

	for _, n := range numberList {
		for i := 0; i < len(boards); i++ {
			marked := markBoard(&boards[i], n)
			if marked && !bingos[i] {
				bingo, total := isBingo(boards[i])
				if bingo {
					fmt.Printf("BINGO! boardnum: %d, total: %d, finalnum: %d, product: %d\n", i, total, n, n*total)
					bingos[i] = true
					numBingos++
					if numBingos == len(boards) {
						fmt.Printf("This is the last bingo!\n")
						return
					}
					// Remove this board
				}
			}
		}
	}
}
