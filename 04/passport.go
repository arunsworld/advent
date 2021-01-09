package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// PassportBatch a batch of passports
type PassportBatch []Passport

// Passport for an individual
type Passport struct {
	byr, iyr, eyr string
	hgt, hcl, ecl string
	pid, cid      string
}

// NewPassportBatch parses input data to create a batch
func NewPassportBatch(r io.Reader) (PassportBatch, error) {
	bs := bufio.NewScanner(r)
	result := PassportBatch{}
	passportDataLines := []string{}
	for bs.Scan() {
		line := strings.TrimSpace(bs.Text())
		switch {
		case line == "" && len(passportDataLines) > 0:
			p, err := NewPassport(passportDataLines)
			if err != nil {
				return PassportBatch{}, err
			}
			result = append(result, p)
			passportDataLines = []string{}
		case line != "":
			passportDataLines = append(passportDataLines, line)
		}
	}
	if len(passportDataLines) > 0 {
		p, err := NewPassport(passportDataLines)
		if err != nil {
			return PassportBatch{}, err
		}
		result = append(result, p)
	}
	return result, nil
}

// NewPassport creates a new passport from raw data
func NewPassport(lines []string) (Passport, error) {
	passport := Passport{}
	for _, line := range lines {
		fields := strings.Split(line, " ")
		for _, field := range fields {
			field = strings.TrimSpace(field)
			kv := strings.Split(field, ":")
			if len(kv) != 2 {
				return Passport{}, fmt.Errorf("field %s in line %s could not be correctly parsed", field, line)
			}
			key := kv[0]
			value := kv[1]
			switch key {
			case "byr":
				passport.byr = value
			case "iyr":
				passport.iyr = value
			case "eyr":
				passport.eyr = value
			case "hgt":
				passport.hgt = value
			case "hcl":
				passport.hcl = value
			case "ecl":
				passport.ecl = value
			case "pid":
				passport.pid = value
			case "cid":
				passport.cid = value
			default:
				return Passport{}, fmt.Errorf("encountered unknown key %s in line %s", key, line)
			}
		}
	}
	return passport, nil
}

// ValidPassportCount counts valid passports in batch
func (pb PassportBatch) ValidPassportCount() int {
	count := 0
	for _, passport := range pb {
		if passport.IsValid() {
			count++
		}
	}
	return count
}

// StrictlyValidPassportCount counts strictly valid passports in batch
func (pb PassportBatch) StrictlyValidPassportCount() int {
	count := 0
	for _, passport := range pb {
		if passport.IsStrictlyValid() {
			count++
		}
	}
	return count
}

// IsValid checks if a passport is valid
func (p Passport) IsValid() bool {
	switch {
	case p.byr == "":
		return false
	case p.iyr == "":
		return false
	case p.eyr == "":
		return false
	case p.hgt == "":
		return false
	case p.hcl == "":
		return false
	case p.ecl == "":
		return false
	case p.pid == "":
		return false
	}
	return true
}

// IsStrictlyValid checks if passport is strictly valid
func (p Passport) IsStrictlyValid() bool {
	switch {
	case !p.IsValid():
		return false
	case !p.isBYRValid():
		return false
	case !p.isIYRValid():
		return false
	case !p.isEYRValid():
		return false
	case !p.isHGTValid():
		return false
	case !p.isHCLValid():
		return false
	case !p.isECLValid():
		return false
	case !p.isPIDValid():
		return false
	}
	return true
}

func (p Passport) isBYRValid() bool {
	return validateYearInRange(p.byr, 1920, 2002)
}

func (p Passport) isIYRValid() bool {
	return validateYearInRange(p.iyr, 2010, 2020)
}

func (p Passport) isEYRValid() bool {
	return validateYearInRange(p.eyr, 2020, 2030)
}

func validateYearInRange(input string, min, max int) bool {
	v, err := strconv.Atoi(input)
	switch {
	case err != nil:
		return false
	case v < min:
		return false
	case v > max:
		return false
	}
	return true
}

func (p Passport) isHGTValid() bool {
	if len(p.hgt) < 3 {
		return false
	}
	valueStr := p.hgt[:len(p.hgt)-2]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return false
	}
	units := p.hgt[len(p.hgt)-2:]
	switch {
	case units == "cm" && value >= 150 && value <= 193:
		return true
	case units == "in" && value >= 59 && value <= 76:
		return true
	default:
		return false
	}
}

var hclCodeParser = regexp.MustCompile(`[a-f0-9]{6}`)

func (p Passport) isHCLValid() bool {
	if len(p.hcl) != 7 {
		return false
	}
	if p.hcl[0] != '#' {
		return false
	}
	code := p.hcl[1:]
	if !hclCodeParser.MatchString(code) {
		return false
	}
	return true
}

func (p Passport) isECLValid() bool {
	for _, acceptable := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if p.ecl == acceptable {
			return true
		}
	}
	return false
}

func (p Passport) isPIDValid() bool {
	if len(p.pid) != 9 {
		return false
	}
	_, err := strconv.Atoi(p.pid)
	if err != nil {
		return false
	}
	return true
}
