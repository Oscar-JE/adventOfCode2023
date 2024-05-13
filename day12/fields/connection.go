package fields

type Connection struct {
	// positiv direction is to the right
	nrOfMoved     int  //negative means moved in opposite direction , positive direction is to the right
	partInboarder bool // kan vi säga med säkerhet att den kommer från vänster? Ja om vi expanderar bitarna på rätt sätt
	boarderIndex  int  // only relevant if one part is inside the boarder part in boarder always taken from the left
}

func FittingConnector(c Connection) Connection {
	c.nrOfMoved = -c.nrOfMoved // här måste vi ta hänsyn till om det finns någon i väggen också ?
	// dp,

	return c
}
