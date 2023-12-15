package main

import (
	"bufio"
	"day7/hand"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	plays := parse("input.txt")
	for _, play := range plays {
		fmt.Println(play)
	}
	sort.SliceStable(plays, func(i int, j int) bool {
		return less(plays[i], plays[j])
	})
	println("____ sorted ____")
	for _, play := range plays {
		fmt.Println(play)
	}
	score := 0
	for i, play := range plays {
		score += (i + 1) * play.bid
	}
	fmt.Println("answer:")
	fmt.Println(score)
}

func less(p1 Play, p2 Play) bool {
	return hand.Lower(p1.hand, p2.hand)
}

type Play struct {
	hand hand.Hand
	bid  int
}

func parse(path string) []Play {
	file, fileError := os.Open(path)
	if fileError != nil {
		panic("file not found")
	}
	scanner := bufio.NewScanner(file)
	plays := []Play{}
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, " ")
		handLine := splitted[0]
		bidLine := splitted[1]
		hand := hand.Init(handLine)
		bid, nrError := strconv.Atoi(bidLine)
		if nrError != nil {
			panic("unable to convert bid")
		}
		plays = append(plays, Play{hand: hand, bid: bid})
	}
	return plays
}
