# Advent of Code in Go

* Solution attempts to https://adventofcode.com/2020
* Each directory is independent of each other for simplicity
* Showcase features of Go documented in `README.md` within each directory

# How to execute

* `go test -v .` to run tests
* `go run .`     to run code

## Code coverage
* `go test -coverprofile=coverage.out`
* `go tool cover -func=coverage.out`
* `go tool cover -html=coverage.out`