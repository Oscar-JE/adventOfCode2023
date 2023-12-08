package parse

import (
	"fmt"
	"testing"
)

func TestSeedParse(t *testing.T) {
	seedLine := "seeds: 79 14 55 13"
	seeds := parseSeed(seedLine)
	fmt.Println(seeds)
}

func TestThripletParse(t *testing.T) {
	line := "50 98 2"
	trip := parseTriple(line)
	fmt.Println(trip)
}

func TestParse(t *testing.T) {
	path := "../input_short.txt"
	seeds, triplets := Parse(path)
	fmt.Println(seeds)
	fmt.Println(triplets)
}
