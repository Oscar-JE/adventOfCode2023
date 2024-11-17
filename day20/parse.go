package main

import (
	"day20/network"
	"day20/relays"
	"os"
	"strings"
)

func parseFile(filename string) network.Network {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("file reading failed")
	}
	network := parseExample(string(content))
	return network
}

func parseExample(content string) network.Network {
	lines := strings.Split(content, "\r\n")
	idAndConnectionsList := []idAndConnections{}
	for _, line := range lines {
		idAndConnectionsList = append(idAndConnectionsList, emptyWithId(line))
	}
	for _, line := range lines {
		connections := cennections(line)
	}
	return network.Network{}
}

func emptyWithId(line string) idAndConnections {
	splitted := strings.Split(line, " -> ")
	name := splitted[0]
	return idAndConnections{name, relays.Relay{}}
}

func connections(line string) []connection {
	connections := []connection{}
	splitted := strings.Split(line, " -> ")
	fromId := splitted[0]
	toIds := toIds(splitted[1])
}

type idAndConnections struct {
	id         string
	connection relays.Relay
}

type connection struct {
	fromId string
	toId   string
}
