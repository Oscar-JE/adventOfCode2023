package main

import (
	"bufio"
	"day12/field"
	"day12/order"
	"day12/parse"
	"day12/parttwo"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("file not found ")
	}
	scanner := bufio.NewScanner(file)
	inputs := []FieldAndOrder{}
	for scanner.Scan() {
		f, o := parse.Line(scanner.Text())
		inputs = append(inputs, FieldAndOrder{f, o})
	}
	start := time.Now()
	sum := 0
	linenumber := 1
	for _, input := range inputs {
		fmt.Println("workin on line number d%", linenumber)
		sum += parttwo.NumberOfCombination(input.F, input.O)
		linenumber++
	}
	fmt.Println(sum)
	fmt.Println(time.Since(start))
}

type FieldAndOrder struct {
	F field.Field
	O order.Order
}
