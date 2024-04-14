package looper

import (
	"day10/field"
	"day10/point"
	"testing"
)

func TestSimpleLoop(t *testing.T) {
	field, err := field.Init([]string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		"....."})
	if err != nil {
		t.Errorf(err.Error())
	}
	l := findLoop(field)
	if l != 7 {
		t.Errorf("loopLength failed in simple case")
	}
}

func TestFurthest(t *testing.T) {
	field, err := field.Init([]string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		"....."})
	if err != nil {
		t.Errorf(err.Error())
	}
	l := NrOfSrepsToFurthestPoint(field)
	if l != 4 {
		t.Errorf("loopLength failed in simple case")
	}
}

func TestSimpleLoop2(t *testing.T) {
	field, err := field.Init([]string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF"})
	if err != nil {
		t.Errorf(err.Error())
	}
	l := findLoop(field)
	if l != 7 {
		t.Errorf("loopLength failed in simple case")
	}
}

func TestComplexLoop(t *testing.T) {
	field, err := field.Init([]string{
		"7-F7-",
		".FJ|7",
		"SJLL7",
		"|F--J",
		"LJ.LJ"})
	if err != nil {
		t.Errorf(err.Error())
	}
	l := findLoop(field)
	if l != 15 {
		t.Errorf("loopLength failed in simple case")
	}
}

func TestComplexLoopNrOfSteps(t *testing.T) {
	field, err := field.Init([]string{
		"7-F7-",
		".FJ|7",
		"SJLL7",
		"|F--J",
		"LJ.LJ"})
	if err != nil {
		t.Errorf(err.Error())
	}
	l := NrOfSrepsToFurthestPoint(field)
	if l != 8 {
		t.Errorf("loopLength failed in simple case")
	}
}

func TestFindConnectedSimpleLoopAtStart(t *testing.T) {
	field, err := field.Init([]string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		"....."})
	if err != nil {
		t.Errorf(err.Error())
	}
	res := findConnected(field, point.Init(1, 1))
	expected := point.InitVecSet([]point.Vec{point.Init(1, 2), point.Init(2, 1)})
	if !expected.Eq(res) {
		t.Errorf("Undersök FindConnected Simple loop vid start")
	}
}

func TestFindConnectedSimpleLoopAtLink(t *testing.T) {
	field, err := field.Init([]string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		"....."})
	if err != nil {
		t.Errorf(err.Error())
	}
	res := findConnected(field, point.Init(2, 1))
	expected := point.InitVecSet([]point.Vec{point.Init(1, 1), point.Init(3, 1)})
	if !expected.Eq(res) {
		t.Errorf("Undersök FindConnected Simple loop vid länk")
	}
}
