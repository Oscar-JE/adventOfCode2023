package main

import (
	"day15/hashmap"
	"day15/hashone"
	"fmt"
	"os"
	"strings"
)

func main() {
	instuctions := parse("input.txt")
	myMap := hashmap.Init()
	for _, inst := range instuctions {
		myMap.PerformInstruction(inst)
	}
	fmt.Println(myMap.FocusingPower())
}

func parse(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("no file found")
	}
	content := string(file)
	return strings.Split(content, ",")
}

func partOne() {
	sum := 0
	codes := parse("input.txt")
	for _, code := range codes {
		sum += hashone.Hash(code)
	}
	fmt.Println(sum)
}
