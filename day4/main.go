package main

import (
	"day4/parse"
	"day4/ticket"
	"fmt"
)

func main() {
	fmt.Println(part2("input.txt"))
}

func part1(path string) int {
	lines := parse.ReadInput(path)
	sum := 0
	for _, line := range lines {
		winning, holding := ticket.Parse(line)
		matches := ticket.NrOfMatches(winning, holding)
		score := ticket.Winnings(matches)
		sum += score
	}
	return sum
}

func part2(path string) int {
	lines := parse.ReadInput(path)
	tally := ticket.ParseTally(lines)
	tally.Multiply()
	return tally.SumOfTally()
}
