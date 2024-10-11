package main

import "testing"

func TestPart1Short(t *testing.T) {
	res := part1Internal("short.txt")
	if res != 62 {
		t.Errorf("investigate short example")
	}
}
