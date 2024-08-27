package mirror

import (
	"day16/mirrors"
	"day16/vec"
)

type Empty struct {
}

// det finns ingen anledning till att Relfect ska känna till classen Bean när vi enbart jobbar med vectorer

func (e Empty) Reflect(v vec.Vec2d) []vec.Vec2d {
	return []vec.Vec2d{v}
}

func (e Empty) Rep() string {
	return "."
}

type Splitter struct {
	splitDirection vec.Vec2d
}

func (s Splitter) Reflect(b vec.Vec2d) []vec.Vec2d { // får läsa på om det är rätt beteende
	if b.Orthogonal(s.splitDirection) {
		dir1 := s.splitDirection
		dir2 := s.splitDirection.Flip()
		return []vec.Vec2d{dir1, dir2}
	}
	return []vec.Vec2d{b}
}

func (s Splitter) Rep() string {
	if vec.Init(0, 1) == s.splitDirection {
		return "-"
	}
	return "|"
}

type Mirror struct {
	normal vec.Vec2d // får se till att längden av normalen är en int eller om jag kan lösa den på annat sätt
}

func (m Mirror) Reflect(v vec.Vec2d) []vec.Vec2d {
	infallande := v
	korrektNormal := m.korrektnNormal(infallande)
	nyRiktning := spegla(korrektNormal, infallande)
	return []vec.Vec2d{nyRiktning}
}

func (m Mirror) Rep() string {
	if m.normal == vec.Init(-1, 1) {
		return "\\"
	}
	return "/"
}

func (m Mirror) korrektnNormal(infallande vec.Vec2d) vec.Vec2d {
	korrektnNormal := m.normal
	orginalRiktning := infallande
	if korrektnNormal.InnerProduct(orginalRiktning) > 0 {
		korrektnNormal = korrektnNormal.Flip()
	}
	return korrektnNormal
}

func spegla(normal vec.Vec2d, infallande vec.Vec2d) vec.Vec2d {
	normalLengthSquared := normal.InnerProduct(normal)
	infallande = infallande.ScalarMultiplication(normalLengthSquared) // känns som att det ska kunna förkortas
	normalAntiRiktning := normal.ScalarMultiplication(-normal.InnerProduct(infallande) / normalLengthSquared)
	toLong := vec.Add(infallande, normalAntiRiktning.ScalarMultiplication(2)) // nej går inte då måste man muttiplicera här istället
	return toLong.ScalarDevision(normalLengthSquared)
}

func Init(r rune) mirrors.Reflector {
	if r == '|' {
		return Splitter{vec.Init(1, 0)}
	} else if r == '-' {
		return Splitter{vec.Init(0, 1)}
	} else if r == '\\' {
		return Mirror{vec.Init(-1, 1)}
	} else if r == '/' {
		return Mirror{vec.Init(1, 1)}
	}
	return Empty{}
}
