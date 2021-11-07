package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	l := parse(input)
	if len(l) != 9 {
		t.Fatalf("expected 9 reindeers, got %d", len(l))
	}

}

func TestEval(t *testing.T) {
	s := evalDistance(1000, Record{Name: "Comet", Speed: 14, SpeedLength: 10, RestLength: 127})
	if s != 1120 {
		t.Fatalf("expected 1120 km, got %d", s)
	}
}

func TestScore(t *testing.T) {
	r1 := Record{Name: "Comet", Speed: 14, SpeedLength: 10, RestLength: 127}
	r2 := Record{Name: "Dancer", Speed: 16, SpeedLength: 11, RestLength: 162}
	s := scoreDistance(1000, []Record{r1, r2})
	if s != 689 {
		t.Fatalf("expected 689 points, got %d", s)
	}
}
