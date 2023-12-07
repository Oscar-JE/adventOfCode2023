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

func boarderOffset(row int, col int) []Position {
	positions := []Position{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}
	offset := Position{row, col}
	for i := range positions {
		positions[i] = positions[i].Add(offset)
	}
	return positions
}

func (s Schematic) perimiter(row int, col int) string {
	positions := boarderOffset(row, col)
	retStr := ""
	for _, pos := range positions {
		retStr += s.getSymbol(pos.row, pos.col)
	}
	return retStr
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

func (s Schematic) perimiter2(row int, col int) (Position, bool) {
	symbol := "."
	positions := boarderOffset(row, col)
	for _, pos := range positions {
		symbol = s.getSymbol(pos.row, pos.col)
		if isSpecial(symbol) {
			return pos, true
		}
	}
	return Position{0, 0}, false

}

type PosAndNumber struct {
	pos    Position
	number int
}

func (s Schematic) Score2() int {
	positionsAndNumbers := []PosAndNumber{}
	for row := 1; row < len(s.lines)-1; row++ {
		cache := 0
		shouldAdd := false
		pos := Position{-1, -1}
		for col := 1; col < len(s.lines[0])-1; col++ {
			symbol := s.getSymbol(row, col)
			number, err := strconv.Atoi(symbol)
			if err != nil {
				if shouldAdd {
					positionsAndNumbers = append(positionsAndNumbers, PosAndNumber{pos: pos, number: cache})
				}
				cache = 0
				shouldAdd = false
			} else {
				cache = cache*10 + number
				pos2, isSpec := s.perimiter2(row, col)
				if isSpec {
					pos = pos2
				}

				shouldAdd = shouldAdd || isSpec
			}
		}
		if shouldAdd {
			positionsAndNumbers = append(positionsAndNumbers, PosAndNumber{pos: pos, number: cache})
		}
	}
	//fmt.Println(positionsAndNumbers)
	multi := 0
	for i := range positionsAndNumbers {
		numberOfmatches := 0
		for j := i + 1; j < len(positionsAndNumbers); j++ {
			if positionsAndNumbers[i].pos == positionsAndNumbers[j].pos {
				multi += positionsAndNumbers[i].number * positionsAndNumbers[j].number
				numberOfmatches++
			}
		}
		if numberOfmatches > 1 {
			fmt.Println("Wtf matches : %d", numberOfmatches)
		}
	}
	return multi
}
