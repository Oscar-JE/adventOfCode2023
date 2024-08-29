package main

import (
	"day16/beam"
	"day16/matrix"
	"day16/mirror"
	"day16/mirrors"
	"day16/vec"
	"fmt"
	"os"
	"strings"
)

func main() {
	partTwo()
}

func partTwo() {
	mat := parse("long.txt")
	maximum := 0
	for i := 0; i < mat.GetRows(); i++ {
		starterBeam := beam.Init(vec.Init(i, -1), vec.Init(0, 1))
		res := partOneInternal(mat, starterBeam)
		if res > maximum {
			maximum = res
		}
	}
	for i := 0; i < mat.GetRows(); i++ {
		starterBeam := beam.Init(vec.Init(i, mat.GetCols()), vec.Init(0, -1))
		res := partOneInternal(mat, starterBeam)
		if res > maximum {
			maximum = res
		}
	}
	for j := 0; j < mat.GetCols(); j++ {
		starterBeam := beam.Init(vec.Init(-1, j), vec.Init(1, 0))
		res := partOneInternal(mat, starterBeam)
		if res > maximum {
			maximum = res
		}
	}
	for j := 0; j < mat.GetCols(); j++ {
		starterBeam := beam.Init(vec.Init(mat.GetRows(), j), vec.Init(-1, 0))
		res := partOneInternal(mat, starterBeam)
		if res > maximum {
			maximum = res
		}
	}
	fmt.Println(maximum)
}

func partOne() {
	mat := parse("long.txt")
	startBeam := beam.Init(vec.Init(0, -1), vec.Init(0, 1))
	fmt.Println(partOneInternal(mat, startBeam))
}

func partOneInternal(mat matrix.Matrix[mirrors.Reflector], startBeam beam.Beam) int {
	setting := mirrors.Init(mat, startBeam)
	return setting.NrOfEnergised()
}

func parse(file string) matrix.Matrix[mirrors.Reflector] {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic("file was not read correctly")
	}
	content := string(bytes)
	lines := strings.Split(content, "\r\n")
	rows := len(lines)
	if len(lines) == 0 {
		panic("Was content realy discovered?")
	}
	cols := len([]rune(lines[0]))
	reflectors := []mirrors.Reflector{}
	contentNoLines := strings.Replace(content, "\r\n", "", -1)
	for _, r := range contentNoLines {
		reflectors = append(reflectors, mirror.Init(r))
	}
	return matrix.Init(reflectors, rows, cols)
}
