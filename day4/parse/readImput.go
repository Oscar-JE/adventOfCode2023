package parse

import (
	"bufio"
	"os"
)

func ReadInput(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		return []string{}
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
