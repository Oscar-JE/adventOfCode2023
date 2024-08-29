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
	"time"
)

func main() {
	mat := parse("long.txt")
	startBeam := beam.Init(vec.Init(0, -1), vec.Init(0, 1))
	fmt.Println(mat)
	setting := mirrors.Init(mat, startBeam)
	start := time.Now()
	res := setting.NrOfEnergised()
	t := time.Since(start)
	fmt.Println(res)
	fmt.Println(t)
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
