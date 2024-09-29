package main

import (
	"day18/excavation"
	"day18/hole"
	"day18/instruction"
	"fmt"
	"os"
)

func main() {
	part1()
}

func part1() {
	content, err := os.ReadFile("short.txt")
	if err != nil {
		panic("file reading failed")
	}
	rep := string(content)
	instructions := instruction.ParseInstructions(rep, "\r\n")
	hole := hole.Init(excavation.Excavate(instructions))
	fmt.Println(hole)
}
