package main

import (
	"day14/dish"
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("short.txt")
	if err != nil {
		panic("file not found")
	}
	text := string(content)
	fmt.Println(partTwo(text))
}

func partOne(stringRep string) int {
	d := dish.InitFromText(stringRep)
	for i := range 20 {
		fmt.Printf("after %d cycle\n", i)
		fmt.Println(d)
		fmt.Println(d.Score())
		d.Cycle()
	}

	return 0
}

func partTwo(stringRep string) int {
	d := dish.InitFromText(stringRep)
	return d.Calculate2()
}
