package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

// Aircraft has seats
type Aircraft struct {
	Seats []Seat
}

// NewAircraft creates an aircraft seats from a batch of boarding passes
func NewAircraft(r io.Reader) (Aircraft, error) {
	csvr := csv.NewReader(r)
	records, err := csvr.ReadAll()
	if err != nil {
		return Aircraft{}, err
	}
	seats := make([]Seat, 0, len(records))
	for _, record := range records {
		seat, err := NewSeat(record[0])
		if err != nil {
			return Aircraft{}, err
		}
		seats = append(seats, seat)
	}
	return Aircraft{Seats: seats}, nil
}

func (a Aircraft) maxSeatID() int {
	max := 0
	for _, seat := range a.Seats {
		if seat.identifier() > max {
			max = seat.identifier()
		}
	}
	return max
}

func (a Aircraft) missingSeatID() int {
	seatIDs := make([]int, 0, len(a.Seats))
	for _, seat := range a.Seats {
		seatIDs = append(seatIDs, seat.identifier())
	}
	sort.Ints(seatIDs)
	firstSeatID := 0
	for i, v := range seatIDs {
		if i == 0 {
			firstSeatID = v
			continue
		}
		if v != firstSeatID+i {
			return firstSeatID + i
		}
	}
	return 0
}

// Seat on a plane
type Seat struct {
	Row, Column int
}

// NewSeat gives a seat based on boarding pass code
func NewSeat(code string) (Seat, error) {
	if len(code) != 10 {
		return Seat{}, fmt.Errorf("code should be 10 digits long")
	}
	seat := Seat{}
	row, err := seat.findRow(code[:8])
	if err != nil {
		return Seat{}, err
	}
	seat.Row = row

	col, err := seat.findColumn(code[7:])
	if err != nil {
		return Seat{}, err
	}
	seat.Column = col
	return seat, nil
}

func (seat Seat) findRow(code string) (int, error) {
	calc := binaryCalculator{zero: 'F', one: 'B'}
	v, _, err := calc.resolve(code, 0, 127)
	if err != nil {
		return 0, fmt.Errorf("problem finding row from code: %s: %v", code, err)
	}
	return v, nil
}

func (seat Seat) findColumn(code string) (int, error) {
	calc := binaryCalculator{zero: 'L', one: 'R'}
	v, _, err := calc.resolve(code, 0, 7)
	if err != nil {
		return 0, fmt.Errorf("problem finding column from code: %s: %v", code, err)
	}
	return v, nil
}

func (seat Seat) identifier() int {
	return seat.Row*8 + seat.Column
}

type binaryCalculator struct {
	zero, one rune
}

func (calc binaryCalculator) resolve(code string, low, high int) (int, int, error) {
	switch {
	case high-low == 1 && rune(code[0]) == calc.zero:
		return low, low, nil
	case high-low == 1 && rune(code[0]) == calc.one:
		return high, high, nil
	case high-low == 1:
		return 0, 0, fmt.Errorf("code %s is neither zero or one", string(code[0]))
	case rune(code[0]) == calc.zero:
		newHigh := (low + high - 1) / 2
		return calc.resolve(code[1:], low, newHigh)
	case rune(code[0]) == calc.one:
		newLow := (low + high + 1) / 2
		return calc.resolve(code[1:], newLow, high)
	default:
		return 0, 0, fmt.Errorf("code %s is neither zero or one", string(code[0]))
	}
}
