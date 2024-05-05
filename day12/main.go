package main

import (
	"bufio"
	"day12/parse"
	"day12/parttwo"
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
	linenumber := 1
	for scanner.Scan() {
		fmt.Println("workin on line number d%", linenumber)
		linenumber++
		f, o := parse.Line(scanner.Text())
		sum += parttwo.NumberOfCombination(f, o)
	}
	fmt.Println(sum)
}
