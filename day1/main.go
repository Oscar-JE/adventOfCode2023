package main

import (
	"day1/logic"
	"day1/parse"
	"fmt"
)

func main() {
	//fmt.Println("part one")
	part1()
	fmt.Println("part two")
	part2()
}

func part1() {
	filePath := "input.txt"
	lines := parse.ReadInput(filePath)
	sum := 0
	for _, line := range lines {
		sum += logic.ValuFromDigits(line)
	}
	//fmt.Println(sum)
}

func part2() {
	filepath := "input.txt"
	lines := parse.ReadInput(filepath)
	sum := 0
	for _, line := range lines {
		sum += logic.ValueFromDigitAndWords(line)
	}
	fmt.Println(sum)
}
