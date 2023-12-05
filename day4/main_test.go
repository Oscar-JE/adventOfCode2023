package main

import "testing"

func TestPart1(t *testing.T) {
	res := part1("input_short.txt")
	expected := 13
	if res != expected {
		t.Errorf("wtf")
	}
}

func TestPart2(t *testing.T) {
	part2("input_short.txt")
}
