package field

import (
	"day10/point"
	"day10/tile"
	"errors"
)

type Field struct {
	runes   [][]rune
	rows    int
	columns int
}

func (f Field) GetRows() int {
	return f.rows
}

func (f Field) GetCols() int {
	return f.columns
}

func Init(rows []string) (Field, error) {
	runes := [][]rune{}
	nrRows := len(rows)
	if nrRows == 0 {
		return Field{}, errors.New("empty field")
	}
	nrCols := len([]rune(rows[0]))
	for i := range rows {
		runesInrow := []rune(rows[i])
		runes = append(runes, runesInrow)
	}

	return Field{runes: runes, rows: nrRows, columns: nrCols}, nil
}

func (f Field) GetByVec(v point.Vec) tile.Tile {
	return f.GetByInt(v.GetX(), v.GetY())
}

func (f Field) GetByInt(x int, y int) tile.Tile {
	p := point.Init(x, y)
	return tile.Init(f.runes[x][y], p)
}

func (f Field) FindStartingPosition() point.Vec {
	for i := range f.runes {
		for j := range f.columns {
			if f.runes[i][j] == 'S' {
				return point.Init(i, j)
			}
		}
	}
	return point.Init(0, 0)
}

func (f Field) Outside(vec point.Vec) bool {
	return vec.GetX() < 0 || vec.GetX() >= f.rows || vec.GetY() < 0 || vec.GetY() >= f.columns
}

func (f Field) String() string {
	retString := ""
	for i := range f.rows {
		for j := range f.columns {
			retString += string(f.runes[i][j])
		}
		retString += "\n"
	}
	return retString
}

func (f *Field) RemoveEverythingExcept(vecs point.VecList) {
	for i := range f.rows {
		for j := range f.columns {
			if !vecs.Has(point.Init(i, j)) {
				f.Set(i, j, '.')
			}
		}
	}
}

func (f *Field) Set(row int, col int, r rune) {
	f.runes[row][col] = r
}

func (f *Field) FillTheseWith(vecs point.VecList, r rune) {
	for _, vec := range vecs.GetList() {
		f.Set(vec.GetX(), vec.GetY(), r)
	}
}
