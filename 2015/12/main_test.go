package main

import (
	"encoding/json"
	"testing"
)

func TestOne(t *testing.T) {
	s := Sum("[1,2,3]")
	if s != 6 {
		t.Errorf("expected 6 got %d", s)
	}
}

func TestTwo(t *testing.T) {
	s := Sum("[-1,{\"a\":1}]")
	if s != 0 {
		t.Errorf("expected 0 got %d", s)
	}
}

func TestTwoTwo(t *testing.T) {

	s := RecurseSum("[-1,{\"a\":1}]")
	if s != 0 {
		t.Errorf("expected 0 got %d", s)
	}
	var tree interface{}
	json.Unmarshal([]byte(`{"d":"red","e":[1,2,3,4],"f":5}`), &tree)
	u := RecurseSum(tree)
	if u != 0 {
		t.Errorf("expected 0 got %d", s)
	}
}

func TestTwoThre(t *testing.T) {
	var tree interface{}
	json.Unmarshal([]byte(`[1,"red",5]`), &tree)
	s := RecurseSum(tree)
	if s != 6 {
		t.Errorf("expected 6 got %d", s)
	}
}
