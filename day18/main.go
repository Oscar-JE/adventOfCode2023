package main

import (
	"day18/excavation"
	"day18/hole"
	"day18/instruction"
	"fmt"
	"os"
)

func main() {
	part2()
}

func part1() {
	fmt.Println(part1Internal("long.txt"))
}

func part1Internal(fileName string) int {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic("file reading failed")
	}
	rep := string(content)
	instructions := instruction.ParseInstructions(rep, "\r\n")
	hole := hole.Init(excavation.ExcavateInSegments(instructions))
	return hole.NrExcavated()
}

func part2() {
	content, err := os.ReadFile("short.txt")
	if err != nil {
		panic("file reading failed")
	}
	rep := string(content)
	instructions := instruction.ParseInstructions2(rep, "\r\n")
	hole := hole.Init(excavation.ExcavateInSegments(instructions))
	fmt.Println(hole)
	fmt.Println(hole.NrExcavated())
}
