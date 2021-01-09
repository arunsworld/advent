package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// PwdValidator contains list of Policy and Data
type PwdValidator []*PolicyAndData

// NewPwdValidator creates a new PwdValidator from input data
func NewPwdValidator(input io.Reader) (PwdValidator, error) {
	csvr := csv.NewReader(input)
	csvr.Comma = ':'
	records, err := csvr.ReadAll()
	if err != nil {
		return PwdValidator{}, fmt.Errorf("input data format error: %v", err)
	}
	result := make(PwdValidator, 0, len(records))
	for _, rec := range records {
		policy, err := NewPolicy(rec[0])
		if err != nil {
			return PwdValidator{}, fmt.Errorf("unable to create Policy from %s: %v", rec[0], err)
		}
		result = append(result, &PolicyAndData{
			Policy: policy,
			Data:   strings.TrimSpace(rec[1]),
		})
	}
	return result, nil
}

// PolicyStandard identifies which standard to apply
type PolicyStandard int

const (
	UnknownPolicyStandard = iota
	OldPolicyStandard
	NewPolicyStandard
)

// Validate validates all data against policies
func (validator PwdValidator) Validate(std PolicyStandard) {
	for _, pd := range validator {
		switch std {
		case OldPolicyStandard:
			pd.IsValid = pd.Policy.ValidateWithOldPolicy(pd.Data)
		case NewPolicyStandard:
			pd.IsValid = pd.Policy.ValidateWithNewPolicy(pd.Data)
		}
	}
}

// Count counts valid policies
func (validator PwdValidator) Count() int {
	var result int
	for _, pd := range validator {
		if pd.IsValid {
			result++
		}
	}
	return result
}

// PolicyAndData holds policy info and data record
type PolicyAndData struct {
	Policy  Policy
	Data    string
	IsValid bool
}

var policyParser = regexp.MustCompile(`(\d+)-(\d+) (\w)`)

// NewPolicy creates a new policy from the given string
func NewPolicy(policyStr string) (Policy, error) {
	policyStr = strings.TrimSpace(policyStr)
	parsedPolicy := policyParser.FindAllStringSubmatch(policyStr, -1)
	if len(parsedPolicy) == 0 {
		return Policy{}, fmt.Errorf("unable to find a valid policy in %s", policyStr)
	}
	letter := rune(parsedPolicy[0][3][0])
	min, err := strconv.Atoi(parsedPolicy[0][1])
	if err != nil {
		return Policy{}, fmt.Errorf("unable to parse min value (%s) as integer", parsedPolicy[0][0])
	}
	max, err := strconv.Atoi(parsedPolicy[0][2])
	if err != nil {
		return Policy{}, fmt.Errorf("unable to parse max value (%s) as integer", parsedPolicy[0][1])
	}
	return Policy{
		letter: letter,
		min:    min,
		max:    max,
	}, nil
}

// Policy defines a password policy
type Policy struct {
	letter   rune
	min, max int
}

// ValidateWithOldPolicy validates given data based on counts
func (p Policy) ValidateWithOldPolicy(data string) bool {
	count := strings.Count(data, string(p.letter))
	switch {
	case count < p.min:
		return false
	case count > p.max:
		return false
	default:
		return true
	}
}

// ValidateWithNewPolicy validates given data based on position
func (p Policy) ValidateWithNewPolicy(data string) bool {
	letterOne := rune(data[p.min-1])
	letterTwo := rune(data[p.max-1])

	switch {
	case letterOne == p.letter && letterTwo != p.letter:
		return true
	case letterOne != p.letter && letterTwo == p.letter:
		return true
	default:
		return false
	}
}
