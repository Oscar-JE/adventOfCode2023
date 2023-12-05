package game

import "testing"

func TestParseIden(t *testing.T) {
	id := parseIdentiy("Game 2")
	if id != 2 {
		t.Errorf("identity failed parse")
	}
}

func TestParseColor(t *testing.T) {
	rgb := []string{"3 blue", "4 red"}
	blue := parseColor(rgb, "blue")
	green := parseColor(rgb, "green")
	red := parseColor(rgb, "red")
	if red != 4 || green != 0 || blue != 3 {
		t.Errorf("wrong in color parse")
	}
}

func TestParseShowing(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green"
	id, sets := ParseShowing(input)
	if id != 1 || sets[0].Blue != 3 || sets[1].Red != 1 {
		t.Errorf("oh lordylordy")
	}
}
