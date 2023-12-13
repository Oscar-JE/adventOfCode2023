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
	seeds, tripletsGs := parse.Parse("input_short.txt")
	mappings := []maps.Transform{}
	for _, tripG := range tripletsGs {
		subTrans := []maps.SubTransform{}
		for _, trip := range tripG.Group {
			subTrans = append(subTrans, maps.InitSubTransform(trip.Second, trip.First, trip.Last))
		}
		mappings = append(mappings, maps.InitTransform(subTrans))

	}
	for _, seed := range seeds {
		fmt.Println(seed)
		fmt.Println(maps.Atlas(mappings, seed))
	}
}
