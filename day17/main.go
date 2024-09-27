package main

import (
	"day16/vec"
	"day17/dijkstras"
	"day17/environment"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part2())
}

func part1() int {
	env, rows, cols := parseEnv("big.txt")
	startState := environment.StartState()
	cash := environment.InitStateCach(startState, rows, cols)
	//cash := environment.InitSimpleCach()
	dijk := dijkstras.Init[environment.State](env, &cash)
	winningStates := environment.WinningStates(vec.Init(rows-1, cols-1))
	distance := dijk.SmallestDist(winningStates, startState, 0)
	//visualizeShortestDist(rows, cols, cash)
	return distance
}

func part2() int {
	env, rows, cols := parseEnv("simpler.txt")
	startState := environment.StartState2()
	cash := environment.InitSimpleCach()
	//cash := environment.InitStateCach(startState, rows, cols)
	dijk := dijkstras.Init[environment.State](env, &cash)
	winningStates := environment.WinningStates2(vec.Init(rows-1, cols-1))
	distance := dijk.SmallestDist(winningStates, startState, 0)
	visualizeShortestDist(rows, cols, cash)
	return distance
}

func parseEnv(filename string) (environment.Environment, int, int) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic("input file was not found or was corrupted")
	}
	content := string(file)
	lines := strings.Split(content, "\r\n")
	rows := len(lines)
	cols := len([]rune(lines[0]))
	data := []int{}
	contentNoLines := strings.Replace(content, "\r\n", "", -1)
	for _, r := range contentNoLines {
		value, err := strconv.Atoi(string(r))
		if err != nil {
			panic("pars of costs failed")
		}
		data = append(data, value)
	}
	env := environment.InitEnv(data, rows, cols)
	return env, rows, cols
}

func visualizeShortestDist(rows int, cols int, cach environment.SimpleCach) {
	for i := 0; i < rows; i++ {
		rowRep := ""
		for j := 0; j < cols; j++ {
			cost := cach.FindLowestWithPosition(vec.Init(i, j))
			rowRep += CostRep(cost)
		}
		fmt.Println(rowRep)
	}
}

func CostRep(v int) string {
	separator := " "
	if v >= 100 {
		return separator + "nan"
	}
	if v >= 10 {
		return separator + separator + fmt.Sprint(v)
	}
	return separator + separator + separator + fmt.Sprint(v)
}
