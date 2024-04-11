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
	part2()
}

func part1() {
	fReader, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fReader)
	sum := 0
	for scanner.Scan() {
		series := parse(scanner.Text())
		sum += extrapolation.Extrapolate(series)
	}
	fmt.Printf("Sum of all extrapolations is %d", sum)
}

func part2() {
	fReader, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fReader)
	sum := 0
	for scanner.Scan() {
		series := parse(scanner.Text())
		sum += extrapolation.Backwards(series)
	}
	fmt.Printf("Sum of all extrapolations is %d", sum)
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
