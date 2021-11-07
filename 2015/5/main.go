package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func niceString(st string) bool {
	vowel := 0
	for _, c := range st {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			vowel++
		}
	}
	if vowel < 3 {
		return false
	}
	last := '_'
	double := false
	for _, c := range st {
		if c == last {
			double = true
			break
		}
		last = c
	}
	if !double {
		return false
	}
	a := [4]string{"ab", "cd", "pq", "xy"}
	for _, seq := range a {
		if strings.Contains(st, seq) {
			return false
		}
	}

	return true
}

func newNiceString(st string) bool {
	pair := false
	for i := 0; i < len(st)-3; i++ {
		sub := st[i : i+2]
		rest := st[i+2:]

		if strings.Contains(rest, sub) {
			pair = true
			break
		}
	}
	if !pair {
		return false
	}
	for i := 0; i < len(st)-2; i++ {
		if st[i] == st[i+2] {
			return true
		}
	}
	return false
}

func main() {

	if !niceString("ugknbfddgicrmopn") {
		fmt.Printf("Error")
	}

	if !niceString("aaa") {
		fmt.Printf("Error")
	}

	if niceString("jchzalrnumimnmhp") {
		fmt.Printf("Error")
	}

	if niceString("haegwjzuvuyypxyu") {
		fmt.Printf("Error")
	}

	if niceString("dvszwmarrgswjxmb") {
		fmt.Printf("Error")
	}

	input, _ := os.Open("input.txt")
	line, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal("readall")
	}
	lines := strings.Split(string(line), "\n")
	count := 0
	for _, l := range lines {
		if niceString(l) {
			count++
		}
	}
	fmt.Printf("number %d\n", count)

	if !newNiceString("qjhvhtzxzqqjkmpb") {
		log.Fatal("qjhvhtzxzqqjkmpb")
	}
	if !newNiceString("qthvhtzxzqqjkmpbpb") {
		log.Fatal("qjhvhtzxzqqjkmpb")
	}
	if !newNiceString("xxyxx") {
		log.Fatal("sec")
	}

	if newNiceString("uurcxstgmygtbstg") {
		log.Fatal("thrd")
	}
	if newNiceString("ieodomkazucvgmuy") {
		log.Fatal("fourth")
	}

	if !newNiceString("ieodmkazucvgmuyodo") {
		log.Fatal("fourth")
	}

	count = 0
	for _, l := range lines {
		if newNiceString(l) {
			count++
		}
	}
	fmt.Printf("new number %d\n", count)
}
