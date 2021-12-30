package main

import "fmt"

func main() {

	// No puzzle input needed here yet - we know the positions
	//input := "input20.txt"

	//lines := aocutil.LoadStringArray(input)

	// First line is the enhancement algorithm

	p1 := Player{position: 4, totalScore: 0, playerNum: 0}
	p2 := Player{position: 2, totalScore: 0, playerNum: 1}

	play(p1, p2)

	fmt.Printf("Total Wins for Player 1: %d, Player 2: %d\n", totalWins[0], totalWins[1])

}
