package item

type Item struct {
	cool        int
	musical     int
	aerodynamic int
	shiny       int
}

func Init(cool int, musical int, aerodynamic int, shiny int) Item {
	return Item{cool: cool, musical: musical, aerodynamic: aerodynamic, shiny: shiny}
}

func (it Item) Score() int {
	return it.cool + it.musical + it.aerodynamic + it.shiny
}

type Demand interface {
	Check(Item) bool
}

type NoDemand struct {
}

func (n NoDemand) Check(it Item) bool {
	return true
}

type Cooler struct {
	CoolDemand int
}

func (c Cooler) Check(it Item) bool {
	return c.CoolDemand < it.cool
}

type LessCool struct {
	CoolDemand int
}

func (lc LessCool) Check(it Item) bool {
	return lc.CoolDemand > it.cool
}

type MoreMusical struct {
	Musical int
}

func (mm MoreMusical) Check(it Item) bool {
	return mm.Musical < it.musical
}

type LessMusical struct {
	Musical int
}

func (lm LessMusical) Check(it Item) bool {
	return lm.Musical > it.musical
}

type MoreAerodynamic struct {
	Aerodynamic int
}

func (ma MoreAerodynamic) Check(it Item) bool {
	return ma.Aerodynamic < it.aerodynamic
}

type LessAerodynamic struct {
	Aerodynamic int
}

func (la LessAerodynamic) Check(it Item) bool {
	return la.Aerodynamic > it.aerodynamic
}

type MoreShiny struct {
	Shiny int
}

func (ms MoreShiny) Check(it Item) bool {
	return ms.Shiny < it.shiny
}

type LessShiny struct {
	Shiny int
}

func (ls LessShiny) Check(it Item) bool {
	return ls.Shiny > it.shiny
}
