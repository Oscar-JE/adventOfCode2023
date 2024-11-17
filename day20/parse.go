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
	cons := []connection{}
	for _, line := range lines {
		cons = append(cons, connections(line)...)
	}
	for _, con := range cons {
		var from *relays.Relay = nil
		for _, el := range idAndConnectionsList {
			if con.fromId == el.id {
				from = &el.relay
			}
		}
		var to *relays.Relay = nil
		for _, el := range idAndConnectionsList {
			loopId := strings.Trim(el.id, "%&")
			if loopId == con.toId {
				to = &el.relay
			}
		}
		from.AppendOutgoing(to)
		to.AppendIngoing(from)
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
	for _, toId := range toIds {
		connections = append(connections, connection{fromId: fromId, toId: toId})
	}
	return connections
}

func toIds(destinationRep string) []string {
	splitted := strings.Split(destinationRep, ", ")
	return splitted
}

type idAndConnections struct {
	id    string
	relay relays.Relay
}

type connection struct {
	fromId string
	toId   string
}
