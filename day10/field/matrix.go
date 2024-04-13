package field

import (
	"day10/point"
	"day10/tile"
	"errors"
)

type Field struct {
	rows    []string
	columns int
}

func Init(rows []string) (Field, error) {
	if len(rows) == 0 {
		return Field{}, errors.New("empty field")
	}
	length := len(rows[0])
	for i := 1; i < len(rows); i++ {
		if len(rows[i]) != length {
			return Field{}, errors.New("inconsisten row length in Init input")
		}
	}
	return Field{rows: rows, columns: length}, nil
}

func (f Field) GetByVec(v point.Vec) tile.Tile {
	return f.GetByInt(v.GetX(), v.GetY())
}

func (f Field) GetByInt(x int, y int) tile.Tile {
	p := point.Init(x, y)
	s := f.rows[x]
	return tile.Init(runeAt(s, y), p)
}

func runeAt(s string, index int) rune {
	return ([]rune(s))[index]
}
