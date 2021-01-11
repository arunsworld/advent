package main

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

func Test_Instructions_Can_Be_Parsed_Into_A_Computer(t *testing.T) {
	sample := `nop +0
				acc +1
				jmp -4`
	comp, err := newComputer(strings.NewReader(sample))
	if err != nil {
		t.Fatal(err)
	}
	expected := &computer{
		instructions: []instruction{
			instruction{code: nop, value: 0},
			instruction{code: acc, value: 1},
			instruction{code: jmp, value: -4},
		},
	}
	if !reflect.DeepEqual(comp, expected) {
		log.Println(comp)
		t.Fatal("computer didn't match expected")
	}
	t.Run("and acc/jmp swapped", func(t *testing.T) {
		comp, v := comp.copyWithInstructionSwapAfter(1)
		expected := &computer{
			instructions: []instruction{
				instruction{code: nop, value: 0},
				instruction{code: acc, value: 1},
				instruction{code: nop, value: -4},
			},
		}
		if !reflect.DeepEqual(comp, expected) {
			log.Println(comp)
			t.Fatal("computer didn't match expected")
		}
		if v != 2 {
			t.Fatal("swap location didn't match expected")
		}
	})
}

func Test_A_Computer(t *testing.T) {
	input := `nop +0
				acc +1
				jmp +4
				acc +3
				jmp -3
				acc -99
				acc +1
				jmp -4
				acc +6`
	computer, err := newComputer(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	t.Run("can process instructions and stop when an infinite loop is detected", func(t *testing.T) {
		computer.executeUntilInfiniteLoop()
		if computer.accumulator != 5 {
			t.Fatalf("expected 5 got: %d", computer.accumulator)
		}
	})
	t.Run("can fix instruction causing infinite loop", func(t *testing.T) {
		result, newC, err := computer.instructionToFixToAvoidInfiniteLoop()
		if err != nil {
			t.Fatal(err)
		}
		if result != 7 {
			t.Fatalf("expected instruction 7, got: %d", result)
		}
		if newC.accumulator != 8 {
			t.Fatalf("expected accumulator 8, got: %d", newC.accumulator)
		}
	})
}
