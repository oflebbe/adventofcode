package main

import (
	_ "embed"
	"testing"
)

//go:embed testinput.txt
var testinput string

func TestParse(t *testing.T) {
	r := Parse("Alice would gain 54 happiness units by sitting next to Bob.")
	if len(r) != 1 {
		t.Errorf("expected len to be 1 is %d", len(r))
	}
	if r[Pair{A: "Alice", B: "Bob"}] != 54 {
		t.Errorf("expected len to be 54 but is %d", r[Pair{A: "Alice", B: "Bob"}])
	}
}

func TestParseInput(t *testing.T) {
	r := Parse(testinput)
	if len(r) != 12 {
		t.Errorf("expected len to be 12 is %d", len(r))
	}
	names := ReturnNames(r)
	if len(names) != 4 {
		t.Errorf("expected 4 names but got %d", len(names))
	}

	ring := CreateRing(names)
	if ring.Len() != len(names) {
		t.Errorf("expected %d names but got %d", len(names), ring.Len())
	}

	sum := Eval(ring, r)
	if sum != 330 {
		t.Errorf("expected %d sum but got %d", 330, sum)
	}
}

type Count struct {
	Counter int
}

func (c *Count) Inc(_ []int) {
	c.Counter++
}

func TestPermutations(t *testing.T) {
	c := Count{}
	v := make([]int, 5)
	Permutations(0, v, c.Inc)
	if c.Counter != 5*4*3*2 {
		t.Errorf("Permutations wrong")
	}
}

func TestRunAll(t *testing.T) {
	r := Parse(testinput)
	if len(r) != 12 {
		t.Errorf("expected len to be 12 is %d", len(r))
	}
	names := ReturnNames(r)
	if len(names) != 4 {
		t.Errorf("expected 4 names but got %d", len(names))
	}

	b := RunAll(names, r)
	if b != 330 {

		t.Errorf("expected 330 names but got %d", b)
	}

}
