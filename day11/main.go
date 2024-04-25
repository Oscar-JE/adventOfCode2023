package main

import (
	"bufio"
	"day11/linearalgebra"
	"day11/universe"
	"fmt"
	"os"
)

func main() {
	positions := parseInput()
	u := universe.Init(positions)
	fmt.Println(u.SillyExpandAndSumDistance(1000000))
}

func parseInput() []linearalgebra.Vec {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("file not found")
	}
	scanner := bufio.NewScanner(file)
	rowNr := 0
	positions := []linearalgebra.Vec{}
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		for j, r := range runes {
			if r == '#' {
				positions = append(positions, linearalgebra.InitVec(rowNr, j))
			}
		}
		rowNr++
	}
	return positions
}
