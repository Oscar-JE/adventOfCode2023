package main

import (
	"day13/parsestr"
	"testing"
)

func TestReflectionRow(t *testing.T) {
	m := parsestr.ParseFromString("#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.")
	colRef := reflectionCol(m)
	if colRef != 5 {
		t.Errorf("investigate first example from instructions")
	}
}

func TestReflectionScore1(t *testing.T) {
	m := parsestr.ParseFromString("#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.")
	refScore := reflectionScore(m)
	if refScore != 5 {
		t.Errorf("investigate first example from instructions")
	}
}

func TestReflectionScore2(t *testing.T) {
	m := parsestr.ParseFromString("#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#")
	refScore := reflectionScore(m)
	if refScore != 400 {
		t.Errorf("investigate first example from instructions")
	}
}
