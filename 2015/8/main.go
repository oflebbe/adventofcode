package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	NIX = iota
	ESCAPE
	EINS
	ZWEI
)

func encodeLine(st string) int {
	l := len(st) + 6
	cc := strings.Count(st, "\\")
	cc += strings.Count(st, "\"")
	return l + cc

}

func handleLine(st string) (int, int) {
	st = st[1 : len(st)-1]
	val := ""
	state := NIX
	result := ""
	for _, c := range st {
		if state == ESCAPE {
			if c == '\\' {
				result += string(c)
				state = NIX
			} else if c == '"' {
				result += string(c)
				state = NIX
			} else if c == 'x' {
				state = EINS
			} else {
				log.Fatal("esc")
			}
		} else if state == EINS {
			val += string(c)
			state = ZWEI
		} else if state == ZWEI {
			val += string(c)
			var w [1]byte
			fmt.Sscanf(val, "%x", &w[0])
			if 1 != len(string(w[:])) {
				log.Fatalf("len 1 %d x%sx", w, string(w[:]))
			}
			result += string(w[:])
			val = ""
			state = NIX
		} else if state == NIX {
			if c == '\\' {
				state = ESCAPE
			} else {
				result += string(c)
			}
		} else {
			log.Fatal("state")
		}
	}
	fmt.Printf("%d\n", encodeLine(st))
	return len(st) + 2 - len(result), encodeLine(st) - (len(st) + 2)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("open")
	}
	in, _ := ioutil.ReadAll(f)
	count := 0
	other := 0
	for _, s := range strings.Split(string(in), "\n") {
		a, b := handleLine(s)
		count += a
		other += b
	}
	fmt.Printf("%d %d\n", count, other)
}
