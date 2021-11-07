package main

import "testing"

func TestMain(t *testing.T) {
	str := "R8,U5,L5,D3"
	grid := Grid{}
	grid.updateExtension(str)
	if grid.Xmin != 0 || grid.Xmax != 8 || grid.Ymin != 0 || grid.Ymax != 5 {
		t.Error("Main")
	}
	str2 := "U7,R6,D4,L4"
	grid.updateExtension(str2)
	if grid.Xmin != 0 || grid.Xmax != 8 || grid.Ymin != 0 || grid.Ymax != 7 {
		t.Error("Main")
	}
	grid.createGrid()

	grid.drawWire(str)
	grid.drawWire(str2)

	if grid.minCrossing != 6 {
		t.Error("Main")
	}
}

func TestMain2(t *testing.T) {
	/*	if support("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83").minCrossing != 159 {
		t.Error("Main21")
	} */
	if support("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7").minCrossing != 135 {
		t.Error("Main22")
	}

}
