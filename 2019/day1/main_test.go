package main

import "testing"

func TestRocket(t *testing.T) {
	if rocket(1969) != 1969+966 {
		t.Error("1969")
	}
	if rocket(100756) != 100756+50346 {
		t.Error("100756")
	}

}
