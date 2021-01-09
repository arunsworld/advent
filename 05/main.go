package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("boarding_passes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	aircraft, err := NewAircraft(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n *** MAX SEAT ID: %d *** \n\n", aircraft.maxSeatID())

	fmt.Printf("\n *** MISSING SEAT ID: %d *** \n\n", aircraft.missingSeatID())
}
