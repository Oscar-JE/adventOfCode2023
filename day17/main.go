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
	fmt.Println(part1())
}

func part1() int {
	file, err := os.ReadFile("small.txt")
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
	cash := environment.InitStateCach(rows, cols)
	dijk := dijkstras.Init[environment.State](env, &cash)
	winningStates := environment.WinnigStates(vec.Init(rows-1, cols-1))
	startState := environment.StartState()
	return dijk.SmallestDist(winningStates, startState, 0)
}
