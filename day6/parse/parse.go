package parse

import (
	"bufio"
	"day6/solution"
	"os"
	"strconv"
	"strings"
)

func Parse(path string) []solution.Race {
	timeAndDistaces := []solution.Race{}
	lines := getLines(path)
	timeLine := lines[0]
	distLine := lines[1]
	times := parseLine(timeLine)
	distances := parseLine(distLine)
	for i := range times {
		timeAndDistaces = append(timeAndDistaces, solution.Race{Time: times[i], Distance: distances[i]})
	}
	return timeAndDistaces
}

func getLines(path string) []string {
	lines := []string{}
	file, err := os.Open(path)
	if err != nil {
		panic("cant find the supplied file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func parseLine(line string) []int {
	times := []int{}
	numberRep := strings.Trim(strings.Split(line, ":")[1], " ")
	numbers := strings.Split(numberRep, " ")
	for _, numRep := range numbers {
		if numRep == "" {
			continue
		}
		t, err := strconv.Atoi(numRep)
		if err != nil {
			panic("fan i string parsing")
		}
		times = append(times, t)
	}
	return times
}

func Parse2(path string) solution.Race {
	lines := getLines(path)
	timeLine := lines[0]
	distLine := lines[1]
	time := parseLine2(timeLine)
	dist := parseLine2(distLine)
	return solution.Race{Time: time, Distance: dist}
}

func parseLine2(line string) int {
	numberRep := strings.ReplaceAll(strings.Split(line, ":")[1], " ", "")
	number, err := strconv.Atoi(numberRep)
	if err != nil {
		panic("parseLine2 failed to convert number representation to number")
	}
	return number
}
