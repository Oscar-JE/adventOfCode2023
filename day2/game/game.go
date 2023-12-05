package game

import (
	"strconv"
	"strings"
)

type Game struct {
	identity    int
	restriction Set
	played      []Set
}

func (g *Game) SetIdentity(id int) {
	g.identity = id
}

func (g *Game) SetPlayed(sets []Set) {
	g.played = sets
}

func Init(restriction Set) Game {
	return Game{restriction: restriction, played: []Set{}, identity: 0}
}

func (g Game) Score() int {
	if g.possible() {
		return g.identity
	}
	return 0
}

func (g Game) possible() bool {
	possible := true
	for _, set := range g.played {
		possible = possible && g.restriction.contains(set)
	}
	return possible
}

func ParseShowing(line string) (identity int, sets []Set) {
	identityAndRest := strings.Split(line, ": ")
	id := parseIdentiy(identityAndRest[0])
	setsReps := strings.Split(identityAndRest[1], "; ")
	played := []Set{}
	for _, setRep := range setsReps {
		played = append(played, parseSet(setRep))
	}
	return id, played
}

func parseIdentiy(identity string) int {
	scoreStr := strings.Split(identity, " ")[1]
	score, err := strconv.Atoi(scoreStr)
	if err != nil {
		panic("parse of identity failed")
	}
	return score
}

func parseSet(setRep string) Set {
	rgb := strings.Split(setRep, ", ")
	red := parseColor(rgb, "red")
	green := parseColor(rgb, "green")
	blue := parseColor(rgb, "blue")
	return Set{Red: red, Green: green, Blue: blue}
}

func parseColor(rgb []string, ask string) int {
	for _, rep := range rgb {
		intAndWord := strings.Split(rep, " ")
		if intAndWord[1] == ask {
			number, err := strconv.Atoi(intAndWord[0])
			if err != nil {
				panic("getColor failed to parse int")
			}
			return number
		}
	}
	return 0
}
