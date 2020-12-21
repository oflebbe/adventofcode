package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type meal struct {
	ingredients []string
	allergens   []string
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
		m := meal{ingredients: strings.Fields(sides[0]), allergens: strings.Split(sides[1][:len(sides[1])-1], ", ")}
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

func solve(meals []meal) []meal {
	allergens := make(map[string]string)
	for _, meal := range meals {
		for _, al := range meal.allergens {
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
				if stringInSlice(al, m.allergens) {
					for _, i := range m.ingredients {
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
				allergens[al] = potential
				println(al)

				for i := 0; i < len(meals); i++ {
					meals[i].ingredients = removeFromSlice(meals[i].ingredients, potential)
					meals[i].allergens = removeFromSlice(meals[i].allergens, al)
				}
			}
		}
	}
	var newmeals []meal
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
			if stringInSlice(al, m.allergens) {
				for _, i := range m.ingredients {
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
		var newingredients []string
		for _, i := range m.ingredients {
			_, ok := potentials[i]
			if ok {
				newingredients = append(newingredients, i)
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
				ifound = m.ingredients[0]
				afound = m.allergens[0]
				allergens[afound] = ifound
				break
			}
		}
		if ifound != "" {
			for i := 0; i < len(meals); i++ {
				meals[i].ingredients = removeFromSlice(meals[i].ingredients, ifound)
				meals[i].allergens = removeFromSlice(meals[i].allergens, afound)
			}
			// Now return empty allergens
			newmeals = []meal{}
			for i := 0; i < len(meals); i++ {
				if len(meals[i].allergens) != 0 {
					newmeals = append(newmeals, meals[i])
				}
			}
		}

	}
	return newmeals
}

func main() {
	fh, _ := os.Open("input.txt")
	buf, _ := ioutil.ReadAll(fh)
	m := readin(string(buf))
	rest := solve(m)
	for _, m := range rest {
		fmt.Printf("%+v\n", m)
	}

}
