package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example string

func TestPart1(t *testing.T) {
	p, r := Parse(example)
	for i := 0; i < 10; i++ {
		p = Iter(p, r)
		l := Length(p)
		println(l)
	}
	s := Score(p)
	if s != 1588 {
		t.Errorf("Wrong score %d", s)
	}

}
