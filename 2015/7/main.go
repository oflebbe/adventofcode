package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	Lshift = iota
	Rshift
	Not
	Assign
	Connect
	UniAnd
)

var network map[string]Wire

var calculated map[string]*uint16

type Wire interface {
	result() uint16
}

type Uni struct {
	which       int
	howmuch     uint16
	source      string
	destination string
}

func (s Uni) result() uint16 {

	v, ok := calculated[s.destination]
	if ok && v == nil {
		log.Fatal("already started to calc")
	}
	if ok {
		return *v
	}
	calculated[s.destination] = nil
	fmt.Printf("U: %v\n", s)

	n, ok := network[s.source]
	if !ok && s.which != Assign {
		log.Fatal("Problem")
	}

	var res uint16
	if s.which == Assign {
		res = s.howmuch
	} else {

		res = n.result()
		if s.which == Lshift {
			res <<= s.howmuch
		} else if s.which == Rshift {
			res >>= s.howmuch
		} else if s.which == Not {
			res = ^res
		} else if s.which == UniAnd {
			res = s.howmuch & res
		} else if s.which == Connect {
			// res = res
		} else {
			log.Fatal("shift")
		}
	}
	calculated[s.destination] = &res
	return res
}

const (
	And = iota
	Or
)

type Binary struct {
	which       int
	source1     string
	source2     string
	destination string
}

func (b Binary) result() uint16 {
	v, ok := calculated[b.destination]
	if ok && v == nil {
		log.Fatal("already started to calc")
	}
	if ok {
		return *v
	}
	calculated[b.destination] = nil

	fmt.Printf("B: %v\n", b)
	n1, ok := network[b.source1]
	if !ok {
		log.Fatal("ok1")
	}
	n2, ok := network[b.source2]
	if !ok {
		log.Fatal("ok2")
	}
	s1 := n1.result()
	s2 := n2.result()
	var res uint16
	if b.which == And {
		res = s1 & s2
	} else if b.which == Or {
		res = s1 | s2
	} else {
		log.Fatal("binar")
	}
	calculated[b.destination] = &res
	return res
}

func main() {
	input, _ := os.Open("input.txt")
	line, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal("readall")
	}

	lines := strings.Split(string(line), "\n")
	network = make(map[string]Wire)

	for _, l := range lines {
		var s, d string
		var v uint16
		_, err := fmt.Sscanf(l, "NOT %s -> %s", &s, &d)
		if err == nil {
			network[d] = Uni{which: Not, source: s, destination: d}
			continue
		}
		var t string
		_, err = fmt.Sscanf(l, "%s RSHIFT %d -> %s", &s, &v, &d)
		if err == nil {

			network[d] = Uni{which: Rshift, howmuch: v, source: s, destination: d}
			continue
		}

		_, err = fmt.Sscanf(l, "%s LSHIFT %d -> %s", &s, &v, &d)
		if err == nil {

			network[d] = Uni{which: Lshift, howmuch: v, source: s, destination: d}
			continue
		}

		_, err = fmt.Sscanf(l, "%d AND %s -> %s", &v, &s, &d)
		if err == nil {
			network[d] = Uni{which: UniAnd, howmuch: v, source: s, destination: d}
			continue
		}
		var s2 string
		_, err = fmt.Sscanf(l, "%s %s %s -> %s", &s, &t, &s2, &d)
		if err == nil {
			var typ int
			if t == "AND" {
				typ = And
			} else if t == "OR" {
				typ = Or
			} else {
				log.Fatal("Bin")
			}
			network[d] = Binary{which: typ, source1: s, source2: s2, destination: d}
			continue
		}
		_, err = fmt.Sscanf(l, "%d -> %s", &v, &d)
		if err == nil {
			network[d] = Uni{which: Assign, howmuch: v, destination: d}
			continue
		}

		_, err = fmt.Sscanf(l, "%s -> %s", &s, &d)
		if err == nil {
			network[d] = Uni{which: Connect, source: s, destination: d}
			continue
		}

		log.Fatal("could not scan")
	}
	calculated = make(map[string]*uint16)
	fmt.Printf("Result %d", network["a"].result())

}
