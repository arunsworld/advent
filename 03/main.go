package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("forest.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	forest, err := NewForest(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n *** TREE COUNT: %d *** \n\n", forest.CountTrees(Slope{Right: 3, Down: 1}))

	multiSlopeProduct := forest.CountTrees(Slope{Right: 1, Down: 1}) *
		forest.CountTrees(Slope{Right: 3, Down: 1}) *
		forest.CountTrees(Slope{Right: 5, Down: 1}) *
		forest.CountTrees(Slope{Right: 7, Down: 1}) *
		forest.CountTrees(Slope{Right: 1, Down: 2})

	fmt.Printf("\n *** MULTI SLOPE PRODUCT: %d *** \n\n", multiSlopeProduct)
}
