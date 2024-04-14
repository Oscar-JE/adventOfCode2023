package looper

import (
	"day10/field"
	"day10/point"
)

func FindLoop(f field.Field) int {
	loopLen := 0
	start := f.FindStartingPosition()
	possibleNextpossitions := FindConnected(f, start)
	previsousPosition := start
	currentPOssition := possibleNextpossitions.GetVectors()[0]
	for currentPOssition != start {
		loopLen++
		//fmt.Println(currentPOssition)
		possibleNextpossitions = FindConnected(f, currentPOssition)
		for _, el := range possibleNextpossitions.GetVectors() {
			if el != previsousPosition {
				previsousPosition = currentPOssition
				currentPOssition = el
				break
			}
		}
	}
	return loopLen
}

func furthestAway(f field.Field) int {
	loopLen := FindLoop(f) // stämmer inte måste tänka på jämn och ojämn
	return loopLen / 2
}

func FindConnected(f field.Field, p point.Vec) point.VecSet {
	tile := f.GetByVec(p)
	possibleNextpossitions := []point.Vec{}
	possitionCandidate := tile.Ways()
	for _, pos := range possitionCandidate.GetVectors() {
		loopTile := f.GetByVec(pos)
		if loopTile.Ways().Has(p) {
			possibleNextpossitions = append(possibleNextpossitions, pos)
		}
	}
	return point.InitVecSet(possibleNextpossitions)
}
