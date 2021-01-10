package main

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

func Test_Customs_Declaration_Data_Can_Be_Parsed(t *testing.T) {
	input := `abc

				a
				b
				c
				
				ab
				ac`
	gs, err := newGroups(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	expectedGS := groups{
		group{
			size:    1,
			answers: answers{answer('a'): count(1), answer('b'): count(1), answer('c'): count(1)},
		},
		group{
			size:    3,
			answers: answers{answer('a'): count(1), answer('b'): count(1), answer('c'): count(1)},
		},
		group{
			size:    2,
			answers: answers{answer('a'): count(2), answer('b'): count(1), answer('c'): count(1)},
		},
	}
	if !reflect.DeepEqual(gs, expectedGS) {
		log.Println(gs)
		t.Fatal("groups didn't match expected")
	}
}

func Test_A_Customs_Declaration(t *testing.T) {
	input := `abc

				a
				b
				c
				
				ab
				ac
				
				a
				a
				a
				a
				
				b`
	customs, err := newGroups(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}
	t.Run("can be counted for unique answers that anyone gave", func(t *testing.T) {
		if customs.totalAnyAnswers() != 11 {
			t.Fatalf("expected 11 got: %d", customs.totalAnyAnswers())
		}
	})
	t.Run("can be counted for answers everyone in the group cave", func(t *testing.T) {
		if customs.totalAgreedAnswers() != 6 {
			t.Fatalf("expected 6 got: %d", customs.totalAgreedAnswers())
		}
	})
}
