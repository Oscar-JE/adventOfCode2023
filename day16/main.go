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
	mat := parse("short.txt")
	startBeam := beam.Init(vec.Init(0, -1), vec.Init(0, 1))
	fmt.Println(mat.Get(7, 1).Rep())
	setting := mirrors.Init(mat, startBeam)
	fmt.Println(setting.NrOfEnergised())
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
	for _, r := range content {
		reflectors = append(reflectors, mirror.Init(r))
	}
	return matrix.Init(reflectors, rows, cols)
}
