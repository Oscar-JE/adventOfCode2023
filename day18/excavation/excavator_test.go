package excavation

import (
	"day16/vec"
	"day18/instruction"
	"testing"
)

func TestExcevatorVoidInstruction(t *testing.T) {
	exec, _ := initExcavator()
	instruction1 := instruction.Init(0, vec.Init(0, 0))
	holes := exec.dig(instruction1)
	if len(holes) != 0 {
		t.Errorf("why are you digging?!?")
	}
}

func TestExtractorSimpleInstruction(t *testing.T) {
	exec, _ := initExcavator()
	instruct := instruction.Init(1, vec.Init(0, 1))
	holes1 := exec.dig(instruct)
	if len(holes1) != 1 || holes1[0] != vec.Init(0, 1) {
		t.Errorf("can you not follow simple instructions?")
	}
	holes2 := exec.dig(instruct)
	if len(holes2) != 1 || holes1[0] != vec.Init(0, 2) {
		t.Errorf("You need to remember where you left off")
	}

}
