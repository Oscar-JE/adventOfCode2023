package main

import (
	"day5/maps"
	"day5/parse"
	"fmt"
)

func main() {
	fmt.Println(part2())
}

func part1() int {
	seeds, tripletsGs := parse.Parse("input.txt")
	mappings := InitMappings(tripletsGs)
	min := maps.Atlas(mappings, seeds[0])

	for i := 1; i < len(seeds); i++ {
		location := maps.Atlas(mappings, seeds[i])
		min = minimum(min, location)
	}
	return min
}

func part2() int {
	seedNumbers, tripletsGs := parse.Parse("input.txt")
	set := createSet(seedNumbers)
	mappings := InitMappings(tripletsGs)
	for i := 0; i < 100000000; i++ {
		seedAlternative := maps.AtlasInv(mappings, i)
		if set.Contains(seedAlternative) {
			return i
		}
	}
	fmt.Println(seedNumbers)
	fmt.Println(set)
	return 5
}

func createSet(seedNr []int) maps.Set {
	intervals := []maps.Interval{}
	for i := 0; i < len(seedNr)-1; i += 2 {
		intervals = append(intervals, maps.InitInterval(seedNr[i], seedNr[i+1]))
	}
	return maps.InitSeedSet(intervals)
}

func InitMappings(groups []parse.TripletGroup) []maps.Transform {
	mappings := []maps.Transform{}
	for _, tripG := range groups {
		subTrans := []maps.SubTransform{}
		for _, trip := range tripG.Group {
			subTrans = append(subTrans, maps.InitSubTransform(trip.Second, trip.First, trip.Last))
		}
		mappings = append(mappings, maps.InitTransform(subTrans))

	}
	return mappings
}

func minimum(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
