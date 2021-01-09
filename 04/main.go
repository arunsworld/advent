package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("passport-data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pb, err := NewPassportBatch(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n *** VALID PASSPORTS: %d *** \n\n", pb.ValidPassportCount())

	fmt.Printf("\n *** STRICTLY VALID PASSPORTS: %d *** \n\n", pb.StrictlyValidPassportCount())
}
