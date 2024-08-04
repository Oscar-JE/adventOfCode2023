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
		sum += reflectionScore2(m)
	}
	fmt.Println(sum)
}

func reflectionCol(m matrix.Matrix) int {
	if m.GetNrRows() == 0 {
		return 0
	}
	firstRow := m.GetRow(0)
	commonReflec := reflection.FindPossibleRef(firstRow)
	for i := 1; i < m.GetNrRows(); i++ {
		nextRow := m.GetRow(i)
		commonReflec.Align(nextRow)
		if !commonReflec.HasReflection() {
			return 0
		}
	}
	return commonReflec.GetIndex() + 1
}

func reflectionCol2(m matrix.Matrix) int {
	if m.GetNrRows() == 0 {
		return 0
	}
	firstRow := m.GetRow(0)
	nrColumns := len(firstRow)
	count := reflection.InitIndexCount(nrColumns)
	for i := 0; i < m.GetNrRows(); i++ {
		nextRow := m.GetRow(i)
		count.Align(nextRow)
	}
	index, err := count.IndexWithNrMatches(m.GetNrRows() - 1)
	if err != nil {
		return 0
	}
	for i := 0; i < m.GetNrRows(); i++ {
		values := m.GetRow(i)
		if !reflection.IsReflektionIndex(index, values) {
			// Är på kandidatrad för ny reflektion
			if reflection.IsReflectionWithOneException(index, values) {
				return index + 1
			}
		}
	}
	return 0
}

func reflectionScore(m matrix.Matrix) int {
	nrColsToTheLeft := reflectionCol(m)
	m.Transpose()
	nrRowsAbove := reflectionCol(m)
	return nrColsToTheLeft + 100*nrRowsAbove
}

func reflectionScore2(m matrix.Matrix) int {
	nrColsToTheLeft := reflectionCol2(m)
	m.Transpose()
	nrRowsAbove := reflectionCol2(m)
	return nrColsToTheLeft + 100*nrRowsAbove
}
