package main

import (
	"bufio"
	"day12/parse"
	"day12/partone"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("file not found ")
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		f, o := parse.Line(scanner.Text())
		sum += partone.NumberOfCombination(f, o)
	}
	fmt.Println(sum)
}
