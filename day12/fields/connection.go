package fields

type Connection struct {
	nrOfMoved     int //negative means moved in opposite direction , positive direction is to the right
	partInboarder bool
	boarderIndex  int // only relevant if one part is inside the boarder part in boarder always taken from the left
}

func FittingConnector(c Connection) Connection {
	c.nrOfMoved = -c.nrOfMoved
	return c
}
