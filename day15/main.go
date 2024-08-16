package main

import (
	"day15/hashone"
	"fmt"
	"os"
	"strings"
)

func main() {
	sum := 0
	codes := parse("input.txt")
	for _, code := range codes {
		sum += hashone.Hash(code)
	}
	fmt.Println(sum)
}

func parse(fileName string) []string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic("no file found")
	}
	content := string(file)
	return strings.Split(content, ",")
}
