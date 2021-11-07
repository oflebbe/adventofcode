package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	coord := Coord{1, 1}
	c2 := follow(coord, 1)
	c3 := Coord{1, 1}
	if coord != c3 {
		t.Error("bla")
	}
	c3.y++
	if c2 != c3 {
		t.Error("bla")
	}
	fmt.Printf("%+v\n", coord)
	fmt.Printf("%+v\n", c2)
}
