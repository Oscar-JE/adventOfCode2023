package main

import (
	"day14/dish"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("big.txt")
	if err != nil {
		panic("file not found")
	}
	text := string(content)
	fmt.Println(partOne(text))
}

func partOne(stringRep string) int {
	d := dish.InitFromText(stringRep)
	return d.Calculate()
}
