package main

import "testing"

var testInput = `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`

func TestParse(t *testing.T) {
	li := parse(testInput)
	if len(li) != 2 {
		t.Errorf("zwei")
	}
	if li[0][capacity-1] != -1 || li[1][flavor-1] != -2 {
		t.Errorf("values")
	}

	score := eval(li, []int{44, 56})
	if score != 62842880 {
		t.Errorf("example score")
	}

	best := findIngredients(li, []int{0, 0}, 0)
	if best != 62842880 {
		t.Errorf("example inge")
	}

}
