package solution

type Connection struct {
	blockLenght    int
	blockPlacement int
}

func NoConnectionDDemand() Connection {
	return Connection{blockLenght: 0, blockPlacement: 0}
}

func InitConnection(block int, placement int) Connection {
	return Connection{blockLenght: block, blockPlacement: placement}
}

func (c Connection) LeftDemand() int {
	if c.blockLenght <= 0 {
		return -1
	}
	return c.blockPlacement
}

func (c Connection) RightDemand() int {
	if c.blockLenght <= 0 {
		return -1
	}
	return c.blockLenght - 1 - c.blockLenght
}
