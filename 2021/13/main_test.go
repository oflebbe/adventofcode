package main

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed "example.txt"
var example string

func TestReader(t *testing.T) {
	a, b := parse(example)
	fmt.Printf("%v\n%v\n", a, b)

}
