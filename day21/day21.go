package main

var games []Game

var totalWins [2]int64

func play(p1 Player, p2 Player) {

	var winner int

	var done bool
	for d1 := 1; d1 <= 3; d1++ {
		for d2 := 1; d2 <= 3; d2++ {
			for d3 := 1; d3 <= 3; d3++ {

				done = false
				for _, g := range games {

					winner = g.match(p1, p2)

					if winner != -1 {
						totalWins[winner]++
						done = true
						break
					}
				}

				if done {
					continue
				}

				newP1 := p1
				newP1.position = (p1.position+d1+d2+d3-1)%10 + 1
				newP1.totalScore += newP1.position

				if newP1.totalScore >= 21 {
					// This ends the game with p1 as the winner

					totalWins[p1.playerNum]++
					g := Game{pos1: p1.position,
						score1: p1.totalScore,
						pos2:   p2.position,
						score2: p2.totalScore,
						winner: p1.playerNum}
					games = append(games, g)
					continue
				}

				// Let's see if the new arrangement matches
				done = false
				for _, g := range games {

					winner = g.match(p2, newP1)

					if winner != -1 {
						totalWins[winner]++
						// And add the earlier value with the opposite winner

						g := Game{pos1: p1.position,
							score1: p1.totalScore,
							pos2:   p2.position,
							score2: p2.totalScore,
							winner: winner}

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

				newP2 := p2

				play(newP2, newP1)
			}
		}
	}
}

func main() {

	// No puzzle input needed here yet - we know the positions
	//input := "input20.txt"

	//lines := aocutil.LoadStringArray(input)

	// First line is the enhancement algorithm

	p1 := Player{startposition: 4, startscore: 0, position: 4, totalScore: 0, playerNum: 0}
	p2 := Player{startposition: 2, startscore: 0, position: 2, totalScore: 0, playerNum: 1}

	play(p1, p2)

}
