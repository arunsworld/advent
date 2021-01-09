package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_A_Small_Forest(t *testing.T) {
	forestLayout := `..##.......
				     #...#...#..
				     .#....#..#.`
	forest, err := NewForest(strings.NewReader(forestLayout))
	if err != nil {
		t.Fatal(err)
	}
	expectedForest := Forest{
		Width: 11, Height: 3,
		Trees: map[Location]struct{}{
			Location{X: 2, Y: 0}: struct{}{},
			Location{X: 3, Y: 0}: struct{}{},

			Location{X: 0, Y: 1}: struct{}{},
			Location{X: 4, Y: 1}: struct{}{},
			Location{X: 8, Y: 1}: struct{}{},

			Location{X: 1, Y: 2}: struct{}{},
			Location{X: 6, Y: 2}: struct{}{},
			Location{X: 9, Y: 2}: struct{}{},
		},
	}
	if !reflect.DeepEqual(forest, expectedForest) {
		t.Fatal("forest not as expected")
	}
	t.Run("given a slope", func(t *testing.T) {
		slope := Slope{Right: 3, Down: 1}
		t.Run("traces a path and is able to count the trees in it", func(t *testing.T) {
			trees := forest.CountTrees(slope)
			if trees != 1 {
				t.Fatalf("expected 1 tree, got: %d", trees)
			}

		})
	})
}

func Test_A_Bigger_Forest(t *testing.T) {
	fullForest := `..##.......
							#...#...#..
							.#....#..#.
							..#.#...#.#
							.#...##..#.
							..#.##.....
							.#.#.#....#
							.#........#
							#.##...#...
							#...##....#
							.#..#...#.#`
	forest, err := NewForest(strings.NewReader(fullForest))
	if err != nil {
		t.Fatal(err)
	}
	t.Run("is able to count trees for multiple slopes", func(t *testing.T) {
		type testData struct {
			slope Slope
			count int
		}
		data := []testData{
			testData{slope: Slope{Right: 1, Down: 1}, count: 2},
			testData{slope: Slope{Right: 3, Down: 1}, count: 7},
			testData{slope: Slope{Right: 5, Down: 1}, count: 3},
			testData{slope: Slope{Right: 7, Down: 1}, count: 4},
			testData{slope: Slope{Right: 1, Down: 2}, count: 2},
		}
		for _, td := range data {
			c := forest.CountTrees(td.slope)
			if td.count != c {
				t.Fatalf("count of trees for slope %v not as expected. expected: %d. got: %d", td.slope, td.count, c)
			}
		}
	})
}
