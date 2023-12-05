package schematic

import (
	"fmt"
	"strconv"
	"strings"
)

type Schematic struct {
	lines []string
}

func Init(lines []string) Schematic {
	for i := range lines {
		lines[i] = "." + lines[i] + "."
	}
	emptyLine := ""
	for i := 0; i < len(lines[0]); i++ {
		emptyLine += "."
	}
	lines = append([]string{emptyLine}, lines...)
	lines = append(lines, emptyLine)
	return Schematic{lines}
}

func (s Schematic) String() string {
	retV := ""
	for _, line := range s.lines {
		retV = retV + fmt.Sprintln(line)
	}
	return retV
}

func (s Schematic) getSymbol(row int, col int) string {
	return string(s.lines[row][col])
}

func (s Schematic) perimiter(row int, col int) string {
	return s.getSymbol(row-1, col) +
		s.getSymbol(row-1, col-1) +
		s.getSymbol(row, col-1) +
		s.getSymbol(row+1, col-1) +
		s.getSymbol(row+1, col) +
		s.getSymbol(row+1, col+1) +
		s.getSymbol(row, col+1) +
		s.getSymbol(row-1, col+1)
}

func isSpecial(symbol string) bool {
	return len(symbol) == 1 && !strings.Contains("1234567890.", symbol)
}

func containsSpecial(symbols string) bool {
	contains := false
	for i := 0; i < len(symbols); i++ {
		contains = contains || isSpecial(string(symbols[i]))
	}
	return contains
}

func (s Schematic) Score() int {
	sum := 0
	for row := 1; row < len(s.lines)-1; row++ {
		cache := 0
		shouldAdd := false
		for col := 1; col < len(s.lines[0])-1; col++ {
			symbol := s.getSymbol(row, col)
			number, err := strconv.Atoi(symbol)
			if err != nil {
				if shouldAdd {
					sum += cache
				}
				cache = 0
				shouldAdd = false
			} else {
				cache = cache*10 + number
				shouldAdd = shouldAdd || containsSpecial(s.perimiter(row, col))
			}
		}
		if shouldAdd {
			sum += cache
		}
	}
	return sum
}

func (s Schematic) numberAtPosition(row int, col int) (NumberAndPosition, error) {
	return
}
