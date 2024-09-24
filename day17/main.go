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
	env, rows, cols := parseEnv("small.txt")
	startState := environment.StartState()
	cash := environment.InitStateCach(startState, rows, cols)
	dijk := dijkstras.Init[environment.State](env, &cash)
	winningStates := environment.WinningStates(vec.Init(rows-1, cols-1))
	return dijk.SmallestDist(winningStates, startState, 0)
}

func part2() int {
	env, rows, cols := parseEnv("simpler.txt")
	startState := environment.StartState2()
	cash := environment.InitStateCach(startState, rows, cols)
	dijk := dijkstras.Init[environment.State](env, &cash)
	winningStates := environment.WinningStates2(vec.Init(rows-1, cols-1))
	return dijk.SmallestDist(winningStates, startState, 0)
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
