package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func Sum(line string) int {
	numStarted := false
	start := 0
	sumValues := 0
	for i := 0; i < len(line); i++ {
		if !numStarted {
			if (line[i] >= '0' && line[i] <= '9') || line[i] == '-' {
				numStarted = true
				start = i
			}
		} else if line[i] >= '0' && line[i] <= '9' {

		} else {
			b := line[start:i]
			v, err := strconv.Atoi(b)
			if err != nil {
				panic("shouldnt happen")
			}
			sumValues += v
			numStarted = false
		}
	}
	return sumValues
}

func RecurseSum(tree interface{}) int {
	sum := 0
	switch tree := tree.(type) {
	case bool, string:
	case float64:
		sum += int(tree)
	case []interface{}:
		for _, v := range tree {
			sum += RecurseSum(v)
		}
	case map[string]interface{}:
		ignore := false
		for _, v := range tree {
			st, ok := v.(string)
			if ok && st == "red" {
				ignore = true
			}
			sum += RecurseSum(v)
		}
		if ignore {
			sum = 0
		}
	default:
		panic("shouldn't happen")
	}
	return sum
}

func main() {
	a := Sum(input)
	fmt.Printf("%d\n", a)

	decoder := json.NewDecoder(strings.NewReader(input))
	var tree interface{}
	decoder.Decode(&tree)

	fmt.Printf("%d\n", RecurseSum(tree))
}
