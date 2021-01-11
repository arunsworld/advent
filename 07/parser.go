package main

import (
	"bufio"
	"io"
	"strings"
)

func splitLines(r io.Reader) ([]string, error) {
	result := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
