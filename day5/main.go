package main

import (
	"day5/maps"
	"day5/parse"
	"fmt"
)

func main() {
	part1()
}

func part1() {
	seeds, triplets := parse.Parse("input_short.txt")
	mappings := []maps.Mapping{}
	for _, trip := range triplets {
		mappings = append(mappings, maps.InitMapping(trip.Second, trip.First, trip.Last))
	}
	for _, seed := range seeds {
		fmt.Println(seed)
		fmt.Println(maps.Atlas(mappings, seed))
	}
}
