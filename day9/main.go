package main

import (
	"bufio"
	"day9/extrapolation"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	fReader, err := os.Open("input_short.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fReader)
	sum := 0
	for scanner.Scan() {
		series := parse(scanner.Text())
		sum += extrapolation.Extrapolate(series)
	}
	fmt.Println("Sum of all extrapolations is %d", sum)
}

func parse(inputString string) []int {
	strNumbers := strings.Split(inputString, " ")
	series := []int{}
	for _, strNumber := range strNumbers {
		nr, err := strconv.Atoi(strNumber)
		if err != nil {
			panic(err)
		}
		series = append(series, nr)
	}
	return series
}
