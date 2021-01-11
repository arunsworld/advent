package main

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

func Test_Bag_Rules_Can_Be_Parsed(t *testing.T) {
	rules := `light red bags contain 1 bright white bag, 2 muted yellow bags.
				dark orange bags contain 3 bright white bags, 4 muted yellow bags.`
	expected := bags{
		"light red": bag{
			id:       "light red",
			contents: map[color]count{"bright white": 1, "muted yellow": 2},
			within:   map[color]struct{}{},
		},
		"bright white": bag{
			id:     "bright white",
			within: map[color]struct{}{"light red": struct{}{}, "dark orange": struct{}{}},
		},
		"muted yellow": bag{
			id:     "muted yellow",
			within: map[color]struct{}{"light red": struct{}{}, "dark orange": struct{}{}},
		},
		"dark orange": bag{
			id:       "dark orange",
			contents: map[color]count{"bright white": 3, "muted yellow": 4},
			within:   map[color]struct{}{},
		},
	}
	bags, err := newBags(strings.NewReader(rules))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(bags, expected) {
		log.Println(bags)
		t.Fatal("bags not as expected")
	}
}

func Test_Within_Bags(t *testing.T) {
	bagRules := `light red bags contain 1 bright white bag, 2 muted yellow bags.
					dark orange bags contain 3 bright white bags, 4 muted yellow bags.
					bright white bags contain 1 shiny gold bag.
					muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
					shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
					dark olive bags contain 3 faded blue bags, 4 dotted black bags.
					vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
					faded blue bags contain no other bags.
					dotted black bags contain no other bags.`
	bags, err := newBags(strings.NewReader(bagRules))
	if err != nil {
		t.Fatal(err)
	}
	t.Run("we can retrieve bag colors that can eventually contain a given bag color", func(t *testing.T) {
		result := bags.allContainersOf("shiny gold")
		if len(result) != 4 {
			log.Println(result)
			t.Fatalf("expected count to be 4, got: %d", len(result))
		}
	})
	bagRules = `shiny gold bags contain 2 dark red bags.
				dark red bags contain 2 dark orange bags.
				dark orange bags contain 2 dark yellow bags.
				dark yellow bags contain 2 dark green bags.
				dark green bags contain 2 dark blue bags.
				dark blue bags contain 2 dark violet bags.
				dark violet bags contain no other bags.`
	bags, err = newBags(strings.NewReader(bagRules))
	if err != nil {
		t.Fatal(err)
	}
	t.Run("we can count how many bags a given bag must contain", func(t *testing.T) {
		result := bags.allContentsOf("shiny gold")
		if result != 126 {
			t.Fatalf("expected count to be 126, got: %d", result)
		}
	})
}
