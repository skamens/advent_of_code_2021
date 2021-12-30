package main

import "fmt"

var games []Game

var totalWins [2]int64

func play(turnPlayer Player, otherPlayer Player) {

	var winner int

	var done bool
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {

				done = false

				// See if the current combo matches a game in our cache
				for _, g := range games {

					winner = g.match(turnPlayer, otherPlayer)

					if winner != -1 {
						switch winner {
						case 0:
							totalWins[turnPlayer.playerNum]++
						case 1:
							totalWins[otherPlayer.playerNum]++
						}

						done = true
						break
					}
				}

				if done {
					continue
				}

				// Now see if after taking their turn, the turn player will win.

				newTurnPlayer := turnPlayer
				newTurnPlayer.position = (turnPlayer.position+d1+d2+d3-1)%10 + 1
				newTurnPlayer.totalScore += newTurnPlayer.position

				if newTurnPlayer.totalScore >= 21 {
					// This ends the game with the turnplayer as the winner
					// So that means we want to add the values from before
					// they take their turn to the cache

					totalWins[turnPlayer.playerNum]++
					g := Game{turnPlayer: turnPlayer,
						otherPlayer: otherPlayer,
						winner:      0} // Indicates that the player whose turn it is will win
					games = append(games, g)
					continue
				}

				// Now, we are going to move to the next turn, but first let's see if
				// the opposite combo will win
				done = false
				for _, g := range games {

					winner = g.match(otherPlayer, newTurnPlayer)

					if winner != -1 {
						switch winner {
						case 0:
							totalWins[otherPlayer.playerNum]++
						case 1:
							totalWins[newTurnPlayer.playerNum]++
						}
						totalWins[winner]++

						// And add the earlier value with the opposite winner

						g := Game{turnPlayer: turnPlayer,
							otherPlayer: otherPlayer,
							winner:      1} // Indicates that in this case the "other" plaer will win

						games = append(games, g)
						done = true
						break
					}
				}

				if done {
					continue
				}

				// It doesn't end the game, so essentially start a new game with a new
				// setup.

				newOtherPlayer := otherPlayer

				play(newOtherPlayer, newTurnPlayer)
			}
		}
	}
}

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
