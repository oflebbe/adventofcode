package main

import "testing"

func TestOne(t *testing.T) {
	input := []rune("hijklmmn")
	if Valid(input) {
		t.Fatalf("One")
	}
}

func TestTwo(t *testing.T) {
	input := []rune("abbceffg")
	if Valid(input) {
		t.Fatalf("Two")
	}

	input = []rune("abbcegjk")
	if Valid(input) {
		t.Fatalf("Three")
	}
}

func TestFout(t *testing.T) {
	input := []rune("abcdffaa")
	if !Valid(input) {
		t.Fatalf("Four")
	}
}

func TestFive(t *testing.T) {
	input := []rune("ghjaabcc")
	if !Valid(input) {
		t.Fatalf("Five")
	}
}



func TestNext(t *testing.T) {
	input := []rune("abcdefgh")
	wanted := "abcdffaa"

	if string(Next(input)) != wanted {
		t.Fatalf("Next1")
	}
}

func TestNext2(t *testing.T) {
	input := []rune("ghijklmn")
	wanted := "ghjaabcc"

	if string(Next(input)) != wanted {
		t.Fatalf("Next2")
	}
}
