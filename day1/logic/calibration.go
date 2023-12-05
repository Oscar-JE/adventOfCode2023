package logic

import (
	"errors"
	"strconv"
	"strings"
)

var didgitWords []string = []string{"zero",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func ValueFromDigitAndWords(input string) int {
	return diditFromWordLeft(input)*10 + diditFromWordRight(input)
}

func diditFromWordLeft(input string) int {
	i := 0
	wordCache := ""
	for i < len(input) {
		character := string(input[i])
		digit, errDidit := strconv.Atoi(character)
		wordCache = wordCache + character
		digitFromWord, errWord := didgitFromWord(wordCache)
		if errDidit == nil {
			return digit
		}
		if errWord == nil {
			return digitFromWord
		}
		i++
	}
	return 0
}

func diditFromWordRight(input string) int {
	i := len(input) - 1
	wordCache := ""
	for i >= 0 {
		character := string(input[i])
		digit, errDidit := strconv.Atoi(character)
		wordCache = character + wordCache
		digitFromWord, errWord := didgitFromWord(wordCache)
		if errDidit == nil {
			return digit
		}
		if errWord == nil {
			return digitFromWord
		}
		i--
	}
	return 0
}

func didgitFromWord(input string) (int, error) {
	for i, word := range didgitWords {
		if strings.Contains(input, word) {
			return i, nil
		}
	}
	return 0, errors.New("no matching")
}

func ValuFromDigits(input string) int {
	return diditFromLeft(input)*10 + digitFromRight(input)
}

func diditFromLeft(input string) int {
	i := 0
	for i < len(input) {
		digit, err := strconv.Atoi(string(input[i]))
		if err == nil {
			return digit
		}
		i++
	}
	return 0
}

func digitFromRight(input string) int {
	i := len(input) - 1
	for i >= 0 {
		digit, err := strconv.Atoi(string(input[i]))
		if err == nil {
			return digit
		}
		i--
	}
	return 0
}
