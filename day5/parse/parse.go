package parse

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Triplet struct {
	First  int
	Second int
	Last   int
}

func Parse(path string) ([]int, []Triplet) {
	file, err := os.Open(path)
	if err != nil {
		panic("cant find the supplied file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	seedLine := scanner.Text()
	seeds := parseSeed(seedLine)
	triplets := []Triplet{}
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, ":") && line != "" {
			triple := parseTriple(line)
			triplets = append(triplets, triple)
		}
	}
	return seeds, triplets
}

func parseSeed(line string) []int {
	numberLine := strings.Trim(strings.Split(line, ":")[1], " ")
	numbers := strings.Split(numberLine, " ")
	seeds := []int{}
	for _, rep := range numbers {
		seed, err := strconv.Atoi(rep)
		if err != nil {
			panic("misslykande med seed parsning")
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

func parseTriple(line string) Triplet {
	numbers := strings.Split(line, " ")
	first, err1 := strconv.Atoi(numbers[0])
	second, err2 := strconv.Atoi(numbers[1])
	third, err3 := strconv.Atoi(numbers[2])
	if err1 != nil || err2 != nil || err3 != nil {
		panic("failed to parse tripet")
	}
	return Triplet{First: first, Second: second, Last: third}
}
