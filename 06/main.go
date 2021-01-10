package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("customs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gs, err := newGroups(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n *** TOTAL ANSWERS: %d *** \n", gs.totalAnyAnswers())

	fmt.Printf("\n *** TOTAL AGREED ANSWERS: %d *** \n\n", gs.totalAgreedAnswers())
}
