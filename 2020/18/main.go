package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func token(st string) []string {
	st = strings.Replace(st, "(", "( ", -1)
	st = strings.Replace(st, ")", " )", -1)

	return strings.Fields(st)
}

func eval(toks []string) int {
	lvalue := 0
	lpos := 1
	if toks[0] == "(" {
		braces := 1
	Loop:
		for i := 1; i < len(toks); i++ {
			switch toks[i] {
			case "(":
				braces++
			case ")":
				braces--
				if braces <= 0 {
					lvalue = eval(toks[1:i])
					lpos = i + 1
					break Loop
				}
			}
		}
	} else {
		lvalue, _ = strconv.Atoi(toks[0])
	}
	for {
		if len(toks) == lpos {
			return lvalue
		}
		op := toks[lpos]

		rvalue := 0
		if toks[lpos+1] == "(" {
			braces := 0
		OtherLoop:
			for i := lpos + 2; i < len(toks); i++ {
				switch toks[i] {
				case "(":
					braces++
				case ")":
					braces--
					if braces < 0 {
						rvalue = eval(toks[lpos+2 : i])
						lpos = i + 1
						break OtherLoop
					}
				}
			}
		} else {
			rvalue, _ = strconv.Atoi(toks[lpos+1])
			lpos = lpos + 2
		}
		switch op {
		case "+":
			lvalue = lvalue + rvalue
		case "*":
			lvalue = lvalue * rvalue
		default:
			panic("op unknown")
		}
	}
}

func main() {
	type results struct {
		st  string
		res int
	}
	r := []results{{st: "2 * 3 + (4 * 5)", res: 26},
		{st: "5 + (8 * 3 + 9 + 3 * 4 * 3)", res: 437},
		{st: "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", res: 13632}}
	for _, k := range r {
		if eval(token(k.st)) != k.res {
			fmt.Printf("%s = %d should be %dn", k.st, eval(token(k.st)), k.res)
		}
	}
	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	lines := strings.Split(string(buf), "\n")
	sum := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		sum += eval(token(line))
	}
	println(sum)
}
