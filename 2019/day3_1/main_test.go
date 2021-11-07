package main

import "testing"

func TestMain(t *testing.T) {
	str := "R8,U5,L5,D3"

	wire1 := Wire{}
	wire1.walk(str, func(x, y int) {})
	str2 := "U7,R6,D4,L4"
	wire2 := Wire{}
	wire1.nearCrossing = 10000
	wire2.walk(str2, func(x, y int) {
		wire1.testPoint(x, y)
	})
	if wire1.nearCrossing != 6 {
		t.Error("Main")
	}
}

func TestMainLen(t *testing.T) {
	str := "R8,U5,L5,D3"

	wire1 := Wire{}
	wire1.walk(str, func(x, y int) {})
	str2 := "U7,R6,D4,L4"
	wire2 := Wire{}
	wire1.nearCrossing = 10000
	wire2.walk(str2, func(x, y int) {
		wire1.testPointLen(x, y)
	})
	if wire1.nearCrossing != 30 {
		t.Error("Main")
	}
}

func TestMain2(t *testing.T) {

	str := "R75,D30,R83,U83,L12,D49,R71,U7,L72"

	wire1 := Wire{}
	wire1.walk(str, func(x, y int) {})
	str2 := "U62,R66,U55,R34,D71,R55,D58,R83"
	wire2 := Wire{}
	wire1.nearCrossing = 10000
	wire2.walk(str2, func(x, y int) {
		wire1.testPoint(x, y)
	})
	if wire1.nearCrossing != 159 {
		t.Error("Main")
	}

	/*
		if support("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7").minCrossing != 135 {
			t.Error("Main22")
		} */

}

func TestMain2Len(t *testing.T) {

	str := "R75,D30,R83,U83,L12,D49,R71,U7,L72"

	wire1 := Wire{}
	wire1.walk(str, func(x, y int) {})
	str2 := "U62,R66,U55,R34,D71,R55,D58,R83"
	wire2 := Wire{}
	wire1.nearCrossing = 10000
	wire2.walk(str2, func(x, y int) {
		wire1.testPointLen(x, y)
	})
	if wire1.nearCrossing != 610 {
		t.Error("Main")
	}

	/*
		if support("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7").minCrossing != 135 {
			t.Error("Main22")
		} */

}

func TestMain3(t *testing.T) {

	str := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"

	wire1 := Wire{}
	wire1.walk(str, func(x, y int) {})
	str2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	wire2 := Wire{}
	wire1.nearCrossing = 10000
	wire2.walk(str2, func(x, y int) {
		wire1.testPointLen(x, y)
	})
	if wire1.nearCrossing != 410 {
		t.Error("Main")
	}

}
