package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Rule interface {
	Matches(st []uint8) (int, error)
}

type TermRule struct {
	num  int
	char uint8
}

func (r TermRule) Matches(in []uint8) (int, error) {
	if len(in) == 0 {
		return 0, errors.New("eol")
	}
	if in[0] == r.char {
		return 1, nil
	} else {
		return 0, errors.New("bla")
	}
}

type ConcatRule struct {
	num        int
	sub_rule_1 []*Rule
}

func matches(in []uint8, rules []*Rule) (int, error) {
	if len(in) == 0 {
		return 0, errors.New("TOO")
	}
	adv := 0
	tin := in
	var err error
	sum := 0
	for _, r := range rules {
		tin = tin[adv:]
		adv, err = (*r).Matches(tin)
		sum += adv
		if err != nil {
			break
		}
	}
	if err != nil {
		return 0, err
	}
	return sum, nil
}

func (r ConcatRule) Matches(in []uint8) (int, error) {
	return matches(in, r.sub_rule_1)
}

type AlternativeRule struct {
	num        int
	sub_rule_1 []*Rule
	sub_rule_2 []*Rule
}

func (r AlternativeRule) Matches(in []uint8) (int, error) {
	adv1, err1 := matches(in, r.sub_rule_1)
	adv2, err2 := matches(in, r.sub_rule_2)

	if err1 != nil && err2 != nil {
		return 0, err1
	}
	if err1 == nil && err2 == nil {
		if adv1 != adv2 {
			panic("ungleich")
		}
	}
	if err1 == nil {
		return adv1, nil
	}
	return adv2, nil
}

func CreateRules(input string) Rule {
	stage := make(map[int]Rule) // xx : "a b", "c d"
	for _, line := range strings.Split(input, "\n") {
		tok := strings.SplitN(line, ": ", 2)
		num, _ := strconv.Atoi(tok[0])
		rules := strings.SplitN(tok[1], " | ", 2)
		if len(rules) == 2 {
			stage[num] = &AlternativeRule{num: num}
		} else if rules[0][0] == '"' {
			stage[num] = &TermRule{num: num}
		} else {
			stage[num] = &ConcatRule{num: num}
		}
	}
	for _, line := range strings.Split(input, "\n") {
		tok := strings.SplitN(line, ": ", 2)
		num, _ := strconv.Atoi(tok[0])
		r := stage[num]
		var a, b, c, d, e int
		n, err := fmt.Sscanf(tok[1], "%d %d | %d %d %d", &a, &b, &c, &d, &e)
		if err == nil && n == 5 {
			alter := r.(*AlternativeRule)
			ar := stage[a]
			br := stage[b]
			alter.sub_rule_1 = []*Rule{&ar, &br}
			cr := stage[c]
			dr := stage[d]
			er := stage[e]
			alter.sub_rule_2 = []*Rule{&cr, &dr, &er}
			continue
		}
		n, err := fmt.Sscanf(tok[1], "%d %d | %d %d", &a, &b, &c, &d)
		if err == nil && n == 4 {
			alter := r.(*AlternativeRule)
			ar := stage[a]
			br := stage[b]
			alter.sub_rule_1 = []*Rule{&ar, &br}
			cr := stage[c]
			dr := stage[d]
			alter.sub_rule_2 = []*Rule{&cr, &dr}
			continue
		}
		n, err = fmt.Sscanf(tok[1], "%d | %d %d", &a, &b, &c)
		if err == nil && n == 2 {
			alter := r.(*AlternativeRule)
			ar := stage[a]
			alter.sub_rule_1 = []*Rule{&ar}
			br := stage[b]
			cr := stage[c]
			alter.sub_rule_2 = []*Rule{&br, &cr}
			continue
		}
		n, err = fmt.Sscanf(tok[1], "%d | %d", &a, &b)
		if err == nil && n == 2 {
			alter := r.(*AlternativeRule)
			ar := stage[a]
			alter.sub_rule_1 = []*Rule{&ar}
			br := stage[b]
			alter.sub_rule_2 = []*Rule{&br}
			continue
		}
		n, err = fmt.Sscanf(tok[1], "%d %d %d", &a, &b, &c)
		if err == nil && n == 3 {
			ar := stage[a]
			br := stage[b]
			cr := stage[c]
			conc := r.(*ConcatRule)
			conc.sub_rule_1 = []*Rule{&ar, &br, &cr}
			continue
		}
		n, err = fmt.Sscanf(tok[1], "%d %d", &a, &b)
		if err == nil && n == 2 {
			ar := stage[a]
			br := stage[b]
			conc := r.(*ConcatRule)
			conc.sub_rule_1 = []*Rule{&ar, &br}
			continue
		}
		n, err = fmt.Sscanf(tok[1], "%d", &a)
		if err == nil && n == 1 {
			ar := stage[a]
			conc := r.(*ConcatRule)
			conc.sub_rule_1 = []*Rule{&ar}
			continue
		}
		term := r.(*TermRule)
		term.char = tok[1][1]
	}
	return stage[0]
}

func parseInput(st string) (Rule, string) {
	parts := strings.Split(st, "\n\n")
	zero := CreateRules(parts[0])
	text := parts[1]
	return zero, text
}

func testLines(zero Rule, input string) int {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		buf := []uint8(line)
		l, err := zero.Matches(buf)
		if err == nil && l == len(buf) {
			count++
		}
	}
	return count
}

func main() {
	/*	input := `0: 1 2
		1: "a"
		2: 1 3 | 3 1
		3: "b"

		aab
		aba
		aaa`
			zero, text := parseInput(string(input))
			println(testLines(zero, text))
	*/
	/* input := `0: 4 1 5
	1: 2 3 | 3 2
	2: 4 4 | 5 5
	3: 4 5 | 5 4
	4: "a"
	5: "b"

	ababbb
	bababa
	abbbab
	aaabbb
	aaaabbb`
	zero, text := parseInput(string(input))
	println(testLines(zero, text))*/

	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	zero, text := parseInput(string(buf))
	println(testLines(zero, text))

	fh, _ = os.Open("modinput.txt")
	buf, _ = ioutil.ReadAll(fh)
	zero, text = parseInput(string(buf))
	println(testLines(zero, text))

}
