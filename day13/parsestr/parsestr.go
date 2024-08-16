package parsestr

import (
	"day13/matrix"
	"strings"
	"unicode/utf8"
)

func ParseFromString(inputString string) matrix.Matrix {
	if inputString == "" {
		return matrix.InitZero(0, 0)
	}
	lines := strings.Split(inputString, "\n")
	rows := len(lines)
	cols := utf8.RuneCountInString(lines[0])
	elements := strings.Replace(inputString, "\n", "", -1)
	els := []int{}
	for _, r := range elements {
		el := 0
		if r == '#' {
			el = 1
		}
		els = append(els, el)
	}
	return matrix.InitRowOrder(rows, cols, els)
}
