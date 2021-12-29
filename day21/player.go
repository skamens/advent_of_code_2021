package main

// Player - the player
type Player struct {
	startposition int
	startscore    int
	position      int
	totalScore    int

	playerNum int
}

func (p *Player) turn(rolls int) {

	p.position = ((p.position + rolls - 1) % 10) + 1
	p.totalScore += p.position

}
