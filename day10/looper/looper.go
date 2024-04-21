package looper

import (
	"day10/field"
	"day10/middle"
	"day10/point"
	"fmt"
)

func NrOfSrepsToFurthestPoint(f field.Field) int {
	loopLen := findLoopExcludingStart(f)
	return middle.MiddlePointFromZero(loopLen)
}

func findLoopExcludingStart(f field.Field) int {
	boundary := findBoundaryOfLoop(f)
	return boundary.Len() - 1
}

func findBoundaryOfLoop(f field.Field) point.VecList {
	boundary := []point.Vec{}
	start := f.FindStartingPosition()
	possibleNextpossitions := findConnected(f, start)
	previsousPosition := start
	boundary = append(boundary, start)
	currentPOssition := possibleNextpossitions.GetVectors()[0]
	for currentPOssition != start {
		boundary = append(boundary, currentPOssition)
		possibleNextpossitions = findConnected(f, currentPOssition)
		for _, el := range possibleNextpossitions.GetVectors() {
			if el != previsousPosition {
				previsousPosition = currentPOssition
				currentPOssition = el
				break
			}
		}
	}

	return point.InitVecList(boundary)
}

func NrOfInteriorPoints(f field.Field) int {
	boundary := findBoundaryOfLoop(f)
	startingPoints := findStartingPoints(boundary)
	filled := flodFillDetailjerad(f, startingPoints)
	f.RemoveEverythingExcept(boundary)
	f.FillTheseWith(point.InitVecList(filled.GetVectors()), 'I')
	fmt.Println(f)
	return filled.Len()
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

func flodFill(f field.Field, startingPoints point.VecList) int {
	return flodFillDetailjerad(f, startingPoints).Len()
}

func flodFillDetailjerad(f field.Field, startingPoints point.VecList) point.VecSet {
	boundary := findBoundaryOfLoop(f)
	visitied := point.InitVecSet([]point.Vec{})
	evaluationOrder := startingPoints
	for !evaluationOrder.IsEmpty() {
		evalute := evaluationOrder.Popp()
		if f.Outside(evalute) || boundary.Has(evalute) || visitied.Has(evalute) {
			continue
		}
		visitied.Add(evalute)
		connected := cross(evalute)
		for _, el := range connected {
			evaluationOrder.Append(el)
		}
	}
	return visitied
}

func cross(v point.Vec) []point.Vec {
	up := point.Init(-1, 0)
	right := point.Init(0, 1)
	down := point.Init(1, 0)
	left := point.Init(0, -1)
	return []point.Vec{v.Add(up), v.Add(right), v.Add(down), v.Add(left)}
}

func findStartingPoints(bondary point.VecList) point.VecList {
	currentVec, indexOfClosest := bondary.FindCLosestTo(point.Init(0, 0))
	nextVector := bondary.Get(indexOfClosest + 1)
	loopDirection := nextVector.Subtract(currentVec)
	nrOfRotations := 1
	inwards := point.Rotate90counterclockwise(loopDirection)
	if point.InnerProduct(inwards, currentVec) < 0 {
		nrOfRotations = 3
	}
	startingPoints := point.InitVecList([]point.Vec{})
	insideCandidate := nextVector.Add(point.Rotate90counterclockwiseMultipleTimes(loopDirection, nrOfRotations))
	if !bondary.Has(insideCandidate) {
		startingPoints.Append(insideCandidate)
	}
	for i := 1; i < bondary.Len(); i++ {
		currentVec := bondary.Get(i)
		nextVector = bondary.Get(i + 1)
		loopDirection = nextVector.Subtract(currentVec)
		insideCandidate = nextVector.Add(point.Rotate90counterclockwiseMultipleTimes(loopDirection, nrOfRotations))
		insideCandidate2 := currentVec.Add(point.Rotate90counterclockwiseMultipleTimes(loopDirection, nrOfRotations))
		if !bondary.Has(insideCandidate) && !startingPoints.Has(insideCandidate) {
			startingPoints.Append(insideCandidate)
		}
		if !bondary.Has(insideCandidate2) && !startingPoints.Has(insideCandidate2) {
			startingPoints.Append(insideCandidate)
		}

	}
	return startingPoints
}
