package main

import (
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type meal struct {
	ingredients map[string]interface{}
	allergens   map[string]interface{}
}

// mapping ingredient allergen
type dict map[string]string

func readin(lines string) []meal {
	var meals []meal
	for _, line := range strings.Split(lines, "\n") {
		if line == "" {
			continue
		}
		sides := strings.Split(line, "(contains ")
		var m meal
		m.ingredients = make(map[string]interface{})
		m.allergens = make(map[string]interface{})
		for _, k := range strings.Fields(sides[0]) {
			m.ingredients[k] = nil
		}
		for _, k := range strings.Split(sides[1][:len(sides[1])-1], ", ") {
			m.allergens[k] = nil
		}
		meals = append(meals, m)
	}
	return meals
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func removeFromSlice(s []string, el string) []string {
	var i int
	var b string
	for i, b = range s {
		if b == el {
			s[len(s)-1], s[i] = s[i], s[len(s)-1]
			return s[:len(s)-1]
		}
	}

	return s
}

func solve(meals []meal) (int, string) {

	allergens := make(map[string]string)
	for _, meal := range meals {
		for al := range meal.allergens {
			allergens[al] = ""
		}
	}

	found := true
	for found {
		found = false

		for al := range allergens {
			count := make(map[string]int)
			counter := 0
			var potential string
			if allergens[al] != "" {
				continue
			}

			for _, m := range meals {
				_, ok := m.allergens[al]
				if ok {
					for i := range m.ingredients {
						count[i]++
					}
					counter++
				}
			}
			for k, v := range count {
				if v == counter {
					if potential == "" {
						// found potential allergen producing ingredient
						potential = k
					} else {
						// found other potential allergen producing ingredient
						potential = "impossible!"
						break
					}
				}
			}

			if potential != "" && potential != "impossible!" {
				found = true
				allergens[al] = potential
				println(al)

				for i := 0; i < len(meals); i++ {
					delete(meals[i].ingredients, potential)
					delete(meals[i].allergens, al)
				}
			}
		}
	}
	/*	var newmeals []meal
		// Now return empty allergens
		for i := 0; i < len(meals); i++ {
			if len(meals[i].allergens) != 0 {
				newmeals = append(newmeals, meals[i])
			}
		}
		meals = newmeals

		// mark all potentials
		potentials := make(map[string]interface{})
		for al := range allergens {
			count := make(map[string]int)
			counter := 0
			if allergens[al] != "" {
				continue
			}

			for _, m := range meals {
				_, ok := m.allergens[al]
				if ok {
					for i := range m.ingredients {
						count[i]++
					}
					counter++
				}
			}
			for k, v := range count {
				if v == counter {
					potentials[k] = nil
				}
			}
		}

		newmeals = []meal{}
		// remove everything else from meals
		for _, m := range meals {
			newingredients := make(map[string]interface{})
			for i := range m.ingredients {
				_, ok := potentials[i]
				if ok {
					newingredients[i] = nil
				}
			}
			newmeals = append(newmeals, meal{ingredients: newingredients, allergens: m.allergens})
		}
		meals = newmeals

		for _, m := range newmeals {
			fmt.Printf("%+v\n", m)
		}

		for len(meals) > 0 {
			ifound := ""
			afound := ""
			for _, m := range meals {
				if len(m.allergens) == 1 && len(m.ingredients) == 1 {
					for ifound = range m.ingredients {
						break
					}
					for afound = range m.allergens {
						break
					}
					allergens[afound] = ifound
					break
				}
			}
			if ifound != "" {
				for i := 0; i < len(meals); i++ {
					delete(meals[i].ingredients, ifound)
					delete(meals[i].allergens, afound)
				}
				// Now remove empty allergens
				newmeals = []meal{}
				for i := 0; i < len(meals); i++ {
					if len(meals[i].allergens) != 0 {
						newmeals = append(newmeals, meals[i])
					}
				}
			}

		}
		return newmeals*/
	counter := 0
	for _, m := range meals {
		counter += len(m.ingredients)
	}
	type pair struct {
		i string
		a string
	}
	var canonical []pair
	for k, v := range allergens {
		canonical = append(canonical, pair{i: v, a: k})
	}
	sort.Slice(canonical, func(i, j int) bool {
		return canonical[i].a < canonical[j].a
	})
	st := ""
	for _, i := range canonical {
		if st != "" {
			st += ","
		}
		st += i.i
	}
	return counter, st
}

func main() {
	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	m := readin(string(buf))
	rest, can := solve(m)
	println(rest, can)
}
