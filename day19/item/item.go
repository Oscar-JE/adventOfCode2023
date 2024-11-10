package item

import "day19/set"

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

type ItemSet struct {
	cool        set.Set
	musical     set.Set
	aerodynamic set.Set
	shiny       set.Set
}

func InitItemSet(cool set.Set, musical set.Set, aerodynamic set.Set, shiny set.Set) ItemSet {
	return ItemSet{cool: cool, musical: musical, aerodynamic: aerodynamic, shiny: shiny}
}

func EmptyItemSet() ItemSet {
	return ItemSet{cool: set.Set{}, musical: set.Set{}, aerodynamic: set.Set{}, shiny: set.Set{}}
}

func Union(A ItemSet, B ItemSet) ItemSet {
	return ItemSet{cool: set.Union(A.cool, B.cool), musical: set.Union(A.musical, B.musical),
		aerodynamic: set.Union(A.aerodynamic, B.aerodynamic), shiny: set.Union(A.shiny, B.shiny)}
}

type Demand interface {
	Check(Item) bool
	CheckItemSet(ItemSet) ItemSet
}

type NoDemand struct {
}

func (n NoDemand) Check(it Item) bool {
	return true
}

func (n NoDemand) CheckItemSet(it ItemSet) ItemSet {
	return it
}

type Cooler struct {
	CoolDemand int
}

func (c Cooler) Check(it Item) bool {
	return c.CoolDemand < it.cool
}

func (c Cooler) CheckItemSet(it ItemSet) ItemSet {
	newCoolSet := it.cool.DemandBelow(c.CoolDemand)
	it.cool = newCoolSet
	return it
}

type LessCool struct {
	CoolDemand int
}

func (lc LessCool) Check(it Item) bool {
	return lc.CoolDemand > it.cool
}

func (lc LessCool) CheckItemSet(it ItemSet) ItemSet {
	newCoolSet := it.cool.DemandAbove(lc.CoolDemand)
	it.cool = newCoolSet
	return it
}

type MoreMusical struct {
	Musical int
}

func (mm MoreMusical) Check(it Item) bool {
	return mm.Musical < it.musical
}

func (mm MoreMusical) CheckItemSet(it ItemSet) ItemSet {
	newMusical := it.musical.DemandBelow(mm.Musical)
	it.musical = newMusical
	return it
}

type LessMusical struct {
	Musical int
}

func (lm LessMusical) Check(it Item) bool {
	return lm.Musical > it.musical
}

func (lm LessMusical) CheckItemSet(it ItemSet) ItemSet {
	newMusical := it.musical.DemandAbove(lm.Musical)
	it.musical = newMusical
	return it
}

type MoreAerodynamic struct {
	Aerodynamic int
}

func (ma MoreAerodynamic) Check(it Item) bool {
	return ma.Aerodynamic < it.aerodynamic
}

func (ma MoreAerodynamic) CheckItemSet(it ItemSet) ItemSet {
	newAero := it.aerodynamic.DemandBelow(ma.Aerodynamic)
	it.aerodynamic = newAero
	return it
}

type LessAerodynamic struct {
	Aerodynamic int
}

func (la LessAerodynamic) Check(it Item) bool {
	return la.Aerodynamic > it.aerodynamic
}

func (la LessAerodynamic) CheckItemSet(it ItemSet) ItemSet {
	newAero := it.aerodynamic.DemandAbove(la.Aerodynamic)
	it.aerodynamic = newAero
	return it
}

type MoreShiny struct {
	Shiny int
}

func (ms MoreShiny) Check(it Item) bool {
	return ms.Shiny < it.shiny
}

func (ms MoreShiny) CheckItemSet(it ItemSet) ItemSet {
	newShiny := it.shiny.DemandBelow(ms.Shiny)
	it.shiny = newShiny
	return it
}

type LessShiny struct {
	Shiny int
}

func (ls LessShiny) Check(it Item) bool {
	return ls.Shiny > it.shiny
}

func (ls LessShiny) CheckItemSet(it ItemSet) ItemSet {
	newShiny := it.shiny.DemandAbove(ls.Shiny)
	it.shiny = newShiny
	return it
}
