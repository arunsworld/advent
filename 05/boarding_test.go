package main

import (
	"testing"
)

func Test_Boarding_Pass_Codes_Are_Parsed_To_Seats(t *testing.T) {
	type testData struct {
		code string
		id   int
	}
	data := []testData{
		testData{code: "FBFBBFFRLR", id: 357},
		testData{code: "BFFFBBFRRR", id: 567},
		testData{code: "FFFBBBFRRR", id: 119},
		testData{code: "BBFFBBFRLL", id: 820},
	}
	for _, v := range data {
		seat, err := NewSeat(v.code)
		if err != nil {
			t.Fatal(err)
		}
		if seat.identifier() != v.id {
			t.Fatalf("for seat %s id %d didn't match expected %d", v.code, seat.identifier(), v.id)
		}
	}
}
