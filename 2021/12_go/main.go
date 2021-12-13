package main

import (
	_ "embed"
	"regexp"
	"strings"
)

//go:embed "input.txt"
var input string

var RE, _ = regexp.Compile("[[:upper:]]")

func parse(input string) map[string][]string {
	cave := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		toks := strings.SplitN(line, "-", 2)
		cave[toks[0]] = append(cave[toks[0]], toks[1])
		cave[toks[1]] = append(cave[toks[1]], toks[0])
	}
	return cave
}

func descend(cave map[string][]string, smallVisited map[string]struct{}, twice, pos string) uint {
	found := uint(0)
	if pos == "end" {
		return 1
	}
	for _, option := range cave[pos] {
		if option == "start" {
			continue
		}

		if _, ok := smallVisited[option]; ok {
			if twice == "" {
				found += descend(cave, smallVisited, option, option)
			} else {
				continue
			}
		} else {
			if !RE.MatchString(option) {
				smallVisited[option] = struct{}{}
				found += descend(cave, smallVisited, twice, option)
				delete(smallVisited, option)
			} else {
				found += descend(cave, smallVisited, twice, option)
			}
		}
	}
	return found
}

func main() {
	cave := parse(input)
	visited := make(map[string]struct{})
	visited["start"] = struct{}{}
	println(descend(cave, visited, "start", "start"))
	println(descend(cave, visited, "", "start"))
}
