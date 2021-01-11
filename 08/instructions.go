package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type computer struct {
	instructions []instruction
	// computer state during execution
	pos                  int // where we are on the instructions
	accumulator          int
	executedInstructions map[int]struct{}
}

type instruction struct {
	code  code
	value int
}

type code int

const (
	nop code = iota
	acc
	jmp
)

func newCode(c string) code {
	switch c {
	case "acc":
		return acc
	case "jmp":
		return jmp
	default:
		return nop
	}
}

func newComputer(r io.Reader) (*computer, error) {
	csvr := csv.NewReader(r)
	csvr.Comma = ' '
	records, err := csvr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to parse instructions: %v", err)
	}
	instructions := make([]instruction, 0, len(records))
	for _, record := range records {
		codeStr, vStr := record[0], record[1]
		codeStr = strings.TrimSpace(codeStr)
		vStr = strings.TrimSpace(vStr)
		v, err := strconv.Atoi(vStr)
		if err != nil {
			return nil, fmt.Errorf("instruction value %s not a number", vStr)
		}
		instructions = append(instructions, instruction{
			code:  newCode(codeStr),
			value: v,
		})
	}
	return &computer{instructions: instructions}, nil
}

// registerInstructionExecution returns false if it's a duplication
func (c *computer) registerInstructionExecution(i int) bool {
	_, instructionRepeated := c.executedInstructions[i]
	switch instructionRepeated {
	case true:
		return false
	default:
		c.executedInstructions[i] = struct{}{}
		return true
	}
}

// executeUntilInfiniteLoop returns true if exection stops on detection of infinite loop (otherwise false)
func (c *computer) executeUntilInfiniteLoop() bool {
	i := 0 // start with the first instruction & reset state
	c.executedInstructions = make(map[int]struct{})
	c.accumulator = 0
	for i < len(c.instructions) {
		if !c.registerInstructionExecution(i) {
			return true
		}
		instruction := c.instructions[i]
		switch instruction.code {
		case nop:
			i++
		case acc:
			c.accumulator += instruction.value
			i++
		case jmp:
			i += instruction.value
		}
	}
	return false
}

func (c *computer) copyWithInstructionSwapAfter(numberOfInstructions int) (*computer, int) {
	var swapDone bool
	var swapPosition int
	result := make([]instruction, 0, len(c.instructions))
	for i, instruction := range c.instructions {
		switch {
		case i < numberOfInstructions:
		case swapDone:
		case instruction.code == jmp:
			instruction.code = nop
			swapDone = true
			swapPosition = i
		case instruction.code == nop:
			instruction.code = jmp
			swapDone = true
			swapPosition = i
		default:
		}
		result = append(result, instruction)
	}
	return &computer{instructions: result}, swapPosition
}

func (c *computer) instructionToFixToAvoidInfiniteLoop() (int, *computer, error) {
	if !c.executeUntilInfiniteLoop() {
		return 0, c, nil
	}
	skipCount := 0
	for skipCount < len(c.instructions) {
		newC, newSkip := c.copyWithInstructionSwapAfter(skipCount)
		if !newC.executeUntilInfiniteLoop() {
			return newSkip, newC, nil
		}
		skipCount = newSkip + 1
	}
	return 0, nil, fmt.Errorf("impossible to fix infinite loop")
}
