package instruction

import "strings"

func parseInstructions(content string, lineEnding string) []Instruction {
	lines := strings.Split(content, lineEnding)
	instructs := []Instruction{}
	for _, line := range lines {
		instructs = append(instructs, parse(line))
	}
	return instructs
}
