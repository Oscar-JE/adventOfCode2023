package item

type Item struct {
	cool        int
	musical     int
	aerodynamic int
	shiny       int
}

type Demand interface {
	Check(Item) bool
}

type Cooler struct {
	coolDemand int
}

func (c Cooler) Check(it Item) bool {
	return c.coolDemand < it.cool
}

type LessCool struct {
	coolDemand int
}

func (lc LessCool) Check(it Item) bool {
	return lc.coolDemand > it.cool
}

type MoreMusical struct {
	musical int
}

func (mm MoreMusical) Check(it Item) bool {
	return mm.musical < it.musical
}

type LessMusical struct {
	musical int
}

func (lm LessMusical) Check(it Item) bool {
	return lm.musical > it.musical
}

type MoreAerodynamic struct {
	aerodynamic int
}

func (ma MoreAerodynamic) Check(it Item) bool {
	return ma.aerodynamic < it.aerodynamic
}

type LessAerodynamic struct {
	aerodynamic
}

func (la LessAerodynamic) Check(it Item) bool {
	return la.aerodynamic > it.aerodynamic
}

type MoreShiny struct {
	shiny int
}

func (ms MoreShiny) Check(it Item) bool {
	return ms.shiny < it.shiny
}

type LessShiny struct {
	shiny int
}

func (ls LessShiny) Check(it Item) bool {
	return ls.shiny > it.shiny
}
