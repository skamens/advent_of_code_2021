package main

import (
	"aoc21/aocinput"
	"fmt"
)

func main() {

	input := "input6.txt"

	vals := aocinput.LoadIntArrayLine(input)

	var counts [6]int

	for _, v := range vals {
		counts[v]++
	}

	var fish [9]int

	fish[1] = 1
	totalDays := 256
	total := 0
	for day := 0; day < totalDays; day++ {
		fishLen := 0
		for _, cnt := range fish {
			fishLen += cnt
		}
		fmt.Printf("Day %d, fishLen: %d\n", day, fishLen)
		fmt.Println(fish)

		newFish := 0
		for i := 0; i <= 8; i++ {
			if i == 0 {
				newFish = fish[i]
			} else {
				fish[i-1] = fish[i]
			}
		}
		fish[8] = newFish
		fish[6] += newFish

		if (totalDays - day) < 6 {
			fishLen := 0
			for _, cnt := range fish {
				fishLen += cnt
			}
			total = total + counts[totalDays-day]*fishLen
			fmt.Printf("Total: %d\n", total)
		}

	}

}
