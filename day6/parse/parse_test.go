package parse

import (
	"fmt"
	"testing"
)

func TestParseTimes(t *testing.T) {
	line := "Time:      7  15   30"
	times := parseLine(line)
	fmt.Println(times)
}

func TestParse(t *testing.T) {
	path := "../input_short.txt"
	timesAndDistances := Parse(path)
	fmt.Println(timesAndDistances)
}
