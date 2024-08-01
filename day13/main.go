package main

import (
	"day13/matrix"
	"day13/parsestr"
	"day13/reflection"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileContent, err := os.ReadFile("inputbig.txt")
	if err != nil {
		panic("no file found with that seekedname")
	}
	text := string(fileContent)
	matrixReps := strings.Split(text, "\n\n")
	matrixs := []matrix.Matrix{}
	for _, rep := range matrixReps {
		matrixs = append(matrixs, parsestr.ParseFromString(rep))
	}
	sum := 0
	for _, m := range matrixs {
		sum += reflectionScore(m)
	}
	fmt.Println(sum)
}

func reflectionCol(m matrix.Matrix) int {
	if m.GetNrRows() == 0 {
		return 0
	}
	firstRow := m.GetRow(0)
	commonReflec := reflection.FindPossibleRef(firstRow)
	for i := 1; i < m.GetNrRows(); i++ { // stor bugg
		nextRow := m.GetRow(i)
		commonReflec.Align(nextRow)
		if !commonReflec.HasReflection() {
			return 0
		}
	}
	return commonReflec.GetIndex() + 1
}

func reflectionScore(m matrix.Matrix) int { // hmmm
	nrColsToTheLeft := reflectionCol(m)
	m.Transpose()
	nrRowsAbove := reflectionCol(m)
	return nrColsToTheLeft + 100*nrRowsAbove
}
