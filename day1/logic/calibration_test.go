package logic

import "testing"

func TestDigitFromRight1(t *testing.T) {
	input := "1abc2"
	if digitFromRight(input) != 2 {
		t.Errorf("didgit from right failed")
	}
}

func TestDigitFromLeft1(t *testing.T) {
	input := "1abc2"
	res := diditFromLeft(input)
	if res != 1 {
		t.Errorf("didgit from left failed, should be %d : was %d", 1, res)
	}
}

func TestDigitFromWordRight1(t *testing.T) {
	input := "two1nine"
	res := diditFromWordRight(input)
	if res != 9 {
		t.Errorf("didgit from left right, should be %d : was %d", 9, res)
	}
}

func TestDigitFromWordLeft1(t *testing.T) {
	input := "two1nine"
	res := diditFromWordLeft(input)
	if res != 2 {
		t.Errorf("didgit from left failed, should be %d : was %d", 2, res)
	}
}
