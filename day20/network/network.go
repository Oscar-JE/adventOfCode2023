package network

import "day20/relays"

type Network struct { // bra att ha en bild av strukturen man ska bygga
	nodes       []relays.Relay
	broadcaster relays.Relay
}
