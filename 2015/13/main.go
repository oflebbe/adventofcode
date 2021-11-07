package main

import (
	"container/ring"
	_ "embed"
	"fmt"
	"log"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

type Pair struct {
	A string
	B string
}

func Parse(st string) map[Pair]int {
	st = strings.ReplaceAll(st, ".", "")
	rules := make(map[Pair]int)
	lines := strings.Split(st, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var A, B, gainLoss string
		var happy int
		num, err := fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &A, &gainLoss, &happy, &B)

		if num != 4 || err != nil {
			log.Fatalf("Error sscanf %d %v", num, err)
		}
		if gainLoss == "lose" {
			happy = -happy
		}
		rules[Pair{A: A, B: B}] = happy
	}
	return rules
}

func ReturnNames(rules map[Pair]int) []string {
	names := make(map[string]struct{})
	for k, _ := range rules {
		names[k.A] = struct{}{}
		names[k.B] = struct{}{}
	}
	nameList := make([]string, 0)
	for n, _ := range names {
		nameList = append(nameList, n)
	}
	sort.Strings(nameList)
	return nameList
}

func CreateRing(names []string) *ring.Ring {
	N := len(names)
	r := ring.New(len(names))
	for i := 0; i < N; i++ {
		r.Value = names[i]
		r = r.Next()
	}
	return r
}

func Eval(r *ring.Ring, rules map[Pair]int) int {
	sum := 0
	for i := 0; i < r.Len(); i++ {
		sum += rules[Pair{A: r.Value.(string), B: r.Prev().Value.(string)}]
		sum += rules[Pair{A: r.Value.(string), B: r.Next().Value.(string)}]
		r = r.Next()
	}
	return sum
}

func Permutations(depth int, num []int, Call func(nums []int)) {
	N := len(num)
	for i := 0; i < N; i++ {
		iFound := false
		for j := 0; j < depth; j++ {
			if num[j] == i {
				iFound = true
				break
			}
		}
		if !iFound {
			num[depth] = i
			if depth == N-1 {
				Call(num)
			} else {
				Permutations(depth+1, num, Call)
			}
		}
	}
}

type Storage struct {
	names       []string
	rules       map[Pair]int
	best        int
	bestSeating []string
}

func (s *Storage) EvalPermutation(places []int) {
	table := make([]string, len(s.names))
	for i := 0; i < len(places); i++ {
		table[i] = s.names[places[i]]
	}
	val := Eval(CreateRing(table), s.rules)
	if val > s.best {
		s.best = val
		s.bestSeating = table
	}
}

func RunAll(names []string, rules map[Pair]int) int {
	s := Storage{names: names, rules: rules, best: 0}
	Permutations(0, make([]int, len(names)), s.EvalPermutation)
	fmt.Printf("Best Seating %+v", s.bestSeating)
	return s.best
}

func main() {
	r := Parse(input)
	names := ReturnNames(r)

	b := RunAll(names, r)
	fmt.Printf("%d\n", b)

	names = append(names, "olaf")
	b = RunAll(names, r)
	fmt.Printf("%d\n", b)

}
