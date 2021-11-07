package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

const (
	capacity = iota + 1
	durability
	flavor
	texture
	calories
)

var re = regexp.MustCompile(`\w+: capacity (-?\d+), durability (-?\d+), flavor (-?\d+), texture (-?\d+), calories (-?\d+)`)

func parse(st string) [][]int {
	ret := [][]int{}

	for _, line := range re.FindAllStringSubmatch(st, -1) {
		prop := make([]int, calories)
		for i := capacity; i <= calories; i++ {
			var err error
			prop[i-1], err = strconv.Atoi(line[i])
			if err != nil {
				panic(err)
			}
		}
		ret = append(ret, prop)
	}
	return ret
}

func eval(prop [][]int, qu []int) int {
	score := 1
	for p := 0; p < texture; p++ {
		sum := 0
		for ingred := 0; ingred < len(prop); ingred++ {
			sum += prop[ingred][p] * qu[ingred]
		}
		if sum < 0 {
			sum = 0
		}
		score *= sum
	}
	return score
}

func evalCal500(prop [][]int, qu []int) int {

	score := 1
	for p := 0; p < calories; p++ {
		sum := 0
		for ingred := 0; ingred < len(prop); ingred++ {
			sum += prop[ingred][p] * qu[ingred]
		}
		if sum < 0 {
			sum = 0
		}
		if p == calories-1 {
			if sum == 500 {
				sum = 1
			} else {
				sum = 0
			}
		}
		score *= sum
	}

	return score
}

func findIngredients(prop [][]int, teaspoons []int, level int, f func([][]int, []int) int) int {
	num := len(prop)
	alreadySum := 0
	best := 0
	for i := 0; i < level; i++ {
		alreadySum += teaspoons[i]
	}
	if level == num-1 {
		teaspoons[level] = 100 - alreadySum
		val := findIngredients(prop, teaspoons, level+1, f)
		return val
	}
	if level == num {
		return f(prop, teaspoons)
	}
	for i := 1; i < 100-alreadySum; i++ {
		teaspoons[level] = i
		val := findIngredients(prop, teaspoons, level+1, f)
		if val > best {
			best = val
		}
	}
	return best
}

func main() {
	para := parse(input)
	start := make([]int, len(para))
	val := findIngredients(para, start, 0, eval)
	fmt.Printf("%d\n", val)

	val = findIngredients(para, start, 0, evalCal500)
	fmt.Printf("%d\n", val)
}
