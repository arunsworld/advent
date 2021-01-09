package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// Forest holds details about a forest
type Forest struct {
	Trees         map[Location]struct{}
	Width, Height int
}

// Location are coordinates
type Location struct {
	X, Y int
}

// NewForest creates a new forest from layout
func NewForest(r io.Reader) (Forest, error) {
	csvr := csv.NewReader(r)
	records, err := csvr.ReadAll()
	if err != nil {
		return Forest{}, fmt.Errorf("unable to read layout to generate forest")
	}
	forest := Forest{}
	trees := make(map[Location]struct{})
	for i, rec := range records {
		row := strings.TrimSpace(rec[0])
		w, err := forest.validatedWidth(len(row))
		if err != nil {
			return Forest{}, err
		}
		forest.Width = w
		for j, v := range row {
			if v == '#' {
				trees[Location{X: j, Y: i}] = struct{}{}
			}
		}
	}
	return Forest{
		Trees:  trees,
		Height: len(records),
		Width:  forest.Width,
	}, nil
}

func (f Forest) validatedWidth(width int) (int, error) {
	switch {
	case f.Width == 0:
		return width, nil
	case width == f.Width:
		return width, nil
	}
	return 0, fmt.Errorf("width of forest not consistent. got: %d expected: %d", width, f.Width)
}

// Slope indicates the slop of travel in a top-down direction
type Slope struct {
	Right, Down int
}

// TopLeft is the top-left location of a forest
var TopLeft = Location{X: 0, Y: 0}

// CountTrees counts the trees while traversing the forest at a given slope
func (f Forest) CountTrees(s Slope) int {
	location := TopLeft
	result := 0
	for f.IsLocationInsideForest(location) {
		if f.IsTreeInLocation(location) {
			result++
		}
		location = f.Relocate(location, s)
	}
	return result
}

// IsLocationInsideForest indicates whether the location is still within (not fallen out from bottom)
func (f Forest) IsLocationInsideForest(location Location) bool {
	if location.Y > f.Height-1 {
		return false
	}
	return true
}

// IsTreeInLocation checks if a tree is in the given location
func (f Forest) IsTreeInLocation(location Location) bool {
	_, treeInLocation := f.Trees[location]
	return treeInLocation
}

// Relocate returns a new location from current location using Slope
func (f Forest) Relocate(location Location, s Slope) Location {
	newX := location.X + s.Right
	newY := location.Y + s.Down
	// wrap around X
	if newX > f.Width-1 {
		newX = newX - f.Width
	}
	return Location{X: newX, Y: newY}
}
