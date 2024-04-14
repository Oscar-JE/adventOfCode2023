package tile

import "day10/point"

type Tile struct {
	movements MovementsGen
	position  point.Vec
}

func InitTile(position point.Vec, movements MovementsGen) Tile {
	return Tile{position: position, movements: movements}
}

func (t Tile) Ways() point.VecSet {
	moves := t.movements.moves()
	return moves.Translate(t.position)
}

func (t Tile) GetPos() point.Vec {
	return t.position
}

type MovementsGen interface {
	moves() point.VecSet
}

func Connected(t1 Tile, t2 Tile) bool {
	return t2.Ways().Has(t1.GetPos()) && t1.Ways().Has(t2.GetPos())
}

func Init(r rune, p point.Vec) Tile {
	if r == '.' {
		return InitTile(p, empty{})
	}
	if r == '|' {
		return InitTile(p, straigth{standing: true})
	}
	if r == '-' {
		return InitTile(p, straigth{standing: false})
	}
	if r == 'L' {
		return InitTile(p, NortEastBend())
	}
	if r == 'J' {
		return InitTile(p, NortWestBend())
	}
	if r == '7' {
		return InitTile(p, SoutWestBend())
	}
	if r == 'F' {
		return InitTile(p, EastSouthBend())
	}

	return InitTile(p, empty{})
}
