package main

// Game - Representing a single game along with its outcome
type Game struct {
	pos1   int
	score1 int
	pos2   int
	score2 int

	winner int
}

// Match a game to a winner?
func (g Game) match(p1 Player, p2 Player) int {
	if g.pos1 == p1.position && g.score1 == p1.totalScore &&
		g.pos2 == p2.position && g.score2 == p2.totalScore {
		return g.winner
	}

	return -1
}
