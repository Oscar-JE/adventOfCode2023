package looper

import (
	"day10/field"
	"day10/middle"
	"day10/point"
)

func NrOfSrepsToFurthestPoint(f field.Field) int {
	loopLen := findLoop(f)
	return middle.MiddlePointFromZero(loopLen)
}

func findLoop(f field.Field) int {
	loopLen := 0
	start := f.FindStartingPosition()
	possibleNextpossitions := findConnected(f, start)
	previsousPosition := start
	currentPOssition := possibleNextpossitions.GetVectors()[0]
	for currentPOssition != start {
		loopLen++
		//fmt.Println(currentPOssition)
		possibleNextpossitions = findConnected(f, currentPOssition)
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

func findBoundaryOfLoop(f field.Field) point.VecSet {
	boundary := []point.Vec{}
	start := f.FindStartingPosition()
	possibleNextpossitions := findConnected(f, start)
	previsousPosition := start
	currentPOssition := possibleNextpossitions.GetVectors()[0]
	for currentPOssition != start {
		boundary = append(boundary, currentPOssition)
		//fmt.Println(currentPOssition)
		possibleNextpossitions = findConnected(f, currentPOssition)
		for _, el := range possibleNextpossitions.GetVectors() {
			if el != previsousPosition {
				previsousPosition = currentPOssition
				currentPOssition = el
				break
			}
		}
	}
	return point.InitVecSet(boundary)
}

func findConnected(f field.Field, p point.Vec) point.VecSet {
	tile := f.GetByVec(p)
	possibleNextpossitions := []point.Vec{}
	possitionCandidate := tile.Ways()
	for _, pos := range possitionCandidate.GetVectors() {
		if f.Outside(pos) {
			continue
		}
		loopTile := f.GetByVec(pos)
		if loopTile.Ways().Has(p) {
			possibleNextpossitions = append(possibleNextpossitions, pos)
		}
	}
	return point.InitVecSet(possibleNextpossitions)
}
