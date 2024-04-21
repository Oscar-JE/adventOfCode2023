package main

import (
	"bufio"
	"day10/field"
	"day10/looper"
	"fmt"
	"os"
)

func main() {
	lines := readInput()
	f, err := field.Init(lines)
	if err != nil {
		panic("error creating field")
	}
	fmt.Println(looper.NrOfSrepsToFurthestPoint(f))
	fmt.Println(looper.NrOfInteriorPoints(f))
}

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("file not found")
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
