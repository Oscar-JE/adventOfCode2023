package main

import (
	"day6/parse"
	"day6/solution"
	"fmt"
)

func main() {
	part2()
}

func part1() {
	path := "input.txt"
	races := parse.Parse(path)
	mult := 1
	for _, race := range races {
		wins := solution.NrOfWins(race)
		mult *= wins
	}
	fmt.Println(mult)
}

func part2() {
	path := "input.txt"
	race := parse.Parse2(path)
	fmt.Println(solution.NrOfWins(race))
}
