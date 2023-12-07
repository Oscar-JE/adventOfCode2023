package main

import (
	"day3/parse"
	"day3/schematic"
	"fmt"
)

func main() {
	lines := parse.ReadInput("input.txt")
	schem := schematic.Init(lines)
	fmt.Println(schem)
	fmt.Println(schem.Score2())
}
