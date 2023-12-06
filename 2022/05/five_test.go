package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestInput(t *testing.T) {
	buf, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatalf("read test.txt: %v", err)
	}
	top, bot := parse(string(buf))
	fmt.Printf("%+v\n", top)

	fmt.Printf("%+v\n", bot)

	processAll(&top, bot)
	fmt.Printf("%+v\n", top)
	st := result(top)
	if st != "CMZ" {
		t.Errorf("expected CMZ, got %s", st)
	}
}
func TestInput9001(t *testing.T) {
	buf, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatalf("read test.txt: %v", err)
	}
	top, bot := parse(string(buf))
	fmt.Printf("%+v\n", top)

	fmt.Printf("%+v\n", bot)

	processAll9001(&top, bot)
	fmt.Printf("%+v\n", top)
	st := result(top)
	if st != "MCD" {
		t.Errorf("expected MCD, got %s", st)
	}
}
