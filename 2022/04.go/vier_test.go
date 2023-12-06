package main

import (
	"testing"
)

func TestVier(t *testing.T) {
	count1, count2 := find_elfs("test.txt")
	if count1 != 2 {
		t.Errorf("count1 %d != 2", count1)
	}
	if count2 != 4 {
		t.Errorf("count2 %d != 4", count2)
	}
}

func BenchmarkTask(b *testing.B) {
	find_elfs("input.txt")
	//fmt.Printf("%d %d", c1, c2)
	b.ReportAllocs()

}
