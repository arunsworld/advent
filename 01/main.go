package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := flag.String("input", "input.txt", "input file containing expense report")
	flag.Parse()

	rawInput, err := readRawInputFromFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lst, err := expenseReportContentsToList(rawInput)
	if err != nil {
		log.Fatal(err)
	}

	result, err := productOfTwoSum(lst, 2020)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n**** TWO SUM RESULT ****: %v\n\n", result)

	result, err = productOfThreeSum(lst, 2020)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n**** THREE SUM RESULT ****: %v\n\n", result)
}

func readRawInputFromFile(fname string) (string, error) {
	f, err := os.Open(fname)
	if err != nil {
		return "", fmt.Errorf("unable to open %s: %v", fname, err)
	}
	defer f.Close()

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("unable to read %s: %v", fname, err)
	}
	return string(contents), nil
}

func expenseReportContentsToList(lst string) ([]int, error) {
	r := strings.NewReader(lst)
	csvr := csv.NewReader(r)
	records, err := csvr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read list as expected: %v", err)
	}
	output := make([]int, 0, len(records))
	for _, v := range records {
		vv := strings.TrimSpace(v[0])
		vvv, err := strconv.Atoi(vv)
		if err != nil {
			return nil, fmt.Errorf("unable to read list as expected: %v", err)
		}
		output = append(output, vvv)
	}
	return output, nil
}

func productOfTwoSum(lst []int, target int) (int, error) {
	if len(lst) == 1 {
		return 0, fmt.Errorf("no two sum combinations found")
	}
	leftOp := lst[0]
	for i := 1; i < len(lst); i++ {
		rightOp := lst[i]
		if leftOp+rightOp == target {
			return leftOp * rightOp, nil
		}
	}
	return productOfTwoSum(lst[1:], target)
}

func productOfThreeSum(lst []int, target int) (int, error) {
	if len(lst) == 2 {
		return 0, fmt.Errorf("no three sum combinations found")
	}
	leftOp := lst[0]
	rightOp, err := productOfTwoSum(lst[1:], target-leftOp)
	if err == nil {
		return leftOp * rightOp, nil
	}
	return productOfThreeSum(lst[1:], target)
}
