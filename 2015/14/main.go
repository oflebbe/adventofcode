package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Record struct {
	Name        string
	Speed       int
	SpeedLength int
	RestLength  int
}

func parse(lines string) []Record {
	// Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
	result := []Record{}
	for _, line := range strings.Split(input, "\n") {
		var name string
		var speed int
		var speedLength int
		var restLength int
		arg, err := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &speedLength, &restLength)
		if arg != 4 || err != nil {
			break
		}
		result = append(result, Record{Name: name, Speed: speed, SpeedLength: speedLength, RestLength: restLength})
	}
	return result
}

func evalDistance(time int, r Record) int {
	period := r.SpeedLength + r.RestLength
	times := time / period
	remaining := time % period

	regularDistance := times * r.SpeedLength * r.Speed
	if remaining > r.SpeedLength {
		remaining = r.SpeedLength
	}
	return regularDistance + remaining*r.Speed
}

func bestDistance(time int, r []Record) int {
	best := 0
	for _, reindeer := range r {
		s := evalDistance(time, reindeer)
		if s > best {
			best = s
		}
	}
	return best
}

func scoreDistance(untilTime int, r []Record) int {
	score := make(map[int]int)
	for t := 1; t <= untilTime; t++ {
		best := -1
		s := make([]int, len(r))
		for w, reindeer := range r {
			s[w] = evalDistance(t, reindeer)
			if s[w] > best {
				best = s[w]
			}
		}
		for w := 0; w < len(r); w++ {
			if s[w] == best {
				score[w]++
			}
		}
	}
	best := -1
	for _, v := range score {
		if v > best {
			best = v
		}
	}
	return best
}

func main() {
	fmt.Printf("%d\n", bestDistance(2503, parse(input)))

	fmt.Printf("%d\n", scoreDistance(2503, parse(input)))
}
