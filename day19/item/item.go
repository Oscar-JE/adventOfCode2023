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
	cool        set.Interval
	musical     set.Interval
	aerodynamic set.Interval
	shiny       set.Interval
}

func (it ItemSet) Cardinality() int {
	return it.cool.Cardinality() * it.musical.Cardinality() * it.aerodynamic.Cardinality() * it.shiny.Cardinality()
}

func InitItemSet(cool set.Interval, musical set.Interval, aerodynamic set.Interval, shiny set.Interval) ItemSet {
	return ItemSet{cool: cool, musical: musical, aerodynamic: aerodynamic, shiny: shiny}
}

func EmptyItemSet() ItemSet {
	return ItemSet{cool: set.InitInterval(0, 0), musical: set.InitInterval(0, 0), aerodynamic: set.InitInterval(0, 0), shiny: set.InitInterval(0, 0)}
}

func StandardItemSet(lowerIncluded int, upperIncluded int) ItemSet {
	standarInterval := set.InitInterval(lowerIncluded, upperIncluded+1)
	return ItemSet{cool: standarInterval, musical: standarInterval, aerodynamic: standarInterval, shiny: standarInterval}
}

type Demand interface {
	Check(Item) bool
	CheckItemSet(ItemSet) (ItemSet, ItemSet)
}

type NoDemand struct {
}

func (n NoDemand) Check(it Item) bool {
	return true
}

func (n NoDemand) CheckItemSet(it ItemSet) (ItemSet, ItemSet) {
	return it, EmptyItemSet()
}

type Cooler struct {
	CoolDemand int
}

func (c Cooler) Check(it Item) bool {
	return c.CoolDemand < it.cool
}

func (c Cooler) CheckItemSet(it ItemSet) (ItemSet, ItemSet) {
	newCoolSet, coolComplement := it.cool.DemandBelow(c.CoolDemand)
	complement := it
	complement.cool = coolComplement
	it.cool = newCoolSet
	return it, complement
}

type LessCool struct {
	CoolDemand int
}

func (lc LessCool) Check(it Item) bool {
	return lc.CoolDemand > it.cool
}

func (lc LessCool) CheckItemSet(it ItemSet) (ItemSet, ItemSet) {
	newCoolSet, coolSetComplement := it.cool.DemandAbove(lc.CoolDemand)
	complement := it
	complement.cool = coolSetComplement
	it.cool = newCoolSet
	return it, complement
}

type MoreMusical struct {
	Musical int
}

func (mm MoreMusical) Check(it Item) bool {
	return mm.Musical < it.musical
}

func (mm MoreMusical) CheckItemSet(it ItemSet) (ItemSet, ItemSet) {
	newMusical, musicalComplement := it.musical.DemandBelow(mm.Musical)
	complement := it
	complement.musical = musicalComplement
	it.musical = newMusical
	return it, complement
}

type LessMusical struct {
	Musical int
}

func (lm LessMusical) Check(it Item) bool {
	return lm.Musical > it.musical
}

func (lm LessMusical) CheckItemSet(it ItemSet) (ItemSet, ItemSet) {
	newMusical, musicalComplement := it.musical.DemandAbove(lm.Musical)
	complement := it
	complement.musical = musicalComplement
	it.musical = newMusical
	return it, complement
}

type MoreAerodynamic struct {
	Aerodynamic int
}

func (ma MoreAerodynamic) Check(it Item) bool {
	return ma.Aerodynamic < it.aerodynamic
}

func (ma MoreAerodynamic) CheckItemSet(it ItemSet) (ItemSet, ItemSet) {
	newAero, aeroComplement := it.aerodynamic.DemandBelow(ma.Aerodynamic)
	complement := it
	complement.aerodynamic = aeroComplement
	it.aerodynamic = newAero
	return it, complement
}

type LessAerodynamic struct {
	Aerodynamic int
}

func (la LessAerodynamic) Check(it Item) bool {
	return la.Aerodynamic > it.aerodynamic
}

func (la LessAerodynamic) CheckItemSet(it ItemSet) (ItemSet, ItemSet) {
	newAero, aeroComplement := it.aerodynamic.DemandAbove(la.Aerodynamic)
	complement := it
	complement.aerodynamic = aeroComplement
	it.aerodynamic = newAero
	return it, complement
}

type MoreShiny struct {
	Shiny int
}

func (ms MoreShiny) Check(it Item) bool {
	return ms.Shiny < it.shiny
}

func (ms MoreShiny) CheckItemSet(it ItemSet) (ItemSet, ItemSet) {
	newShiny, shinyComplement := it.shiny.DemandBelow(ms.Shiny)
	complement := it
	complement.shiny = shinyComplement
	it.shiny = newShiny
	return it, complement
}

type LessShiny struct {
	Shiny int
}

func (ls LessShiny) Check(it Item) bool {
	return ls.Shiny > it.shiny
}

func (ls LessShiny) CheckItemSet(it ItemSet) (ItemSet, ItemSet) {
	newShiny, shinyComplement := it.shiny.DemandAbove(ls.Shiny)
	complement := it
	complement.shiny = shinyComplement
	it.shiny = newShiny
	return it, complement
}
