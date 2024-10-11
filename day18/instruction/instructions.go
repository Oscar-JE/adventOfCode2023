package instruction

import "strings"

func ParseInstructions(content string, lineEnding string) []Instruction {
	lines := strings.Split(content, lineEnding)
	instructs := []Instruction{}
	for _, line := range lines {
		instructs = append(instructs, parse(line))
	}
	return instructs
}

func ParseInstructions2(content string, lineEnding string) []Instruction {
	lines := strings.Split(content, lineEnding)
	instructs := []Instruction{}
	for _, line := range lines {
		instructs = append(instructs, parse2(line))
	}
	return instructs
}
