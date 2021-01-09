package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("passwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	validator, err := NewPwdValidator(f)
	if err != nil {
		log.Fatal(err)
	}

	validator.Validate(OldPolicyStandard)

	fmt.Printf("\n*** OLD POLICY RESULT: %d ***\n\n", validator.Count())

	validator.Validate(NewPolicyStandard)

	fmt.Printf("\n*** NEW POLICY RESULT: %d ***\n\n", validator.Count())

}
