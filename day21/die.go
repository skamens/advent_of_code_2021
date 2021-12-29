package main

// Die - the die in our game
type Die struct {
	size     int
	lastRoll int
	numRolls int
}

func (d *Die) roll() int {
	d.lastRoll = (d.lastRoll + 1) % d.size
	d.numRolls++
	return d.lastRoll
}
