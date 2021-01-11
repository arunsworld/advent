package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("instructions.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	computer, err := newComputer(f)
	if err != nil {
		log.Fatal(err)
	}
	computer.executeUntilInfiniteLoop()

	fmt.Printf("\n *** VALUE OF ACCUMULATOR: %d *** \n", computer.accumulator)

	_, newComputer, err := computer.instructionToFixToAvoidInfiniteLoop()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n *** VALUE OF ACCUMULATOR AFTER FIX: %d *** \n\n", newComputer.accumulator)
}
