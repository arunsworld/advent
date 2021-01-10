package main

import (
	"fmt"
	"io"
)

type groups []group

type group struct {
	size    count
	answers answers
}

type answers map[answer]count

type answer rune

type count int

func (answs answers) register(ans answer) {
	v, exists := answs[answer(ans)]
	switch exists {
	case true:
		answs[answer(ans)] = v + 1
	default:
		answs[answer(ans)] = 1
	}
}

func (c count) matchedSize(comp count) int {
	switch c == comp {
	case true:
		return 1
	default:
		return 0
	}
}

func newGroups(r io.Reader) (groups, error) {
	records, err := splitOnBlankLines(r)
	if err != nil {
		return groups{}, fmt.Errorf("problem parsing input: %v", err)
	}
	result := groups{}
	for _, record := range records {
		g := newGroup(record)
		result = append(result, g)
	}
	return result, nil
}

func newGroup(input []string) group {
	answers := make(answers)
	for _, row := range input {
		for _, ans := range row {
			answers.register(answer(ans))
		}
	}
	return group{
		size:    count(len(input)),
		answers: answers,
	}
}

func (g groups) totalAnyAnswers() int {
	total := 0
	for _, grp := range g {
		total += grp.anyAnswerCount()
	}
	return total
}

func (g group) anyAnswerCount() int {
	return len(g.answers)
}

func (g groups) totalAgreedAnswers() int {
	total := 0
	for _, grp := range g {
		total += grp.agreedAnswerCount()
	}
	return total
}

func (g group) agreedAnswerCount() int {
	total := 0
	for _, count := range g.answers {
		total += count.matchedSize(g.size)
	}
	return total
}
