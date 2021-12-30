package main

// Game - Representing a single game along with its outcome
type Game struct {
	turnPlayer  Player
	otherPlayer Player

	winner int // A value of 0 indicates the turn player wins; else the other player wins
}

// Match a game to a winner?
func (g Game) match(turnPlayer Player, otherPlayer Player) int {
	if g.turnPlayer.position == turnPlayer.position && g.turnPlayer.totalScore == turnPlayer.totalScore &&
		g.otherPlayer.position == otherPlayer.position && g.otherPlayer.totalScore == otherPlayer.totalScore {
		return g.winner
	}

	return -1
}
