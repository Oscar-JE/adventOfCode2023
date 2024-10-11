package instruction

import (
	"day16/vec"
	"day18/hex"
	"strconv"
	"strings"
)

type Instruction struct {
	distance  int
	direction vec.Vec2d
}

func Init(distance int, direction vec.Vec2d) Instruction {
	return Instruction{distance: distance, direction: direction}
}

func (i Instruction) GetDistance() int {
	return i.distance
}

func (i Instruction) GetDirection() vec.Vec2d {
	return i.direction
}

func parse(rep string) Instruction {
	parts := strings.Split(rep, " ")
	directionCoding := parts[0]
	distRep := parts[1]
	dist, err := strconv.Atoi(distRep)
	if err != nil {
		panic("distance parsing fauled")
	}
	return Instruction{distance: dist, direction: direction(directionCoding)}
}

func parse2(rep string) Instruction {
	parts := strings.Split(rep, " ")
	colorCode := parts[2]
	nummeric := cleanUp(colorCode)
	symbols := []rune(nummeric)
	directionRep := symbols[len(symbols)-1]
	direction := directionMapping2(directionRep)
	hexSymbols := symbols[:len(symbols)-1]
	hexStr := string(hexSymbols)
	nrSteps := hex.Parse(hexStr)
	return Instruction{direction: direction, distance: nrSteps}
}

func direction(directionCoding string) vec.Vec2d {
	r := rune(directionCoding[0])
	return directionMapping(r)
}

func directionMapping(directionEncodeing rune) vec.Vec2d {
	if directionEncodeing == 'U' {
		return vec.Init(-1, 0)
	}
	if directionEncodeing == 'R' {
		return vec.Init(0, 1)
	}
	if directionEncodeing == 'D' {
		return vec.Init(1, 0)
	}
	if directionEncodeing == 'L' {
		return vec.Init(0, -1)
	}
	panic("parsing of direction failed")
}

func directionMapping2(directionEncodeing rune) vec.Vec2d {
	if directionEncodeing == '3' {
		return vec.Init(-1, 0)
	}
	if directionEncodeing == '0' {
		return vec.Init(0, 1)
	}
	if directionEncodeing == '1' {
		return vec.Init(1, 0)
	}
	if directionEncodeing == '2' {
		return vec.Init(0, -1)
	}
	panic("parsing of direction failed")
}

func cleanUp(colorCode string) string {
	ret := strings.Replace(colorCode, "(", "", -1)
	ret = strings.Replace(ret, ")", "", -1)
	ret = strings.Replace(ret, "#", "", -1)
	return ret
}
