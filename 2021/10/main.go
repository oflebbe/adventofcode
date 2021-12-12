package main

import (
	_ "embed"
	"sort"
	"strings"
)

//go:embed "input.txt"
var input string

type ScoreList struct {
	Rune   rune
	Points int
}

var runeMap map[rune]ScoreList

var autocompleteMap map[rune]int

func init() {
	runeMap = make(map[rune]ScoreList)
	runeMap[')'] = ScoreList{Rune: '(', Points: 3}
	runeMap['>'] = ScoreList{Rune: '<', Points: 25137}
	runeMap[']'] = ScoreList{Rune: '[', Points: 57}
	runeMap['}'] = ScoreList{Rune: '{', Points: 1197}

	autocompleteMap = make(map[rune]int)
	autocompleteMap['('] = 1
	autocompleteMap['<'] = 4
	autocompleteMap['['] = 2
	autocompleteMap['{'] = 3
}

func score(line string) int {
	stack := []rune{}

	for _, r := range line {

		switch r {
		case '(', '<', '{', '[':
			stack = append(stack, r)
		case ')', '>', '}', ']':
			last, ok := runeMap[r]
			if !ok {
				panic("wat")
			}
			if last.Rune != stack[len(stack)-1] {
				println(string(r))
				return last.Points
			}
			stack = stack[:len(stack)-1]
		default:
			panic("unknwn char")
		}
	}
	if len(stack) != 0 {
		return 0
	}
	panic("should not happen")
}

func autoCompleteScore(line string) int {
	stack := []rune{}

	for _, r := range line {

		switch r {
		case '(', '<', '{', '[':
			stack = append(stack, r)
		case ')', '>', '}', ']':
			last, ok := runeMap[r]
			if !ok {
				panic("wat")
			}
			if last.Rune != stack[len(stack)-1] {
				return 0
			}
			stack = stack[:len(stack)-1]
		default:
			panic("unknown char")
		}
	}
	if len(stack) == 0 {
		return 0
	}
	score := 0
	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		score  += autocompleteMap[stack[i]]
	}

	return score
}

func parse(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		sum += score(line)
	}
	return sum
}

func parse2(input string) int {
	sum := []int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		s := autoCompleteScore(line)
		if s == 0 {
			continue
		}
		sum = append(sum, s)
	}
	sort.Ints(sum)
	if len(sum)%2 == 0 {
		panic("even")
	}
	return sum[len(sum)/2]
}

func main() {
	println(parse(input))
	println(parse2(input))

}
