package main

import (
	_ "embed"
	"sort"
	"strings"
)

//go:embed "input.txt"
var input string

/// Important
// Assumption only upper case
// and only ASCII chars!
type Replace struct {
	Target      string
	Replacement [2]string
}

func Parse(input string) (map[string]int, []Replace) {
	parts := strings.Split(input, "\n\n")
	replacements := make([]Replace, 0)
	for _, line := range strings.Split(parts[1], "\n") {
		toks := strings.Split(line, " -> ")
		if len(toks[0]) != 2 || len(toks[1]) != 1 {
			panic("invalid input‚")
		}
		z := [2]string{string(toks[0][0]) + string(toks[1][0]), string(toks[1][0]) + string(toks[0][1])}
		replacements = append(replacements, Replace{Target: toks[0], Replacement: z})
	}
	first := make(map[string]int)
	for i := 0; i < len(parts[0])-1; i++ {
		first[parts[0][i:i+2]]++
	}
	first[" "+parts[0][0:1]]++
	first[parts[0][len(parts[0])-1:len(parts[0])]+" "]++
	return first, replacements
}

func Iter(poly map[string]int, replacements []Replace) map[string]int {
	new := make(map[string]int)

	for _, v := range replacements {
		num, ok := poly[v.Target]
		if ok {
			new[v.Replacement[0]] += num
			new[v.Replacement[1]] += num
			delete(poly, v.Target)
		}
	}
	for k, v := range poly {
		new[k] += v
	}
	return new
}

func Length(poly map[string]int) int {
	length := 0
	for _, v := range poly {
		length += v
	}
	return length - 1
}

func Score(poly map[string]int) int {

	count := make(map[rune]int)
	for k, v := range poly {
		for _, ch := range k {
			count[ch] += v
		}
	}
	delete(count, rune(' '))
	countVec := make([]int, 0)
	for k := range count {
		countVec = append(countVec, count[k])
	}
	sort.Ints(countVec)
	return (countVec[len(countVec)-1] - countVec[0]) / 2
}

func main() {
	p, r := Parse(input)
	for i := 0; i < 10; i++ {
		p = Iter(p, r)
	}
	println(Score(p))
	for i := 0; i < 30; i++ {
		p = Iter(p, r)
	}
	println(Score(p))
}
