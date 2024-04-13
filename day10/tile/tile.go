package tile

import "day10/point"

type Tile interface {
	Ways() point.VecSet
	GetPos() point.Vec
}

func Connected(t1 Tile, t2 Tile) bool {
	return t2.Ways().Has(t1.GetPos()) && t1.Ways().Has(t2.GetPos())
}

func Init(r rune, p point.Vec) Tile {
	if r == '.' {
		return empty{}
	}
	if r == '|' {
		return straigth{position: p, standing: true}
	}
	if r == '-' {
		return straigth{position: p, standing: false}
	}
	if r == 'L' {
		return NortEastBend(p)
	}
	if r == 'J' {
		return NortWestBend(p)
	}
	if r == '7' {
		return SoutWestBend(p)
	}
	if r == 'F' {
		return EastSouthBend(p)
	}

	return empty{}
}
