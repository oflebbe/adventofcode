package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func lookup(rules map[string][]string, input string) map[string]interface{} {
	c := make(map[string]interface{})
	for _, v := range rules[input] {
		c[v] = nil
		c2 := lookup(rules, v)
		for k2 := range c2 {
			c[k2] = nil
		}
	}
	return c
}

type pair struct {
	quantity int
	what     string
}

func lookup2(rules map[string][]pair, input string) int {
	c := 1
	for _, v := range rules[input] {
		c += v.quantity * lookup2(rules, v.what)
	}
	return c
}

func problem(input string) (int, int) {
	lines := strings.Split(input, "\n")
	rulesR := make(map[string][]string)
	rulesL := make(map[string][]pair)
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Index(line, "bags contain no other bags.") != -1 {
			continue
		}
		leftRight := strings.Split(line, " bags contain ")
		right := regexp.MustCompile(" bags, | bags.| bag, | bag.").Split(leftRight[1], -1)

		for _, r := range right {
			if r == "" {
				continue
			}
			colors := strings.SplitN(r, " ", 2)
			quant, _ := strconv.Atoi(colors[0])
			rulesR[colors[1]] = append(rulesR[colors[1]], leftRight[0])
			rulesL[leftRight[0]] = append(rulesL[leftRight[0]], pair{quant, colors[1]})
		}
	}

	return len(lookup(rulesR, "shiny gold")), lookup2(rulesL, "shiny gold") - 1

}
func main() {
	test := `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`
	println(problem(string(test)))

	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	fh.Close()

	println(problem(string(buf)))
}
