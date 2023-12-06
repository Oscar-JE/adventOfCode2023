package main

import (
	"day6/parse"
	"fmt"
)

func main() {
	part1()
}

func part1() {
	path := "input_short.txt"
	races := parse.Parse(path)
	fmt.Println(races)
}
