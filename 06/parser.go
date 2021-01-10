package main

import (
	"bufio"
	"io"
	"strings"
)

func splitOnBlankLines(r io.Reader) ([][]string, error) {
	result := [][]string{}
	scanner := bufio.NewScanner(r)
	buffer := []string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		switch {
		case line == "" && len(buffer) > 0:
			result = append(result, buffer)
			buffer = []string{}
		case line == "": // ignore multiple entirely blank lines
		default:
			buffer = append(buffer, line)
		}
	}
	if len(buffer) > 0 {
		result = append(result, buffer)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
