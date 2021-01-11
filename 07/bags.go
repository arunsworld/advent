package main

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type bags map[color]bag

type bag struct {
	id       color
	contents map[color]count
	within   map[color]struct{} // within which immediate bag can this be contained
}

type color string

type count int

func (bs bags) bagWithColor(c color) bag {
	b, isPresent := bs[c]
	if isPresent {
		return b
	}
	return bag{
		id:     c,
		within: make(map[color]struct{}),
	}
}

// updateWithin updates bags contained within a parent bag with this info
func (bs bags) updateWithin(within color, contents map[color]count) {
	for c := range contents {
		b := bs.bagWithColor(c)
		b.within[within] = struct{}{}
		bs[c] = b
	}
}

func (bs bags) allContainersOf(c color) map[color]struct{} {
	bag := bs[c]
	if len(bag.within) == 0 {
		// if a bag isn't within anything then it's the top level container and contains the bag
		return map[color]struct{}{
			c: struct{}{},
		}
	}
	result := map[color]struct{}{}
	for immediateC := range bag.within {
		result[immediateC] = struct{}{} // bag is within it's immediate container
		// and all containers of it
		for c := range bs.allContainersOf(immediateC) {
			result[c] = struct{}{}
		}
	}
	return result
}

func (bs bags) allContentsOf(c color) int {
	bag := bs[c]
	if len(bag.contents) == 0 {
		// bag is empty of other bags
		return 0
	}
	result := 0
	for c, cnt := range bag.contents {
		result += int(cnt)
		result += int(cnt) * bs.allContentsOf(c)
	}
	return result
}

func newBags(r io.Reader) (bags, error) {
	lines, err := splitLines(r)
	if err != nil {
		return nil, fmt.Errorf("problem parsing rules as individual lines: %v", err)
	}
	bags := make(bags)
	for _, line := range lines {
		c, contents, err := parseRule(line)
		if err != nil {
			return nil, fmt.Errorf("problem parsing %s: %v", line, err)
		}
		b := bags.bagWithColor(c)
		if len(b.contents) > 0 {
			return nil, fmt.Errorf("bag with color: %s duplicated", c)
		}
		b.contents = contents
		bags[c] = b
		bags.updateWithin(c, contents) // update all the bags in contents to be within bag with color c
	}
	return bags, nil
}

func parseRule(rule string) (color, map[color]count, error) {
	if c, isTerminal := isTerminalBag(rule); isTerminal {
		return c, nil, nil
	}
	return parseAsRegularBag(rule)
}

var noOtherBags = regexp.MustCompile(`(.*?) bags contain no other bags.`)

func isTerminalBag(rule string) (color, bool) {
	result := noOtherBags.FindAllStringSubmatch(rule, -1)
	if result == nil {
		return "", false
	}
	return color(result[0][1]), true
}

var bagParser = regexp.MustCompile(`(.*?) bag[s]?`)
var bagParserWithCount = regexp.MustCompile(`(\d+) (.*?) bag[s]?`)

func parseAsRegularBag(rule string) (color, map[color]count, error) {
	ruleSplit := strings.Split(rule, " contain ")
	if len(ruleSplit) != 2 {
		return "", nil, fmt.Errorf("rule doesn't include the string contain")
	}
	mainBag := strings.TrimSpace(ruleSplit[0])
	result := bagParser.FindAllStringSubmatch(mainBag, -1)
	if result == nil {
		return "", nil, fmt.Errorf("%s couldn't be parsed as a bag", mainBag)
	}
	mainBagColor := color(result[0][1])

	contentLines := strings.Split(ruleSplit[1], ",")
	contents := make(map[color]count)
	for _, line := range contentLines {
		line = strings.TrimSpace(line)
		result := bagParserWithCount.FindAllStringSubmatch(line, -1)
		if result == nil {
			return "", nil, fmt.Errorf("%s couldn't be parsed as a bag with count", line)
		}
		c, err := strconv.Atoi(result[0][1])
		if err != nil {
			return "", nil, fmt.Errorf("bag count not a integer in %s", line)
		}
		contents[color(result[0][2])] = count(c)
	}

	return mainBagColor, contents, nil
}
