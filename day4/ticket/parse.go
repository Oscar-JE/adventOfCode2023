package ticket

import (
	"strconv"
	"strings"
)

func Parse(line string) ([]int, []int) {
	arrays := strings.Split(line, ": ")[1]
	arr1And2 := strings.Split(arrays, " | ")
	left := parseArr(arr1And2[0])
	right := parseArr(arr1And2[1])
	return left, right
}

func parseArr(line string) []int {
	symbols := strings.Split(line, " ")
	ints := []int{}
	for _, symbol := range symbols {
		if symbol == "" {
			continue
		}
		number, err := strconv.Atoi(symbol)
		if err != nil {
			panic("error in parsing line numbers")
		}
		ints = append(ints, number)
	}
	return ints
}

func ParseTally(lines []string) TicketTally {
	tickets := []Ticket{}
	for _, line := range lines {
		winning, holding := Parse(line)
		tickets = append(tickets, Ticket{winning: winning, holding: holding})
	}
	return InitTally(tickets)
}
