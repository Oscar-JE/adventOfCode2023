package main

import (
	"day2/game"
	"day2/parse"
	"fmt"
)

func main() {
	part2()
}

func part2() {
	lines := parse.ReadInput("input.txt")
	sum := 0
	for _, line := range lines {
		_, showed := game.ParseShowing(line)
		minimum := game.FindMinimum(showed)
		power := minimum.Power()
		sum += power
	}
	fmt.Println(sum)
}

func part1() {
	lines := parse.ReadInput("input.txt")
	restriction := game.Set{Red: 12, Green: 13, Blue: 14}
	sum := 0
	g := game.Init(restriction)
	for _, line := range lines {
		id, showed := game.ParseShowing(line)
		g.SetIdentity(id)
		g.SetPlayed(showed)
		sum += g.Score()
	}
	fmt.Println(sum)
}
