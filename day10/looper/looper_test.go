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
	l := findLoopExcludingStart(field)
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
	l := findLoopExcludingStart(field)
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
	l := findLoopExcludingStart(field)
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

func TestNrOfEnclosedPoints(t *testing.T) {
	field, err := field.Init([]string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"..........."})
	if err != nil {
		t.Errorf(err.Error())
	}
	l := NrOfInteriorPoints(field)
	if l != 4 {
		t.Errorf("Interior points counting failed")
	}
}

func TestNrOfEnclosedPointsNoCorridor(t *testing.T) {
	field, err := field.Init([]string{
		"..........",
		".S------7.",
		".|F----7|.",
		".||....||.",
		".||....||.",
		".|L-7F-J|.",
		".|..||..|.",
		".L--JL--J.",
		".........."})

	if err != nil {
		t.Errorf(err.Error())
	}
	l := NrOfInteriorPoints(field)
	if l != 4 {
		t.Errorf("Interior points counting failed")
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

func TestFloodFillInsiderSimple(t *testing.T) {
	field, err := field.Init([]string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		"....."})
	if err != nil {
		t.Errorf(err.Error())
	}
	res := flodFill(field, point.InitVecList([]point.Vec{point.Init(2, 2)}))
	expected := 1
	if expected != res {
		t.Errorf("Undersök FlodFill i det enklaste caset")
	}
}

func TestFloodFillOutsideSimple(t *testing.T) {
	field, err := field.Init([]string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		"....."})
	if err != nil {
		t.Errorf(err.Error())
	}
	boarder := findBoundaryOfLoop(field)
	startingPoints := findStartingPoints(boarder)
	res := startingPoints.Len()
	expected := 1
	if expected != res {
		t.Errorf("Undersök find startingPoints")
	}
}

func TestFindStartingPointsSimple(t *testing.T) {
	field, err := field.Init([]string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		"....."})
	if err != nil {
		t.Errorf(err.Error())
	}
	res := flodFill(field, point.InitVecList([]point.Vec{point.Init(0, 0)}))
	expected := 16
	if expected != res {
		t.Errorf("Undersök FlodFill i det enklaste caset utanför")
	}
}
