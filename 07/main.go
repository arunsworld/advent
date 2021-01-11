package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("bag_rules.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bags, err := newBags(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n *** ALL CONTAINERS FOR shiny gold: %d *** \n", len(bags.allContainersOf("shiny gold")))

	fmt.Printf("\n *** ALL CONTENTS OF shiny gold: %d *** \n\n", bags.allContentsOf("shiny gold"))
}
